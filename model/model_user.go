package model

type User struct {
	UID       string `bson:"uid" json:"uid"`
	Email     string `bson:"email" json:"email"`
	CreatedAt int64  `bson:"createdAt" json:"createdAt"`

	Password string `bson:"password" json:"-"`
}
