package cat

import (
	"ego/src/common"
	"fmt"
)

func selByIdDao(id int) (t *TbItemCat) {
	rows, err := common.Dql("select * from tb_item_cat where id=?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if rows.Next() {
		t = new(TbItemCat)
		_ = rows.Scan(&t.Id, &t.ParentId, &t.Name, &t.Status, &t.SortOrder, &t.IsParent, &t.Created, &t.Update)
		return t
	}
	common.CloseConn()
	return
}
