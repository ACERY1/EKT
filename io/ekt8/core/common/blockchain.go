package common

import (
	"../../db"
	"github.com/EducationEKT/EKT/io/ekt8/core/common/block"
)

type BlockChain struct {
	db db.EKTDB
}

func (blockChain *BlockChain) addBlock(block block.Block) {
}
