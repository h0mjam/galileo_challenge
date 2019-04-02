package types

import "time"

type Measure struct {
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}
