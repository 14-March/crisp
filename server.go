/*
Package crisp
@Author：14March
@File：server.go
*/
package main

import (
	"fmt"
	"os"
	"path"

	"github.com/hcraM41/crisp/comm/clog"
)

func main() {
	fmt.Println("hello crisp~")

	ex, err := os.Executable()

	if nil != err {
		panic(err)
	}

	clog.Config(path.Dir(ex) + "/log/crisp.log")
	clog.Info("hihihihihi")

}
