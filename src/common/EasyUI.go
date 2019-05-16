package common

type DataGrid struct {
	//当前页显示的数据
	Row interface{} `json:"rows"`
	//总个数
	Total int `json:"total"`
}
