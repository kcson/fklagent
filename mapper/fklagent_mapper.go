package mapper

import (
	"fasoo.com/fklagent/bean"
	"fasoo.com/fklagent/db"
	"fasoo.com/fklagent/util/log"
)

func SelectCenterCodeANDTableIdByPath(path string) (*bean.Attr, error) {
	attr := new(bean.Attr)
	err := db.DB.Get(attr,
		`SELECT
					cent_cd,
					tbl_id
				FROM
					bbp_prd_tabown.data_fle_sttr
				WHERE
					fle_path = $1`, path)
	if err != nil {
		log.ERROR(err.Error())
		return attr, err
	}
	return attr, nil
}

func SelectQISA(centerCode, tableId string) ([]bean.Attr, error) {
	var attrs []bean.Attr
	err := db.DB.Select(&attrs,
		`SELECT 
					attr_dlmt,
					fld_nm,
					COALESCE(fld_eng_nm,'') as fld_eng_nm
				FROM
					bbp_prd_tabown.data_attr_tbl_sttr
				WHERE
					cent_cd = $1 AND tbl_id = $2`, centerCode, tableId)
	if err != nil {
		log.ERROR(err.Error())
		return attrs, err
	}

	return attrs, nil
}

func InsertKLResult(result *bean.KLResult) error {
	_, err := db.DB.NamedExec(
		`INSERT INTO 
				bbp_prd_tabown.msrm_rslt_tbl_sttr (
					strd_date,
					cent_cd,
					tbl_id,
					msrm_dv,
					msrm_rslt
				)
				VALUES (
				    :strd_date,
					:cent_cd,
					:tbl_id,
					:msrm_dv,
					:msrm_rslt
				)`,
		result)
	if err != nil {
		log.ERROR(err.Error())
		return err
	}
	return nil
}
