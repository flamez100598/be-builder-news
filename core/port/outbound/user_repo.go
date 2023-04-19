package outbound

// import (
// 	"lykafe/news/core/data/model"
// )

type UserRepo interface {
	UpdateUser(id, username, email, name, avatar string) error
}