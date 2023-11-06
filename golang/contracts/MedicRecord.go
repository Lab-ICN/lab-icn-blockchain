// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package medic

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MedicRecordPatient is an auto generated low-level Go binding around an user-defined struct.
type MedicRecordPatient struct {
	Id            common.Address
	NoKtp         *big.Int
	Nama          string
	Alamat        string
	JenisKelamin  string
	GolonganDarah string
	NomorTelepon  *big.Int
	NoBpjs        *big.Int
	Records       []MedicRecordRecord
}

// MedicRecordRecord is an auto generated low-level Go binding around an user-defined struct.
type MedicRecordRecord struct {
	Cid       string
	PatientId common.Address
	DoctorId  common.Address
	TimeAdded *big.Int
}

// MedicMetaData contains all meta data concerning the Medic contract.
var MedicMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"doctorId\",\"type\":\"address\"}],\"name\":\"DoctorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hospitaId\",\"type\":\"address\"}],\"name\":\"HospitalAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"patientId\",\"type\":\"address\"}],\"name\":\"InsuranceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"patientId\",\"type\":\"address\"}],\"name\":\"PatientAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"patientId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"doctorId\",\"type\":\"address\"}],\"name\":\"RecordAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nama\",\"type\":\"string\"}],\"name\":\"addDoctor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hospitalId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_namaRs\",\"type\":\"string\"}],\"name\":\"addHospital\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_insuranceId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_namaInsurance\",\"type\":\"string\"}],\"name\":\"addInsurance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_patientId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_nama\",\"type\":\"string\"}],\"name\":\"addPatient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_patientId\",\"type\":\"address\"}],\"name\":\"addRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"doctors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nama\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_patientId\",\"type\":\"address\"}],\"name\":\"getPatientData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"noKtp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nama\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"alamat\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"jenisKelamin\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"golonganDarah\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nomorTelepon\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noBpjs\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"patientId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"doctorId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeAdded\",\"type\":\"uint256\"}],\"internalType\":\"structMedicRecord.Record[]\",\"name\":\"records\",\"type\":\"tuple[]\"}],\"internalType\":\"structMedicRecord.Patient\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_patientId\",\"type\":\"address\"}],\"name\":\"getPatientExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_patientId\",\"type\":\"address\"}],\"name\":\"getRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"patientId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"doctorId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeAdded\",\"type\":\"uint256\"}],\"internalType\":\"structMedicRecord.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSenderRole\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hospitals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nama\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"insurances\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nama\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"patients\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"noKtp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nama\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"alamat\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"jenisKelamin\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"golonganDarah\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nomorTelepon\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noBpjs\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061308d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063457bcea91161008c578063a0076e8411610066578063a0076e841461022a578063a9583c221461025a578063db66c0601461028b578063f093cf80146102a7576100cf565b8063457bcea9146101ae57806348e81852146101ca5780637bc769ed146101fa576100cf565b80630869cfbc146100d45780630be617001461010b578063172847b51461012757806326b5440a1461014357806327bae60a1461015f5780633a88162e14610190575b600080fd5b6100ee60048036038101906100e991906122b7565b6102d8565b60405161010298979695949392919061239c565b60405180910390f35b6101256004803603810190610120919061256b565b610560565b005b610141600480360381019061013c91906125b4565b61073b565b005b61015d60048036038101906101589190612610565b610a88565b005b610179600480360381019061017491906122b7565b610d34565b60405161018792919061266c565b60405180910390f35b610198610e00565b6040516101a5919061269c565b60405180910390f35b6101c860048036038101906101c39190612610565b610fe1565b005b6101e460048036038101906101df91906122b7565b61128d565b6040516101f1919061292c565b60405180910390f35b610214600480360381019061020f91906122b7565b611819565b6040516102219190612969565b60405180910390f35b610244600480360381019061023f91906122b7565b611982565b6040516102519190612a0a565b60405180910390f35b610274600480360381019061026f91906122b7565b611da2565b60405161028292919061266c565b60405180910390f35b6102a560048036038101906102a09190612610565b611e6e565b005b6102c160048036038101906102bc91906122b7565b612117565b6040516102cf92919061266c565b60405180910390f35b60006020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101549080600201805461032790612a5b565b80601f016020809104026020016040519081016040528092919081815260200182805461035390612a5b565b80156103a05780601f10610375576101008083540402835291602001916103a0565b820191906000526020600020905b81548152906001019060200180831161038357829003601f168201915b5050505050908060030180546103b590612a5b565b80601f01602080910402602001604051908101604052809291908181526020018280546103e190612a5b565b801561042e5780601f106104035761010080835404028352916020019161042e565b820191906000526020600020905b81548152906001019060200180831161041157829003601f168201915b50505050509080600401805461044390612a5b565b80601f016020809104026020016040519081016040528092919081815260200182805461046f90612a5b565b80156104bc5780601f10610491576101008083540402835291602001916104bc565b820191906000526020600020905b81548152906001019060200180831161049f57829003601f168201915b5050505050908060050180546104d190612a5b565b80601f01602080910402602001604051908101604052809291908181526020018280546104fd90612a5b565b801561054a5780601f1061051f5761010080835404028352916020019161054a565b820191906000526020600020905b81548152906001019060200180831161052d57829003601f168201915b5050505050908060060154908060070154905088565b3373ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610630576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062790612ad8565b60405180910390fd5b33600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010190816107009190612ca4565b507f69908b2e63170aa89ef89719bbe3b4c92cccc38653968ab25f4464eff7c28ac8336040516107309190612d76565b60405180910390a150565b3373ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461080b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080290612ddd565b60405180910390fd5b808073ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146108db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d290612e49565b60405180910390fd5b600060405180608001604052808581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020014281525090506000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060080181908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000190816109ac9190612ca4565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015550507f79439a93ba0b8067e556b3942a78ac58d991c58e2e23041c4e7aa1c446fc1188848433604051610a7a93929190612e69565b60405180910390a150505050565b3373ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4f90612ddd565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff16600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610c28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1f90612ef3565b60405180910390fd5b81600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001019081610cf89190612ca4565b507f0d017ca7e087763e37d963098c041704b8f461afbeb90695705d44a74f6639cb82604051610d289190612d76565b60405180910390a15050565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054610d7d90612a5b565b80601f0160208091040260200160405190810160405280929190818152602001828054610da990612a5b565b8015610df65780601f10610dcb57610100808354040283529160200191610df6565b820191906000526020600020905b815481529060010190602001808311610dd957829003601f168201915b5050505050905082565b60603373ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610ed4576040518060400160405280600681526020017f646f63746f7200000000000000000000000000000000000000000000000000008152509050610fde565b3373ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610fa5576040518060400160405280600781526020017f70617469656e74000000000000000000000000000000000000000000000000008152509050610fde565b6040518060400160405280600781526020017f756e6b6e6f776e0000000000000000000000000000000000000000000000000081525090505b90565b3373ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146110b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110a890612ddd565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611181576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161117890612f5f565b60405180910390fd5b81600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010190816112519190612ca4565b507f98b422b4248851ae9e727f5fd491b255462f5872bd0493c1ddf7308ccedfa1a2826040516112819190612d76565b60405180910390a15050565b6112956121e3565b3373ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611365576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161135c90612ddd565b60405180910390fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806101200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201805461141f90612a5b565b80601f016020809104026020016040519081016040528092919081815260200182805461144b90612a5b565b80156114985780601f1061146d57610100808354040283529160200191611498565b820191906000526020600020905b81548152906001019060200180831161147b57829003601f168201915b505050505081526020016003820180546114b190612a5b565b80601f01602080910402602001604051908101604052809291908181526020018280546114dd90612a5b565b801561152a5780601f106114ff5761010080835404028352916020019161152a565b820191906000526020600020905b81548152906001019060200180831161150d57829003601f168201915b5050505050815260200160048201805461154390612a5b565b80601f016020809104026020016040519081016040528092919081815260200182805461156f90612a5b565b80156115bc5780601f10611591576101008083540402835291602001916115bc565b820191906000526020600020905b81548152906001019060200180831161159f57829003601f168201915b505050505081526020016005820180546115d590612a5b565b80601f016020809104026020016040519081016040528092919081815260200182805461160190612a5b565b801561164e5780601f106116235761010080835404028352916020019161164e565b820191906000526020600020905b81548152906001019060200180831161163157829003601f168201915b50505050508152602001600682015481526020016007820154815260200160088201805480602002602001604051908101604052809291908181526020016000905b8282101561180a57838290600052602060002090600402016040518060800160405290816000820180546116c390612a5b565b80601f01602080910402602001604051908101604052809291908181526020018280546116ef90612a5b565b801561173c5780601f106117115761010080835404028352916020019161173c565b820191906000526020600020905b81548152906001019060200180831161171f57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152505081526020019060010190611690565b50505050815250509050919050565b60003373ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146118eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118e290612ddd565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16149050919050565b60603373ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480611aac57503373ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b611aeb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ae290612fcb565b60405180910390fd5b818073ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611bbb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bb290612e49565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600801805480602002602001604051908101604052809291908181526020016000905b82821015611d965783829060005260206000209060040201604051806080016040529081600082018054611c4f90612a5b565b80601f0160208091040260200160405190810160405280929190818152602001828054611c7b90612a5b565b8015611cc85780601f10611c9d57610100808354040283529160200191611cc8565b820191906000526020600020905b815481529060010190602001808311611cab57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152505081526020019060010190611c1c565b50505050915050919050565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054611deb90612a5b565b80601f0160208091040260200160405190810160405280929190818152602001828054611e1790612a5b565b8015611e645780601f10611e3957610100808354040283529160200191611e64565b820191906000526020600020905b815481529060010190602001808311611e4757829003601f168201915b5050505050905082565b3373ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611f3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f3590612ddd565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361200d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161200490613037565b60405180910390fd5b816000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190816120db9190612ca4565b507fa58cff07db0adda3ab7a40bd2f18ab57210bf38391a52ffd2e8eee6b8bdec71a8260405161210b9190612d76565b60405180910390a15050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461216090612a5b565b80601f016020809104026020016040519081016040528092919081815260200182805461218c90612a5b565b80156121d95780601f106121ae576101008083540402835291602001916121d9565b820191906000526020600020905b8154815290600101906020018083116121bc57829003601f168201915b5050505050905082565b604051806101200160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081526020016060815260200160608152602001606081526020016000815260200160008152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061228482612259565b9050919050565b61229481612279565b811461229f57600080fd5b50565b6000813590506122b18161228b565b92915050565b6000602082840312156122cd576122cc61224f565b5b60006122db848285016122a2565b91505092915050565b6122ed81612279565b82525050565b6000819050919050565b612306816122f3565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561234657808201518184015260208101905061232b565b60008484015250505050565b6000601f19601f8301169050919050565b600061236e8261230c565b6123788185612317565b9350612388818560208601612328565b61239181612352565b840191505092915050565b6000610100820190506123b2600083018b6122e4565b6123bf602083018a6122fd565b81810360408301526123d18189612363565b905081810360608301526123e58188612363565b905081810360808301526123f98187612363565b905081810360a083015261240d8186612363565b905061241c60c08301856122fd565b61242960e08301846122fd565b9998505050505050505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61247882612352565b810181811067ffffffffffffffff8211171561249757612496612440565b5b80604052505050565b60006124aa612245565b90506124b6828261246f565b919050565b600067ffffffffffffffff8211156124d6576124d5612440565b5b6124df82612352565b9050602081019050919050565b82818337600083830152505050565b600061250e612509846124bb565b6124a0565b90508281526020810184848401111561252a5761252961243b565b5b6125358482856124ec565b509392505050565b600082601f83011261255257612551612436565b5b81356125628482602086016124fb565b91505092915050565b6000602082840312156125815761258061224f565b5b600082013567ffffffffffffffff81111561259f5761259e612254565b5b6125ab8482850161253d565b91505092915050565b600080604083850312156125cb576125ca61224f565b5b600083013567ffffffffffffffff8111156125e9576125e8612254565b5b6125f58582860161253d565b9250506020612606858286016122a2565b9150509250929050565b600080604083850312156126275761262661224f565b5b6000612635858286016122a2565b925050602083013567ffffffffffffffff81111561265657612655612254565b5b6126628582860161253d565b9150509250929050565b600060408201905061268160008301856122e4565b81810360208301526126938184612363565b90509392505050565b600060208201905081810360008301526126b68184612363565b905092915050565b6126c781612279565b82525050565b6126d6816122f3565b82525050565b600082825260208201905092915050565b60006126f88261230c565b61270281856126dc565b9350612712818560208601612328565b61271b81612352565b840191505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000608083016000830151848203600086015261276f82826126ed565b915050602083015161278460208601826126be565b50604083015161279760408601826126be565b5060608301516127aa60608601826126cd565b508091505092915050565b60006127c18383612752565b905092915050565b6000602082019050919050565b60006127e182612726565b6127eb8185612731565b9350836020820285016127fd85612742565b8060005b85811015612839578484038952815161281a85826127b5565b9450612825836127c9565b925060208a01995050600181019050612801565b50829750879550505050505092915050565b60006101208301600083015161286460008601826126be565b50602083015161287760208601826126cd565b506040830151848203604086015261288f82826126ed565b915050606083015184820360608601526128a982826126ed565b915050608083015184820360808601526128c382826126ed565b91505060a083015184820360a08601526128dd82826126ed565b91505060c08301516128f260c08601826126cd565b5060e083015161290560e08601826126cd565b5061010083015184820361010086015261291f82826127d6565b9150508091505092915050565b60006020820190508181036000830152612946818461284b565b905092915050565b60008115159050919050565b6129638161294e565b82525050565b600060208201905061297e600083018461295a565b92915050565b600082825260208201905092915050565b60006129a082612726565b6129aa8185612984565b9350836020820285016129bc85612742565b8060005b858110156129f857848403895281516129d985826127b5565b94506129e4836127c9565b925060208a019950506001810190506129c0565b50829750879550505050505092915050565b60006020820190508181036000830152612a248184612995565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680612a7357607f821691505b602082108103612a8657612a85612a2c565b5b50919050565b7f5468697320646f63746f7220616c7265616479206578697374732e0000000000600082015250565b6000612ac2601b83612317565b9150612acd82612a8c565b602082019050919050565b60006020820190508181036000830152612af181612ab5565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302612b5a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612b1d565b612b648683612b1d565b95508019841693508086168417925050509392505050565b6000819050919050565b6000612ba1612b9c612b97846122f3565b612b7c565b6122f3565b9050919050565b6000819050919050565b612bbb83612b86565b612bcf612bc782612ba8565b848454612b2a565b825550505050565b600090565b612be4612bd7565b612bef818484612bb2565b505050565b5b81811015612c1357612c08600082612bdc565b600181019050612bf5565b5050565b601f821115612c5857612c2981612af8565b612c3284612b0d565b81016020851015612c41578190505b612c55612c4d85612b0d565b830182612bf4565b50505b505050565b600082821c905092915050565b6000612c7b60001984600802612c5d565b1980831691505092915050565b6000612c948383612c6a565b9150826002028217905092915050565b612cad8261230c565b67ffffffffffffffff811115612cc657612cc5612440565b5b612cd08254612a5b565b612cdb828285612c17565b600060209050601f831160018114612d0e5760008415612cfc578287015190505b612d068582612c88565b865550612d6e565b601f198416612d1c86612af8565b60005b82811015612d4457848901518255600182019150602085019450602081019050612d1f565b86831015612d615784890151612d5d601f891682612c6a565b8355505b6001600288020188555050505b505050505050565b6000602082019050612d8b60008301846122e4565b92915050565b7f53656e646572206973206e6f74206120646f63746f7200000000000000000000600082015250565b6000612dc7601683612317565b9150612dd282612d91565b602082019050919050565b60006020820190508181036000830152612df681612dba565b9050919050565b7f50617469656e7420646f6573206e6f7420657869737400000000000000000000600082015250565b6000612e33601683612317565b9150612e3e82612dfd565b602082019050919050565b60006020820190508181036000830152612e6281612e26565b9050919050565b60006060820190508181036000830152612e838186612363565b9050612e9260208301856122e4565b612e9f60408301846122e4565b949350505050565b7f5468697320696e737572616e636520616c7265616479206578697374732e0000600082015250565b6000612edd601e83612317565b9150612ee882612ea7565b602082019050919050565b60006020820190508181036000830152612f0c81612ed0565b9050919050565b7f5468697320686f73706974616c20616c7265616479206578697374732e000000600082015250565b6000612f49601d83612317565b9150612f5482612f13565b602082019050919050565b60006020820190508181036000830152612f7881612f3c565b9050919050565b7f53656e64657220646f6573206e6f742065786973740000000000000000000000600082015250565b6000612fb5601583612317565b9150612fc082612f7f565b602082019050919050565b60006020820190508181036000830152612fe481612fa8565b9050919050565b7f546869732070617469656e7420616c7265616479206578697374732e00000000600082015250565b6000613021601c83612317565b915061302c82612feb565b602082019050919050565b6000602082019050818103600083015261305081613014565b905091905056fea2646970667358221220a8f7aa37156d9ff8e73cbeff36a7c9153df3e4986ce7fc2ddfdbb0ab77a6815d64736f6c63430008120033",
}

