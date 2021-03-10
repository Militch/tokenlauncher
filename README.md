# TokenLauncher

## Build

```shell
solc --optimize --abi --bin contract/erc20_token.sol -o build --overwrite
solc --optimize --abi --bin contract/crowd_sale.sol -o build --overwrite
```

```shell
abigen --abi build/ERC20Token.abi --bin build/ERC20Token.bin --pkg contract --type ERC20Token --out contract/erc20_token.go
abigen --abi build/CrowdSale.abi --bin build/CrowdSale.bin --pkg contract --type CrowdSale --out contract/crowd_sale.go
```