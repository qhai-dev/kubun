package domain

import "time"

type UserEntity struct {
	ID             int64     // 用户ID
	Name           string    // 用户名
	Nickname       string    // 用户昵称
	Email          string    // 用户邮箱
	Phone          string    // 用户手机号
	Avatar         string    // 用户头像
	HashPassword   string    // 用户密码
	EnterPriseList []string  // 关联企业列表
	DepartmentList []string  // 关联部门列表
	CreatedAt      time.Time // 创建时间
	UpdatedAt      time.Time // 更新时间
}