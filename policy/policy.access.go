package policy

type accessPolicy struct {
	resource string
	list     []*Policy
}
