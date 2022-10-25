package factories

import (
	"gohub/app/models/user"
	"gohub/pkg/helpers"

	"github.com/bxcodec/faker/v3"
)

func makeUsers(times int) []user.User {
	var objs []user.User

	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		models := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, models)
	}
	return objs
}
