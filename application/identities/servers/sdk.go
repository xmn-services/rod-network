package servers

// Application represents a server application
type Application interface {
	Current() Current
}

// Current represents the current application
type Current interface {
	Start() error
	Stop() error
}
