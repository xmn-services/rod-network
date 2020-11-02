package addresses

import (
	"errors"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	sender           *hash.Hash
	recipients       []hash.Hash
	subject          *hash.Hash
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		sender:           nil,
		recipients:       nil,
		subject:          nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithSender adds a sender to the builder
func (app *builder) WithSender(sender hash.Hash) Builder {
	app.sender = &sender
	return app
}

// WithRecipients add recipients to the builder
func (app *builder) WithRecipients(recipients []hash.Hash) Builder {
	app.recipients = recipients
	return app
}

// WithSubject adds a subject to the builder
func (app *builder) WithSubject(subject hash.Hash) Builder {
	app.subject = &subject
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Address instance
func (app *builder) Now() (Address, error) {
	if len(app.recipients) <= 0 {
		app.recipients = nil
	}

	data := [][]byte{}
	if app.sender != nil {
		data = append(data, app.sender.Bytes())
	}

	if app.recipients != nil {
		for _, oneRecipient := range app.recipients {
			data = append(data, oneRecipient.Bytes())
		}
	}

	if app.subject != nil {
		data = append(data, app.subject.Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.sender != nil && app.recipients != nil && app.subject != nil {
		return createAddressWithSenderAndRecipientsAndSubject(immutable, app.sender, app.recipients, app.subject), nil
	}

	if app.sender != nil && app.recipients != nil {
		return createAddressWithSenderAndRecipients(immutable, app.sender, app.recipients), nil
	}

	if app.sender != nil && app.subject != nil {
		return createAddressWithSenderAndSubject(immutable, app.sender, app.subject), nil
	}

	if app.recipients != nil && app.subject != nil {
		return createAddressWithRecipientsAndSubject(immutable, app.recipients, app.subject), nil
	}

	if app.sender != nil {
		return createAddressWithSender(immutable, app.sender), nil
	}

	if app.recipients != nil {
		return createAddressWithRecipients(immutable, app.recipients), nil
	}

	if app.subject != nil {
		return createAddressWithSubject(immutable, app.subject), nil
	}

	return nil, errors.New("the Address is invalid")
}
