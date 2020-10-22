package transactions

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/signature"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	pkFactory        signature.PrivateKeyFactory
	amountRingKeys   uint
	content          Content
	signature        signature.RingSignature
	pk               signature.PrivateKey
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
	pkFactory signature.PrivateKeyFactory,
	amountRingKeys uint,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		pkFactory:        pkFactory,
		amountRingKeys:   amountRingKeys,
		content:          nil,
		signature:        nil,
		pk:               nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder, app.pkFactory, app.amountRingKeys)
}

// WithContent adds content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// WithSignature adds a ring signature to the builder
func (app *builder) WithSignature(signature signature.RingSignature) Builder {
	app.signature = signature
	return app
}

// WithPrivateKey adds a privateKey to the builder
func (app *builder) WithPrivateKey(pk signature.PrivateKey) Builder {
	app.pk = pk
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Transaction instance")
	}

	if app.pk != nil {
		rand.Seed(time.Now().UnixNano())
		pubKeyIndex := rand.Intn(int(app.amountRingKeys))
		ringKeys := []signature.PublicKey{}
		for i := 0; i < int(app.amountRingKeys); i++ {
			ringKeys = append(ringKeys, app.pkFactory.Create().PublicKey())
			if i == pubKeyIndex {
				ringKeys = append(ringKeys, app.pk.PublicKey())
			}
		}

		msg := app.content.Hash().String()
		sig, err := app.pk.RingSign(msg, ringKeys)
		if err != nil {
			return nil, err
		}

		app.signature = sig
	}

	if app.signature == nil {
		return nil, errors.New("the ring signature is mandatory in order to build a Transaction instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		[]byte(app.signature.String()),
		[]byte(strconv.Itoa(int(app.createdOn.UnixNano()))),
	})

	if err != nil {
		return nil, err
	}

	msg := app.content.Hash().String()
	if !app.signature.Verify(msg) {
		str := fmt.Sprintf("the ring signature could not be validated against the content hash (%s)", msg)
		return nil, errors.New(str)
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createTransaction(immutable, app.content, app.signature), nil
}
