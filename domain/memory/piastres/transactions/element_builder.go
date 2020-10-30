package transactions

import (
	"errors"

	"github.com/xmn-services/rod-network/domain/memory/piastres/cancels"
	"github.com/xmn-services/rod-network/libs/hash"
)

type elementBuilder struct {
	cancel cancels.Cancel
	bucket *hash.Hash
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		cancel: nil,
		bucket: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithCancel adds a cancel to the builder
func (app *elementBuilder) WithCancel(cancel cancels.Cancel) ElementBuilder {
	app.cancel = cancel
	return app
}

// WithBucket adds a bucket to the builder
func (app *elementBuilder) WithBucket(bucket hash.Hash) ElementBuilder {
	app.bucket = &bucket
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.cancel != nil {
		return createElementWithCancel(app.cancel), nil
	}

	if app.bucket != nil {
		return createElementWithBucket(app.bucket), nil
	}

	return nil, errors.New("the Element is invalid")
}
