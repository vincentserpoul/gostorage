package ethereum

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Config is a conf for the ethereum node
type Config struct {
	Host            string
	JSONRPCPort     string
	JSONRPCProtocol string
	WSProtocol      string
	WSPort          string
	NetworkID       int64
}

// NewRPCCli connects to ethereum node and return a connection backend
func NewRPCCli(ethConf Config) (*ethclient.Client, error) {
	cli, err := ethclient.Dial(
		ethConf.JSONRPCProtocol +
			"://" + ethConf.Host +
			":" + ethConf.JSONRPCPort,
	)
	if err != nil {
		return nil, fmt.Errorf("NewRPCCli: rpc.Dial %v", err)
	}

	networkID, errNID := cli.NetworkID(context.Background())
	if errNID != nil {
		return nil, fmt.Errorf("NewRPCCli: NetworkID %v", errNID)
	}
	if networkID.Int64() != ethConf.NetworkID {
		return nil,
			fmt.Errorf(
				"NewRPCCli: NetworkID %d "+
					"different from wanted NetworkID %d",
				networkID, ethConf.NetworkID,
			)
	}

	return cli, nil
}

// NewWSCli connects to ethereum node and return a connection backend
func NewWSCli(ethConf Config) (*ethclient.Client, error) {
	cli, err := ethclient.Dial(
		ethConf.WSProtocol +
			"://" + ethConf.Host +
			":" + ethConf.WSPort,
	)
	if err != nil {
		return nil, fmt.Errorf("NewWSCli: rpc.Dial %v", err)
	}

	networkID, errNID := cli.NetworkID(context.Background())
	if errNID != nil {
		return nil, fmt.Errorf("NewWSCli: NetworkID %v", errNID)
	}
	if networkID.Int64() != ethConf.NetworkID {
		return nil,
			fmt.Errorf(
				"NewWSCli: NetworkID %d "+
					"different from wanted NetworkID %d",
				networkID, ethConf.NetworkID,
			)
	}

	return cli, nil
}
