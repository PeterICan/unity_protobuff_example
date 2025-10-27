package server

// Server defines the interface for a network server.
type Server interface {
	StartServer() error
}
