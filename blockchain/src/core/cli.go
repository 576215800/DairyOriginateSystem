package core

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	Bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" addblock -data BLOCK_DATA -add a block to the blockchain")
	fmt.Println("printchain - print all the blocks of the blockchain")
}
func (cli *CLI) valiateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.valiateArgs()
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printBlockCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block Data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}

		cli.addBlock(*addBlockData)
	}
	if printBlockCmd.Parsed() {
		cli.printChain()
	}
}
