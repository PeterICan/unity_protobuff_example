package main

// Server defines the interface for a network server.
type Server interface {
	Start() error
}
