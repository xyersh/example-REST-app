package user

import "context"

type Storage interface {

	//Создание пользователя
	Create(ctx *context.Context, user User) (string, error)

	// поиск пользователя по ID
	FindOne(ctx *context.Context, id string) (User, error)

	// изменение пользователя
	Update(ctx *context.Context, user User) error

	// удаление пользователя
	Delete(ctx *context.Context, id string) error
}
