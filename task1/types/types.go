package types

//Stat statistics ds
type Stat struct {
	ID     int     `json:"id" validate:"nonzero"`
	Date   string  `json:"date" validate:"nonzero"`
	Shows  int     `json:"shows" validate:"nonzero"`
	Clicks int     `json:"clicks" validate:"nonzero"`
	Costs  float32 `json:"costs" validate:"nonzero"`
}

//Stats slice of Stat
type Stats []Stat
