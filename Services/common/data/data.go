package data

type LoginReq struct {
	LoginId  string `json:"loginId"`
	Password string `json:"password"`
}

type LoginRes struct {
	Result    string `json:"result"`
	Code      string `json:"code"`
	SessionId string `json:"sessionId"`
	UserType  string `json:"userType"`
	ReqId     string `json:"reqId"`
	ReqPass   string `json:"reqPass"`
}

type GetSessionReq struct {
	SessionId string `json:"sessionId"`
}

type CheckSessionReq struct {
	SessionId string `json:"sessionId"`
}

type CheckSessionRes struct {
	Result string `json:"result"`
}
