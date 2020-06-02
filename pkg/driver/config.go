package driver

// Config is the config for the CSI driver
type Config struct {
	Port int
	// Name is the name of the CSI driver
	Name string
	// StorjVersion is the version of the underlying storage
	Version string
}
