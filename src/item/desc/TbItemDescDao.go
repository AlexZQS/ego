package desc

import "ego/src/common"

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
