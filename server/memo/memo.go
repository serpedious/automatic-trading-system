package memo

import "time"

type Memo struct {
	ID         int       `json:"id"`
	User_id    int       `json:"user_id"`
	Content    string    `json:"content"`
	Done       bool      `json:"done"`
	Delete     bool      `json:"delete"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type DoneMemo struct {
	ID   int  `json:"id"`
	Done bool `json:"done"`
}

type DeleteMemo struct {
	ID     int  `json:"id"`
	Delete bool `json:"delete"`
}

type GetMemo struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
