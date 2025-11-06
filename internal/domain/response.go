package domain

type ConjugationResponse struct {
	Verb         string                       `json:"verb"`
	Conjugations map[string]map[string]string `json:"conjugations"`
}
