package canonical

import "time"

type User struct {
	Id        string    `bson:"_id"`
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"created_at"`
}
