package handlers

import "github.com/iraunchy/dyel/db"

// ListProgramsInput is empty because GET /programs takes no payload.
type ListProgramsInput struct{}

// GetProgramInput holds the :id param for GET /programs/:id.
type GetProgramInput struct {
	ID string `uri:"id" binding:"required"`
}

// CreateProgramInput maps the JSON body for POST /programs.
type CreateProgramInput struct {
	Name     string   `json:"name" binding:"required"`
	SharedBy string   `json:"shared_by" binding:"required"`
	Days     []db.Day `json:"days" binding:"required"`
}

// ToModel converts CreateProgramInput → *db.Program
func (in CreateProgramInput) ToModel() *db.Program {
	return &db.Program{
		Name:     in.Name,
		SharedBy: in.SharedBy,
		Days:     in.Days,
	}
}

// UpdateProgramInput merges ID+body and already has ToModel()
type UpdateProgramInput struct {
	ID       string
	Name     string
	SharedBy string
	Days     []db.Day
}

func (in UpdateProgramInput) ToModel() *db.Program {
	return &db.Program{
		ID:       in.ID,
		Name:     in.Name,
		SharedBy: in.SharedBy,
		Days:     in.Days,
	}
}

// UpdateProgramURI only for URI binding
type UpdateProgramURI struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// UpdateProgramJSON only for JSON binding
type UpdateProgramJSON struct {
	Name     string   `json:"name"      binding:"required"`
	SharedBy string   `json:"shared_by" binding:"required"`
	Days     []db.Day `json:"days"`
}

// Merge combines them into your full DTO
func (u UpdateProgramURI) Merge(j UpdateProgramJSON) UpdateProgramInput {
	return UpdateProgramInput{
		ID:       u.ID,
		Name:     j.Name,
		SharedBy: j.SharedBy,
		Days:     j.Days,
	}
}
