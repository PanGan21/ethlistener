test:
	cd manager && go test -v -cover

ganache:
	ganache-cli --account="0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef,1000000000000000000000000000000000000000" --mnemonic="range pear quit paddle harvest glory insect tissue erupt spray sport child" 

compile-contract:
	solc --optimize --abi ./contracts/EventExample.sol -o build

generate-client-contract:
	abigen --sol=contracts/EventExample.sol --pkg=contracts --out=contracts/EventExample.go

.PHONY: test, ganache, compile-contract, generate-client-contract