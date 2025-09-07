package v1

import (
	"time"
)

type User struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	ObjectMeta `json:"metadata,omitempty"`

	Status int `json:"status" gorm:"column:status" validate:"omitempty"`

	// Required: true
	Nickname string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`

	// Required: true
	Password string `json:"password,omitempty" gorm:"column:password" validate:"required"`

	// Required: true
	Email string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`

	Phone string `json:"phone" gorm:"column:phone" validate:"omitempty"`

	IsAdmin int `json:"isAdmin,omitempty" gorm:"column:isAdmin" validate:"omitempty"`

	TotalPolicy int64 `json:"totalPolicy" gorm:"-" validate:"omitempty"`

	LoginedAt time.Time `json:"loginedAt,omitempty" gorm:"column:loginedAt"`
}

type UserList struct {
	// +optional
	ListMeta `json:",inline"`
	Items    []*User `json:"items"`
}

func (u *User) TableName() string {
	return "user"
}
