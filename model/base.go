package model

import (
	"errors"
	"gin-scaffold/pkg/uuid"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Common struct {
	ID        string    `json:"id" gorm:"type:varchar(32);primaryKey;not null;comment:'主键ID'"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:'创建时间'"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:'更新时间'"`
	DeletedAt time.Time `json:"deleted_at" gorm:"comment:'软删除时间'"`
}

// Paginate 数据分页
func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		if pageSize > 100 {
			pageSize = 100
		} else if pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// buildWhere .
func buildWhere(rawQuery string, db *gorm.DB) (*gorm.DB, error) {
	if rawQuery != "" {
		queryList := strings.Split(rawQuery, ",")
		for _, query := range queryList {
			oneQuery := strings.Split(query, ":")
			if len(oneQuery) != 2 && len(oneQuery) != 3 {
				return db, errors.New("parseRawQuery error, rawQuery should like: 'title:=:golang,name:like:%jason%,id:<:100' , if the whereType is '=', you can omitted it: title:golang, notice: '%' after encode is %25")
			}
			if len(oneQuery) == 2 {
				field := oneQuery[0]
				whereType := "="
				value := oneQuery[1]
				db = db.Where(field+" "+whereType+" "+"?", value)
			}
			if len(oneQuery) == 3 {
				field := oneQuery[0]
				whereType := oneQuery[1]
				value := oneQuery[2]
				db = db.Where(field+" "+whereType+" "+"?", value)
			}
		}
	}
	return db, nil
}

// buildOrder .
func buildOrder(rawOrder string, db *gorm.DB) (*gorm.DB, error) {
	if rawOrder != "" {
		orders := strings.Split(rawOrder, ",")
		for _, order := range orders {
			oneOrder := strings.Split(order, ":")
			if len(oneOrder) != 1 && len(oneOrder) != 2 {
				return db, errors.New("parse rawOrder error, rawOrder should like:'created:desc,id:asc,name', orderType default is asc")
			}

			if len(oneOrder) == 1 {
				field := oneOrder[0]
				db = db.Order(field)
			}
			if len(oneOrder) == 2 {
				field := oneOrder[0]
				orderType := oneOrder[1]
				db = db.Order(field + " " + orderType)
			}
		}
	}
	return db, nil
}

func (c *Common) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.UUID()
	return nil
}
