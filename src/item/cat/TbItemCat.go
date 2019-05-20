package cat

//商品类目
type TbItemCat struct {
	Id        int
	ParentId  int
	Name      string
	Status    byte
	SortOrder int8
	IsParent  bool
	Created   string
	Update    string
}

type EasyUITree struct {
	Id    int    `json:"id"`
	Text  string `json:"text"`
	State string `json:"state"`
}
