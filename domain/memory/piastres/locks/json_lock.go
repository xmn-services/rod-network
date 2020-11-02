package locks

import (
	"time"
)

// JSONLock represents a JSON lock instance
type JSONLock struct {
	PublicKeys []string  `json:"pubkeys"`
	CreatedOn  time.Time `json:"created_on"`
}

func createJSONLockFromLock(ins Lock) *JSONLock {
	publicKeys := []string{}
	lst := ins.PublicKeys()
	for _, onePubKey := range lst {
		publicKeys = append(publicKeys, onePubKey.String())
	}

	createdOn := ins.CreatedOn()
	return createJSONLock(publicKeys, createdOn)
}

func createJSONLock(
	publicKeys []string,
	createdOn time.Time,
) *JSONLock {
	out := JSONLock{
		PublicKeys: publicKeys,
		CreatedOn:  createdOn,
	}

	return &out
}
