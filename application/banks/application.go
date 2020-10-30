package banks

import (
	"time"

	"github.com/xmn-services/rod-network/domain/memory/accounts/banks"
	"github.com/xmn-services/rod-network/domain/memory/accounts/banks/branches"
	"github.com/xmn-services/rod-network/domain/memory/addresses"
	"github.com/xmn-services/rod-network/libs/hash"
)

type application struct {
	addressRepository addresses.Repository
	builder           banks.Builder
	repository        banks.Repository
	service           banks.Service
	branchBuilder     branches.Builder
}

func createApplication(
	addressRepository addresses.Repository,
	builder banks.Builder,
	repository banks.Repository,
	service banks.Service,
	branchBuilder branches.Builder,
) Application {
	out := application{
		addressRepository: addressRepository,
		builder:           builder,
		repository:        repository,
		service:           service,
		branchBuilder:     branchBuilder,
	}

	return &out
}

// Retrieve retrieves a bank by hash
func (app *application) Retrieve(hash hash.Hash) (banks.Bank, error) {
	return app.repository.Retrieve(hash)
}

// RetrieveAll retrieves all banks
func (app *application) RetrieveAll() ([]banks.Bank, error) {
	return app.repository.RetrieveAll()
}

// New creates a new bank
func (app *application) New(name string, institutionNumber string) error {
	createdOn := time.Now().UTC()
	bank, err := app.builder.Create().WithName(name).WithInstitutionNumber(institutionNumber).CreatedOn(createdOn).Now()
	if err != nil {
		return err
	}

	return app.service.Insert(bank)
}

// AddBranch adds a new branch to a bank
func (app *application) AddBranch(bankHash hash.Hash, transitNumber string, addressHash hash.Hash) error {
	bank, err := app.repository.Retrieve(bankHash)
	if err != nil {
		return err
	}

	address, err := app.addressRepository.Retrieve(addressHash)
	if err != nil {
		return err
	}

	newBranch, err := app.branchBuilder.Create().WithTransitNumber(transitNumber).WithAddress(address).Now()
	if err != nil {
		return err
	}

	branches := []branches.Branch{}
	if bank.HasBranches() {
		branches = bank.Branches()
	}

	lastUpdatedOn := time.Now().UTC()
	branches = append(branches, newBranch)

	name := bank.Name()
	institutionNumber := bank.InstitutionNumber()
	createdOn := bank.CreatedOn()
	updatedBank, err := app.builder.Create().WithName(name).WithInstitutionNumber(institutionNumber).WithBranches(branches).CreatedOn(createdOn).LastUpdatedOn(lastUpdatedOn).Now()
	if err != nil {
		return err
	}

	return app.service.Update(bank, updatedBank)
}
