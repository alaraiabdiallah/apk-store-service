package models

type VersionDS struct {
	Id 	 interface{} `bson:"_id" json:"id"`
	Flag string `bson:"flag" json:"flag" form:"flag"`
	Version string `bson:"version" json:"version" form:"version"`
}