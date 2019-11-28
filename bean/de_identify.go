package bean

type Attr struct {
	CenterCode    string `db:"cent_cd"`
	TableId       string `db:"tbl_id"`
	AttrDelimiter string `db:"attr_dlmt"`
	FieldName     string `db:"fld_nm"`
	FieldEngName  string `db:"fld_eng_nm"`
}
