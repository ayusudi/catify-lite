package models

type SaveFactRequest struct {
	Fact string `json:"fact" example:"Cats have five toes on their front paws, but only four on the back ones."`
}