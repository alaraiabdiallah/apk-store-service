package models

import "time"

type VersionDS struct {
	Id        interface{} `bson:"_id" json:"id"`
	Flag      string      `bson:"flag" json:"flag" form:"flag"`
	Version   string      `bson:"version" json:"version" form:"version"`
	BuildCode uint64      `bson:"build_code" json:"build_code"`
	CreatedAt time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at" json:"updated_at"`
}
