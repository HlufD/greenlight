package data

import "time"

// order matters on whe using directives
// 1 name of the key
// the second is like omit empty

type Movie struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"-"`
	Year      int32     `json:"year,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Runtime   int32     `json:"runtime,omitempty"`
	Version   string    `json:"version"`
}
