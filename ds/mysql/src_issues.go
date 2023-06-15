package mysql

const srcIssuesTableName = "issues"

type SrcIssue struct {
	TgChatID    int64 `gorm:"column:tlgrm_chat_id"`
	TgMessageID int64 `gorm:"column:tlgrm_message_id"`
	RdmnIssueID int64 `gorm:"column:rdmn_issue_id"`
}

func (SrcIssue) TableName() string {
	return srcIssuesTableName
}

func (m *MySQL) SrcIssuesGet() ([]SrcIssue, error) {

	issues := []SrcIssue{}

	r := m.client.
		Find(&issues)
	if r.Error != nil {
		return nil, r.Error
	}

	return issues, nil
}
