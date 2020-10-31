package miners

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	"github.com/xmn-services/rod-network/domain/memory/piastres/chains"
	"github.com/xmn-services/rod-network/libs/hash"
)

// difficulty calculates the next block difficulty
func difficulty(chain chains.Chain, amountTrx uint) uint {
	blockDifficulty := chain.Genesis().Difficulty().Block()
	base := float64(blockDifficulty.Base())
	increasePerTrx := blockDifficulty.IncreasePerTrx()

	sum := float64(0)
	for i := 0; i < int(amountTrx); i++ {
		index := float64(i + 1)
		sum += (index * increasePerTrx)
	}

	return uint(sum + base)
}

// prefix returns the prefix based on the difficulty
func prefix(difficulty uint) ([]byte, error) {
	// calculate the requested data:
	var data = []interface{}{}
	for i := 0; i < int(difficulty); i++ {
		data = append(data, int8(miningBeginValue))
	}

	// create the begin bytes buffer:
	buf := new(bytes.Buffer)
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			return nil, err
		}
	}

	// create the requested prefix:
	return buf.Bytes(), nil
}

func mine(
	hashAdapter hash.Adapter,
	difficulty uint,
	hsh hash.Hash,
) (string, error) {
	// create the requested prefix:
	requestedPrefix, err := prefix(difficulty)
	if err != nil {
		return "", err
	}

	// execute the mining:
	return mineRecursively(
		hashAdapter,
		requestedPrefix,
		hsh.Bytes(),
		"",
	)
}

func mineRecursively(
	hashAdapter hash.Adapter,
	requestedPrefix []byte,
	baseData []byte,
	baseTries string,
) (string, error) {
	baseWithTries := [][]byte{
		baseData,
	}

	if baseTries != "" {
		baseWithTries = append(baseWithTries, []byte(baseTries))
	}

	for i := uint(0); i <= maxMiningValue; i++ {
		try := baseWithTries
		try = append(try, []byte(strconv.Itoa(int(i))))
		res, err := hashAdapter.FromMultiBytes(try)
		if err != nil {
			return "", err
		}

		if bytes.HasPrefix(res.Bytes(), requestedPrefix) {
			baseTries = fmt.Sprintf("%s%d", baseTries, i)
			return baseTries, nil
		}

	}

	// none of the tries work, so try with an additionral base try:
	for i := uint(0); i < maxMiningTries; i++ {
		additional := fmt.Sprintf("%s%d", baseTries, i)
		results, err := mineRecursively(hashAdapter, requestedPrefix, baseData, additional)
		if err != nil {
			return "", err
		}

		if results != "" {
			return results, nil
		}
	}

	return "", errors.New("the mining was impossible")
}
