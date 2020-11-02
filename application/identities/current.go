package identities

import "github.com/xmn-services/rod-network/domain/memory/identities"

type current struct {
	identityBuilder    identities.Builder
	identityRepository identities.Repository
	identityService    identities.Service
	name               string
	password           string
	seed               string
}

func createCurrent(
	identityBuilder identities.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
	name string,
	password string,
	seed string,
) Current {
	out := current{
		identityBuilder:    identityBuilder,
		identityRepository: identityRepository,
		identityService:    identityService,
		name:               name,
		password:           password,
		seed:               seed,
	}

	return &out
}

// Update updates the identity
func (app *current) Update(update Update) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	// retrieve the identity:
	seed := identity.Seed()
	name := identity.Name()
	root := identity.Root()
	newPassword := app.password
	builder := app.identityBuilder.Create().WithSeed(seed).WithName(name).WithRoot(root)
	if update.HasSeed() {
		uSeed := update.Seed()
		builder.WithSeed(uSeed)
	}

	if update.HasName() {
		uName := update.Name()
		builder.WithName(uName)
	}

	if update.HasRoot() {
		uRoot := update.Root()
		builder.WithRoot(uRoot)
	}

	if update.HasPassword() {
		newPassword = update.Password()
	}

	updatedIdentity, err := builder.Now()
	if err != nil {
		return err
	}

	err = app.identityService.Update(
		identity.Hash(),
		updatedIdentity,
		app.password,
		newPassword,
	)

	if err != nil {
		return err
	}

	app.password = newPassword
	return nil
}

// Retrieve retrieves the identity
func (app *current) Retrieve() (identities.Identity, error) {
	return app.identityRepository.Retrieve(app.name, app.password, app.seed)
}

// Delete deletes the identity
func (app *current) Delete() error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	return app.identityService.Delete(identity, app.password)
}
