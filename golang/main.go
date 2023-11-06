package main

import (
	"math/big"

	medic "github.com/Lab-ICN/lab-icn-blockchain/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	wallet, err := medic.NewWallet()
	if err != nil {
		panic(err.Error())
	}
	client, err := ethclient.Dial("http://10.34.4.171:8545")
	if err != nil {
		panic(err.Error())
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, big.NewInt(8989))
	if err != nil {
		panic(err.Error())
	}
	medicRecordContract, err := medic.NewMedic(common.HexToAddress("0xf2140d33eb893bbb03395dffa4b67ef24924b708"), client)
	if err != nil {
		panic(err.Error())
	}
	medicRecordContract.AddPatient(transactor, common.HexToAddress("0x9ceFB7f35f6EF5E9b0021B13b6dd15e3b19d5CB8"), "Ananda Fitra")
}
