package policy

import "time"

type PolicyEffect int

const (
	Allow PolicyEffect = 1 << iota
	Deny
	Unauthorized
)

type Policy struct {
	ID        int64
	Request   string
	Resource  string
	Audience  []string
	Effect    PolicyEffect
	ValidFrom *time.Time
	ValidTo   *time.Time
}
