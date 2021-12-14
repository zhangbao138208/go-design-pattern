package model

type PostRequest struct {
	BlogName string `json:"blog_name" xml:"blog_name1" form:"blog_name1" uri:"blog_name3" binding:"required"`
	BlogId   int `json:"blog_id" xml:"blog_id1" form:"blog_id1" uri:"blog_id3" binding:"required"`
}
