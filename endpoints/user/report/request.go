package userreport

import (
	"encoding/json"
)

type UserReportRequest struct {
	ID 		string
}

func (userReportRequest *UserReportRequest) ToString() string {
	bytes, _ := json.Marshal(userReportRequest)
	return string(bytes)
}