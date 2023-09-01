// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

import './interfaces/Ownable.sol';

struct DataPasien {
  uint256 noKtp;
  string  nama;
  string  alamat;
  string  jenisKelamin;
  string  golonganDarah;
  uint256 nomorTelepon;
  uint256 noBpjs;
}

contract MedicRecord {
  mapping(uint256 => DataPasien) private dataPasien;

  constructor() {

  }

  function getDataPasien(uint256 _noKtp) public view returns (DataPasien memory) {
    return dataPasien[_noKtp];
  }
}
