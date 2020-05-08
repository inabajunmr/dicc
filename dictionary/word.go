package dictionary

type FunctionalLavel int

type Result struct {
	SearchWord  string
	Definitions []Definition
}

type Definition struct {
	Descriptions    []string
	FunctionalLabel string
}
