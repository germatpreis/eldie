package culprits

type Food struct {
	Name  string          `json:"name"`
	Alias map[string]bool `json:"alias"`
}

type Culprit struct {
	Food       Food `json:"food"`
	Likelihood int  `json:"likelihood"`
}
