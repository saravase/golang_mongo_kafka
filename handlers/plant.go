package handlers

import (
	"log"
)

// Initialize struct type Plant with properties
type Plant struct {
	logger *log.Logger
}

// Initialize the Plant struct properties
func NewPlant(logger *log.Logger) *Plant {
	return &Plant{
		logger: logger,
	}
}
