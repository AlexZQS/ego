package param

import (
	"ego/src/common"
	c "ego/src/item/cat"
	"strconv"
	"strings"
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

//删除规格参数
func delByIdsService(ids string) (e common.EgoResult) {
	split := strings.Split(ids, ",")
	idInt := make([]int, 0)
	for _, e := range split {
		id, _ := strconv.Atoi(e)
		idInt = append(idInt, id)
	}
	count := delByIdDao(idInt)
	if count > 0 {
		e.Status = 200
	}

	return
}

//根据类目id查询规格参数是否已经添加
func selByCatIdService(catId int) (e common.EgoResult) {
	param := selByCatIdDao(catId)
	if param == nil {
		e.Status = 200
	}
	return
}
