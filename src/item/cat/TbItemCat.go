package cat

//商品类目
type TbItemCat struct {
	Id        int
	ParentId  int
	Name      string
	Status    byte
	SortOrder int8
	IsParent  byte
	Created   string
	Update    string
}
