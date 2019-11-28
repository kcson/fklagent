package process

import (
	"fasoo.com/fklagent/util/config"
	"fasoo.com/fklagent/util/log"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	if filepath.Ext(path) != ".csv" {
		return nil
	}
	dir, file := filepath.Split(path)
	if filepath.Base(dir) != "recv" {
		return nil
	}
	log.DEBUG(dir)
	log.DEBUG(file)

	fileId := strings.TrimSuffix(file, filepath.Ext(file))
	successFile := filepath.Join(dir, fileId+".success")
	log.DEBUG(successFile)
	//success 파일이 있는 경우 처리가 끝난 파일 이므로 skip
	if _, err = os.Stat(successFile); err == nil {
		return nil
	}

	//센터 코드와 파일 id를 가져오기 위해 file path 생성
	sep := strings.Split(successFile, "_")
	if len(sep) < 5 {
		return nil
	}
	filePath := strings.Join(sep[:4], "_")
	log.DEBUG(filePath)

	runR(filePath, path)

	return nil
}
