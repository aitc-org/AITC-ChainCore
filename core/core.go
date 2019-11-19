package main

import (
	"fmt"
	"lurcury/db"
	"lurcury/core/block"
	"lurcury/core/genesis"
	"lurcury/params"
	"lurcury/http"
	"lurcury/types"
	"time"
	"flag"
/*
        "os"
        "os/signal"
        "syscall"
*/
)

func main(){
        fmt.Println("-datadir string, -init bool , -port string, -chain string, max params three")
        //datadir := flag.String("datadir", "../dbdata", "Data dir")
        datadir := flag.String("datadir", "../dbdata", "Data dir")
        port := flag.String("port", "9000", "port ")
        chain := flag.String("chain", "BNN", "chain name ")
	init := flag.Bool("init", false, "init ")
        flag.Parse()

	fmt.Println(*init)
	config := params.Chain()
	config.Port =  *port
	config.Datadir = *datadir
	config.ChainName = *chain
	//fmt.Println("run",config)
        core_arg := &types.CoreStruct{}
	core_arg.Model = "0"
	core_arg.Config = config
        core_arg.Db = db.OpenDB("../../ccore"+config.Datadir)//*datadir)
	core_arg.NameDb = db.OpenDB("../../ccore"+config.Datadir+"/Name")
	go http.Server(core_arg)
	var tmpBlock types.BlockJson

	num := db.BlockNumberGet(core_arg.Db, "kaman")

	if (num == ""){
		genesis.InitAccount(*core_arg, genesis.GenesisBlock(config.Datadir))
		tmpBlock = genesis.GenesisBlock(config.Datadir)
        	db.BlockHexPut(core_arg.Db, tmpBlock.Hash, tmpBlock)
        	db.BlockNumberPut(core_arg.Db, "0", tmpBlock.Hash)
		db.BlockNumberPut(core_arg.Db, "kaman", "0")
	}else{
                num := db.BlockNumberGet(core_arg.Db, "kaman")
		hex := db.BlockNumberGet(core_arg.Db, num)
                tmpBlock = db.BlockHexGet(core_arg.Db, hex)
	}
	pri:=""
		for i:=0;i!=-1;{
			if(len(core_arg.PendingTransaction)>0){
				time.Sleep(1 * time.Second)
				fmt.Println(len(core_arg.PendingTransaction))
				time.Sleep(1 * time.Second)
				ff := time.Now()
				tmpBlock = block.CreateBlockPOA(core_arg, tmpBlock, pri)
				fmt.Println(time.Now().Sub(ff))
				//time.Sleep(1 * time.Second)
			}else{
				fmt.Println("no transaction")
				time.Sleep(1 * time.Second)
			}
		}
}
