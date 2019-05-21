package param

import (
	"ego/src/common"
	c "ego/src/item/cat"
)

//显示规格参数
func showParamService(page, rows int) (d common.DataGrid) {
	t := selByPageDao(page, rows)
	d.Total = selCount()
	cats := make([]TbItemParamCat, 0)
	for i := 0; i < len(t); i++ {
		var cat TbItemParamCat
		cat.Id = t[i].Id
		cat.Updated = t[i].Updated
		cat.Created = t[i].Created
		cat.ParamData = t[i].ParamData
		cat.ItemCatId = t[i].ItemCatId
		cat.CatName = c.ShowCatByIdService(t[i].ItemCatId).Name
		cats = append(cats, cat)

	}
	d.Row = cats
	return
}
