package controllers

import (
	"wxapi.credit/models"
	"wxapi.credit/util/mysql"
	"wxapi.credit/common"
)

// Operations about Consume
type ConsumeController struct {
	BaseController
}

// @Title AddCreditCard
// @Description User AddCreditCard
// @Param cardId query int false
// @Param expenseMoney query int true
// @Param expenseType query int true
// @Param expenseDate query string true
// @Param remark query string true
// @Success 200 {int}
// @Failure 403
// @router /add-expense-record [post]
func (this *ConsumeController) AddExpenseRecord() {
	this.MustParams("cardId", "expenseMoney", "expenseType", "expenseDate", "remark")

	if _, ok := models.TypeList[common.Str2Int(this.Params["expenseType"])]; !ok {
		this.Abort666("add failed")
	}

	money := common.Float2Money(common.Str2Float(this.Params["expenseMoney"]))

	if cardId, ok := this.Params["cardId"]; ok && cardId != "" {
		f, err := models.FindCard(cardId)
		if f == nil || err != nil {
			this.Abort666("add failed", map[string]interface{}{
				"errMsg": err.Error(),
			})
		}

		models.InsertConsumeRecord(mysql.MapModel{
			"type": common.Str2Int(this.Params["expenseType"]),
			"date": common.Str2Int(this.Params["expenseDate"]),
			"remark": this.Params["remark"],
			"money": money,
			"user_id": this.User.Id,
			"card_id": f.Id,
		})


		this.User.Update(mysql.MapModel{
			"total_debt": money + this.User.TotalDebt,
		})
	}

	this.JsonSucc("add success")
}

// @Title GetCreditCard
// @Description User GetCreditCard
// @Success 200 {int}
// @Failure 403
// @router /get-consume-list [post]
func (this *ConsumeController) GetConsumeList() {
	list := models.GetConvertTypeList()

	this.JsonSucc("get success", map[string]interface{}{
		"consumeList": list,
	})
}
