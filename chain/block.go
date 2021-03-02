package chain

import (
	"time"
)

const VERSION = 0x00

/**
 * 区块的数据结构
 */
type Block struct {
	Height  int64 //高度
	Version int64 //版本号
	PreHash [32]byte
	//默克尔根
	Timestemp int64
	//Difficulty int64
	Nonce int64
	Data  []byte //区块体
}

/**
 * 创建新区块的函数
 */
func CreateBlock(height int64, preHash [32]byte, data []byte) Block {
	block := Block{}
	block.Height = height + 1
	block.PreHash = preHash
	block.Version = VERSION
	block.Timestemp = time.Now().Unix()
	block.Data = data

	return block
}

/**
 * 该函数用于生产创世区块
 */
func CreateGenesisBlock(data []byte) Block {
	genesis := Block{}
	genesis.Height = 0
	genesis.PreHash = [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	genesis.Version = VERSION
	genesis.Timestemp = time.Now().Unix()
	genesis.Data = data
	return genesis
}
