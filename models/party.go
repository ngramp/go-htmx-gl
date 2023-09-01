package models

import (
	"gorm.io/gorm"
)

type Party struct {
	gorm.Model
	Name string
	// Other common attributes
}

type Person struct {
	PartyID uint
	Party
	Age int
	// Additional attributes specific to a person
}

type Company struct {
	PartyID uint
	Party
	RegistrationNumber string
	// Additional attributes specific to a company
}

type Relationship struct {
	gorm.Model
	PartyID        uint
	RelatedPartyID uint
	// Other relationship attributes
	Topic     string
	Sentiment string
}
