package dao

import "myproject/internal/model"

func (d *Dao) GetUser(id int) (u model.User) {
	d.forgeDb.Table(model.User{}.TableName()).
		Where("id=? ", id).
		First(&u)
	return
}
