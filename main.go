package main

import (
	"XianfengChain/chain"
	"fmt"
)

func main()  {
	fmt.Println("hello word")

	block := chain.CreateGenesisBlock([]byte(string("第一个区块")))
	fmt.Println(block)

}
