package addresses

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type address struct {
	immutable  entities.Immutable
	sender     *hash.Hash
	recipients []hash.Hash
	subject    *hash.Hash
}

func createAddressFromJSON(js *JSONAddress) (Address, error) {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder().Create().CreatedOn(js.CreatedOn)
	if js.Sender != "" {
		sender, err := hashAdapter.FromString(js.Sender)
		if err != nil {
			return nil, err
		}

		builder.WithSender(*sender)
	}

	if len(js.Recipients) > 0 {
		recipients := []hash.Hash{}
		for _, oneRecipient := range js.Recipients {
			ins, err := hashAdapter.FromString(oneRecipient)
			if err != nil {
				return nil, err
			}

			recipients = append(recipients, *ins)
		}

		builder.WithRecipients(recipients)
	}

	if js.Subject != "" {
		subject, err := hashAdapter.FromString(js.Subject)
		if err != nil {
			return nil, err
		}

		builder.WithSubject(*subject)
	}

	return builder.Now()
}

func createAddressWithSender(
	immutable entities.Immutable,
	sender *hash.Hash,
) Address {
	return createAddressInternally(immutable, sender, nil, nil)
}

func createAddressWithRecipients(
	immutable entities.Immutable,
	recipients []hash.Hash,
) Address {
	return createAddressInternally(immutable, nil, recipients, nil)
}

func createAddressWithSubject(
	immutable entities.Immutable,
	subject *hash.Hash,
) Address {
	return createAddressInternally(immutable, nil, nil, subject)
}

func createAddressWithSenderAndRecipients(
	immutable entities.Immutable,
	sender *hash.Hash,
	recipients []hash.Hash,
) Address {
	return createAddressInternally(immutable, sender, recipients, nil)
}

func createAddressWithSenderAndSubject(
	immutable entities.Immutable,
	sender *hash.Hash,
	subject *hash.Hash,
) Address {
	return createAddressInternally(immutable, sender, nil, subject)
}

func createAddressWithRecipientsAndSubject(
	immutable entities.Immutable,
	recipients []hash.Hash,
	subject *hash.Hash,
) Address {
	return createAddressInternally(immutable, nil, recipients, subject)
}

func createAddressWithSenderAndRecipientsAndSubject(
	immutable entities.Immutable,
	sender *hash.Hash,
	recipients []hash.Hash,
	subject *hash.Hash,
) Address {
	return createAddressInternally(immutable, sender, recipients, subject)
}

func createAddressInternally(
	immutable entities.Immutable,
	sender *hash.Hash,
	recipients []hash.Hash,
	subject *hash.Hash,
) Address {
	out := address{
		immutable:  immutable,
		sender:     sender,
		recipients: recipients,
		subject:    subject,
	}

	return &out
}

// Hash returns the hash
func (obj *address) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// CreatedOn returns the creation time
func (obj *address) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasSender returns true if there is a sender, false otherwise
func (obj *address) HasSender() bool {
	return obj.sender != nil
}

// Sender returns the sender, if any
func (obj *address) Sender() *hash.Hash {
	return obj.sender
}

// HasRecipients returns true if there is recipients, false otherwise
func (obj *address) HasRecipients() bool {
	return obj.recipients != nil
}

// Recipients returns the recipients, if any
func (obj *address) Recipients() []hash.Hash {
	return obj.recipients
}

// HasSubject returns true if there is a subject, false otherwise
func (obj *address) HasSubject() bool {
	return obj.subject != nil
}

// Subject returns the subject, if any
func (obj *address) Subject() *hash.Hash {
	return obj.subject
}

// MarshalJSON converts the instance to JSON
func (obj *address) MarshalJSON() ([]byte, error) {
	ins := createJSONAddressFromAddress(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *address) UnmarshalJSON(data []byte) error {
	ins := new(JSONAddress)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createAddressFromJSON(ins)
	if err != nil {
		return err
	}

	insAddress := pr.(*address)
	obj.immutable = insAddress.immutable
	obj.sender = insAddress.sender
	obj.recipients = insAddress.recipients
	obj.subject = insAddress.subject
	return nil
}
