// Package entity holds all the entities that
package entity

import "github.com/google/uuid"

type Item struct {
	// ID an the identifier of entity
	ID          uuid.UUID
	Name        string
	Description string
}
