package core

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db" // 文件就是一个本地数据库
const blocksBucket = "blocks"

//tip 存储最后一个块的hash
//在链的末端可能出现短暂分叉的情况,所以选择tip其实也是选择了哪条链
//db 存储数据库链接
//Blockchain keeps a sequence of Blocks
type Blockchain struct {
	tip []byte
	Db  *bolt.DB
}

var bc *Blockchain

func NewBlockchain() {
	var tip []byte

	//打开一个BoltDB文件
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error { // 向文件中提交数据
		b := tx.Bucket([]byte(blocksBucket)) //有没有这么一个key-value的数组

		// 如果数据库中不存在区块链就创建一个,否则直接读取最后一个块的哈希
		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize()) // key value  字节数组   结构体序列化
			if err != nil {
				log.Panic(err)
			}
			err = b.Put([]byte("l"), genesis.Hash) //leader 创世区块的哈希
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l")) //tip存领头的哈希值
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc = &Blockchain{tip, db}

}

//AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	//首先获取最后一个块的哈希用于生成新的哈希
	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	newBlock := NewBlock(data, lastHash)

	err = bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash

		return nil
	})
}

//BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	currentHash []byte
	Db          *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.Db}
	return bci
}

//返回链中的下一个块
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)
		return nil

	})
	if err != nil {
		log.Panic(err)
	}
	i.currentHash = block.PrevBlockHash
	return block
}
