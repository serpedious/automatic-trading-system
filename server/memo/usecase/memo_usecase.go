package usecase

import (
	"log"

	"github.com/serpedious/automatic-trading-system/server/memo"
	"github.com/serpedious/automatic-trading-system/server/tool"
)

type Memo memo.Memo
type DoneMemo memo.DoneMemo
type DeleteMemo memo.DeleteMemo
type GetMemo memo.GetMemo

func (m *Memo) Validate() string {
	if m.Content == "" {
		msg := "white space is invalid"
		return msg
	}

	return ""
}

func (m *Memo) InsertMemo(userid int) error {
	sql_query := "INSERT INTO MEMOS(USER_ID, CONTENT, DONE, DELETE) VALUES($1, $2, $3, $4) RETURNING id;"

	Db := tool.NewDb()
	defer Db.Close()

	err := Db.QueryRow(sql_query, userid, m.Content, false, false).Scan(&m.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *Memo) CreateMemo(userid int) error {
	err := m.InsertMemo(userid)
	if err != nil {
		return err
	}
	return nil
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

func (m *DeleteMemo) UpdateDelete() error {
	Db := tool.NewDb()
	defer Db.Close()

	sql_query := "SELECT delete FROM MEMOS WHERE id = $1;"
	rows, err := Db.Query(sql_query, m.ID)
	if err != nil {
		return err
	}

	for rows.Next() {
		var isDelete bool
		if err := rows.Scan(&isDelete); err != nil {
			log.Fatal(err)
		}
		sql_query := "UPDATE memos SET delete = true WHERE id = $1;"
		Db.QueryRow(sql_query, m.ID).Scan(&m.ID)
	}
	return nil
}

func (m *DeleteMemo) DeleteMemo() error {
	err := m.UpdateDelete()
	if err != nil {
		return err
	}
	return nil
}

func (m *GetMemo) GetAll(userid int) ([]GetMemo, error) {
	Db := tool.NewDb()
	defer Db.Close()

	sql_query := "SELECT id, content, done, delete FROM MEMOS WHERE delete = false and user_id = $1;"
	rows, err := Db.Query(sql_query, userid)
	if err != nil {
		return nil, err
	}
	var memos []GetMemo

	for rows.Next() {
		var memo GetMemo
		if err := rows.Scan(&memo.ID, &memo.Content, &memo.Done, &memo.Delete); err != nil {
			log.Fatal(err)
		}
		memos = append(memos, GetMemo{ID: memo.ID, Content: memo.Content, Done: memo.Done, Delete: memo.Delete})
	}

	return memos, nil
}

func (m *GetMemo) GetAllMemo(userid int) ([]GetMemo, error) {
	memos, err := m.GetAll(userid)
	if err != nil {
		return nil, err
	}
	return memos, nil
}
