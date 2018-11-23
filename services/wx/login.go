package wx

import (
	"github.com/astaxie/beego"
	"wxapi.credit/util/https"
	"encoding/json"
	"errors"
)

type LoginResp struct {
	OpenId 		string	`json:"openid"`
	SessionKey 	string	`json:"session_key"`
	UnionId 	string	`json:"unionid"`
}

func Login(code string) (*LoginResp, error) {
	params := map[string]string {
		"appid": beego.AppConfig.String("wxconfig::appid"),
		"secret": beego.AppConfig.String("wxconfig::secret"),
		"grant_type": beego.AppConfig.String("wxconfig::grant_type"),
		"js_code": code,
	}

	body, err := https.Get(beego.AppConfig.String("wxconfig::url"), params)
	if err != nil {
		return nil, err
	}

	var resp LoginResp
	json.Unmarshal(body, &resp)

	if resp.OpenId == "" {
		return nil, errors.New("result json decode failed")
	}

	return &resp, nil
}
