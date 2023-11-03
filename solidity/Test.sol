// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Test {
    int64 private counter = 0;

    function count() public returns(bool) {
        counter++;
        return true;
    }

    function getCount() public view returns(int64) {
        return counter;
    }
}
