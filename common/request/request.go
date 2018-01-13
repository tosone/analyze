package request

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Unknwon/com"
	"github.com/parnurzeal/gorequest"
	"github.com/tosone/worklyzer/common/constants"
	"github.com/tosone/worklyzer/models"
)

// GetSpecialRange 获得特定时间段的信息
func GetSpecialRange(user *models.User) (summaryItems []constants.SummaryItem, err error) {
	startString := user.LastTravel.Format(time.RFC3339)
	endString := time.Now().Format(time.RFC3339)

	var url = "https://wakatime.com/api/v1/users/current/summaries"
	var query = "start=" + startString[:10] + "&end=" + endString[:10]
	var errs []error
	var body string
	var resp gorequest.Response
	var unauthorized constants.Unauthorized
	var summaries constants.Summaries
	var start, end time.Time

	resp, body, errs = gorequest.New().
		Timeout(time.Second*30).
		Get(url).
		Query(query).
		Set("Authorization", "Basic "+com.Base64Encode(user.ApiKey)).
		End()

	defer func() {
		var logger = new(models.Log)
		logger.Url = url
		logger.QueryString = query
		logger.Method = constants.GET
		logger.Type = constants.SummaryRequest
		if err != nil {
			logger.Status = constants.Error
			logger.Msg = err.Error()
			logger.Create()
			return
		}
		logger.Status = constants.Success
		logger.Body = []byte(body)
		logger.Start = start
		logger.End = end
		logger.Create()
	}()

	if len(errs) != 0 {
		err = errs[len(errs)-1]
		return
	}
	if resp == nil {
		err = fmt.Errorf("response is timeout or something error")
	}

	err = json.Unmarshal([]byte(body), &unauthorized)
	if err != nil {
		return
	}
	if unauthorized.Error == constants.UnauthorizedStr {
		err = fmt.Errorf("apikey is error or network error")
		return
	}

	err = json.Unmarshal([]byte(body), &summaries)
	if err != nil {
		return
	}
	summaryItems = summaries.Data
	start, err = time.Parse(time.RFC3339, summaries.Start)
	if err != nil {
		return
	}
	end, err = time.Parse(time.RFC3339, summaries.End)
	if err != nil {
		return
	}

	return
}
