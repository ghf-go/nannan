package es_driver

type EsResponseBaseDoc struct {
	Index string `json:"_index"`
	Type  string `json:"_type"`
	Id    string `json:"_id"`
	Found bool   `json:"found"`
}
type EsResponseDocsBaseDoc struct {
	*EsResponseBaseDoc
	Version     int `json:"_version"`
	SeqNo       int `json:"_seq_no"`
	PrimaryTerm int `json:"_primary_term"`
}
type EsResponseBase struct {
	ErrorMsg  interface{} `json:"error"`
	ErrorCode int         `json:"status"`
}
type EsResponseUpdate struct {
	*EsResponseBase
	*EsResponseBaseDoc
	Version int    `json:"_version"`
	Result  string `json:"result"`
	Shards  struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	SeqNo       int `json:"_seq_no"`
	PrimaryTerm int `json:"_primary_term"`
}

type T struct {
	*EsResponseBase
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
}
type EsResponseHitBaseHits struct {
	Total struct {
		Value    int    `json:"value"`
		Relation string `json:"relation"`
	} `json:"total"`
	MaxScore float64 `json:"max_score"`
}
type EsResponseHitBaseDoc struct {
	*EsResponseBaseDoc
	Score float64 `json:"_score"`
}

func (r *EsResponseUpdate) IsSuccess() bool {
	return r.Result == "deleted" || r.Result == "updated" || r.Result == "created"
}
func (r *EsResponseUpdate) IsDelOK() bool {
	return r.Result == "deleted"
}
func (r *EsResponseUpdate) IsUpdateOK() bool {
	return r.Result == "updated"
}
func (r *EsResponseUpdate) IsCreateOK() bool {
	return r.Result == "created"
}
