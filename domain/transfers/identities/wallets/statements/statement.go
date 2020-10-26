package statements

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/rod-network/libs/entities"
	"github.com/xmn-services/rod-network/libs/hash"
)

type statement struct {
	immutable entities.Immutable
	incoming  []hash.Hash
	outgoing  []hash.Hash
}

func createStatementFromJSON(ins *jsonStatement) (Statement, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	builder := NewBuilder().Create().
		WithHash(*hsh).
		CreatedOn(ins.CreatedOn)

	if len(ins.Incoming) > 0 {
		incoming := []hash.Hash{}
		for _, oneIncoming := range ins.Incoming {
			hsh, err := hashAdapter.FromString(oneIncoming)
			if err != nil {
				return nil, err
			}

			incoming = append(incoming, *hsh)
		}

		builder.WithIncoming(incoming)
	}

	if len(ins.Outgoing) > 0 {
		outgoing := []hash.Hash{}
		for _, oneOutgoing := range ins.Outgoing {
			hsh, err := hashAdapter.FromString(oneOutgoing)
			if err != nil {
				return nil, err
			}

			outgoing = append(outgoing, *hsh)
		}

		builder.WithOutgoing(outgoing)
	}

	return builder.Now()
}

func createStatement(
	immutable entities.Immutable,
) Statement {
	return createStatementInternally(immutable, nil, nil)
}

func createStatementWithIncoming(
	immutable entities.Immutable,
	incoming []hash.Hash,
) Statement {
	return createStatementInternally(immutable, incoming, nil)
}

func createStatementWithOutgoing(
	immutable entities.Immutable,
	outgoing []hash.Hash,
) Statement {
	return createStatementInternally(immutable, nil, outgoing)
}

func createStatementWithIncomingAndOutgoing(
	immutable entities.Immutable,
	incoming []hash.Hash,
	outgoing []hash.Hash,
) Statement {
	return createStatementInternally(immutable, incoming, outgoing)
}

func createStatementInternally(
	immutable entities.Immutable,
	incoming []hash.Hash,
	outgoing []hash.Hash,
) Statement {
	out := statement{
		immutable: immutable,
		incoming:  incoming,
		outgoing:  outgoing,
	}

	return &out
}

// Hash returns the hash
func (obj *statement) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// HasIncoming returns true if there is incoming, false otherwise
func (obj *statement) HasIncoming() bool {
	return obj.incoming != nil
}

// Incoming returns the incoming, if any
func (obj *statement) Incoming() []hash.Hash {
	return obj.incoming
}

// HasOutgoing returns true if there is outgoing, false otherwise
func (obj *statement) HasOutgoing() bool {
	return obj.outgoing != nil
}

// Outgoing returns the outgoing, if any
func (obj *statement) Outgoing() []hash.Hash {
	return obj.outgoing
}

// CreatedOn returns the creation time
func (obj *statement) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *statement) MarshalJSON() ([]byte, error) {
	ins := createJSONStatementFromStatement(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *statement) UnmarshalJSON(data []byte) error {
	ins := new(jsonStatement)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createStatementFromJSON(ins)
	if err != nil {
		return err
	}

	insStatement := pr.(*statement)
	obj.immutable = insStatement.immutable
	obj.incoming = insStatement.incoming
	obj.outgoing = insStatement.outgoing
	return nil
}
