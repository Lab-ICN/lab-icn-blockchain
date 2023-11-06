package main

import (
	"context"
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
	caller := &bind.CallOpts{
		From:    transactor.From,
		Context: context.Background(),
	}

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

	handleRequests(medicRecordContract, transactor, caller)
}

func handleRequests(medicRecordContract *medic.Medic, transactor *bind.TransactOpts, caller *bind.CallOpts) {
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
			"message": fmt.Sprint("Doctor  has been added, txHash:", tx.Hash().Hex()),
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

		tx, err := medicRecordContract.AddRecord(transactor, "Record-02", common.HexToAddress(address), transactor.From, temperature, sistol, diastol)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("TxHash:", tx.Hash().Hex())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprint("Record has been added, txHash:", tx.Hash().Hex()),
		})
	})

	// still status 500
	r.GET("/get-sender-role", func(c *gin.Context) {
		// fmt.Println(caller.From)
		// role, err := medicRecordContract.GetSenderRole(caller)
		role, err := medicRecordContract.GetSenderRole(caller)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"sender-role": role,
			})
		}
	})

	r.POST("/add-insurance", func(c *gin.Context) {
		address := c.PostForm("address")
		namaInsurance := c.PostForm("insuranceName")

		tx, err := medicRecordContract.AddInsurance(transactor, common.HexToAddress(address), namaInsurance)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("TxHash:", tx.Hash().Hex())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprint("Insurance ", namaInsurance, " has been added, txHash: ", tx.Hash().Hex()),
		})
	})

	r.POST("/get-records", func(c *gin.Context) {
		patientId := c.PostForm("patientId")
		records, err := medicRecordContract.GetRecords(caller, common.HexToAddress(patientId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {

			c.JSON(http.StatusOK, gin.H{
				"records": records,
			})
		}
	})
	r.Run()
}
