package mysql

const dstFeedbackIssuesTableName = "feedback_issues"

type DstFeedbackIssue struct {
	TgUserID    int64 `gorm:"column:tlgrm_userid"`
	RdmnIssueID int64 `gorm:"column:rdmn_issue_id"`
}

type DstFeedbackIssueInsertData struct {
	TgUserID    int64 `gorm:"column:tlgrm_userid"`
	RdmnIssueID int64 `gorm:"column:rdmn_issue_id"`
}

func (DstFeedbackIssue) TableName() string {
	return dstFeedbackIssuesTableName
}

func (DstFeedbackIssueInsertData) TableName() string {
	return dstFeedbackIssuesTableName
}

func (m *MySQL) DstFeedbackIssuesSave(issues []DstFeedbackIssueInsertData) error {

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

func (m *MySQL) DstFeedbackIssuesGet() ([]DstFeedbackIssue, error) {

	issues := []DstFeedbackIssue{}

	r := m.client.
		Find(&issues)
	if r.Error != nil {
		return nil, r.Error
	}

	return issues, nil
}

func (m *MySQL) DstFeedbackIssuesDeleteAll() error {

	r := m.client.
		Where("1 = 1").
		Delete(DstFeedbackIssue{})
	if r.Error != nil {
		return r.Error
	}

	return nil
}
