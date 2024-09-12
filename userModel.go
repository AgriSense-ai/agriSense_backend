package main

import (
	"time"
)

// TODO: Read this https://gorm.io/docs/models.html
// TODO: Read this https://www.linkedin.com/pulse/using-gorm-golang-orm-alex-edwards/
// TODO: Read this https://www.calhoun.io/why-you-should-never-use-golangs-default-uuid-package/
// TODO: Read this https://www.linkedin.com/pulse/database-structure-user-roles-management-aj-february/

type User struct {
	UserID         string    `gorm:"type:uuid;primary_key"`
	Username       string    `gorm:"unique;not null"`
	FirstName      string    `gorm:"unique;not null"`
	LastName       string    `gorm:"unique:not null"`
	Email          string    `gorm:"unique;not null"`
	Password       string    `gorm:"not null"`
	DateRegistered time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DateUpdated    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	LastLogin      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	isActive       bool      `gorm:"unique;not null"`
}

type Role struct {
	RoleID      string `gorm:"type:uuid;primary_key"`
	RoleName    string `gorm:"unique;not null"`
	Description string `gorm:"unique;not null"`
	Price       int64  `gorm:"not null"`
}

type UserRole struct {
	UserRoleID string `gorm:"type:uuid;primary_key"`
	UserID     string `gorm:"type:uuid;not null;foreignkey:UserID"`
	RoleID     string `gorm:"type:uuid;not null;foreignkey:RoleID"`
}

type Permission struct {
	PermissionID   string `gorm:"type:uuid;primary_key"`
	PermissionName string `gorm:"unique;not null"`
	Description    string `gorm:"unique;not null"`
}

type RolePermission struct {
	RolePermissionID string `gorm:"type:uuid;primary_key"`
	RoleID           string `gorm:"type:uuid;not null;foreignkey:RoleID"`
	PermissionID     string `gorm:"type:uuid;not null;foreignkey:PermissionID"`
}

type AuditLog struct {
	AuditLogID  string    `gorm:"type:uuid;primary_key"`
	UserID      string    `gorm:"type:uuid;not null;foreignkey:UserID"`
	Activity    string    `gorm:"not null"`
	DateCreated time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type PasswordReset struct {
	PasswordResetID string    `gorm:"type:uuid;primary_key"`
	UserID          string    `gorm:"type:uuid;not null;foreignkey:UserID"`
	Token           string    `gorm:"not null"`
	DateCreated     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ExpiryDate      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
