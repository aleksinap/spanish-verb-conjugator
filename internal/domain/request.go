package domain

type ConjugationRequest struct {
	Verb    string
	Tenses  []string
	Persons []string
}
