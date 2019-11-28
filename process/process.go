package process

import (
	"fasoo.com/fklagent/mapper"
	"fasoo.com/fklagent/util/config"
	"fasoo.com/fklagent/util/log"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func runR(filePath, dataFilePath string) {
	attr, err := mapper.SelectCenterCodeANDTableIdByPath(filePath)
	if err != nil {
		return
	}
	log.DEBUG(attr.CenterCode)
	log.DEBUG(attr.TableId)

	rCmd := config.Cfg.RCmd
	rScript := config.Cfg.RScriptPath

	cmd := exec.Command(rCmd, rScript)
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

	//if err = cmd.Wait(); err != nil {
	//	log.ERROR(err.Error())
	//	return nil
	//}

}
