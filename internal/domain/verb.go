package domain

type ConjugationType string

const (
	Regular   ConjugationType = "regular"
	Irregular ConjugationType = "irregular"
)

type ConjugationGroup string

const (
	AR ConjugationGroup = "ar"
	ER ConjugationGroup = "er"
	IR ConjugationGroup = "ir"
)

type Verb struct {
	Name             string
	ConjugationGroup ConjugationGroup
	ConjugationType  ConjugationType
}
