package dto

type ContactReq struct {
	Email string
	Phone string
	Name string
	Subject string
	Msg string
  Offset int
  Limit int
}