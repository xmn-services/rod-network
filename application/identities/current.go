package identities

import "github.com/xmn-services/rod-network/domain/memory/identities"

type current struct {
	identityBuilder identities.Builder
	identityService identities.Service
	identity        identities.Identity
	password        string
}

func createCurrent(
	identityBuilder identities.Builder,
	identityService identities.Service,
	identity identities.Identity,
	password string,
) Current {
	out := current{
		identityBuilder: identityBuilder,
		identityService: identityService,
		identity:        identity,
		password:        password,
	}

	return &out
}

// Update updates the identity
func (app *current) Update(update Update) error {
	seed := app.identity.Seed()
	name := app.identity.Name()
	root := app.identity.Root()
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
		updatedIdentity,
		name,
		seed,
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
func (app *current) Retrieve() identities.Identity {
	return app.identity
}

// Delete deletes the identity
func (app *current) Delete() error {
	name := app.identity.Name()
	seed := app.identity.Seed()
	return app.identityService.Delete(name, seed, app.password)
}
