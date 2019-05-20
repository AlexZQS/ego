package cat

//根据Id查询类目
func ShowCatByIdService(id int) *TbItemCat {
	return selByIdDao(id)
}

func showCatByPidService(pid int) (tree []EasyUITree) {
	cats := selByPid(pid)
	tree = make([]EasyUITree, 0)
	for _, n := range cats {

		state := "open"
		if n.IsParent {
			state = "closed"
		}

		tree = append(tree, EasyUITree{
			n.Id,
			n.Name,
			state,
		})
	}
	return
}
