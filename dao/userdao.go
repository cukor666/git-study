package dao

type UserDao struct{}

func (u *UserDao) GetUser(id int) string {
	return "UserDao.GetUser"
}
