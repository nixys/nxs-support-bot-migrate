package dsmigrate

import (
	"fmt"

	"github.com/nixys/nxs-support-bot-migrate/ds/mysql"
	"github.com/nixys/nxs-support-bot-migrate/ds/redis"
)

type Settings struct {
	Src SrcSettings
	Dst DstSettings
}

type SrcSettings struct {
	MySQL mysql.MySQL
	Redis redis.Redis
}

type DstSettings struct {
	MySQL mysql.MySQL
}

type Migrate struct {
	s srcCtx
	d dstCtx
}

type srcCtx struct {
	m mysql.MySQL
	r redis.Redis
}

type dstCtx struct {
	m mysql.MySQL
}

func Init(s Settings) Migrate {

	return Migrate{
		s: srcCtx{
			m: s.Src.MySQL,
			r: s.Src.Redis,
		},
		d: dstCtx{
			s.Dst.MySQL,
		},
	}
}

type user struct {
	tgID   int64
	rdmnID int64
}

const langDefault = "en"

func (m *Migrate) Migrate() error {

	if err := m.users(); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	if err := m.issues(); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	if err := m.feedbackIssues(); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	return nil
}

func (m *Migrate) users() error {

	if err := m.d.m.DstUsersDeleteAll(); err != nil {
		return fmt.Errorf("migrate users: %w", err)
	}

	srcUsers, err := m.s.m.SrcIDsGet()
	if err != nil {
		return fmt.Errorf("migrate users: %w", err)
	}

	users := []mysql.DstUserInsertData{}
	for _, usr := range srcUsers {

		// Skip incorrect data
		if usr.TgUserID == 0 || len(usr.TgUserName) == 0 || usr.RdmnUserID == 0 {
			continue
		}

		if len(users)%100 == 0 {
			if err := m.d.m.DstUsersSave(users); err != nil {
				return fmt.Errorf("migrate users: %w", err)
			}
			users = []mysql.DstUserInsertData{}
		}

		users = append(
			users,
			mysql.DstUserInsertData{
				TgUserID:   usr.TgUserID,
				RdmnUserID: usr.RdmnUserID,
				Lang:       langDefault,
			},
		)
	}

	if len(users) > 0 {
		if err := m.d.m.DstUsersSave(users); err != nil {
			return fmt.Errorf("migrate users: %w", err)
		}
	}

	return nil
}

func (m *Migrate) issues() error {

	if err := m.d.m.DstIssuesBanchDeleteAll(); err != nil {
		return fmt.Errorf("migrate issues: %w", err)
	}

	srcIssues, err := m.s.m.SrcIssuesGet()
	if err != nil {
		return fmt.Errorf("migrate issues: %w", err)
	}

	// Delete duplicates (if exist)
	srcIssuesClean := make(map[string]mysql.SrcIssue)
	for _, i := range srcIssues {
		srcIssuesClean[fmt.Sprintf("%d-%d-%d", i.TgChatID, i.TgMessageID, i.RdmnIssueID)] = i
	}

	diss := []mysql.DstIssueBanchInsertData{}
	for _, i := range srcIssuesClean {

		// Skip incorrect data
		if i.TgChatID == 0 || i.TgMessageID == 0 || i.RdmnIssueID == 0 {
			continue
		}

		if len(diss)%100 == 0 {
			if err := m.d.m.DstIssuesBanchSave(diss); err != nil {
				return fmt.Errorf("migrate issues: %w", err)
			}
			diss = []mysql.DstIssueBanchInsertData{}
		}

		diss = append(
			diss,
			mysql.DstIssueBanchInsertData{
				TgChatID:    i.TgChatID,
				TgMessageID: i.TgMessageID,
				RdmnIssueID: i.RdmnIssueID,
			},
		)
	}

	if len(diss) > 0 {
		if err := m.d.m.DstIssuesBanchSave(diss); err != nil {
			return fmt.Errorf("migrate issues: %w", err)
		}
	}

	return nil
}

func (m *Migrate) feedbackIssues() error {

	if err := m.d.m.DstFeedbackIssuesDeleteAll(); err != nil {
		return fmt.Errorf("migrate feedback issues: %w", err)
	}

	iss, err := m.s.r.SrcPresalesGet()
	if err != nil {
		return fmt.Errorf("migrate feedback issues: %w", err)
	}

	issues := []mysql.DstFeedbackIssueInsertData{}
	for k, v := range iss {

		// Skip incorrect data
		if k == 0 || v == 0 {
			continue
		}

		if len(issues)%100 == 0 {
			if err := m.d.m.DstFeedbackIssuesSave(issues); err != nil {
				return fmt.Errorf("migrate feedback issues: %w", err)
			}
			issues = []mysql.DstFeedbackIssueInsertData{}
		}

		issues = append(
			issues,
			mysql.DstFeedbackIssueInsertData{
				TgUserID:    k,
				RdmnIssueID: v,
			},
		)
	}

	if len(issues) > 0 {
		if err := m.d.m.DstFeedbackIssuesSave(issues); err != nil {
			return fmt.Errorf("migrate feedback issues: %w", err)
		}
	}

	return nil
}
