package buckets

// Application represents a bucket application
type Application interface {
	Start() error
	Stop() error
}
