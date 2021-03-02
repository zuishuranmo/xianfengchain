package main

import (
	"XianfengChain/chain"
	"fmt"
)

func main()  {
	fmt.Println("hello word")

	gensisblock := chain.CreateGenesisBlock([]byte("第一个区块"))
	fmt.Println("创世区块：",gensisblock)
	block1 := chain.CreateBlock(gensisblock.Height,gensisblock.Hash,nil)
	fmt.Println("区块1：",block1)
	block2 := chain.CreateBlock(block1.Height,block1.Hash,nil)
	fmt.Println("区块2：",block2)


}
