package main

import (
	"math/big"

	medic "github.com/Lab-ICN/lab-icn-blockchain/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	gin "github.com/gin-gonic/gin"

	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	wallet, err := medic.NewWallet()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Initalization Wallet Successful")
	}
	client, err := ethclient.Dial("http://10.34.4.172:8545")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Initalization ETH Client Successful")
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, big.NewInt(8989))
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Initalization Transactor Successful")
	}
	medicRecordContract, err := medic.NewMedic(common.HexToAddress("0xf2140d33eb893bbb03395dffa4b67ef24924b708"), client)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Initalization Smart Contract Successful")
	}
	medicRecordContract.AddPatient(transactor, common.HexToAddress("0x9ceFB7f35f6EF5E9b0021B13b6dd15e3b19d5CB8"), "Ananda Fitra")

	handleRequests(medicRecordContract, transactor)
}

func handleRequests(medicRecordContract *medic.Medic, transactor *bind.TransactOpts) {
	r := gin.Default()

	r.GET("/add-patient", func(c *gin.Context) {
		name := c.PostForm("name")

		medicRecordContract.AddPatient(transactor, common.HexToAddress("0x9ceFB7f35f6EF5E9b0021B13b6dd15e3b19d5CB8"), name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Patient  has been added",
		})
	})

	r.GET("/add-doctor", func(c *gin.Context) {
		name := c.PostForm("name")

		medicRecordContract.AddPatient(transactor, common.HexToAddress("0x9ceFB7f35f6EF5E9b0021B13b6dd15e3b19d5CB8"), name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Patient  has been added",
		})
	})

	r.GET("/add-record", func(c *gin.Context) {
		temperatureUint, err := strconv.ParseUint(c.PostForm("temperature"), 10, 8)
		if err != nil {
			log.Fatal(err)
		}
		sistolUint, err := strconv.ParseUint(c.PostForm("sistol"), 10, 8)
		if err != nil {
			log.Fatal(err)
		}
		diastolUint, err := strconv.ParseUint(c.PostForm("diastol"), 10, 8)
		if err != nil {
			log.Fatal(err)
		}

		temperature := uint8(temperatureUint)
		sistol := uint8(sistolUint)
		diastol := uint8(diastolUint)

		medicRecordContract.AddRecord(transactor, "Record-01", common.HexToAddress("0x9ceFB7f35f6EF5E9b0021B13b6dd15e3b19d5CB8"), common.HexToAddress("0x9ceFB7f35f6EF5E9b0021B13b6dd15e3b19d5CB8"), temperature, sistol, diastol)
		c.JSON(http.StatusOK, gin.H{
			"message": "Record has been added",
		})
	})
	r.Run()
}
