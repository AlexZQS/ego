package item

import "ego/src/common"

func showItemService(page, rows int) (e *common.DataGrid) {
	items := selByPageDao(rows, page)
	if items != nil {
		e = new(common.DataGrid)
		e.Row = items //当前页显示的数据
		e.Total = selCount()
		return
	}
	return nil
}
