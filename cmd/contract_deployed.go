package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func processContractDeployedIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := logs.NewContractDeployedIndexer()
	return processLogIndexer(client, processor, blockNumber)
}