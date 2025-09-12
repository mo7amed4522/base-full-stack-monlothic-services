package system

// GetInitializers returns all registered initializers
func GetInitializers() initSlice {
	// Make a copy of the initializers to avoid modifying the original
	initsCopy := make(initSlice, len(initializers))
	copy(initsCopy, initializers)
	return initsCopy
}