// MedicABI is the input ABI used to generate the binding from.
// Deprecated: Use MedicMetaData.ABI instead.
var MedicABI = MedicMetaData.ABI

// MedicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MedicMetaData.Bin instead.
var MedicBin = MedicMetaData.Bin

// DeployMedic deploys a new Ethereum contract, binding an instance of Medic to it.
func DeployMedic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Medic, error) {
	parsed, err := MedicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MedicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Medic{MedicCaller: MedicCaller{contract: contract}, MedicTransactor: MedicTransactor{contract: contract}, MedicFilterer: MedicFilterer{contract: contract}}, nil
}

// Medic is an auto generated Go binding around an Ethereum contract.
type Medic struct {
	MedicCaller     // Read-only binding to the contract
	MedicTransactor // Write-only binding to the contract
	MedicFilterer   // Log filterer for contract events
}

// MedicCaller is an auto generated read-only Go binding around an Ethereum contract.
type MedicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MedicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MedicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MedicSession struct {
	Contract     *Medic            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MedicCallerSession struct {
	Contract *MedicCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MedicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MedicTransactorSession struct {
	Contract     *MedicTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedicRaw is an auto generated low-level Go binding around an Ethereum contract.
type MedicRaw struct {
	Contract *Medic // Generic contract binding to access the raw methods on
}

// MedicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MedicCallerRaw struct {
	Contract *MedicCaller // Generic read-only contract binding to access the raw methods on
}

// MedicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MedicTransactorRaw struct {
	Contract *MedicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMedic creates a new instance of Medic, bound to a specific deployed contract.
func NewMedic(address common.Address, backend bind.ContractBackend) (*Medic, error) {
	contract, err := bindMedic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Medic{MedicCaller: MedicCaller{contract: contract}, MedicTransactor: MedicTransactor{contract: contract}, MedicFilterer: MedicFilterer{contract: contract}}, nil
}

// NewMedicCaller creates a new read-only instance of Medic, bound to a specific deployed contract.
func NewMedicCaller(address common.Address, caller bind.ContractCaller) (*MedicCaller, error) {
	contract, err := bindMedic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MedicCaller{contract: contract}, nil
}

// NewMedicTransactor creates a new write-only instance of Medic, bound to a specific deployed contract.
func NewMedicTransactor(address common.Address, transactor bind.ContractTransactor) (*MedicTransactor, error) {
	contract, err := bindMedic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MedicTransactor{contract: contract}, nil
}

// NewMedicFilterer creates a new log filterer instance of Medic, bound to a specific deployed contract.
func NewMedicFilterer(address common.Address, filterer bind.ContractFilterer) (*MedicFilterer, error) {
	contract, err := bindMedic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MedicFilterer{contract: contract}, nil
}

// bindMedic binds a generic wrapper to an already deployed contract.
func bindMedic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MedicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Medic *MedicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Medic.Contract.MedicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Medic *MedicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medic.Contract.MedicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Medic *MedicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Medic.Contract.MedicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Medic *MedicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Medic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Medic *MedicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Medic *MedicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Medic.Contract.contract.Transact(opts, method, params...)
}

// Doctors is a free data retrieval call binding the contract method 0xa9583c22.
//
// Solidity: function doctors(address ) view returns(address id, string nama)
func (_Medic *MedicCaller) Doctors(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id   common.Address
	Nama string
}, error) {
	var out []interface{}
	err := _Medic.contract.Call(opts, &out, "doctors", arg0)

	outstruct := new(struct {
		Id   common.Address
		Nama string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Nama = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Doctors is a free data retrieval call binding the contract method 0xa9583c22.
//
// Solidity: function doctors(address ) view returns(address id, string nama)
func (_Medic *MedicSession) Doctors(arg0 common.Address) (struct {
	Id   common.Address
	Nama string
}, error) {
	return _Medic.Contract.Doctors(&_Medic.CallOpts, arg0)
}

// Doctors is a free data retrieval call binding the contract method 0xa9583c22.
//
// Solidity: function doctors(address ) view returns(address id, string nama)
func (_Medic *MedicCallerSession) Doctors(arg0 common.Address) (struct {
	Id   common.Address
	Nama string
}, error) {
	return _Medic.Contract.Doctors(&_Medic.CallOpts, arg0)
}

// GetPatientData is a free data retrieval call binding the contract method 0x48e81852.
//
// Solidity: function getPatientData(address _patientId) view returns((address,uint256,string,string,string,string,uint256,uint256,(string,address,address,uint256)[]))
func (_Medic *MedicCaller) GetPatientData(opts *bind.CallOpts, _patientId common.Address) (MedicRecordPatient, error) {
	var out []interface{}
	err := _Medic.contract.Call(opts, &out, "getPatientData", _patientId)

	if err != nil {
		return *new(MedicRecordPatient), err
	}

	out0 := *abi.ConvertType(out[0], new(MedicRecordPatient)).(*MedicRecordPatient)

	return out0, err

}

// GetPatientData is a free data retrieval call binding the contract method 0x48e81852.
//
// Solidity: function getPatientData(address _patientId) view returns((address,uint256,string,string,string,string,uint256,uint256,(string,address,address,uint256)[]))
func (_Medic *MedicSession) GetPatientData(_patientId common.Address) (MedicRecordPatient, error) {
	return _Medic.Contract.GetPatientData(&_Medic.CallOpts, _patientId)
}

// GetPatientData is a free data retrieval call binding the contract method 0x48e81852.
//
// Solidity: function getPatientData(address _patientId) view returns((address,uint256,string,string,string,string,uint256,uint256,(string,address,address,uint256)[]))
func (_Medic *MedicCallerSession) GetPatientData(_patientId common.Address) (MedicRecordPatient, error) {
	return _Medic.Contract.GetPatientData(&_Medic.CallOpts, _patientId)
}

// GetPatientExists is a free data retrieval call binding the contract method 0x7bc769ed.
//
// Solidity: function getPatientExists(address _patientId) view returns(bool)
func (_Medic *MedicCaller) GetPatientExists(opts *bind.CallOpts, _patientId common.Address) (bool, error) {
	var out []interface{}
	err := _Medic.contract.Call(opts, &out, "getPatientExists", _patientId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPatientExists is a free data retrieval call binding the contract method 0x7bc769ed.
//
// Solidity: function getPatientExists(address _patientId) view returns(bool)
func (_Medic *MedicSession) GetPatientExists(_patientId common.Address) (bool, error) {
	return _Medic.Contract.GetPatientExists(&_Medic.CallOpts, _patientId)
}

// GetPatientExists is a free data retrieval call binding the contract method 0x7bc769ed.
//
// Solidity: function getPatientExists(address _patientId) view returns(bool)
func (_Medic *MedicCallerSession) GetPatientExists(_patientId common.Address) (bool, error) {
	return _Medic.Contract.GetPatientExists(&_Medic.CallOpts, _patientId)
}

// GetRecords is a free data retrieval call binding the contract method 0xa0076e84.
//
// Solidity: function getRecords(address _patientId) view returns((string,address,address,uint256)[])
func (_Medic *MedicCaller) GetRecords(opts *bind.CallOpts, _patientId common.Address) ([]MedicRecordRecord, error) {
	var out []interface{}
	err := _Medic.contract.Call(opts, &out, "getRecords", _patientId)

	if err != nil {
		return *new([]MedicRecordRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]MedicRecordRecord)).(*[]MedicRecordRecord)

	return out0, err

}

// GetRecords is a free data retrieval call binding the contract method 0xa0076e84.
//
// Solidity: function getRecords(address _patientId) view returns((string,address,address,uint256)[])
func (_Medic *MedicSession) GetRecords(_patientId common.Address) ([]MedicRecordRecord, error) {
	return _Medic.Contract.GetRecords(&_Medic.CallOpts, _patientId)
}

// GetRecords is a free data retrieval call binding the contract method 0xa0076e84.
//
// Solidity: function getRecords(address _patientId) view returns((string,address,address,uint256)[])
func (_Medic *MedicCallerSession) GetRecords(_patientId common.Address) ([]MedicRecordRecord, error) {
	return _Medic.Contract.GetRecords(&_Medic.CallOpts, _patientId)
}

// GetSenderRole is a free data retrieval call binding the contract method 0x3a88162e.
//
// Solidity: function getSenderRole() view returns(string)
func (_Medic *MedicCaller) GetSenderRole(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Medic.contract.Call(opts, &out, "getSenderRole")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSenderRole is a free data retrieval call binding the contract method 0x3a88162e.
//
// Solidity: function getSenderRole() view returns(string)
func (_Medic *MedicSession) GetSenderRole() (string, error) {
	return _Medic.Contract.GetSenderRole(&_Medic.CallOpts)
}

// GetSenderRole is a free data retrieval call binding the contract method 0x3a88162e.
//
// Solidity: function getSenderRole() view returns(string)
func (_Medic *MedicCallerSession) GetSenderRole() (string, error) {
	return _Medic.Contract.GetSenderRole(&_Medic.CallOpts)
}

// Hospitals is a free data retrieval call binding the contract method 0xf093cf80.
//
// Solidity: function hospitals(address ) view returns(address id, string nama)
func (_Medic *MedicCaller) Hospitals(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id   common.Address
	Nama string
}, error) {
	var out []interface{}
	err := _Medic.contract.Call(opts, &out, "hospitals", arg0)

	outstruct := new(struct {
		Id   common.Address
		Nama string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Nama = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Hospitals is a free data retrieval call binding the contract method 0xf093cf80.
//
// Solidity: function hospitals(address ) view returns(address id, string nama)
func (_Medic *MedicSession) Hospitals(arg0 common.Address) (struct {
	Id   common.Address
	Nama string
}, error) {
	return _Medic.Contract.Hospitals(&_Medic.CallOpts, arg0)
}

// Hospitals is a free data retrieval call binding the contract method 0xf093cf80.
//
// Solidity: function hospitals(address ) view returns(address id, string nama)
func (_Medic *MedicCallerSession) Hospitals(arg0 common.Address) (struct {
	Id   common.Address
	Nama string
}, error) {
	return _Medic.Contract.Hospitals(&_Medic.CallOpts, arg0)
}

// Insurances is a free data retrieval call binding the contract method 0x27bae60a.
//
// Solidity: function insurances(address ) view returns(address id, string nama)
func (_Medic *MedicCaller) Insurances(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id   common.Address
	Nama string
}, error) {
	var out []interface{}
	err := _Medic.contract.Call(opts, &out, "insurances", arg0)

	outstruct := new(struct {
		Id   common.Address
		Nama string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Nama = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Insurances is a free data retrieval call binding the contract method 0x27bae60a.
//
// Solidity: function insurances(address ) view returns(address id, string nama)
func (_Medic *MedicSession) Insurances(arg0 common.Address) (struct {
	Id   common.Address
	Nama string
}, error) {
	return _Medic.Contract.Insurances(&_Medic.CallOpts, arg0)
}

// Insurances is a free data retrieval call binding the contract method 0x27bae60a.
//
// Solidity: function insurances(address ) view returns(address id, string nama)
func (_Medic *MedicCallerSession) Insurances(arg0 common.Address) (struct {
	Id   common.Address
	Nama string
}, error) {
	return _Medic.Contract.Insurances(&_Medic.CallOpts, arg0)
}

// Patients is a free data retrieval call binding the contract method 0x0869cfbc.
//
// Solidity: function patients(address ) view returns(address id, uint256 noKtp, string nama, string alamat, string jenisKelamin, string golonganDarah, uint256 nomorTelepon, uint256 noBpjs)
func (_Medic *MedicCaller) Patients(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id            common.Address
	NoKtp         *big.Int
	Nama          string
	Alamat        string
	JenisKelamin  string
	GolonganDarah string
	NomorTelepon  *big.Int
	NoBpjs        *big.Int
}, error) {
	var out []interface{}
	err := _Medic.contract.Call(opts, &out, "patients", arg0)

	outstruct := new(struct {
		Id            common.Address
		NoKtp         *big.Int
		Nama          string
		Alamat        string
		JenisKelamin  string
		GolonganDarah string
		NomorTelepon  *big.Int
		NoBpjs        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NoKtp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Nama = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Alamat = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.JenisKelamin = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.GolonganDarah = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.NomorTelepon = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.NoBpjs = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Patients is a free data retrieval call binding the contract method 0x0869cfbc.
//
// Solidity: function patients(address ) view returns(address id, uint256 noKtp, string nama, string alamat, string jenisKelamin, string golonganDarah, uint256 nomorTelepon, uint256 noBpjs)
func (_Medic *MedicSession) Patients(arg0 common.Address) (struct {
	Id            common.Address
	NoKtp         *big.Int
	Nama          string
	Alamat        string
	JenisKelamin  string
	GolonganDarah string
	NomorTelepon  *big.Int
	NoBpjs        *big.Int
}, error) {
	return _Medic.Contract.Patients(&_Medic.CallOpts, arg0)
}

// Patients is a free data retrieval call binding the contract method 0x0869cfbc.
//
// Solidity: function patients(address ) view returns(address id, uint256 noKtp, string nama, string alamat, string jenisKelamin, string golonganDarah, uint256 nomorTelepon, uint256 noBpjs)
func (_Medic *MedicCallerSession) Patients(arg0 common.Address) (struct {
	Id            common.Address
	NoKtp         *big.Int
	Nama          string
	Alamat        string
	JenisKelamin  string
	GolonganDarah string
	NomorTelepon  *big.Int
	NoBpjs        *big.Int
}, error) {
	return _Medic.Contract.Patients(&_Medic.CallOpts, arg0)
}

// AddDoctor is a paid mutator transaction binding the contract method 0x0be61700.
//
// Solidity: function addDoctor(string _nama) returns()
func (_Medic *MedicTransactor) AddDoctor(opts *bind.TransactOpts, _nama string) (*types.Transaction, error) {
	return _Medic.contract.Transact(opts, "addDoctor", _nama)
}

// AddDoctor is a paid mutator transaction binding the contract method 0x0be61700.
//
// Solidity: function addDoctor(string _nama) returns()
func (_Medic *MedicSession) AddDoctor(_nama string) (*types.Transaction, error) {
	return _Medic.Contract.AddDoctor(&_Medic.TransactOpts, _nama)
}

// AddDoctor is a paid mutator transaction binding the contract method 0x0be61700.
//
// Solidity: function addDoctor(string _nama) returns()
func (_Medic *MedicTransactorSession) AddDoctor(_nama string) (*types.Transaction, error) {
	return _Medic.Contract.AddDoctor(&_Medic.TransactOpts, _nama)
}

// AddHospital is a paid mutator transaction binding the contract method 0x457bcea9.
//
// Solidity: function addHospital(address _hospitalId, string _namaRs) returns()
func (_Medic *MedicTransactor) AddHospital(opts *bind.TransactOpts, _hospitalId common.Address, _namaRs string) (*types.Transaction, error) {
	return _Medic.contract.Transact(opts, "addHospital", _hospitalId, _namaRs)
}

// AddHospital is a paid mutator transaction binding the contract method 0x457bcea9.
//
// Solidity: function addHospital(address _hospitalId, string _namaRs) returns()
func (_Medic *MedicSession) AddHospital(_hospitalId common.Address, _namaRs string) (*types.Transaction, error) {
	return _Medic.Contract.AddHospital(&_Medic.TransactOpts, _hospitalId, _namaRs)
}

// AddHospital is a paid mutator transaction binding the contract method 0x457bcea9.
//
// Solidity: function addHospital(address _hospitalId, string _namaRs) returns()
func (_Medic *MedicTransactorSession) AddHospital(_hospitalId common.Address, _namaRs string) (*types.Transaction, error) {
	return _Medic.Contract.AddHospital(&_Medic.TransactOpts, _hospitalId, _namaRs)
}

// AddInsurance is a paid mutator transaction binding the contract method 0x26b5440a.
//
// Solidity: function addInsurance(address _insuranceId, string _namaInsurance) returns()
func (_Medic *MedicTransactor) AddInsurance(opts *bind.TransactOpts, _insuranceId common.Address, _namaInsurance string) (*types.Transaction, error) {
	return _Medic.contract.Transact(opts, "addInsurance", _insuranceId, _namaInsurance)
}

// AddInsurance is a paid mutator transaction binding the contract method 0x26b5440a.
//
// Solidity: function addInsurance(address _insuranceId, string _namaInsurance) returns()
func (_Medic *MedicSession) AddInsurance(_insuranceId common.Address, _namaInsurance string) (*types.Transaction, error) {
	return _Medic.Contract.AddInsurance(&_Medic.TransactOpts, _insuranceId, _namaInsurance)
}

// AddInsurance is a paid mutator transaction binding the contract method 0x26b5440a.
//
// Solidity: function addInsurance(address _insuranceId, string _namaInsurance) returns()
func (_Medic *MedicTransactorSession) AddInsurance(_insuranceId common.Address, _namaInsurance string) (*types.Transaction, error) {
	return _Medic.Contract.AddInsurance(&_Medic.TransactOpts, _insuranceId, _namaInsurance)
}

// AddPatient is a paid mutator transaction binding the contract method 0xdb66c060.
//
// Solidity: function addPatient(address _patientId, string _nama) returns()
func (_Medic *MedicTransactor) AddPatient(opts *bind.TransactOpts, _patientId common.Address, _nama string) (*types.Transaction, error) {
	return _Medic.contract.Transact(opts, "addPatient", _patientId, _nama)
}

// AddPatient is a paid mutator transaction binding the contract method 0xdb66c060.
//
// Solidity: function addPatient(address _patientId, string _nama) returns()
func (_Medic *MedicSession) AddPatient(_patientId common.Address, _nama string) (*types.Transaction, error) {
	return _Medic.Contract.AddPatient(&_Medic.TransactOpts, _patientId, _nama)
}

// AddPatient is a paid mutator transaction binding the contract method 0xdb66c060.
//
// Solidity: function addPatient(address _patientId, string _nama) returns()
func (_Medic *MedicTransactorSession) AddPatient(_patientId common.Address, _nama string) (*types.Transaction, error) {
	return _Medic.Contract.AddPatient(&_Medic.TransactOpts, _patientId, _nama)
}

// AddRecord is a paid mutator transaction binding the contract method 0x172847b5.
//
// Solidity: function addRecord(string _cid, address _patientId) returns()
func (_Medic *MedicTransactor) AddRecord(opts *bind.TransactOpts, _cid string, _patientId common.Address) (*types.Transaction, error) {
	return _Medic.contract.Transact(opts, "addRecord", _cid, _patientId)
}

// AddRecord is a paid mutator transaction binding the contract method 0x172847b5.
//
// Solidity: function addRecord(string _cid, address _patientId) returns()
func (_Medic *MedicSession) AddRecord(_cid string, _patientId common.Address) (*types.Transaction, error) {
	return _Medic.Contract.AddRecord(&_Medic.TransactOpts, _cid, _patientId)
}

// AddRecord is a paid mutator transaction binding the contract method 0x172847b5.
//
// Solidity: function addRecord(string _cid, address _patientId) returns()
func (_Medic *MedicTransactorSession) AddRecord(_cid string, _patientId common.Address) (*types.Transaction, error) {
	return _Medic.Contract.AddRecord(&_Medic.TransactOpts, _cid, _patientId)
}

// MedicDoctorAddedIterator is returned from FilterDoctorAdded and is used to iterate over the raw logs and unpacked data for DoctorAdded events raised by the Medic contract.
type MedicDoctorAddedIterator struct {
	Event *MedicDoctorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MedicDoctorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedicDoctorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MedicDoctorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MedicDoctorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedicDoctorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedicDoctorAdded represents a DoctorAdded event raised by the Medic contract.
type MedicDoctorAdded struct {
	DoctorId common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDoctorAdded is a free log retrieval operation binding the contract event 0x69908b2e63170aa89ef89719bbe3b4c92cccc38653968ab25f4464eff7c28ac8.
//
// Solidity: event DoctorAdded(address doctorId)
func (_Medic *MedicFilterer) FilterDoctorAdded(opts *bind.FilterOpts) (*MedicDoctorAddedIterator, error) {

	logs, sub, err := _Medic.contract.FilterLogs(opts, "DoctorAdded")
	if err != nil {
		return nil, err
	}
	return &MedicDoctorAddedIterator{contract: _Medic.contract, event: "DoctorAdded", logs: logs, sub: sub}, nil
}

// WatchDoctorAdded is a free log subscription operation binding the contract event 0x69908b2e63170aa89ef89719bbe3b4c92cccc38653968ab25f4464eff7c28ac8.
//
// Solidity: event DoctorAdded(address doctorId)
func (_Medic *MedicFilterer) WatchDoctorAdded(opts *bind.WatchOpts, sink chan<- *MedicDoctorAdded) (event.Subscription, error) {

	logs, sub, err := _Medic.contract.WatchLogs(opts, "DoctorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedicDoctorAdded)
				if err := _Medic.contract.UnpackLog(event, "DoctorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDoctorAdded is a log parse operation binding the contract event 0x69908b2e63170aa89ef89719bbe3b4c92cccc38653968ab25f4464eff7c28ac8.
//
// Solidity: event DoctorAdded(address doctorId)
func (_Medic *MedicFilterer) ParseDoctorAdded(log types.Log) (*MedicDoctorAdded, error) {
	event := new(MedicDoctorAdded)
	if err := _Medic.contract.UnpackLog(event, "DoctorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MedicHospitalAddedIterator is returned from FilterHospitalAdded and is used to iterate over the raw logs and unpacked data for HospitalAdded events raised by the Medic contract.
type MedicHospitalAddedIterator struct {
	Event *MedicHospitalAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MedicHospitalAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedicHospitalAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MedicHospitalAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MedicHospitalAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedicHospitalAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedicHospitalAdded represents a HospitalAdded event raised by the Medic contract.
type MedicHospitalAdded struct {
	HospitaId common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHospitalAdded is a free log retrieval operation binding the contract event 0x98b422b4248851ae9e727f5fd491b255462f5872bd0493c1ddf7308ccedfa1a2.
//
// Solidity: event HospitalAdded(address hospitaId)
func (_Medic *MedicFilterer) FilterHospitalAdded(opts *bind.FilterOpts) (*MedicHospitalAddedIterator, error) {

	logs, sub, err := _Medic.contract.FilterLogs(opts, "HospitalAdded")
	if err != nil {
		return nil, err
	}
	return &MedicHospitalAddedIterator{contract: _Medic.contract, event: "HospitalAdded", logs: logs, sub: sub}, nil
}

// WatchHospitalAdded is a free log subscription operation binding the contract event 0x98b422b4248851ae9e727f5fd491b255462f5872bd0493c1ddf7308ccedfa1a2.
//
// Solidity: event HospitalAdded(address hospitaId)
func (_Medic *MedicFilterer) WatchHospitalAdded(opts *bind.WatchOpts, sink chan<- *MedicHospitalAdded) (event.Subscription, error) {

	logs, sub, err := _Medic.contract.WatchLogs(opts, "HospitalAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedicHospitalAdded)
				if err := _Medic.contract.UnpackLog(event, "HospitalAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHospitalAdded is a log parse operation binding the contract event 0x98b422b4248851ae9e727f5fd491b255462f5872bd0493c1ddf7308ccedfa1a2.
//
// Solidity: event HospitalAdded(address hospitaId)
func (_Medic *MedicFilterer) ParseHospitalAdded(log types.Log) (*MedicHospitalAdded, error) {
	event := new(MedicHospitalAdded)
	if err := _Medic.contract.UnpackLog(event, "HospitalAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MedicInsuranceAddedIterator is returned from FilterInsuranceAdded and is used to iterate over the raw logs and unpacked data for InsuranceAdded events raised by the Medic contract.
type MedicInsuranceAddedIterator struct {
	Event *MedicInsuranceAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MedicInsuranceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedicInsuranceAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MedicInsuranceAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MedicInsuranceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedicInsuranceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedicInsuranceAdded represents a InsuranceAdded event raised by the Medic contract.
type MedicInsuranceAdded struct {
	PatientId common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInsuranceAdded is a free log retrieval operation binding the contract event 0x0d017ca7e087763e37d963098c041704b8f461afbeb90695705d44a74f6639cb.
//
// Solidity: event InsuranceAdded(address patientId)
func (_Medic *MedicFilterer) FilterInsuranceAdded(opts *bind.FilterOpts) (*MedicInsuranceAddedIterator, error) {

	logs, sub, err := _Medic.contract.FilterLogs(opts, "InsuranceAdded")
	if err != nil {
		return nil, err
	}
	return &MedicInsuranceAddedIterator{contract: _Medic.contract, event: "InsuranceAdded", logs: logs, sub: sub}, nil
}

// WatchInsuranceAdded is a free log subscription operation binding the contract event 0x0d017ca7e087763e37d963098c041704b8f461afbeb90695705d44a74f6639cb.
//
// Solidity: event InsuranceAdded(address patientId)
func (_Medic *MedicFilterer) WatchInsuranceAdded(opts *bind.WatchOpts, sink chan<- *MedicInsuranceAdded) (event.Subscription, error) {

	logs, sub, err := _Medic.contract.WatchLogs(opts, "InsuranceAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedicInsuranceAdded)
				if err := _Medic.contract.UnpackLog(event, "InsuranceAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInsuranceAdded is a log parse operation binding the contract event 0x0d017ca7e087763e37d963098c041704b8f461afbeb90695705d44a74f6639cb.
//
// Solidity: event InsuranceAdded(address patientId)
func (_Medic *MedicFilterer) ParseInsuranceAdded(log types.Log) (*MedicInsuranceAdded, error) {
	event := new(MedicInsuranceAdded)
	if err := _Medic.contract.UnpackLog(event, "InsuranceAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MedicPatientAddedIterator is returned from FilterPatientAdded and is used to iterate over the raw logs and unpacked data for PatientAdded events raised by the Medic contract.
type MedicPatientAddedIterator struct {
	Event *MedicPatientAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MedicPatientAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedicPatientAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MedicPatientAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MedicPatientAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedicPatientAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedicPatientAdded represents a PatientAdded event raised by the Medic contract.
type MedicPatientAdded struct {
	PatientId common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPatientAdded is a free log retrieval operation binding the contract event 0xa58cff07db0adda3ab7a40bd2f18ab57210bf38391a52ffd2e8eee6b8bdec71a.
//
// Solidity: event PatientAdded(address patientId)
func (_Medic *MedicFilterer) FilterPatientAdded(opts *bind.FilterOpts) (*MedicPatientAddedIterator, error) {

	logs, sub, err := _Medic.contract.FilterLogs(opts, "PatientAdded")
	if err != nil {
		return nil, err
	}
	return &MedicPatientAddedIterator{contract: _Medic.contract, event: "PatientAdded", logs: logs, sub: sub}, nil
}

// WatchPatientAdded is a free log subscription operation binding the contract event 0xa58cff07db0adda3ab7a40bd2f18ab57210bf38391a52ffd2e8eee6b8bdec71a.
//
// Solidity: event PatientAdded(address patientId)
func (_Medic *MedicFilterer) WatchPatientAdded(opts *bind.WatchOpts, sink chan<- *MedicPatientAdded) (event.Subscription, error) {

	logs, sub, err := _Medic.contract.WatchLogs(opts, "PatientAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedicPatientAdded)
				if err := _Medic.contract.UnpackLog(event, "PatientAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePatientAdded is a log parse operation binding the contract event 0xa58cff07db0adda3ab7a40bd2f18ab57210bf38391a52ffd2e8eee6b8bdec71a.
//
// Solidity: event PatientAdded(address patientId)
func (_Medic *MedicFilterer) ParsePatientAdded(log types.Log) (*MedicPatientAdded, error) {
	event := new(MedicPatientAdded)
	if err := _Medic.contract.UnpackLog(event, "PatientAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MedicRecordAddedIterator is returned from FilterRecordAdded and is used to iterate over the raw logs and unpacked data for RecordAdded events raised by the Medic contract.
type MedicRecordAddedIterator struct {
	Event *MedicRecordAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MedicRecordAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedicRecordAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MedicRecordAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MedicRecordAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedicRecordAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedicRecordAdded represents a RecordAdded event raised by the Medic contract.
type MedicRecordAdded struct {
	Cid       string
	PatientId common.Address
	DoctorId  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRecordAdded is a free log retrieval operation binding the contract event 0x79439a93ba0b8067e556b3942a78ac58d991c58e2e23041c4e7aa1c446fc1188.
//
// Solidity: event RecordAdded(string cid, address patientId, address doctorId)
func (_Medic *MedicFilterer) FilterRecordAdded(opts *bind.FilterOpts) (*MedicRecordAddedIterator, error) {

	logs, sub, err := _Medic.contract.FilterLogs(opts, "RecordAdded")
	if err != nil {
		return nil, err
	}
	return &MedicRecordAddedIterator{contract: _Medic.contract, event: "RecordAdded", logs: logs, sub: sub}, nil
}

// WatchRecordAdded is a free log subscription operation binding the contract event 0x79439a93ba0b8067e556b3942a78ac58d991c58e2e23041c4e7aa1c446fc1188.
//
// Solidity: event RecordAdded(string cid, address patientId, address doctorId)
func (_Medic *MedicFilterer) WatchRecordAdded(opts *bind.WatchOpts, sink chan<- *MedicRecordAdded) (event.Subscription, error) {

	logs, sub, err := _Medic.contract.WatchLogs(opts, "RecordAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedicRecordAdded)
				if err := _Medic.contract.UnpackLog(event, "RecordAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRecordAdded is a log parse operation binding the contract event 0x79439a93ba0b8067e556b3942a78ac58d991c58e2e23041c4e7aa1c446fc1188.
//
// Solidity: event RecordAdded(string cid, address patientId, address doctorId)
func (_Medic *MedicFilterer) ParseRecordAdded(log types.Log) (*MedicRecordAdded, error) {
	event := new(MedicRecordAdded)
	if err := _Medic.contract.UnpackLog(event, "RecordAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
