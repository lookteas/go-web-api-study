package service

import (
	"errors"
	"go-web-api-study/internal/model"
	"time"
)

// UserService 用户服务接口
type UserService interface {
	CreateUser(req model.CreateUserRequest) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	UpdateUser(id int, req model.UpdateUserRequest) (*model.User, error)
	DeleteUser(id int) error
	Login(req model.LoginRequest) (*model.LoginResponse, error)
}

// userService 用户服务实现
type userService struct {
	// 这里将来会添加数据库连接
	users []model.User // 临时使用内存存储
}

// NewUserService 创建用户服务实例
func NewUserService() UserService {
	return &userService{
		users: make([]model.User, 0),
	}
}

// CreateUser 创建用户
func (s *userService) CreateUser(req model.CreateUserRequest) (*model.User, error) {
	// 检查用户名是否已存在
	for _, user := range s.users {
		if user.Username == req.Username {
			return nil, errors.New("用户名已存在")
		}
		if user.Email == req.Email {
			return nil, errors.New("邮箱已存在")
		}
	}
	
	// 创建新用户
	user := model.User{
		ID:        len(s.users) + 1,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password, // 实际项目中需要加密
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	s.users = append(s.users, user)
	return &user, nil
}

// GetUserByID 根据ID获取用户
func (s *userService) GetUserByID(id int) (*model.User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("用户不存在")
}

// GetUserByUsername 根据用户名获取用户
func (s *userService) GetUserByUsername(username string) (*model.User, error) {
	for _, user := range s.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("用户不存在")
}

// UpdateUser 更新用户信息
func (s *userService) UpdateUser(id int, req model.UpdateUserRequest) (*model.User, error) {
	for i, user := range s.users {
		if user.ID == id {
			if req.Username != "" {
				s.users[i].Username = req.Username
			}
			if req.Email != "" {
				s.users[i].Email = req.Email
			}
			s.users[i].UpdatedAt = time.Now()
			return &s.users[i], nil
		}
	}
	return nil, errors.New("用户不存在")
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(id int) error {
	for i, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return errors.New("用户不存在")
}

// Login 用户登录
func (s *userService) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.GetUserByUsername(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	
	// 简单密码验证（实际项目中需要使用加密验证）
	if user.Password != req.Password {
		return nil, errors.New("用户名或密码错误")
	}
	
	// 生成简单token（实际项目中使用JWT）
	token := "simple_token_" + user.Username
	
	return &model.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}