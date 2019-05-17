package cat

//根据Id查询类目
func ShowCatByIdService(id int) *TbItemCat {
	return selByIdDao(id)
}
