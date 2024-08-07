package canonical

import "time"

type User struct {
	Id        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
}
