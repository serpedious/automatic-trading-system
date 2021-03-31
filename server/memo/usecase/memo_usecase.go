package usecase

import (
	"time"

	"github.com/serpedious/automatic-trading-system/server/tool"
)

type Memo struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	Done       bool      `json:"done"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (m *Memo) Validate() string {
	if m.Content == "" {
		msg := "white space is invalid"
		return msg
	}

	return ""
}

func (m *Memo) CreateMemo() error {
	err := m.InsertMemo()
	if err != nil {
		return err
	}
	return nil
}

func (m *Memo) InsertMemo() error {
	sql_query := "INSERT INTO MEMOS(CONTENT, DONE, USER_ID) VALUES($1, $2, $3) RETURNING id;"

	Db := tool.NewDb()
	defer Db.Close()

	err := Db.QueryRow(sql_query, m.Content, true, 1).Scan(&m.ID)
	if err != nil {
		return err
	}
	return nil
}
