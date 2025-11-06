package domain

type ConjugationRequest struct {
	Verb    string   `json:"verb"`
	Tenses  []string `json:"tenses"`
	Persons []string `json:"persons"`
}
