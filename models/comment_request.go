package models

type SaveCommentRequest struct {
	FactID  int    `json:"fact_id" example:"1"`
	Name    string `json:"name" example:"John Doe"`
	Comment string `json:"comment" example:"Cats are great pets!"`
}