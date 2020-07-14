package models

import (
	"mime/multipart"
	"time"
)

type MediaDS struct {
	Id 	 interface{} `bson:"_id" json:"id"`
	Flag string `bson:"flag" json:"flag"`
	Version string `bson:"version" json:"version"`
	Filename string `bson:"file_name" json:"file_name"`
	Filepath string `bson:"file_path" json:"file_path"`
	Mime string `bson:"mime" json:"mime"`
	CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

type MediaLink struct {
	Url 	 string `json:"url"`
	Flag 	 string `json:"flag"`
	Version  string `json:"version"`
}

type MediaFilter struct {
	OnlyLink bool
	Query interface{}
}

type MediaUploadParams struct {
	Flag string
	Version string
	File *multipart.FileHeader
}