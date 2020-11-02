package locks

import (
	"time"
)

type jsonLock struct {
	Hash       string    `json:"hash"`
	PublicKeys []string  `json:"pubkeys"`
	CreatedOn  time.Time `json:"created_on"`
}

func createJSONLockFromLock(ins Lock) *jsonLock {
	hash := ins.Hash().String()

	pubKeys := []string{}
	list := ins.PublicKeys()
	for _, onePubKey := range list {
		pubKeys = append(pubKeys, onePubKey.String())
	}

	createdOn := ins.CreatedOn()
	return createJSONLock(hash, pubKeys, createdOn)
}

func createJSONLock(
	hash string,
	pubKeys []string,
	createdOn time.Time,
) *jsonLock {
	out := jsonLock{
		Hash:       hash,
		PublicKeys: pubKeys,
		CreatedOn:  createdOn,
	}

	return &out
}
