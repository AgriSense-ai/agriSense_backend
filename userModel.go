package main

import (
	"time"

	"github.com/google/uuid"
)

// TODO: Read this https://gorm.io/docs/models.html
// TODO: Read this https://www.linkedin.com/pulse/using-gorm-golang-orm-alex-edwards/
// TODO: Read this https://www.calhoun.io/why-you-should-never-use-golangs-default-uuid-package/
// TODO: Read this https://www.linkedin.com/pulse/database-structure-user-roles-management-aj-february/

type User struct {
	UserID         uuid.UUID `gorm:"type:uuid;primary_key"`
	Username       string    `gorm:"unique;not null"`
	FirstName      string    `gorm:"unique;not null"`
	LastName       string    `gorm:"unique:not null"`
	Email          string    `gorm:"unique;not null"`
	Password       string    `gorm:"not null"`
	DateRegistered time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DateUpdated    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	LastLogin      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	IsActive       bool      `gorm:"unique;not null"`
}

type Role struct {
	RoleID      uuid.UUID `gorm:"type:uuid;primary_key"`
	RoleName    string    `gorm:"unique;not null"`
	Description string    `gorm:"unique;not null"`
	Price       int64     `gorm:"not null"`
}

type UserRole struct {
	UserRoleID uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID     User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;type:uuid;not null;foreignkey:UserID;references:User"`
	RoleID     Role      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;type:uuid;not null;foreignkey:RoleID;references:Role"`
}

type Permission struct {
	PermissionID   uuid.UUID `gorm:"type:uuid;primary_key"`
	PermissionName string    `gorm:"unique;not null"`
	Description    string    `gorm:"unique;not null"`
}

type RolePermission struct {
	RolePermissionID uuid.UUID  `gorm:"type:uuid;primary_key"`
	RoleID           Role       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;type:uuid;not null;foreignkey:RoleID;references:Role"`
	PermissionID     Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;type:uuid;not null;foreignkey:PermissionID;references:Permission"`
}

type AuditLog struct {
	AuditLogID  uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;type:uuid;not null;foreignkey:UserID;references:User"`
	Activity    string    `gorm:"not null"`
	DateCreated time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type PasswordReset struct {
	PasswordResetID uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID          User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;type:uuid;not null;foreignkey:UserID;references:User"`
	Token           string    `gorm:"not null"`
	DateCreated     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ExpiryDate      time.Time `gorm:"default:time.Now().Add(3 * 24 * time.Hour)"`
}
