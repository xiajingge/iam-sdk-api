package v1

type Secret struct {
	ObjectMeta `       json:"metadata,omitempty"`
	Username   string `json:"username"           gorm:"column:username"  validate:"omitempty"`
	//nolint: tagliatelle
	SecretID  string `json:"secretID"           gorm:"column:secretID"  validate:"omitempty"`
	SecretKey string `json:"secretKey"          gorm:"column:secretKey" validate:"omitempty"`

	// Required: true
	Expires     int64  `json:"expires"     gorm:"column:expires"     validate:"omitempty"`
	Description string `json:"description" gorm:"column:description" validate:"description"`
}

type SecretList struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	ListMeta `json:",inline"`

	// List of secrets
	Items []*Secret `json:"items"`
}

func (s *Secret) TableName() string {
	return "secret"
}
