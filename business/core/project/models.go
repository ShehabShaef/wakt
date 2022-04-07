package projects

import (
	"github.com/AhmedShaef/wakt/business/core/project/db"
	"time"
	"unsafe"
)

// Project represents an individual project.
type Project struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Wid            string        `json:"wid"`
	Cid            string        `json:"cid"`
	Active         bool          `json:"active"`
	IsPrivate      bool          `json:"is_private"`
	Template       bool          `json:"template"`
	TemplateID     string        `json:"template_id"`
	Billable       bool          `json:"billable"`
	AutoEstimates  bool          `json:"auto_estimates"`
	EstimatedHours time.Duration `json:"estimated_hours"`
	DateUpdated    time.Time     `json:"date_updated"`
	Color          int           `json:"color"`
	Rate           float32       `json:"rate"`
	DateCreated    time.Time     `json:"date_created"`
}

// NewProject contains information needed to create a new project.
type NewProject struct {
	Name           string        `json:"name" validate:"required"`
	Wid            string        `json:"wid"`
	Cid            string        `json:"cid"`
	Active         bool          `json:"active"`
	IsPrivate      bool          `json:"is_private"`
	AutoEstimates  bool          `json:"auto_estimates"`
	EstimatedHours time.Duration `json:"estimated_hours"`
	Template       bool          `json:"template"`
	TemplateID     string        `json:"template_id"`
	Billable       bool          `json:"billable"`
	Color          int           `json:"color"`
	Rate           float32       `json:"rate"`
}

// UpdateProject defines what information may be provided to modify an existing
// project. All fields are optional so projects can send just the fields they want
// changed. It uses pointer fields ,so we can differentiate between a field that
// was not provided and a field that was provided as explicitly blank. Normally
// we do not want to use pointers to basic types ,but we make exceptions around
// marshalling/unmarshalling.
type UpdateProject struct {
	Name           *string        `json:"name"`
	Active         *bool          `json:"active"`
	IsPrivate      *bool          `json:"is_private"`
	AutoEstimates  *bool          `json:"auto_estimates"`
	EstimatedHours *time.Duration `json:"estimated_hours"`
	Template       *bool          `json:"template"`
	TemplateID     *string        `json:"template_id"`
	Billable       *bool          `json:"billable"`
	Color          *int           `json:"color"`
	Rate           *float32       `json:"rate"`
}

// =============================================================================

func toProject(dbProject db.Project) Project {
	pu := (*Project)(unsafe.Pointer(&dbProject))
	return *pu
}

func toProjectsSlice(dbProject []db.Project) []Project {
	projects := make([]Project, len(dbProject))
	for i, dbclint := range dbProject {
		projects[i] = toProject(dbclint)
	}
	return projects
}
