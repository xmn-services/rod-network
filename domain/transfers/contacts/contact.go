package contacts

import (
	"time"

	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/rod-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type contact struct {
	immutable   entities.Immutable
	name        string
	description string
	pubKey      public.Key
	privateKey  encryption.PrivateKey
}

func createContact(
	immutable entities.Immutable,
	name string,
	description string,
	pubKey public.Key,
	privateKey encryption.PrivateKey,
) Contact {
	out := contact{
		immutable:   immutable,
		name:        name,
		description: description,
		pubKey:      pubKey,
		privateKey:  privateKey,
	}

	return &out
}

// Hash returns the hash
func (obj *contact) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Name returns the name
func (obj *contact) Name() string {
	return obj.name
}

// Description returns the description
func (obj *contact) Description() string {
	return obj.description
}

// PublicKey returns the publicKey
func (obj *contact) PublicKey() public.Key {
	return obj.pubKey
}

// PrivateKey returns the privateKey
func (obj *contact) PrivateKey() encryption.PrivateKey {
	return obj.privateKey
}

// CreatedOn returns the creation time
func (obj *contact) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
