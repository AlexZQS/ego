package param

import (
	"ego/src/common"
	"fmt"
	"strconv"
)

func selByPageDao(page, rows int) []TbItemParam {
	r, err := common.Dql("select * from tb_item_param limit ?,?", rows*(page-1), rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	t := make([]TbItemParam, 0)
	for r.Next() {
		var param TbItemParam
		_ = r.Scan(&param.Id, &param.ItemCatId, &param.ParamData, &param.Created, &param.Updated)
		t = append(t, param)
	}
	return t
}

//查询总个数
func selCount() int {
	rows, err := common.Dql("select count(*) from tb_item_param")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if rows.Next() {
		var count int
		_ = rows.Scan(&count)

		return count
	}
	return -1
}

//多条删除
func delByIdDao(ids []int) int {
	sql := "delete from tb_item_param where id in ("
	for i := 0; i < len(ids); i++ {
		sql += strconv.Itoa(ids[i])
		if i < len(ids)-1 {
			sql += ","
		}
	}
	sql += ")"
	count, err := common.Dml(sql)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//根据类目Id查询规格参数
func selByCatIdDao(catid int) *TbItemParam {
	rows, err := common.Dql("select * from tb_item_param where item_cat_id=?", catid)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if rows.Next() {
		var tbItemParam = new(TbItemParam)
		_ = rows.Scan(&tbItemParam.Id, &tbItemParam.ItemCatId, &tbItemParam.ParamData, &tbItemParam.Created, &tbItemParam.Updated)
		return tbItemParam
	}

	return nil
}
