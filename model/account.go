package model

import (
	"errors"
	"gin-scaffold/pkg/database"

	"gorm.io/gorm"
)

type Account struct {
	Common
	UserName string `json:"user_name" gorm:"type:varchar(64);not null;default:'';comment:'用户名'"` // 用户名
	Password string `json:"password" gorm:"type:varchar(128);not null;default:'';comment:'密码'"`  // 密码
	Mobile   int    `json:"mobile" gorm:"type:char(11);index;commnet:'手机号'"`                     // 手机号
}

// Insert .
func (a *Account) Insert() error {
	return database.DB.Save(&a).Error
}

// Delete .
func (a Account) Delete() error {
	return database.DB.Delete(&a).Error
}

// Updates .
func (a Account) Updates() error {
	return database.DB.Updates(&a).Error
}

// GetInfoByMobile 通过手机号来获取用户信息
func (a *Account) GetInfoByMobile(mobile string) (bool, error) {
	err := database.DB.Take(&a, "mobile = ?", mobile).Error
	return errors.Is(err, gorm.ErrRecordNotFound), err
}

// List  分页获取所有记录
func (a Account) List(page, pageSize int, query, order string) (list []Account, total int64, err error) {
	db := database.DB.Model(&a)
	db, err = buildOrder(order, db)
	if err != nil {
		return nil, 0, err
	}

	db, err = buildWhere(query, db)
	if err != nil {
		return nil, 0, err
	}

	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}

	if err := db.Scopes(paginate(page, pageSize)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return
}
