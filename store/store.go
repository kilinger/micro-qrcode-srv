package store

import "io"

// Store interface to save qrcode
type Store interface {
	Save(name string, reader io.Reader, contentType string) (string, error)
}

// New create store
func New() (Store, error) {
	return newMinioStore()
}
