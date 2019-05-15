package user

/**
 * @description
 * @time 2019/5/15 22:03
 * @version
 */
//对应数据库中用户表
type TbUser struct {
	//属性首字母大写 1.要转换给Json 2.可能出现跨包访问
	Id       int64
	UserName string
	Password string
	Phone    string
	Email    string
	Create   string
	Update   string
}
