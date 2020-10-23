package models

import "time"

type (
	// Metadata ...
	Metadata map[string]interface{}
	// Thing ...
	Thing struct {
		ID        string    `json:"id,omitempty"`
		Owner     string    `json:"owner"`
		Name      string    `json:"name"`
		Key       string    `json:"key,omitempty"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
		Metadata  Metadata  `json:"metadata,omitempty"`
	}
)
