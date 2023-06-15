package mysql

const srcIDsTableName = "ids"

type SrcID struct {
	TgUserID   int64  `gorm:"column:tlgrm_userid"`
	TgUserName string `gorm:"column:tlgrm_username"`
	RdmnUserID int64  `gorm:"column:rdmn_userid"`
}

func (SrcID) TableName() string {
	return srcIDsTableName
}

func (m *MySQL) SrcIDsGet() ([]SrcID, error) {

	ids := []SrcID{}

	r := m.client.
		Find(&ids)
	if r.Error != nil {
		return nil, r.Error
	}

	return ids, nil
}
