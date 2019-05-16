package item

import (
	"ego/src/common"
	"encoding/json"
	"net/http"
	"strconv"
)

func ItemHandler() {
	common.Router.HandleFunc("/showItem", ShowItemController)
}

func ShowItemController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	dataGrid := showItemService(page, rows)
	bytes, _ := json.Marshal(dataGrid)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(bytes)
}
