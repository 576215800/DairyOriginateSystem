package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 20 //目标位20     4 一个0  8 两个0  20 5个0

//ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofofwork builds and returns a ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	fmt.Println("target:", target)
	target.Lsh(target, uint(256-targetBits)) // 移位操作   前面的字节变为0 20个位变成0
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \" %s \"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		//fmt.Printf("\r %x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r %x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

// 将一个 int64 转化为一个字节数组(byte array)
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}
