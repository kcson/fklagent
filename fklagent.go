package main

import (
	"fasoo.com/fklagent/process"
	"fasoo.com/fklagent/util/log"
	"sync"
)

func main() {
	log.INFO("start processing data!!")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		process.SFTP()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		process.ESB()
	}()

	wg.Wait()

	//root := "D:\\mywork\\go\\src\\fasoo.com"
	//root := "/Users/kcson/mywork/goproject/src/fasoo.com/fklagent"
	//cmd := exec.Command("cmd", "/c", "dir")
	//cmdOut,_ := cmd.StdoutPipe()
	//defer cmdOut.Close()
	//cmd.Start()
	//outBytes,_ := ioutil.ReadAll(cmdOut)
	//
	//var bufs bytes.Buffer
	//writer := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
	//defer writer.Close()
	//
	//writer.Write(outBytes)
	//
	//fmt.Println(string(bufs.String()))

	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(string(out))
	//filepath.Walk(root, searchDir)
}
