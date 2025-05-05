package db

import (
	"time"
)

// Exercise represents one movement in a Day.
type Exercise struct {
	ID        string    `gorm:"type:text;primaryKey" json:"id"`
	DayID     string    `gorm:"not null;index" json:"day_id"`
	Name      string    `json:"name"`
	Sets      int       `json:"sets"`
	Reps      string    `json:"reps"`
	Rest      string    `json:"rest"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Day is one day in a Program, holding many Exercises.
type Day struct {
	ID        string     `gorm:"type:text;primaryKey" json:"id"`
	ProgramID string     `gorm:"not null;index" json:"program_id"`
	Name      string     `json:"name"`
	Exercises []Exercise `gorm:"constraint:OnDelete:CASCADE" json:"exercises"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Program is a collection of Days, identified by a UUID.
type Program struct {
	ID        string    `gorm:"type:text;primaryKey" json:"id"`
	Name      string    `json:"name"`
	SharedBy  string    `json:"shared_by"`
	Days      []Day     `gorm:"constraint:OnDelete:CASCADE" json:"days"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
