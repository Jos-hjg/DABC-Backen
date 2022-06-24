package contract

import (
	"dabc/config"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strconv"
)

var Client *DABC

func Init()  {
	conn, err := ethclient.Dial(config.C.Chain.Rpc + ":" + config.C.Chain.Port)
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		return
	}
	//prikey, err := crypto.GenerateKey()
	//Pubkey := prikey.Public()
	contract := common.HexToAddress(config.C.Chain.Contract.Address)
	client, err := NewDABC(contract, conn)
	Client = client
	ts, err := client.GetTimestamp(nil)
	if err != nil {
		log.Fatalln(err)
	}
	tt, _ := strconv.Atoi(ts.String())
	log.Println("current", tt)
}

