package api

type Convert struct {
	Shifter *Shifter       `json:"shifter"`
	Items   []*ConvertItem `json:"items"`
}
