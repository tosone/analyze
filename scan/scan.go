package scan

import (
	"time"

	"github.com/spf13/viper"
	"github.com/tosone/worklyzer/logging"
	"github.com/tosone/worklyzer/models"
)

func Run() {
	var err error
	for _, m := range viper.Get("UserMgr").([]interface{}) {
		var user = new(models.User)
		user.Name = m.(map[interface{}]interface{})["Name"].(string)
		user.ApiKey = m.(map[interface{}]interface{})["ApiKey"].(string)
		user.LastTravel = time.Now()
		_, err = user.FindByName()
		if err != nil {
			err = user.Create()
			if err != nil {
				logging.Error(err.Error())
			}
		}
	}
}
