package states

import (
	"path/filepath"

	"github.com/xmn-services/rod-network/libs/hash"
)

func filePath(chainHash hash.Hash, blockIndex uint) string {
	return filepath.Join(chainHash.String(), string(blockIndex))
}
