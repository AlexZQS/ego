package param

import (
	"ego/src/common"
	"encoding/json"
	"net/http"
	"strconv"
)

func ParamHandler() {
	common.Router.HandleFunc("/item/param/show", showParamController)
	common.Router.HandleFunc("/item/param/delete", delByIdsController)
	common.Router.HandleFunc("/item/param/iscat", isCatController)
}

func isCatController(w http.ResponseWriter, r *http.Request) {
	catId, _ := strconv.Atoi(r.FormValue("catid"))
	e := selByCatIdService(catId)
	bytes, _ := json.Marshal(e)
	w.Header().Set(common.HEADER_CONTENT_TYPE, common.JSON_HEADER)
	w.Write(bytes)
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

//删除规格参数
func delByIdsController(w http.ResponseWriter, r *http.Request) {
	result := delByIdsService(r.FormValue("ids"))
	bytes, _ := json.Marshal(result)
	w.Header().Set(common.HEADER_CONTENT_TYPE, common.JSON_HEADER)
	w.Write(bytes)
}
