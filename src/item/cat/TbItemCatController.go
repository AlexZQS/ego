package cat

import (
	"ego/src/common"
	"encoding/json"
	"net/http"
	"strconv"
)

func ItemCatHandle() {
	common.Router.HandleFunc("/item/cat/show", showItemCatController)
}

func showItemCatController(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		id = "0"
	}
	i, _ := strconv.Atoi(id)
	tree := showCatByPidService(i)
	bytes, _ := json.Marshal(tree)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(bytes)
}
