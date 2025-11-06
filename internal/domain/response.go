package domain

type ConjugationResponse struct {
	Verb         string
	Conjugations map[string]map[string]string
}
