package bucket

import (
	transfer_bucket "github.com/xmn-services/rod-network/domain/transfers/buckets"
)

type adapter struct {
	trBuilder transfer_bucket.Builder
}

func createAdapter(
	trBuilder transfer_bucket.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a bucket to a transfer bucket
func (app *adapter) ToTransfer(bucket Bucket) (transfer_bucket.Bucket, error) {
	hsh := bucket.Hash()
	inf := bucket.Information().Hash()
	path := bucket.AbsolutePath()
	pk := bucket.PrivateKey()
	createdOn := bucket.CreatedOn()
	return app.trBuilder.Create().
		WithHash(hsh).
		WithInformation(inf).
		WithAbsolutePath(path).
		WithPrivateKey(pk).
		CreatedOn(createdOn).
		Now()
}
