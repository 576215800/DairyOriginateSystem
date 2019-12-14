package restapi

import (
	"entity"
	_ "fmt"
	"net/http"
	"service"
	"web/framework"
)

type ProcessInfoPutReq struct {
	Process *entity.ProcessInfo
}
type ProcessInfoPutRes struct {
	Result    bool
	ErrorInfo string
}

var processInfoPutAPI = &framework.RESTConfig{
	Path:         API_URL_PROCESSIFNO_PUT,
	Method:       http.MethodPost,
	BodyTemplate: &ProcessInfoPutReq{},
	Callback:     processInfoPutHandleFn,
}

func processInfoPutHandleFn(req *framework.RESTRequest) (int, interface{}, error) {

	body := req.Body.(*ProcessInfoPutReq)
	flag, err := service.AddProcessInfo(body.Process)
	resp := new(ProcessInfoPutRes)
	resp.Result = flag
	if err != nil {

		resp.ErrorInfo = err.Error()
		return http.StatusOK, resp, nil
	}

	resp.ErrorInfo = ""
	return http.StatusOK, resp, nil
}
