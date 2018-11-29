package models

import (
	"wxapi.credit/util/mysql"
	"fmt"
)

const CARD_TABLE = "card"

type Card struct {
	Id       	int		`db:"id"`
	UserId  	int		`db:"user_id"`
	CardNo  	int		`db:"card_no"`
	CardNoFull  string	`db:"card_no_full"`
	BankId  	int		`db:"bank_id"`
	Holder  	string	`db:"holder"`
	BillDate  	int		`db:"bill_date"`
	DueDate  	int		`db:"due_date"`
	CreditLine  int		`db:"credit_line"`
	DebtMoney  	int		`db:"debt_money"`
	Status 		int  	`db:"status"`
	CreatedAt  	int32	`db:"created_at"`
	UpdatedAt  	int32	`db:"updated_at"`
}

func (u *Card) GetTableName() string {
	return CARD_TABLE
}

func FindCardList(userId int) ([]mysql.MapModel, error) {
	sql := fmt.Sprintf("select * from %s where user_id=%d", CARD_TABLE, userId)
	return mysql.FindAll(sql)
}

func FindCard(cardId string) (*Card, error) {
	var info Card
	err := mysql.FindCond(&info, map[string]string{"id": cardId}, "*")

	return &info, err
}

func (u *Card) Update(data mysql.MapModel) bool {
	data.Load(u)
	return mysql.Update(u)
}

func InsertCard(data mysql.MapModel) *Card {
	var info Card
	data.Load(&info)

	if ok := mysql.Insert(&info); ok {
		return &info
	}

	return nil
}

