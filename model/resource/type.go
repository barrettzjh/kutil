package resource

type Resource struct {
	Limits struct {
		CPU string `json:"cpu"`
		Memory string `json:"memory"`
	} `json:"limits"`
	Requests struct {
		CPU string `json:"cpu"`
		Memory string `json:"memory"`
	} `json:"requests"`
}