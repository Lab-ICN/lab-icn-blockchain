package main

import (
	"math/big"

	medic "github.com/Lab-ICN/lab-icn-blockchain/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	gin "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	godotenv.Load(".env")
	wallet, err := medic.NewWallet()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Initalization Wallet Successful")
	}
	client, err := ethclient.Dial("http://10.34.4.171:8545")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Initalization ETH Client Successful")
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, big.NewInt(8989))
	transactor.GasLimit = 3000000

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Initalization Transactor Successful")
	}
	medicRecordContract, err := medic.NewMedic(common.HexToAddress("0x8c57BbFa3e5F3783BFf46A815BCF5Ccb31A1be99"), client)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Initalization Smart Contract Successful")
	}

	handleRequests(medicRecordContract, transactor)
}

func handleRequests(medicRecordContract *medic.Medic, transactor *bind.TransactOpts) {
	r := gin.Default()

	r.POST("/add-patient", func(c *gin.Context) {
		address := c.PostForm("address")
		name := c.PostForm("name")

		tx, err := medicRecordContract.AddPatient(transactor, common.HexToAddress(address), name)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("TxHash:", tx.Hash().Hex())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprint("Patient  has been added, txHash:", tx.Hash().Hex()),
		})
	})

	r.POST("/add-doctor", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println(name)

		tx, err := medicRecordContract.AddDoctor(transactor, name)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("TxHash:", tx.Hash().Hex())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprint("Patient  has been added, txHash:", tx.Hash().Hex()),
		})
	})

	r.POST("/add-record", func(c *gin.Context) {
		address := c.PostForm("address")
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

		fmt.Println(temperatureUint, sistolUint, diastolUint)

		temperature := uint8(temperatureUint)
		sistol := uint8(sistolUint)
		diastol := uint8(diastolUint)

		tx, err := medicRecordContract.AddRecord(transactor, "Record-01", common.HexToAddress(address), transactor.From, temperature, sistol, diastol)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("TxHash:", tx.Hash().Hex())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprint("Record has been added, txHash:", tx.Hash().Hex()),
		})
	})
	r.Run()
}
