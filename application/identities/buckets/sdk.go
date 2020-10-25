package buckets

// Application represents a bucket application
type Application interface {
	Add(path string) error
	Delete(path string) error
}
