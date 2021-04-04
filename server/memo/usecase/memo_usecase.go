package usecase

import (
	"log"
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

func (m *Memo) CreateMemo() error {
	err := m.InsertMemo()
	if err != nil {
		return err
	}
	return nil
}

type DoneMemo struct {
	ID   int  `json:"id"`
	Done bool `json:"done"`
}

func (m *DoneMemo) UpdateMemoStatus() error {
	Db := tool.NewDb()
	defer Db.Close()

	sql_query := "SELECT done FROM MEMOS WHERE id = $1;"
	rows, err := Db.Query(sql_query, m.ID)
	if err != nil {
		return err
	}

	for rows.Next() {
		var isDone bool
		if err := rows.Scan(&isDone); err != nil {
			log.Fatal(err)
		}
		if isDone {
			sql_query := "UPDATE memos SET done = false WHERE id = $1;"
			Db.QueryRow(sql_query, m.ID).Scan(&m.ID)
		} else {
			sql_query := "UPDATE memos SET done = true WHERE id = $1;"
			Db.QueryRow(sql_query, m.ID).Scan(&m.ID)
		}
	}
	return nil
}

func (m *DoneMemo) DoneMemo() error {
	err := m.UpdateMemoStatus()
	if err != nil {
		return err
	}
	return nil
}
