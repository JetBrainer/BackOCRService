package app

type Document interface {
	RuleDocUsage(string)
}

type Docs interface {
	Match(string) string
}
