package identities

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/xmn-services/rod-network/libs/hash"
)

func makeFileName(hsh hash.Hash, seed string) (string, error) {
	in := fmt.Sprintf("%s%s", hsh.String(), seed)
	h := sha512.New()
	_, err := h.Write([]byte(in))
	if err != nil {
		return "", nil
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
