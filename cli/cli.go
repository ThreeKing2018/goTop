package cli

import "github.com/urfave/cli"

func CreateApp() {
	mApp := cli.NewApp() //实例一个对象
	mApp.Name = "ThreeKing2018"
	mApp.Usage = "来自三个大王"
}