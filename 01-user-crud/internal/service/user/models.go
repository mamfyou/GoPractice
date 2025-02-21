package user

import "time"

type RoleType int

const (
	startRoleType RoleType = iota
	Standard
	Manager
	Admin
	endRoleType
)

type User struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"fullname"`
	UserName  string    `json:"username"`
	IsActive  bool      `json:"isActive"`
	Address   string    `json:"address"`
	Role      RoleType  `json:"role" gorm:"column:role_c"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
