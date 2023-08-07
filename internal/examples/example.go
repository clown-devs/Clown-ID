// Struct examples for swagger documentation
package examples

type PublicUserExample struct {
	ID                string `json:"id" example:"1"`
	Username          string `json:"username" example:"aboba"`
	Email             string `json:"email" example:"aboba@gmail.com"`
	Password          string `json:"password,omitempty" swaggerignore:"true"`
	EncryptedPassword string `json:"-"`
}

type CreateUserExample struct {
	ID                string `json:"id" example:"1"`
	Username          string `json:"username" example:"aboba"`
	Email             string `json:"email" example:"aboba@gmail.com"`
	Password          string `json:"password" example:"aboba32"`
	EncryptedPassword string `json:"-"`
}
