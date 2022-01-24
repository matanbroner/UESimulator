/*
 * UE Configuration Factory
 */

package factory

import (
	"fmt"
	"github.com/matanbroner/UESimulator/src/ue/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var UeConfig Config

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		logger.AppLog.Fatal(err)
	}
}

// TODO: Support configuration update from REST api
func InitConfigFactory(f string) {
	content, err := ioutil.ReadFile(f)
	checkErr(err)

	UeConfig = Config{}

	err = yaml.Unmarshal([]byte(content), &UeConfig)
	checkErr(err)

	logger.InitLog.Infof("Successfully initialize configuration %s", f)
}
