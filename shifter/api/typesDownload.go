package api

type Downloads struct {
	Items *[]Download `json:"items"`
}

type Download struct {
	Link        string `json:"link"`
	Name        string `json:"name"`
	Uuid        string `json:"uuid"`
	DisplayName string `json:"displayName"`
}
