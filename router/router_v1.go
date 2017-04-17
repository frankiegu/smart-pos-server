package router

import (
	"net/http"
	"github.com/phonechan/smart-pos-server/handler"
)

func RouterV1() (mux *AppV1ServeMux) {

	mux = NewAppV1ServeMux()

	//激活
	mux.HandleFunc("/app/v1/active", handler.Activate)
	//激活码激活
	mux.HandleFunc("/app/v1/activeCode", handler.Activate)
	//终端参数下载
	mux.HandleFunc("/app/v1/loadTermInfo", handler.Activate)
	//更新（查询版本号）
	mux.HandleFunc("/app/v1/findVersion", handler.Activate)
	//根据激活码获取商户名称
	mux.HandleFunc("/app/v1/getMerName", handler.Activate)
	//交易查询
	mux.HandleFunc("/app/v1/findTrans", handler.Activate)
	//交易汇总信息查询
	mux.HandleFunc("/app/v1/findTransSumInfo", handler.Activate)
	//交易结算（根据批次号查询交易明细）
	mux.HandleFunc("/app/v1/getTransSettle", handler.Activate)

	return mux;
}

type AppV1ServeMux struct {
	http.ServeMux
}

func NewAppV1ServeMux() *AppV1ServeMux {
	return &AppV1ServeMux{*http.NewServeMux()}
}

func (mux *AppV1ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}