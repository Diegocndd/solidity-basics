//SPDX-License-Identifier: MIT

pragma solidity 0.8.30;

contract CustomWallet {
    function getBalance() public view returns (uint) {
        return address(this).balance;
    }

    function withdraw() public {
        address payable to = payable(msg.sender);
        to.transfer(address(this).balance);
    }

    receive() external payable {}
}
