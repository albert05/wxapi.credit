package models

import (
	"wxapi.credit/util/mysql"
)

type FormidInfo struct {
	Id       	int		`db:"id"`
	UserId  	int		`db:"user_id"`
	Formid  	int		`db:"formid"`
	Status 		string  `db:"status"`
	CreatedAt  	int32	`db:"created_at"`
	UpdatedAt  	int32	`db:"updated_at"`
}

func (u *FormidInfo) GetTableName() string {
	return "formid_info"
}

func FindFormidInfo(formid string) (*FormidInfo, error) {
	var info FormidInfo
	err := mysql.FindCond(&info, map[string]string{"formid": formid}, "*")

	return &info, err
}

func (u *FormidInfo) Update(data mysql.MapModel) bool {
	data.Load(u)
	return mysql.Update(u)
}

func InsertFormidInfo(data mysql.MapModel) *FormidInfo {
	var info FormidInfo
	data.Load(&info)

	if ok := mysql.Insert(&info); ok {
		return &info
	}

	return nil
}

