package users

import "xorm.io/xorm"

// Index -  gets a slice of users based on the passed
func Index(db *xorm.Engine, findBy User) (users Users, err error) {
	sess := db.Where("1 = 1")
	if len(findBy.Email) > 0 {
		sess = sess.Where("email = ?", findBy.Email)
	}
	if err = sess.Find(&users); err != nil {
		return
	}
	return
}
