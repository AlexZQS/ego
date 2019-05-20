package desc

import (
	"ego/src/common"
	"fmt"
)

/**
 * @description
 * @time 2019/5/20 23:16
 * @version
 */

//新增描述
func insertDescDao(t TbItemDesc) int {
	count, err := common.Dml("insert into tb_item_desc values(?,?,?,?)", t.ItemId, t.ItemDesc, t.Created, t.Update)
	if err != nil {
		return -1
	}

	return int(count)
}

//根据主键查询描述
func selByIdDao(id int) (t *TbItemDesc) {
	rows, err := common.Dql("select * from tb_item_desc where item_id=?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var desc = new(TbItemDesc)
	if rows.Next() {
		_ = rows.Scan(desc.ItemId, desc.ItemDesc, desc.Created, desc.Update)
		return desc
	}
	return nil
}
