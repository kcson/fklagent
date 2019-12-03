package bean

type Attr struct {
	CenterCode    string `db:"cent_cd"`
	TableId       string `db:"tbl_id"`
	AttrDelimiter string `db:"attr_dlmt"`
	FieldName     string `db:"fld_nm"`
	FieldEngName  string `db:"fld_eng_nm"`
}

type KLResult struct {
	ResultDate string `db:"strd_date"`
	CenterCode string `db:"cent_cd"`
	TableId    string `db:"tbl_id"`
	ResultType string `db:"msrm_dv"`
	Result     string `db:"msrm_rslt"`
}

type RResult struct {
	CenterCode   string `json:"cent_cd"`
	TableId      string `json:"tbl_id"`
	Path         string `json:"path"`
	File         string `json:"file"`
	K            int    `json:"K"`
	LLwcNm       int    `json:"L_lwc_nm"`
	KErrCnt      int    `json:"K_ERR_CNT"`
	LLwcNmErrCnt int    `json:"L_lwc_nm_ERR_CNT"`
}
