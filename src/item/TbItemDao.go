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
			&s, &item.Image, &item.Cid, &item.Status, &item.Create, &item.Update)
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
