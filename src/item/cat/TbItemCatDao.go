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

//根据parent_id 查询所有子类目
func selByPid(pid int) (c []TbItemCat) {
	rows, err := common.Dql("select * from tb_item_cat where parent_id = ?", pid)
	if err != nil {
		fmt.Println(err)
		return
	}

	c = make([]TbItemCat, 0)
	for rows.Next() {
		var item TbItemCat
		_ = rows.Scan(&item.Id, &item.ParentId, &item.Name,
			&item.Status, &item.SortOrder,
			&item.IsParent,
			&item.Created, &item.Update)
		c = append(c, item)
	}
	common.CloseConn()
	return
}
