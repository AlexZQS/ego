package user

import "ego/src/common"

/**
 * @description
 * @time 2019/5/15 22:15
 * @version
 */
func LoginService(un, pwd string) (er common.EgoResult) {
	user := SelByUnPwdDao(un, pwd)
	if user != nil {
		er.Status = 200
	} else {
		er.Status = 400
	}
	return
}
