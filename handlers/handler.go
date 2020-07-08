package handlers

import "log"

// Items handler for getting and updating items
type Items struct {
	l *log.Logger
}

// NewItems retruns a new items handler with the given logger
func NewItems(l *log.Logger) *Items {
	return &Items{l}
}
