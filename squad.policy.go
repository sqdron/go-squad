package squad

import (
	"github.com/sqdron/squad/policy"
)

type actionPolicy struct {
	resource string
	list     []*policy.Policy
}

type IActionPolicy interface {
	Policy(effect policy.PolicyEffect, audience ...string) *actionPolicy
}

func (a *actionPolicy) Policy(effect policy.PolicyEffect, audience ...string) *actionPolicy {
	a.list = append(a.list, &policy.Policy{Effect: effect, Audience: audience, Resource: a.resource})
	return a
}

func ActionPolicy(resource string) *actionPolicy {
	return &actionPolicy{resource: resource, list: []*policy.Policy{}}
}
