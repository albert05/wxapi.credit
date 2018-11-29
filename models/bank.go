package models

// BANK LIST
const BANK_ICBC      =  1
const BANK_ABC   =  2
const BANK_CMB   =  3
const BANK_CCB   =  4
const BANK_BCCB      =  5
const BANK_BJRCB  =  6
const BANK_BOC   =  7
const BANK_COMM      =  8
const BANK_CMBC      =  9
const BANK_BOS   =  10
const BANK_CBHB      =  11
const BANK_CEB  =  12
const BANK_CIB  =  13
const BANK_CITIC    =  14
const BANK_CZB  =  15
const BANK_GDB  =  16
const BANK_HKBEA    =  17
const BANK_HXB  =  18
const BANK_HZCB   =  19
const BANK_NJCB    =  20
const BANK_PINGAN    =  21
const BANK_PSBC  =  22
const BANK_SDB  =  23
const BANK_SPDB      =  24
const BANK_SRCB  =  25

var bankList = map[int]string{
	BANK_ICBC  :  "工商银行",
	BANK_ABC   :  "农业银行",
	BANK_CMB   :  "招商银行",
	BANK_CCB   :  "建设银行",
	BANK_BCCB  :  "北京银行",
	BANK_BJRCB :  "北京农业商业银行",
	BANK_BOC   :  "中国银行",
	BANK_COMM  :  "交通银行",
	BANK_CMBC  :  "民生银行",
	BANK_BOS   :  "上海银行",
	BANK_CBHB  :  "渤海银行",
	BANK_CEB   :  "光大银行",
	BANK_CIB   :  "兴业银行",
	BANK_CITIC :  "中信银行",
	BANK_CZB   :  "浙商银行",
	BANK_GDB   :  "广发银行",
	BANK_HKBEA :  "东亚银行",
	BANK_HXB   :  "华夏银行",
	BANK_HZCB  :  "杭州银行",
	BANK_NJCB  :  "南京银行",
	BANK_PINGAN:  "平安银行",
	BANK_PSBC  :  "邮政储蓄银行",
	BANK_SDB   :  "深圳发展银行",
	BANK_SPDB  :  "浦发银行",
	BANK_SRCB  :  "上海农业商业银行",
}

func GetBankList() map[int]string {
	return bankList
}

func GetConvertBankList() []map[string]interface{} {
	l := len(bankList)

	list := make([]map[string]interface{}, l)
	for bankId, name := range bankList {
		list = append(list, map[string]interface{}{
			"bankId": bankId,
			"bankName": name,
		})
	}

	return list
}

func IsExistsBank(bankId int) bool {
	_, ok := bankList[bankId]
	return ok
}

func GetBankName(bankId int) string {
	if IsExistsBank(bankId) {
		return bankList[bankId]
	}

	return ""
}
