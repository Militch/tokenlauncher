package tokenlauncher

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go/token"
	"io/ioutil"
	"log"
	"net/rpc"
	"reflect"
	"strings"
)

var typeOfError = reflect.TypeOf((*error)(nil)).Elem()
type RPCServerStarter struct {
	rpcserver *rpc.Server
	serviceMap map[string]*service
}

type methodType struct {
	method     reflect.Method
	ArgType    reflect.Type
	ReplyType  reflect.Type
	numCalls   uint
}

type service struct {
	name   string                 // name of service
	rcvr   reflect.Value          // receiver of methods for the service
	typ    reflect.Type           // type of the receiver
	methods map[string]*methodType // registered methods
}

func (s *service) Call(methodName string, params []interface{}) (interface{}, error) {
	mtype := s.methods[methodName]
	function := mtype.method.Func

	argIsValue := false
	var argv reflect.Value
	if mtype.ArgType.Kind() == reflect.Ptr {
		argv = reflect.New(mtype.ArgType.Elem())
	}else {
		argv = reflect.New(mtype.ArgType)
		argIsValue = true
	}
	if argIsValue {
		argv = argv.Elem()
	}
	if len(params) != argv.NumField() {
		return nil, fmt.Errorf("params number is not nill")
	}
	for i:=0; i<argv.NumField(); i++ {
		an := argv.Field(i)

		an.Set(reflect.ValueOf(params[i]))
	}
	replyv := reflect.New(mtype.ReplyType.Elem())
	switch mtype.ReplyType.Elem().Kind() {
	case reflect.Map:
		replyv.Elem().Set(reflect.MakeMap(mtype.ReplyType.Elem()))
	case reflect.Slice:
		replyv.Elem().Set(reflect.MakeSlice(mtype.ReplyType.Elem(), 0, 0))
	}
	returnValues := function.Call([]reflect.Value{s.rcvr, argv, replyv})
	errInter := returnValues[0].Interface()
	if errInter != nil {
		errmsg := errInter.(error).Error()
		return nil, errors.New(errmsg)
	}
	return replyv.Interface(), nil
}


func NewRPCServerStarter() (*RPCServerStarter, error) {
	return &RPCServerStarter{
		serviceMap: make(map[string]*service),
	}, nil
}

func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return token.IsExported(t.Name()) || t.PkgPath() == ""
}
func suitableMethods(typ reflect.Type, reportErr bool) map[string]*methodType {
	methods := make(map[string]*methodType)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mtype := method.Type
		mname := method.Name
		if method.PkgPath != "" { continue }
		if mtype.NumIn() != 3 { continue }
		argType := mtype.In(1)
		if !isExportedOrBuiltinType(argType) { continue }
		replyType := mtype.In(2)
		if replyType.Kind() != reflect.Ptr { continue }
		if !isExportedOrBuiltinType(replyType) { continue }
		if mtype.NumOut() != 1 { continue }
		if returnType := mtype.Out(0); returnType != typeOfError { continue }
		methods[mname] = &methodType{
			method: method,
			ArgType: argType,
			ReplyType: replyType,
		}
	}
	return methods

}
func (server *RPCServerStarter) Register(rcvr interface{}) error  {
	return server.register(rcvr, "", false)
}

func (server *RPCServerStarter) RegisterName(name string,rcvr interface{}) error  {
	return server.register(rcvr, name, true)
}

func (server *RPCServerStarter) register(rcvr interface{}, name string, useName bool) error {
	s := new(service)
	s.typ = reflect.TypeOf(rcvr)
	s.rcvr = reflect.ValueOf(rcvr)
	sname := reflect.Indirect(s.rcvr).Type().Name()
	if useName {
		sname = name
	}
	if sname == "" {
		s := "rpc.Register: no service name for type " + s.typ.String()
		log.Print(s)
		return errors.New(s)
	}
	if !token.IsExported(sname) && !useName {
		s := "rpc.Register: type " + sname + " is not exported"
		log.Print(s)
		return errors.New(s)
	}
	s.name = sname
	s.methods = suitableMethods(s.typ, true)
	server.serviceMap[sname] = s

	return nil
}

func perParseRequestData(c *gin.Context) (map[string]interface{},error) {
	if "POST" != c.Request.Method {
		return nil, errors.New("POST method excepted")
	}
	if nil == c.Request.Body {
		return nil, fmt.Errorf("no POST data")
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading request body")
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if nil != err {
		return nil, fmt.Errorf("error parsing json request")
	}
	return data, nil
}


func (server *RPCServerStarter) handleJsonRPCRequest(c *gin.Context) (*string, interface{}, error)  {
	data, err := perParseRequestData(c)
	if err != nil {
		return nil, nil, err
	}
	id, ok := data["id"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("no or invalid 'id' in request")
	}
	if data["jsonrpc"] != "2.0" {
		return &id, nil, fmt.Errorf("version of jsonrpc is not 2.0")
	}
	method, ok := data["method"].(string)
	mpake := strings.Split(method,".")
	if !ok || len(mpake) != 2  {
		return &id,nil, fmt.Errorf("no or invalid 'method' in request")
	}
	log.Printf("method: %v",method)
	params, ok := data["params"].([]interface{})
	if !ok {
		return &id,nil, fmt.Errorf("no or invalid 'params' in request")
	}

	log.Printf("params: %v",params)
	service := server.serviceMap[mpake[0]]
	if service == nil {
		return &id,nil, fmt.Errorf("no or invalid 'method' in request")
	}
	result, err := service.Call(mpake[1], params)
	if err != nil {
		return &id, nil, err
	}
	return &id, result, nil
}





func (server *RPCServerStarter) Run(addr string) error {

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		id, data,err := server.handleJsonRPCRequest(c)
		if err != nil {
			c.JSON(200, gin.H{
				"jsonrpc": "2.0",
				"error": err.Error(),
				"id": id,
			})
			return
		}
		c.JSON(200, gin.H{
			"jsonrpc": "2.0",
			"id": id,
			"result": data,
		})
	})

	gin.SetMode(gin.ReleaseMode)
	err := r.Run(addr)
	return err
}