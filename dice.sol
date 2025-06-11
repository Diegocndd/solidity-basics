// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

contract DiceContract {
    uint256 public prediction;
    uint256 public random;

    function bet(uint256 pred, bool under) public payable {
        require(
            (pred > 100 || pred < 1),
            "Invalid range for prediction (1 - 100)"
        );
        prediction = pred;

        random = generateRandom();

        bool won = under ? random < prediction : random > prediction;

        if (won) {
            uint256 range = under ? prediction : 100 - prediction;
            address payable to = payable(msg.sender);
            to.transfer(msg.value * (100 / range));
        }
    }

    // [UNSAFE] Use Chainlink VRF on mainnet!
    function generateRandom() private view returns (uint256) {
        uint256 randomNumber = uint256(
            keccak256(
                abi.encodePacked(block.timestamp, block.prevrandao, msg.sender)
            )
        ) % 100;

        return randomNumber;
    }
}
