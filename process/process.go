package process

import (
	"fasoo.com/fklagent/mapper"
	"fasoo.com/fklagent/util/config"
	"fasoo.com/fklagent/util/log"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

func runR(filePath, dataFileFullPath string) {
	attr, err := mapper.SelectCenterCodeANDTableIdByPath(filePath)
	if err != nil {
		log.ERROR(err.Error())
		return
	}
	log.DEBUG(attr.CenterCode)
	log.DEBUG(attr.TableId)
	qisas, err := mapper.SelectQISA(attr.CenterCode, attr.TableId)
	if err != nil {
		log.ERROR(err.Error())
		return
	}

	//"c(\"qi1\",\"qi2\")"
	qi := "c("
	sa := "c("
	for index, qisa := range qisas {
		log.DEBUG(qisa.AttrDelimiter)
		log.DEBUG(qisa.FieldName)
		if qisa.AttrDelimiter == "QI" {
			qi += `"` + qisa.FieldName + `"`
			if index != len(qisas)-1 {
				qi += ","
			}
		} else if qisa.AttrDelimiter == "SA" {
			sa += `"` + qisa.FieldName + `"`
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
}
