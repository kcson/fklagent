package process

import (
	"fasoo.com/fklagent/util/config"
	"fasoo.com/fklagent/util/log"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func SFTP() {
	log.INFO("start process sftp!!")
	path := config.Cfg.SFTPPath
	if err := filepath.Walk(path, naviSFTPDir); err != nil {
		log.ERROR(err.Error())
	}

}

func naviSFTPDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if info.IsDir() {
		return nil
	}
	if filepath.Ext(path) != ".fin" {
		return nil
	}
	dir, file := filepath.Split(path)
	if filepath.Base(dir) != "recv" {
		return nil
	}

	fmt.Println(dir)
	fmt.Println(file)
	fileId := strings.TrimSuffix(file, filepath.Ext(file))
	successFile := filepath.Join(dir, fileId+".success")
	fmt.Println(successFile)
	if _, err = os.Stat(successFile); err == nil {
		return nil
	}
	runR()

	return nil
}
