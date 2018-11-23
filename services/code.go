package services

const SUCCESS_CODE  = 200
const FAILED_CODE  = 100

type Code struct {
	Code int
	Message string
	Data  	map[string]interface{}
}

func SuccRet(message string) Code {
	return Code{
		Code: 		SUCCESS_CODE,
		Message: 	message,
	}
}

func SuccRetEx(message string, ex map[string]interface{}) Code {
	return Code{
		Code: 		SUCCESS_CODE,
		Message: 	message,
		Data: ex,
	}
}

func FailedRet(message string) Code {
	return Code{
		Code: 		FAILED_CODE,
		Message: 	message,
	}
}

func FailedRetEx(message string, ex map[string]interface{}) Code {
	return Code{
		Code: 		FAILED_CODE,
		Message: 	message,
		Data: ex,
	}
}

func CustomRet(code int, message string) Code {
	return Code{
		Code: 		code,
		Message: 	message,
	}
}

func CustomRetEx(code int, message string, ex map[string]interface{}) Code {
	return Code{
		Code: 		code,
		Message: 	message,
		Data: ex,
	}
}
