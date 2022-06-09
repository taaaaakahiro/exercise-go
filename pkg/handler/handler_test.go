package handler

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	// 複数の構造体を一つのtypeの中で宣言
	RespUser struct {
		Id   int    `json:"id" bson:"_id"`
		Name string `json:"name" bson:"_name"`
	}

	RespMessage struct {
		Id   int
		Name string
	}
)

// unit test
func Check(a, b int) int {
	res := a + b
	return res
}

func TestCheck(t *testing.T) {
	t.Run("test handler Check!!!!!!", func(t *testing.T) {
		// seeds := []interface{}{
		// 	User{Id: 0, Name: "test"},
		// 	User{Id: 0, Name: "test"},
		// 	User{Id: 0, Name: "test"},
		// }

		// connect to mongoDB
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// 関数を抜けたらクローズするようにする
		defer cancel()
		// 指定したURIに接続する
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			fmt.Println("connect to monogo", err)
		}
		defer client.Disconnect(ctx)
		// DBにPingする
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			fmt.Println("connection error:", err)
		} else {
			fmt.Println("connection success:")
		}
		fmt.Println("Successfully connected and pinged.")

		res := Check(3, 9)
		if res != 12 {
			assert.Equal(t, res, 12)
		}
	})
}
