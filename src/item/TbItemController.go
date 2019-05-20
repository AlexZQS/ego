package item

import (
	"ego/src/common"
	"encoding/json"
	"net/http"
	"strconv"
)

func ItemHandler() {
	common.Router.HandleFunc("/showItem", ShowItemController)
	common.Router.HandleFunc("/item/delete", delByIdsController)
	common.Router.HandleFunc("/item/instock", instockController)
	common.Router.HandleFunc("/item/offstock", offstockController)
	common.Router.HandleFunc("/item/imageupload", imageUploaderController)
	common.Router.HandleFunc("/item/add", insertController)
}

//商品新增
func insertController(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	result := insertService(request.Form)
	bytes, _ := json.Marshal(result)
	w.Header().Set(common.HEADER_CONTENT_TYPE, common.JSON_HEADER)
	w.Write(bytes)
}

//图片上传
func imageUploaderController(w http.ResponseWriter, request *http.Request) {
	file, header, err := request.FormFile("imgFile")
	if err != nil {
		m := make(map[string]interface{})
		m["error"] = 1
		m["message"] = "接收图片失败"
		b, _ := json.Marshal(m)
		w.Header().Set(common.HEADER_CONTENT_TYPE, common.JSON_HEADER)
		w.Write(b)
		return
	}
	data := imageUploadService(file, header)
	bytes, _ := json.Marshal(data)
	w.Header().Set(common.HEADER_CONTENT_TYPE, common.JSON_HEADER)
	w.Write(bytes)

}

//显示商品信息
func ShowItemController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	dataGrid := showItemService(page, rows)
	bytes, _ := json.Marshal(dataGrid)
	w.Header().Set(common.HEADER_CONTENT_TYPE, common.JSON_HEADER)
	w.Write(bytes)
}

//商品删除
func delByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := delByIdsService(ids)
	bytes, _ := json.Marshal(er)
	w.Header().Set(common.HEADER_CONTENT_TYPE, common.JSON_HEADER)
	w.Write(bytes)
}

//商品上架
func instockController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := instockService(ids)
	bytes, _ := json.Marshal(er)
	w.Header().Set(common.HEADER_CONTENT_TYPE, common.JSON_HEADER)
	w.Write(bytes)
}

//商品下架
func offstockController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := offstockService(ids)
	bytes, _ := json.Marshal(er)
	w.Header().Set(common.HEADER_CONTENT_TYPE, common.JSON_HEADER)
	w.Write(bytes)
}
