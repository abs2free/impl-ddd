// Package entity holds all the entities that
package entity

import "github.com/google/uuid"

type Person struct {
	// ID an the identifier of entity
	ID   uuid.UUID
	Name string
	Age  int
}
