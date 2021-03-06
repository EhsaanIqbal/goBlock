package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

//Difficulty level of problem [static for now]
const Difficulty = 12

//ProofOfWork struct
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

//NewProof
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

//InitData
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(

		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToByte(int64(nonce)),
			ToByte(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

//Validate block
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

//Run Proof-Of-Work
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

//ToByte - utility func
func ToByte(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	Handle(err)

	return buff.Bytes()
}
