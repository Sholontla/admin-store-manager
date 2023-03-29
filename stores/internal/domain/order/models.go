package order

import "github.com/google/uuid"

type Worker2 struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FloatNum    float64   `json:"float_num"`
	CreatedAt   string    `json:"created_at"`
}

type Worker3 struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FloatNum    float64   `json:"float_num"`
	CreatedAt   string    `json:"created_at"`
}

type Worker4 struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FloatNum    float64   `json:"float_num"`
	CreatedAt   string    `json:"created_at"`
}
