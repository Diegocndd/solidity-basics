//SPDX-License-Identifier: MIT

pragma solidity 0.8.30;

contract PayableContract {
    uint public gasReceived;

    receive() external payable {
        gasReceived = gasleft();
    }
}
