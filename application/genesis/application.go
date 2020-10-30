package genesis

type application struct {
	current Current
}

func createApplication(
	current Current,
) Application {
	out := application{
		current: current,
	}

	return &out
}

// Current returns the current application
func (app *application) Current() Current {
	return app.current
}
