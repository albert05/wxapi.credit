package models

import (
	"wxapi.credit/util/mysql"
)

type User struct {
	Id       	int		`db:"id"`
	Gender  	int		`db:"gender"`
	Status  	int		`db:"status"`
	CreatedAt  	int32	`db:"created_at"`
	UpdatedAt  	int32	`db:"updated_at"`
	Openid 		string  `db:"openid"`
	Nickname 	string	`db:"nick_name"`
	AvatarUrl  	string	`db:"avatar_url"`
	Province  	string	`db:"province"`
	City  		string	`db:"city"`
	Country  	string	`db:"country"`
	Language  	string	`db:"language"`
	Remind  	string	`db:"remind"`
	TotalCredit int		`db:"total_credit"`
	TotalDebt  	int		`db:"total_debt"`
}

func (u *User) GetTableName() string {
	return "user"
}

func FindUser(openid string) (*User, error) {
	var u User
	err := mysql.FindCond(&u, map[string]string{"openid": openid}, "*")

	return &u, err
}

func (u *User) Update(data mysql.MapModel) bool {
	data.Load(u)
	return mysql.Update(u)
}

func InsertUser(data mysql.MapModel) *User {
	var u User
	data.Load(&u)

	if ok := mysql.Insert(&u); ok {
		return &u
	}

	return nil
}

