package dto

type InsertAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
