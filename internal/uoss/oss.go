package uoss

import "errors"

var (
	// ErrNewClient represents client creation errors.
	ErrNewClient = errors.New("uoss: fail to create client")

	// ErrPutObject returned when fail to put object to oss.
	ErrPutObject = errors.New("uoss: fail to put object")

	// ErrGetObject returned when fail to get object from oss.
	ErrGetObject = errors.New("uoss: fail to get object")

	// ErrGetObjectAsFile returned when fail to get object as file from oss.
	ErrGetObjectAsFile = errors.New("uoss: fail to get object as file")

	// ErrListObjectsOfCurrentDir returned when fail to list files and directories of current directory
	ErrListObjectsOfCurrentDir = errors.New("uoss: fail to list object")

	// ErrTryNext is a marker error which indicates to try the next entry.
	ErrTryNext = errors.New("try next")

	// ErrNoClientProviderFound means can not find a client provider.
	ErrNoClientProviderFound = errors.New("uoss: no client provider found")
)
