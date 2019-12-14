package core

import (
	"bytes"
	"crypto/sha256"
	_ "crypto/sha256"
	"encoding/gob"
	"log"
	_ "strconv"
	"time"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64  //区块创建的时间戳
	Data          []byte //区块包含的数据
	DataHash      []byte //区块的数据Hash
	PrevBlockHash []byte //前一个区块的哈希值
	Hash          []byte //区块自身的哈希,用于效验区块数据有效
	Nonce         int
}

//Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result) // 创建一个编码器,把go语言的结构体转化为存粹的字节数组
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// 将字节数组反序列化为一个 Block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

//NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	dataHash := getDataHash(data)
	block := &Block{time.Now().Unix(), []byte(data), dataHash, prevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, Hash := pow.Run()
	block.Hash = Hash[:]
	block.Nonce = nonce
	return block
}

// NewGenesisBlock  creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//setHash calculates and sets block hash
// func (b *Block) setHash(block *Block) {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)
// 	b.Hash = hash[:]

// }
func getDataHash(data string) []byte {
	dataHash := sha256.Sum256([]byte(data))
	return dataHash[:]
}
