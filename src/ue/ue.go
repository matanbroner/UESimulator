package main

import (
	"github.com/matanbroner/UESimulator/src/app"
	"github.com/matanbroner/UESimulator/src/ue/logger"
	"github.com/matanbroner/UESimulator/src/ue/ue_service"
	"github.com/matanbroner/UESimulator/src/ue/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var UE = &ue_service.UE{}
var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "ue"
	cliApp.Usage = "Usage: --uecfg config yaml file"
	cliApp.Action = action
	cliApp.Flags = UE.GetCliCmd()

	appLog.Infoln(cliApp.Name)
	appLog.Infoln("UE version: ", version.GetVersion())

	if err := cliApp.Run(os.Args); err != nil {
		logger.AppLog.Errorf("UE Run error: %v", err)
	}
}

func action(c *cli.Context) {
	app.AppInitializeWillInitialize(c.String("free5gccfg"))
	UE.Initialize(c)
	UE.Start()
}
