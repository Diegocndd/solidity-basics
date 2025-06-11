// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

contract DepositNotifier {
    event Notify(address sender, uint amount);

    receive() external payable {
        emit Notify(msg.sender, msg.value);
    }
}
