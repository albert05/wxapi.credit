package models

import (
	"wxapi.credit/util/mysql"
)

type Feedback struct {
	Id       	int		`db:"id"`
	UserId  	int		`db:"user_id"`
	Phone  		int		`db:"phone"`
	Content  	string	`db:"content"`
	Status 		int  	`db:"status"`
	CreatedAt  	int32	`db:"created_at"`
	UpdatedAt  	int32	`db:"updated_at"`
}

func (u *Feedback) GetTableName() string {
	return "feedback"
}

func FindFeedback(userId string) (*Feedback, error) {
	var info Feedback
	err := mysql.FindCond(&info, map[string]string{"user_id": userId}, "*")

	return &info, err
}

func (u *Feedback) Update(data mysql.MapModel) bool {
	data.Load(u)
	return mysql.Update(u)
}

func InsertFeedback(data mysql.MapModel) *Feedback {
	var info Feedback
	data.Load(&info)

	if ok := mysql.Insert(&info); ok {
		return &info
	}

	return nil
}

