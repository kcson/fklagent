package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	root := "D:\\mywork\\go\\src\\fasoo.com"
	cmd := exec.Command("cmd", "/c", "dir")
	cmdOut,_ := cmd.StdoutPipe()
	defer cmdOut.Close()
	cmd.Start()
	outBytes,_ := ioutil.ReadAll(cmdOut)

	var bufs bytes.Buffer
	writer := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
	defer writer.Close()

	writer.Write(outBytes)

	fmt.Println(string(bufs.String()))

	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(string(out))
	filepath.Walk(root, searchDir)
}

func searchDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Print(info.IsDir())
	fmt.Println(path)

	return nil
}
