package mysql

const dstUsersTableName = "users"

type DstUser struct {
	TgUserID   int64  `gorm:"column:tlgrm_userid"`
	RdmnUserID int64  `gorm:"column:rdmn_userid"`
	Lang       string `gorm:"column:lang"`
}

type DstUserInsertData struct {
	TgUserID   int64  `gorm:"column:tlgrm_userid"`
	RdmnUserID int64  `gorm:"column:rdmn_userid"`
	Lang       string `gorm:"column:lang"`
}

func (DstUser) TableName() string {
	return dstUsersTableName
}

func (DstUserInsertData) TableName() string {
	return dstUsersTableName
}

func (m *MySQL) DstUsersSave(users []DstUserInsertData) error {

	if len(users) == 0 {
		return nil
	}

	r := m.client.
		Create(&users)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (m *MySQL) DstUsersGet() ([]DstUser, error) {

	users := []DstUser{}

	r := m.client.
		Find(&users)
	if r.Error != nil {
		return nil, r.Error
	}

	return users, nil
}

func (m *MySQL) DstUsersDeleteAll() error {

	r := m.client.
		Where("1 = 1").
		Delete(DstUser{})
	if r.Error != nil {
		return r.Error
	}

	return nil
}
