package item

import (
	"ego/src/common"
	"ego/src/item/cat"
	"strings"
)

func showItemService(page, rows int) (e *common.DataGrid) {
	items := selByPageDao(rows, page)
	if items != nil {
		itemChildren := make([]TbItemChild, 0)
		for i := 0; i < len(items); i++ {
			var itemChild TbItemChild
			itemChild.Id = items[i].Id
			itemChild.Barcode = items[i].Barcode
			itemChild.Update = items[i].Update
			itemChild.Create = items[i].Create
			itemChild.Status = items[i].Status
			//itemChild.Cid = items[i].Cid
			//itemChild.Image = items[i].Image
			itemChild.Num = items[i].Num
			itemChild.Price = items[i].Price
			itemChild.Image = items[i].Image
			itemChild.SellPoint = items[i].SellPoint
			itemChild.Title = items[i].Title

			itemChild.CategoryName = cat.ShowCatByIdService(items[i].Cid).Name
			itemChildren = append(itemChildren, itemChild)

		}
		e = new(common.DataGrid)
		e.Row = itemChildren //当前页显示的数据
		e.Total = selCount()
		return
	}
	return nil
}

//删除商品
func delByIdsService(ids string) (e common.EgoResult) {
	count := updStatusByIdsDao(strings.Split(ids, ","), 3)
	if count > 0 {
		e.Status = 200

	}
	return
}
