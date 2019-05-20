package desc

/**
 * @description
 * @time 2019/5/20 23:20
 * @version
 */

//新增
func Insert(t TbItemDesc) int {
	return insertDescDao(t)
}

//根据主键查询描述
func SelByIdService(id int) *TbItemDesc {
	return selByIdDao(id)
}
