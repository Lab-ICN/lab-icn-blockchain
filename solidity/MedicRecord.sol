// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

contract MedicRecord {
  struct Patient {
    address id;
    uint256 noKtp;
    string  nama;
    string  alamat;
    string  jenisKelamin;
    string  golonganDarah;
    uint256 nomorTelepon;
    uint256 noBpjs;
    Record[] records;
  }

  struct Record { 
    string cid;
    address patientId;
    address doctorId;
    uint8 temperature;
    uint8 sistol;
    uint8 diastol;
    uint256 timeAdded;
  }

  struct Doctor {
    address id;
    string  nama;
  }

  struct Hospital {
    address id;
    string  nama;
  }

  struct Insurance {
    address id;
    string  nama;
  }

  mapping (address => Patient) public patients;
  mapping (address => Doctor) public doctors;
  mapping (address => Hospital) public hospitals;
  mapping (address => Insurance) public insurances;

  event PatientAdded(address patientId);
  event HospitalAdded(address hospitaId);
  event InsuranceAdded(address patientId);
  event DoctorAdded(address doctorId);
  event RecordAdded(string cid, address patientId, address doctorId); 

  // modifiers

  modifier senderExists {
    require(doctors[msg.sender].id == msg.sender || patients[msg.sender].id == msg.sender, "Sender does not exist");
    _;
  }

  modifier patientExists(address patientId) {
    require(patients[patientId].id == patientId, "Patient does not exist");
    _;
  }

  modifier senderIsDoctor {
    require(doctors[msg.sender].id == msg.sender, "Sender is not a doctor");
    _;
  }

  // functions

  function addPatient(address _patientId, string memory _nama) public senderIsDoctor {
    require(patients[_patientId].id != _patientId, "This patient already exists.");
    patients[_patientId].id = _patientId;
    patients[_patientId].nama = _nama;

    emit PatientAdded(_patientId);
  }

  function addHospital(address _hospitalId, string memory _namaRs) public senderIsDoctor {
    require(hospitals[_hospitalId].id != _hospitalId, "This hospital already exists.");
    hospitals[_hospitalId].id = _hospitalId;
    hospitals[_hospitalId].nama = _namaRs;

    emit HospitalAdded(_hospitalId);
  }
  
  function addInsurance(address _insuranceId, string memory _namaInsurance) public {
    require(insurances[_insuranceId].id != _insuranceId, "This insurance already exists.");
    insurances[_insuranceId].id = _insuranceId;
    insurances[_insuranceId].nama = _namaInsurance;

    emit InsuranceAdded(_insuranceId);
  }

  function addDoctor(string memory _nama) public {
    require(doctors[msg.sender].id != msg.sender, "This doctor already exists.");
    doctors[msg.sender].id = msg.sender;
    doctors[msg.sender].nama = _nama;

    emit DoctorAdded(msg.sender);
  }

  function addRecord(string memory _cid, address _patientId, address _doctorId, uint8 temperature, uint8 diastol, uint8 sistol) public senderIsDoctor patientExists(_patientId) {
    Record memory record = Record(_cid, _patientId, _doctorId, temperature,sistol, diastol,block.timestamp);
    patients[_patientId].records.push(record);

    emit RecordAdded(_cid, _patientId,_doctorId);
  } 

  function getRecords(address _patientId) public view senderExists patientExists(_patientId) returns (Record[] memory) {
    return patients[_patientId].records;
  } 

  function getSenderRole() public view returns (string memory) {
    if (doctors[msg.sender].id == msg.sender) {
      return "doctor";
    } else if (patients[msg.sender].id == msg.sender) {
      return "patient";
    } else {
      return "unknown";
    }
  }

  function getPatientExists(address _patientId) public view senderIsDoctor returns (bool) {
    return patients[_patientId].id == _patientId;
  }

  function getPatientData(address _patientId) public view senderIsDoctor returns (Patient memory) {
    return patients[_patientId];
  }
}
