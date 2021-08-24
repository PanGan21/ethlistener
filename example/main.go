package main

import (
	"context"
	"crypto/ecdsa"
	"ethlistener"
	"ethlistener/contracts"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Oops! There was a problem", err)
	}

	privateKey, err := crypto.HexToECDSA("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	auth := getAuth(privateKey, nonce)
	addr, _, con, err := contracts.DeployEventExample(auth, client)
	contract, _ := ethlistener.NewContract("EventExample", contracts.EventExampleABI, addr)

	var contractsDeployed []*ethlistener.Contract
	contractsDeployed = append(contractsDeployed, contract)

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	eventCh := make(chan *ethlistener.ContractEvent, 1)
	stopCh := make(chan struct{}, 1)

	listener := ethlistener.NewEventListener(client, contractsDeployed)
	go listener.Listen(header.Number, eventCh, stopCh)
	defer close(stopCh)
	for {
		select {
		case event := <-eventCh:
			log.Println("Event name: ", event.Name, " Event data: ", event.Data)
		case <-time.After(3 * time.Second):
			new_nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			new_auth := getAuth(privateKey, new_nonce)
			res, err := con.StoreData(new_auth, big.NewInt(123456), big.NewInt(998908098098))
			if err != nil {
				log.Fatal(err)
			}

			log.Println(res)
		}
	}

}

func getAuth(privateKey *ecdsa.PrivateKey, new_nonce uint64) *bind.TransactOpts {
	new_auth := bind.NewKeyedTransactor(privateKey)
	new_auth.Nonce = big.NewInt(int64(new_nonce))
	new_auth.Value = big.NewInt(0)     // in wei
	new_auth.GasLimit = uint64(300000) // in units
	new_auth.GasPrice = big.NewInt(30000000000)

	return new_auth
}
