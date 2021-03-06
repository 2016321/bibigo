package model

import (
	"database/sql/driver"
	"time"
)

// Archive archive result
type Archive struct {
	ID          int64     `json:"id"`
	AID         int64     `json:"aid"`
	Mid         int64     `json:"mid"`
	TypeID      int16     `json:"typeid"`
	Videos      int       `json:"videos"`
	Title       string    `json:"title"`
	Cover       string    `json:"cover"`
	Content     string    `json:"content"`
	Duration    int       `json:"duration"`
	Attribute   int32     `json:"attribute"`
	Copyright   int8      `json:"copyright"`
	Access      int       `json:"access"`
	PubTime     wocaoTime `json:"pubtime"`
	CTime       wocaoTime `json:"ctime"`
	MTime       wocaoTime `json:"mtime"`
	State       int       `json:"state"`
	MissionID   int64     `json:"mission_id"`
	OrderID     int64     `json:"order_id"`
	RedirectURL string    `json:"redirect_url"`
	Forward     int64     `json:"forward"`
	Dynamic     string    `json:"dynamic"`
}

// Video is
type Video struct {
	AID int64 `json:"aid"`
	CID int64 `json:"cid"`
}

type wocaoTime string

// Scan scan time.
func (jt *wocaoTime) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case time.Time:
		*jt = wocaoTime(sc.Format("2006-01-02 15:04:05"))
	case string:
		*jt = wocaoTime(sc)
	}
	return
}

// Value get time value.
func (jt wocaoTime) Value() (driver.Value, error) {
	return time.Parse("2006-01-02 15:04:05", string(jt))
}
