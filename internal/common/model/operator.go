package model

import "time"

// Operator mirrors the db_operator table structure.
type Operator struct {
	ID            int        `gorm:"column:id;primaryKey"`
	CName         string     `gorm:"column:cname"`
	Phone         string     `gorm:"column:phone"`
	Email         string     `gorm:"column:email"`
	Nickname      string     `gorm:"column:nickname"`
	Password      string     `gorm:"column:password"`
	OpenID        string     `gorm:"column:openId"`
	UnionID       string     `gorm:"column:unionId"`
	UserID        string     `gorm:"column:userid"`
	Department    *int       `gorm:"column:department"`
	Position      string     `gorm:"column:position"`
	LoginNum      *int       `gorm:"column:login_num"`
	CreateTime    *time.Time `gorm:"column:create_time"`
	LastLoginTime *time.Time `gorm:"column:last_login_time"`
	LastIP        string     `gorm:"column:last_ip"`
	Status        int        `gorm:"column:status"`
	Remark        string     `gorm:"column:remark"`
	IsAdmin       string     `gorm:"column:isadmin"`
	Roles         string     `gorm:"column:roles"`
	IsDel         int        `gorm:"column:is_del"`
	Abbr          string     `gorm:"column:abbr"`
	Job           string     `gorm:"column:job"`
}

// TableName returns the table name for Operator.
func (Operator) TableName() string {
	return "db_operator"
}
