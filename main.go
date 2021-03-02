package main

import (
	"XianfengChain/chain"
	"fmt"
)

func main()  {
	fmt.Println("hello word")

	block := chain.CreateGenesisBlock([]byte{})
	fmt.Println(block)

}
