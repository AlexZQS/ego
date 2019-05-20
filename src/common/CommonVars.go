package common

import "github.com/gorilla/mux"

/**
 * @description
 * @time 2019/5/16 0:16
 * @version
 */
var (
	Router              = mux.NewRouter()
	CurrPath            = "http://localhost:80/"
	HEADER_CONTENT_TYPE = "Content-Type"
	JSON_HEADER         = "application/json;charset=utf-8"
)
