package restapi

import (
	"entity"
	_ "fmt"
	"net/http"
	"service"
	"web/framework"
)

type LogisticsInfoPutReq struct {
	Logistics *entity.LogisticsInfo
}
type LogisticsInfoPutRes struct {
	Result    bool
	ErrorInfo string
}

var logisticsInfoPutAPI = &framework.RESTConfig{
	Path:         API_URL_LOGISTICSINFO_PUT,
	Method:       http.MethodPost,
	BodyTemplate: &LogisticsInfoPutReq{},
	Callback:     logisticsInfoPutHandleFn,
}

func logisticsInfoPutHandleFn(req *framework.RESTRequest) (int, interface{}, error) {

	body := req.Body.(*LogisticsInfoPutReq)
	flag, err := service.AddLogisticsInfo(body.Logistics)
	resp := new(LogisticsInfoPutRes)
	resp.Result = flag
	if err != nil {

		resp.ErrorInfo = err.Error()
		return http.StatusOK, resp, nil
	}

	resp.ErrorInfo = ""
	return http.StatusOK, resp, nil
}
