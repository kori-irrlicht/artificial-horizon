package core

type Type uint

type Loader interface {
	Load(string) (interface{}, error)
}

type AssetManager interface {
	// Register adds a Loader for a specific type
	Register(Loader, Type)

	// Load loads the file with the given name and type.
	// If there was no error, nil is returned
	// The loaded asset can be retrieved with Get or Wait
	Load(string, Type) error

	// Get tries to retrieve the named asset.
	// If the asset hasn't been loaded yet, an error will be returned
	Get(string, Type) (interface{}, error)

	// Wait waits until the resources has been loaded and writes the asset to
	// the returned channel
	// If there is no asset, an error will be returned
	Wait(string, Type) chan interface{}
}
