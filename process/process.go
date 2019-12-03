package process

import (
	"encoding/json"
	"fasoo.com/fklagent/bean"
	"fasoo.com/fklagent/mapper"
	"fasoo.com/fklagent/util/config"
	"fasoo.com/fklagent/util/log"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func runR(filePath, dataFileFullPath string) {
	attr, err := mapper.SelectCenterCodeANDTableIdByPath(filePath)
	if err != nil {
		log.ERROR(err.Error() + " : " + filePath)
		return
	}
	log.DEBUG(attr.CenterCode)
	log.DEBUG(attr.TableId)
	qisas, err := mapper.SelectQISA(attr.CenterCode, attr.TableId)
	if err != nil {
		log.ERROR(err.Error() + " : " + attr.CenterCode + " : " + attr.TableId)
		return
	}
	if len(qisas) == 0 {
		log.ERROR("Not found QI SA : " + attr.CenterCode + " : " + attr.TableId)
		return
	}

	//"c(\"qi1\",\"qi2\")"
	qi := "c("
	sa := "c("
	for index, qisa := range qisas {
		log.DEBUG(qisa.AttrDelimiter)
		log.DEBUG(qisa.FieldName)
		if qisa.AttrDelimiter == "QI" {
			qi += `"` + strings.ToLower(qisa.FieldName) + `"`
			if index != len(qisas)-1 {
				qi += ","
			}
		} else if qisa.AttrDelimiter == "SA" {
			sa += `"` + strings.ToLower(qisa.FieldName) + `"`
			if index != len(qisas)-1 {
				sa += ","
			}
		}
	}
	if qi == "c(" {
		qi = `c("")`
	} else {
		qi += ")"
	}
	if sa == "c(" {
		sa = `c("")`
	} else {
		sa += ")"
	}

	//R CMD BATCH --vanilla '--args "/data1/dev/sftp/bbp14/recv" "F_BBP14_00006" c("qi1","qi2") c("sa1","sa2") "BBP14" "TBBP14_ID_06"' /home/fasoo/R/script/r_script_fasoo.R.bak /home/fasoo/R/log/F_BBP14_00006.out
	rCmd := config.Cfg.RCmd
	//rCmdParam := config.Cfg.RCmdParam
	rScript := config.Cfg.RScriptPath
	logPath := config.Cfg.RLogPath

	dataDir, fileId := filepath.Split(filePath)
	logPath = filepath.Join(logPath, fileId+".out")

	//args := fmt.Sprintf("'--args %s %s %s %s %s %s'", `"/data1/dev/sftp/bbp14/recv"`, `"F_BBP14_00006"`, `c("qi1","qi2")`, `c("sa1","sa2")`, `"BBP14"`, `"TBBP14_ID_06"`)
	cmd := exec.Command(rCmd, dataDir, fileId, qi, sa, attr.CenterCode, attr.TableId, rScript, logPath)
	fmt.Println(cmd.String())
	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.ERROR(err.Error())
		return
	}
	defer cmdOut.Close()
	cmdErr, err := cmd.StderrPipe()
	if err != nil {
		log.ERROR(err.Error())
		return
	}
	defer cmdErr.Close()

	if err = cmd.Start(); err != nil {
		log.ERROR(err.Error())
		return
	}
	outBytes, err := ioutil.ReadAll(cmdOut)
	if err != nil {
		log.ERROR(err.Error())
		return
	}
	fmt.Println(string(outBytes))

	errBytes, err := ioutil.ReadAll(cmdErr)
	if err != nil {
		log.ERROR(err.Error())
		return
	}
	fmt.Println(string(errBytes))

	if err = cmd.Wait(); err != nil {
		log.ERROR(err.Error())
		return
	}

	//결과 파일 확인
	resultPath := config.Cfg.RResultPath
	_, dataFile := filepath.Split(dataFileFullPath)
	fileIdWithTime := strings.TrimSuffix(dataFile, filepath.Ext(dataFile))

	resultFileFullPath := filepath.Join(resultPath, fileIdWithTime+".json")
	f, err := ioutil.ReadFile(resultFileFullPath)
	if err != nil {
		log.ERROR("result file error : " + err.Error() + " : " + resultFileFullPath)
		return
	}

	rr := new(bean.RResult)
	err = json.Unmarshal(f, rr)
	if err != nil {
		log.ERROR(err.Error())
		return
	}
	temp := strings.Split(fileIdWithTime, "_")
	date := temp[len(temp)-1]
	log.DEBUG(date)

	result := new(bean.KLResult)
	result.ResultDate = date
	result.CenterCode = attr.CenterCode
	result.TableId = attr.TableId

	//결과 저장
	//K
	result.ResultType = "K"
	result.Result = strconv.Itoa(rr.K)
	err = mapper.InsertKLResult(result)
	if err != nil {
		log.ERROR(err.Error())
	}

	//L_lwc_nm
	result.ResultType = "L_lwc_nm"
	result.Result = strconv.Itoa(rr.LLwcNm)
	err = mapper.InsertKLResult(result)
	if err != nil {
		log.ERROR(err.Error())
	}

	//K_ERR_CNT
	result.ResultType = "K_ERR_CNT"
	result.Result = strconv.Itoa(rr.KErrCnt)
	err = mapper.InsertKLResult(result)
	if err != nil {
		log.ERROR(err.Error())
	}

	//L_lwc_nm_ERR_CNT
	result.ResultType = "L_lwc_nm_ERR_CNT"
	result.Result = strconv.Itoa(rr.LLwcNmErrCnt)
	err = mapper.InsertKLResult(result)
	if err != nil {
		log.ERROR(err.Error())
	}
}
