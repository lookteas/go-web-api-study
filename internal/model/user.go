package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"` // 不在JSON中显示密码
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateUserRequest 创建用户请求结构
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UpdateUserRequest 更新用户请求结构
type UpdateUserRequest struct {
	Username string `json:"username,omitempty" validate:"omitempty,min=3,max=20"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}