package common

import (
	"github.com/EducationEKT/EKT/io/ekt8/db"
)

type BlockChain struct {
	db db.EKTDB
}

func (blockChain *BlockChain) addBlock(block Block) {
}
