package shareholders

import (
	"github.com/xmn-services/rod-network/libs/hash"
	transfer_lock_shareholder "github.com/xmn-services/rod-network/domain/transfers/piastres/locks/shareholders"
)

type repository struct {
	trRepository transfer_lock_shareholder.Repository
	builder      Builder
}

func createRepository(
	trRepository transfer_lock_shareholder.Repository,
	builder Builder,
) Repository {
	out := repository{
		trRepository: trRepository,
		builder:      builder,
	}

	return &out
}

// Retrieve retrieves a shareholder by hash
func (app *repository) Retrieve(hash hash.Hash) (ShareHolder, error) {
	trHolder, err := app.trRepository.Retrieve(hash)
	if err != nil {
		return nil, err
	}

	key := trHolder.Key()
	power := trHolder.Power()
	createdOn := trHolder.CreatedOn()
	return app.builder.Create().WithKey(key).WithPower(power).CreatedOn(createdOn).Now()
}

// RetrieveAll retrieves shareholders by hashes
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]ShareHolder, error) {
	out := []ShareHolder{}
	for _, oneHash := range hashes {
		holder, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, holder)
	}

	return out, nil
}
