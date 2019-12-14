package restapi

import (
	"entity"
	_ "fmt"
	"net/http"
	"service"
	"web/framework"
)

type SellInfoPutReq struct {
	Sell *entity.SellInfo
}
type SellInfoPutRes struct {
	Result    bool
	ErrorInfo string
}

var sellInfoPutAPI = &framework.RESTConfig{
	Path:         API_URL_SELLINFO_PUT,
	Method:       http.MethodPost,
	BodyTemplate: &SellInfoPutReq{},
	Callback:     sellInfoPutHandleFn,
}

func sellInfoPutHandleFn(req *framework.RESTRequest) (int, interface{}, error) {

	body := req.Body.(*SellInfoPutReq)
	flag, err := service.AddSellInfo(body.Sell)
	resp := new(SellInfoPutRes)
	resp.Result = flag
	if err != nil {

		resp.ErrorInfo = err.Error()
		return http.StatusOK, resp, nil
	}

	resp.ErrorInfo = ""
	return http.StatusOK, resp, nil
}
