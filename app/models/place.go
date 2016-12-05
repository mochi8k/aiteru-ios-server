package models

type Place struct {
	id            int
	name          string
	owners        []int
	collaborators []int
	createdAt     string
	createdBy     int
	updatedAt     string
	updatedBy     int
}
