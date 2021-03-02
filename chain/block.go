package chain

import (
	"XianfengChain/consensus"
	"XianfengChain/utils"
	"bytes"
	"crypto/sha256"
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
	Hash    [32]byte
	//默克尔根
	Timestemp int64
	//Difficulty int64
	Nonce int64
	Data  []byte //区块体
}

/**
 * 该方法用于计算区块的hash值
 */
func (block *Block) SetHash()  {
	heightByte,_ := utils.IntToByte(block.Height)
	versionByte,_ := utils.IntToByte(block.Version)
	timeByte,_ := utils.IntToByte(block.Timestemp)
	nonceByte,_ := utils.IntToByte(block.Nonce)

	hashblock := bytes.Join([][]byte{heightByte,versionByte,block.PreHash[:],timeByte,nonceByte,block.Data},[]byte{})
	block.Hash = sha256.Sum256(hashblock)
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

	block.SetHash()//计算本身的hash值
	//尝试给nonce赋值
	//共识机制：pow, pos
	cons := consensus.NewPoW()
	cons.Run()

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
	genesis.SetHash()//计算本身的hash值
	return genesis
}
