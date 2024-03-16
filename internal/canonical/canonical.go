package canonical

type User struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}
