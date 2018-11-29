package controllers

import (
	"wxapi.credit/models"
	"wxapi.credit/util/mysql"
	"wxapi.credit/common"
)

// Operations about Car
type CardController struct {
	BaseController
}

// @Title AddCreditCard
// @Description User AddCreditCard
// @Param cardId query string false
// @Param cardNum query string true
// @Param bankId query string true
// @Param cardholder query string true
// @Param billDate query string true
// @Param dueDate query string true
// @Param lineOfCredit query string true
// @Param debtMoney query string true
// @Success 200 {int}
// @Failure 403
// @router /add-credit-card [post]
func (u *CardController) AddCreditCard() {
	u.MustParams("cardNum", "bankId", "cardholder", "billDate", "dueDate", "lineOfCredit", "debtMoney")

	if !models.IsExistsBank(common.Str2Int(u.Params["bankId"])) {
		u.Abort666("add failed", map[string]interface{}{
			"errMsg": "bankId is not exists",
		})
	}

	isUpt := false
	if cardId, ok := u.Params["cardId"]; ok && cardId != "" {
		f, err := models.FindCard(cardId)
		if err != nil {
			u.Abort666("add failed", map[string]interface{}{
				"errMsg": err.Error(),
			})
		}

		if f != nil && err == nil {
			f.Update(mysql.MapModel{
				"card_no": common.Str2Int(u.Params["cardNum"]),
				"bank_id": common.Str2Int(u.Params["bankId"]),
				"holder": u.Params["cardholder"],
				"bill_date": common.Str2Int(u.Params["billDate"]),
				"due_date": common.Str2Int(u.Params["dueDate"]),
				"credit_line": common.Float2Money(common.Str2Float(u.Params["lineOfCredit"])),
				"debt_money": common.Float2Money(common.Str2Float(u.Params["debtMoney"])),
			})
			isUpt = true
		}
	}

	if !isUpt {
		models.InsertCard(mysql.MapModel{
			"card_no": common.Str2Int(u.Params["cardNum"]),
			"bank_id": common.Str2Int(u.Params["bankId"]),
			"holder": u.Params["cardholder"],
			"bill_date": common.Str2Int(u.Params["billDate"]),
			"due_date": common.Str2Int(u.Params["dueDate"]),
			"credit_line": common.Float2Money(common.Str2Float(u.Params["lineOfCredit"])),
			"debt_money": common.Float2Money(common.Str2Float(u.Params["debtMoney"])),
			"user_id": u.User.Id,
		})
	}

	u.JsonSucc("add success")
}

// @Title DelCreditCard
// @Description User DelCreditCard
// @Param cardId query string true
// @Success 200 {int}
// @Failure 403
// @router /del-credit-card [post]
func (u *CardController) DelCreditCard() {
	u.MustParams("cardId")

	f, err := models.FindCard(u.Params["cardId"])
	if f == nil || err != nil {
		u.Abort666("del failed", map[string]interface{}{
			"errMsg": err.Error(),
		})
	}

	f.Update(mysql.MapModel{
		"status": 1,
	})

	u.JsonSucc("del success")
}

// @Title GetCreditCard
// @Description User GetCreditCard
// @Param cardId query string true
// @Success 200 {int}
// @Failure 403
// @router /get-credit-card [post]
func (u *CardController) GetCreditCard() {
	u.MustParams("cardId")

	f, err := models.FindCard(u.Params["cardId"])
	if f == nil || err != nil {
		u.Abort666("get failed", map[string]interface{}{
			"errMsg": err.Error(),
		})
	}

	u.JsonSucc("get success", map[string]interface{}{
		"cardNum": f.CardNo,
		"bankId": f.BankId,
		"cardholder": f.Holder,
		"billDate": f.BillDate,
		"dueDate": f.DueDate,
		"lineOfCredit": common.Money2Float(f.CreditLine),
		"debtMoney": common.Money2Float(f.DebtMoney),
	})
}

// @Title GetBankList
// @Description User GetBankList
// @Success 200 {int}
// @Failure 403
// @router /get-bank-list [post]
func (u *CardController) GetBankList() {
	list := models.GetConvertBankList()

	u.JsonSucc("get success", map[string]interface{}{
		"bankList": list,
	})
}
