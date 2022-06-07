package persistence

import (
	"context"
	// "fmt"
	// "encoding/json"

	"exercise-go-api/pkg/domain/entity"
	"exercise-go-api/pkg/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	userCollection = "user"
)

type UserRepo struct {
	col *mongo.Collection
}

var _ repository.IUserRepository = (*UserRepo)(nil)

func NewUserRepository(db *mongo.Database) *UserRepo {
	return &UserRepo{
		col: db.Collection(userCollection),
	}
}

func (r UserRepo) ListUsers(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, 0)
	srt := bson.D{
		primitive.E{Key: "id", Value: -1},
	}
	opt := options.Find().SetSort(srt)
	cur, err := r.col.Find(ctx, bson.D{}, opt)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user entity.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// json → Go
	// var usr entity.User
	// jsonString := `{"id":5, "userName":"test"}`

	// if err := json.Unmarshal([]byte(jsonString), &usr); err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// fmt.Printf("%+v\n", usr)
	// users = append(users, usr)

	// Go → json
	// type userStruct struct {
	// 	Id int `json: "id"`
	// 	Name string `json: "name"`
	// }
	// userArray := make([]userStruct, 0)
	// resp := &userStruct{
	// 	Id: 900,
	// 	Name: "userTest",
	// }
	// res, err := json.Marshal(resp)
	// fmt.Println(string(res))

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
