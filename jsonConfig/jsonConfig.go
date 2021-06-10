package jsonConfig

import (
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/spf13/pflag"
)

func ParseConfigFile() []byte{
	var fileLocker sync.Mutex // file locker
	var jsonPath = pflag.StringP("config.army.model.json	-path", "p", "", "Input JsonPath")
	// 把用户传递的命令行参数解析为对应变量的值
	pflag.Parse()
	filename := *jsonPath
	fileLocker.Lock()
	dataJson, err := ioutil.ReadFile(filename) //读取json文件
	fileLocker.Unlock()
	if err != nil {
		fmt.Println("read file error")
		return nil
	}
	return dataJson
}