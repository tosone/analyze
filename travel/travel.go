package travel

import (
	"encoding/json"
	"time"

	"github.com/spf13/viper"
	"github.com/tosone/worklyzer/common/constants"
	"github.com/tosone/worklyzer/common/request"
	"github.com/tosone/worklyzer/logging"
	"github.com/tosone/worklyzer/models"
)

// Travel 遍历所有的人物
func Travel() {
	var err error
	var users []*models.User
	for {
		users, err = new(models.User).FindAll()
		if err != nil {
			logging.Error(err.Error())
			return
		}
		for _, user := range users {
			if time.Since(user.LastTravel).Hours() > float64(viper.GetInt("Setting.Travel")) {
				var summaryItems []constants.SummaryItem
				summaryItems, err = request.GetSpecialRange(user)
				if err != nil {
					break
				}
				for _, item := range summaryItems {
					var summary = new(models.SummaryItem)
					summary.UserID = user.ID
					if summary.Start, err = time.Parse(time.RFC3339, item.Range.Start); err != nil {
						logging.Error(err)
						break
					}
					if summary.End, err = time.Parse(time.RFC3339, item.Range.End); err != nil {
						logging.Error(err)
						break
					}
					if summary.FindByStarter() {
						logging.Error("Record already exist.")
						break
					}
					if summary.Data, err = json.Marshal(item); err != nil {
						logging.Error(err)
						break
					}
					if err = summary.Create(); err != nil {
						logging.Error(err)
						break
					}
				}
			}
		}
		<-time.After(time.Minute * 10)
	}
}
