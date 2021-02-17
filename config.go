package gyro

import (
	"os"
	"path/filepath"
)

// Config defines the configuration required to run the Gyro application
type Config struct {
	EnableHTTP        bool
	HTTPPort          int
	EnableTCP         bool
	TCPPort           int
	EnableGRPC        bool
	GRPCPort          int
	EnableUDP         bool
	UDPPort           int
	EnablePersistence bool
	PersistenceLoc    string
}

// DefaultConfig sets the default config
func DefaultConfig() Config {
	persistenceTmpLoc := filepath.Join(os.TempDir(), "gyro")
	return Config{
		EnableHTTP:        true,
		HTTPPort:          8080,
		EnableTCP:         true,
		TCPPort:           8089,
		EnableGRPC:        true,
		GRPCPort:          8081,
		EnableUDP:         true,
		UDPPort:           8090,
		EnablePersistence: true,
		PersistenceLoc:    persistenceTmpLoc,
	}
}
