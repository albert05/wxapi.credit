package services

type LoginFilter struct {
	nLogin map[string]bool
}

var LFilter LoginFilter

func init() {
	LFilter = LoginFilter{}
	LFilter.nLogin = make(map[string]bool)
}

func (lf *LoginFilter) Add(uri string) {
	lf.nLogin[uri] = true
}

func (lf *LoginFilter) GetNLoginList() map[string]bool {
	return lf.nLogin
}
