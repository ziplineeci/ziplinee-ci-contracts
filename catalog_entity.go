package contracts

import "time"

// CatalogEntity represents any entity stored in the catalog tree
type CatalogEntity struct {
	ID             string                 `json:"id,omitempty"`
	ParentKey      string                 `json:"parentKey,omitempty"`
	ParentValue    string                 `json:"parentValue,omitempty"`
	Key            string                 `json:"key,omitempty"`
	Value          string                 `json:"value,omitempty"`
	LinkedPipeline string                 `json:"linkedPipeline,omitempty"`
	Labels         []Label                `json:"labels,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	InsertedAt     *time.Time             `json:"insertedAt,omitempty"`
	UpdatedAt      *time.Time             `json:"updatedAt,omitempty"`
}
