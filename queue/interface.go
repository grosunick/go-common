package queue

// Base queue interface
type IQueue interface {
	// Add to queue method
	Add(interface{}) bool
	// Remove from queue method
	Remove() (interface{}, bool)
	// Returns queue length
	Size() int
}
