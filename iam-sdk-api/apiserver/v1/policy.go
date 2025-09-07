package v1

import "github.com/ory/ladon"

type AuthzPolicy struct {
	ladon.DefaultPolicy
}

type Policy struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	ObjectMeta `json:"metadata,omitempty"`

	// The user of the policy.
	Username string `json:"username" gorm:"column:username" validate:"omitempty"`

	// AuthzPolicy policy, will not be stored in db.
	Policy AuthzPolicy `json:"policy,omitempty" gorm:"-" validate:"omitempty"`
	// Policy ladon.DefaultPolicy `json:"policy,omitempty" gorm:"-" validate:"omitempty"`

	// The ladon policy content, just a string format of ladon.DefaultPolicy. DO NOT modify directly.
	PolicyShadow string `json:"-" gorm:"column:policyShadow" validate:"omitempty"`
}

type PolicyList struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	ListMeta `json:",inline"`

	// List of policies.
	Items []*Policy `json:"items"`
}

func (p *Policy) TableName() string {
	return "policy"
}
