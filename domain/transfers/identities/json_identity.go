package identities

import "time"

type jsonIdentity struct {
	Hash      string    `json:"hash"`
	Seed      string    `json:"seed"`
	Name      string    `json:"name"`
	Root      string    `json:"root"`
	Wallets   []string  `json:"wallets"`
	Buckets   []string  `json:"buckets"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONIdentityFromIdentity(ins Identity) *jsonIdentity {
	hash := ins.Hash().String()

	seed := ins.Seed()
	name := ins.Name()
	root := ins.Root()

	walletsLst := []string{}
	if ins.HasWallets() {
		wallets := ins.Wallets()
		for _, oneWallet := range wallets {
			walletsLst = append(walletsLst, oneWallet.String())
		}
	}

	bucketsLst := []string{}
	if ins.HasBuckets() {
		buckets := ins.Buckets()
		for _, oneBucket := range buckets {
			bucketsLst = append(bucketsLst, oneBucket.String())
		}
	}

	createdOn := ins.CreatedOn()
	return createJSONIdentity(hash, seed, name, root, walletsLst, bucketsLst, createdOn)
}

func createJSONIdentity(
	hash string,
	seed string,
	name string,
	root string,
	wallets []string,
	buckets []string,
	createdOn time.Time,
) *jsonIdentity {
	out := jsonIdentity{
		Hash:      hash,
		Seed:      seed,
		Name:      name,
		Root:      root,
		Wallets:   wallets,
		Buckets:   buckets,
		CreatedOn: createdOn,
	}

	return &out
}
