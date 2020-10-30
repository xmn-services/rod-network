package buckets

// Application represents the application
type Application interface {
	Current() Current
}

// Current represents a bucket current application
type Current interface {
	Add(absolutePath string) error
	Delete(absolutePath string) error
	Purge(absolutePath string) error
}
