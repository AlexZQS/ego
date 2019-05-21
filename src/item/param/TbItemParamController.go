package param

import (
	"ego/src/common"
	"encoding/json"
	"net/http"
	"strconv"
)

func ParamHandler() {
	common.Router.HandleFunc("/item/param/show", showParamController)
}

//显示规格参数
func showParamController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	dataGrid := showParamService(page, rows)
	bytes, _ := json.Marshal(dataGrid)
	w.Header().Set(common.HEADER_CONTENT_TYPE, common.JSON_HEADER)
	w.Write(bytes)

}
