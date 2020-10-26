package statements

import "time"

type jsonStatement struct {
	Hash      string    `json:"hash"`
	Incoming  []string  `json:"incoming"`
	Outgoing  []string  `json:"outgoing"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONStatementFromStatement(ins Statement) *jsonStatement {
	hash := ins.Hash().String()

	incoming := []string{}
	if ins.HasIncoming() {
		in := ins.Incoming()
		for _, oneIncoming := range in {
			incoming = append(incoming, oneIncoming.String())
		}
	}

	outgoing := []string{}
	if ins.HasOutgoing() {
		out := ins.Outgoing()
		for _, oneOutgoing := range out {
			outgoing = append(outgoing, oneOutgoing.String())
		}
	}

	createdOn := ins.CreatedOn()
	return createJSONStatement(hash, incoming, outgoing, createdOn)
}

func createJSONStatement(
	hash string,
	incoming []string,
	outgoing []string,
	createdOn time.Time,
) *jsonStatement {
	out := jsonStatement{
		Hash:      hash,
		Incoming:  incoming,
		Outgoing:  outgoing,
		CreatedOn: createdOn,
	}

	return &out
}
