package chunks

import (
	"time"
)

type jsonChunk struct {
	Hash        string    `json:"hash"`
	SizeInBytes uint      `json:"size_in_bytes"`
	Data        string    `json:"data"`
	CreatedOn   time.Time `json:"created_on"`
}

func createJSONChunkFromChunk(ins Chunk) *jsonChunk {
	hash := ins.Hash().String()
	sizeInBytes := ins.SizeInBytes()
	data := ins.Data().String()
	createdOn := ins.CreatedOn()
	return createJSONChunk(hash, sizeInBytes, data, createdOn)
}

func createJSONChunk(
	hash string,
	sizeInBytes uint,
	data string,
	createdOn time.Time,
) *jsonChunk {
	out := jsonChunk{
		Hash:        hash,
		SizeInBytes: sizeInBytes,
		Data:        data,
		CreatedOn:   createdOn,
	}

	return &out
}
