package consensus

import "XianfengChain/consensus/pow"

/**
 * consensus作为接口，用于定义共识机制的标准
 */
type Consensus interface {
	Run() interface{}
}

func NewPoW() Consensus {
	proof := pow.POW{}
	return proof
}


