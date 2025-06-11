package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/API-KEY")
	if err != nil {
		log.Fatal(err)
	}

	// Contract address
	contractAddress := common.HexToAddress("0x00e0bed3dd3b2b45af1792e4edf9f55515f15d58")

	contractABI, err := abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"Notify","type":"event"}]`))
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	lastBlock := header.Number

	fmt.Println("Polling for Notify events on Sepolia...")

	for {
		query := ethereum.FilterQuery{
			FromBlock: new(big.Int).Add(lastBlock, big.NewInt(1)),
			ToBlock:   nil,
			Addresses: []common.Address{contractAddress},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Println("Error fetching logs:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, vLog := range logs {
			var event struct {
				Sender common.Address
				Amount *big.Int
			}
			err := contractABI.UnpackIntoInterface(&event, "Notify", vLog.Data)
			if err != nil {
				log.Println("Error unpacking:", err)
				continue
			}
			fmt.Printf("Received deposit from: %s - Amount: %s wei\n", event.Sender.Hex(), event.Amount.String())

			if vLog.BlockNumber > lastBlock.Uint64() {
				lastBlock = big.NewInt(int64(vLog.BlockNumber))
			}
		}

		time.Sleep(5 * time.Second)
	}
}
