package mysql

const dstIssuesBanchTableName = "issues_banch"

type DstIssueBanch struct {
	TgChatID    int64 `gorm:"column:tg_chat_id"`
	TgMessageID int64 `gorm:"column:tg_message_id"`
	RdmnIssueID int64 `gorm:"column:rdmn_issue_id"`
}

type DstIssueBanchInsertData struct {
	TgChatID    int64 `gorm:"column:tg_chat_id"`
	TgMessageID int64 `gorm:"column:tg_message_id"`
	RdmnIssueID int64 `gorm:"column:rdmn_issue_id"`
}

func (DstIssueBanch) TableName() string {
	return dstIssuesBanchTableName
}

func (DstIssueBanchInsertData) TableName() string {
	return dstIssuesBanchTableName
}

func (m *MySQL) DstIssuesBanchSave(issues []DstIssueBanchInsertData) error {

	if len(issues) == 0 {
		return nil
	}

	r := m.client.
		Create(&issues)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (m *MySQL) DstIssuesBanchGet() ([]DstIssueBanch, error) {

	ibs := []DstIssueBanch{}

	r := m.client.
		Find(&ibs)
	if r.Error != nil {
		return nil, r.Error
	}

	return ibs, nil
}

func (m *MySQL) DstIssuesBanchDeleteAll() error {

	r := m.client.
		Where("1 = 1").
		Delete(DstIssueBanch{})
	if r.Error != nil {
		return r.Error
	}

	return nil
}
