package controllers

import "wxapi.credit/models"

// Operations about Car
type IndexController struct {
	BaseController
}

// @Title GetCreditCard
// @Description User GetCreditCard
// @Success 200 {int}
// @Failure 403
// @router /info [post]
func (idx *IndexController) Info() {

	list, err := models.FindCardList(idx.User.Id)
	if err != nil {
		idx.Abort666("get info failed", map[string]interface{}{
			"errMsg": err.Error(),
		})
	}

	idx.JsonSucc("get success", map[string]interface{}{
		"headerUrl": "",
		"totalMoney": idx.User.TotalDebt,
		"buttonInfo": "",
		"recommendCard": "",
		"creditCardList": list, //TODO
		"share": "",
	})
}
