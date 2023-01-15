package tools

import (
	"github.com/go-ini/ini"
	"log"
)

var Cfg *ini.File

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatal("Fail to Load ‘conf/app.ini’:", err)
	}

}
