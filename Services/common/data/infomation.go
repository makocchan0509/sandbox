package data

type InfoReq struct {
	SessionId string `json:"sessionId"`
}

type InfoRes struct {
	Result        string        `json:"result"`
	Code          string        `json:"code"`
	Informartions []EntityInfos `json:"informartions"`
}

type EditInfoReq struct {
	SessionId     string `json:"sessionId"`
	InformationId string `json:"informationId"`
	Title         string `json:"title"`
	Contents      string `json:"contents"`
	EditFlg       string `json:"editFlg"`
}

type EditInfoRes struct {
	Result string `json:"result"`
	Code   string `json:"code"`
}

type EntityInfos struct {
	Information_id string `json:"informationId"`
	Title          string `json:"title"`
	Contents       string `json:"contents"`
	Issue_user     string `json:"issueUser"`
	User_type      string `json:"userType"`
	Update_date    string `json:"updateDate"`
	Create_date    string `json:"createDate"`
}
