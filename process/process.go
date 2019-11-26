package process

import (
	"fasoo.com/fklagent/util/config"
	"fasoo.com/fklagent/util/log"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func runR() {
	rCmd := config.Cfg.RCmd

	cmd := exec.Command(rCmd, "--version")
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
	//outBytes, err := ioutil.ReadAll(cmdOut)
	//if err != nil {
	//	log.ERROR(err.Error())
	//	return nil
	//}
	//fmt.Println(string(outBytes))

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
