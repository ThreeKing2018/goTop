package ini_config

import (
	"github.com/robfig/config"
	"github.com/ThreeKing2018/k3log"
	"fmt"
)

func GetInstro() {
	//读取文件路径
	cfg, err := config.ReadDefault("config.ini")
	if err != nil {
		k3log.Panic("config read file", err)
	}
	//获取[节点] 节点
	name , err := cfg.String("threeKing", "name")
	if err != nil {
		k3log.Panic("get Name", err)
	}
	fmt.Println(name)
	info , err := cfg.String("threeKing", "info")
	if err != nil {
		k3log.Panic("get Name", err)
	}
	fmt.Println(info)
}
