package models

import (
	"wxapi.credit/util/mysql"
	"fmt"
)

const CONSUME_RECORD_TABLE = "consume_record"

type Consume struct {
	Id       	int		`db:"id"`
	UserId  	int		`db:"user_id"`
	CardId  	int		`db:"card_id"`
	Money  		int		`db:"money"`
	Type  		int		`db:"type"`
	Date  		string	`db:"date"`
	Remark  	string	`db:"remark"`
	CreatedAt  	int32	`db:"created_at"`
	UpdatedAt  	int32	`db:"updated_at"`
}

func (u *Consume) GetTableName() string {
	return CONSUME_RECORD_TABLE
}

func FindConsumeRecordList(userId int) ([]mysql.MapModel, error) {
	sql := fmt.Sprintf("select * from %s where user_id=%d", CONSUME_RECORD_TABLE, userId)
	return mysql.FindAll(sql)
}

func FindConsumeRecord(cardId string) (*Consume, error) {
	var info Consume
	err := mysql.FindCond(&info, map[string]string{"id": cardId}, "*")

	return &info, err
}

func (u *Consume) Update(data mysql.MapModel) bool {
	data.Load(u)
	return mysql.Update(u)
}

func InsertConsumeRecord(data mysql.MapModel) *Consume {
	var info Consume
	data.Load(&info)

	if ok := mysql.Insert(&info); ok {
		return &info
	}

	return nil
}

const TYPE_LIFE = 1
const TYPE_SHOP = 2
const TYPE_EAT = 3
const TYPE_REPAY = 4

var TypeList = map[int]string {
	TYPE_LIFE: "生活费",
	TYPE_SHOP: "购物",
	TYPE_EAT: "吃饭",
	TYPE_REPAY: "还款",
}

func GetConvertTypeList() []map[string]interface{} {
	l := len(TypeList)

	list := make([]map[string]interface{}, l)
	for typeId, name := range TypeList {
		list = append(list, map[string]interface{}{
			"consumeId": typeId,
			"consumeName": name,
		})
	}

	return list
}



