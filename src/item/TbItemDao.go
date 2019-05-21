package item

import (
	"database/sql"
	"ego/src/common"
	"fmt"
)

/**
rows : 每页的条数
page : 当前第几页
*/
func selByPageDao(rows, page int) []TbItem {
	//第一个占位表示:从哪条开始查询,0 算起； 第二个: 查询几个
	dql, err := common.Dql("select * from tb_item limit ?,?", rows*(page-1), rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	items := make([]TbItem, 0)

	for dql.Next() {
		var item TbItem
		var s sql.NullString

		//如果直接使用item.barcode由于数据库中列为NULL，导致填充错误
		_ = dql.Scan(&item.Id, &item.Title, &item.SellPoint, &item.Price, &item.Num,
			&s, &item.Image, &item.Cid, &item.Status, &item.Created, &item.Updated)
		item.Barcode = s.String
		items = append(items, item)
	}
	common.CloseConn()
	return items
}

/**
	查询总条数
    如果返回值为< 0 表示查询失败
*/
func selCount() (count int) {
	rows, err := common.Dql("select count(*) from tb_item")
	if err != nil {
		fmt.Println(err)
		return -1
	}

	rows.Next()
	_ = rows.Scan(&count)
	common.CloseConn()

	return
}

/**
  返回值如果小于0表示革新失败
*/
func updStatusByIdsDao(ids []string, status int) int {
	if len(ids) <= 0 {
		return -1
	}

	sql := "update tb_item set status = ? where"
	for i, id := range ids {
		sql += " id= " + id
		if i < len(ids)-1 {
			sql += " or "
		}
	}
	count, err := common.Dml(sql, status)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//商品新增
func insertItemDao(t TbItem) int {

	count, err := common.Dml("insert into tb_item values(?,?,?,?,?,?,?,?,?,?,?)",
		t.Id, t.Title, t.SellPoint, t.Price, t.Num, t.Barcode, t.Image, t.Cid, t.Status, t.Created, t.Updated)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	return int(count)
}

//根据Id 删除
func delById(id int64) int {
	count, err := common.Dml("delete from tb_item where id=?", id)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//根据主键查询内容
func selByIdDao(id int) (t *TbItem) {
	rows, err := common.Dql("select * from tb_item where id=?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if rows.Next() {
		t := new(TbItem)
		var s sql.NullString
		_ = rows.Scan(&t.Id, &t.Title, &t.SellPoint, &t.Price, &t.Num,
			&s, &t.Image, &t.Cid, &t.Status, &t.Created, &t.Updated)
		t.Barcode = s.String
		return t
	}
	return nil
}

//修改商品表数据
func updateItemByIdWithTx(t TbItem) int {
	return common.PrepareWithTx("update tb_item set title =?,"+
		"sell_point=?,price=?,num=?,image=?,cid=?,status=?,updated=? where id=?",
		t.Title, t.SellPoint, t.Price, t.Num, t.Image, t.Cid, t.Status, t.Updated, t.Id)
}
