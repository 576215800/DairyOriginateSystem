package restapi

import (
	"entity"
	"fmt"
	"net/http"
	"service"
	"web/framework"
)

type ProduceInfoPutReq struct {
	Produce *entity.ProduceInfo
}
type ProduceInfoPutRes struct {
	Result    bool
	ErrorInfo string
}

var produceInfoPutAPI = &framework.RESTConfig{
	Path:         API_URL_PRODUCEINFO_PUT,
	Method:       http.MethodPost,
	BodyTemplate: &ProduceInfoPutReq{},
	Callback:     produceInfoPutHandleFn,
}

func produceInfoPutHandleFn(req *framework.RESTRequest) (int, interface{}, error) {
	fmt.Println("进入")
	body := req.Body.(*ProduceInfoPutReq)
	flag, err := service.AddProduceInfo(body.Produce)
	resp := new(ProduceInfoPutRes)
	resp.Result = flag
	if err != nil {

		resp.ErrorInfo = err.Error()
		return http.StatusOK, resp, nil
	}

	resp.ErrorInfo = ""
	return http.StatusOK, resp, nil
}
