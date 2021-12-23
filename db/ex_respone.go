package db

type EsResponse struct {
	DbName    string        `json:"_index"`
	Table     string        `json:"_type"`
	ID        string        `json:"_id"`
	Version   int           `json:"_version"`
	Created   bool          `json:"created"`
	Exists    bool          `json:"exists"`
	Found     bool          `json:"found"`
	Source    interface{}   `json:"_source"`
	ErrorMsg  string        `json:"error"`
	ErrorCode int           `json:"status"`
	Docs      []*EsResponse `json:"docs"`
	HttpCode  int
}

//是否删除成功
func (er *EsResponse) IsDeleteOK() bool {
	return er.HttpCode == 200 && er.Found
}

//是否找到
func (er *EsResponse) IsFindOK() bool {
	return er.HttpCode == 200 && er.Found
}
func (er *EsResponse) IsCreateOK() bool {
	return er.HttpCode == 200 && er.Created
}
func (er *EsResponse) IsUpdateOK() bool {
	return er.HttpCode == 200 && !er.Created
}
