package param

import (
	"ego/src/common"
	"fmt"
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
