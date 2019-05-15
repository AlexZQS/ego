package user

import (
	"ego/src/common"
	"fmt"
)

/**
 * @description
 * @time 2019/5/15 22:06
 * @version
 */
//根据用户和密码查询，如果返回值为nil，表示查询失败，否则成功
func SelByUnPwdDao(userName, password string) *TbUser {
	sql := "select * from tb_user where username =? and password=? or email=? and password = ? or phone = ? and password =?"
	rows, err := common.Dql(sql, userName, password, userName, password, userName, password)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if rows.Next() {
		user := new(TbUser)
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Phone, &user.Email, &user.Create, &user.Update)
		common.CloseConn()
		return user
	}
	return nil
}
