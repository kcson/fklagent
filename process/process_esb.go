package process

import (
	"fasoo.com/fklagent/util/config"
	"fasoo.com/fklagent/util/log"
	"fmt"
	"os"
	"path/filepath"
)

func ESB() {
	log.INFO("start process esb!!")
	path := config.Cfg.ESBPath

	filepath.Walk(path, naviESBDir)

}

func naviESBDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if info.IsDir() {
		return nil
	}
	fmt.Println(path)

	return nil
}
