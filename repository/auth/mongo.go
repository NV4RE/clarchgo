package auth

import (
	"clarchgo/entity/auth"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type userMongoDB struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Username    string             `bson:"username"`
	Credentials *auth.Credentials  `bson:"-"`
	Roles       []auth.Role        `bson:"roles"`
	IsActivated bool               `bson:"is_activated"`
	CreatedAt   time.Time          `bson:"created_at"`
	DeletedAt   time.Time          `bson:"deleted_at"`
	IsDeleted   bool               `bson:"is_deleted"`
}

type tokenMongoDB struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Token     string             `bson:"token"`
	ExpiresBy time.Time          `bson:"expires_by"`
}

func (u userMongoDB) toUser() auth.User {
	return auth.User{
		Username:    u.Username,
		Credentials: u.Credentials,
		Roles:       u.Roles,
		IsActivated: u.IsActivated,
		CreatedAt:   u.CreatedAt,
	}
}

type mongoDB struct {
	userCol  *mongo.Collection
	tokenCol *mongo.Collection
}

func (m mongoDB) CreateUser(ctx context.Context, user auth.User) error {
	u := newUserMongoDB(user)

	_, err := m.userCol.InsertOne(ctx, u)
	if mongo.IsDuplicateKeyError(err) {
		return ErrUserAlreadyExist
	}

	return err
}

func (m mongoDB) GetUserByUsername(ctx context.Context, username string) (auth.User, error) {
	var u userMongoDB

	doc := m.userCol.FindOne(
		ctx,
		bson.M{
			"username":   username,
			"is_deleted": false,
		},
	)

	err := doc.Decode(&u)
	if err != nil {
		return auth.User{}, err
	}

	return u.toUser(), nil
}

func (m mongoDB) DeleteUserByUsername(ctx context.Context, username string) error {
	res, err := m.userCol.UpdateOne(
		ctx,
		bson.M{
			"username":   username,
			"is_deleted": false,
		},
		bson.M{
			"$set": bson.M{
				"is_deleted": true,
				"deleted_at": time.Now(),
			},
		},
	)

	if res.ModifiedCount != 1 || err == mongo.ErrNoDocuments {
		return ErrUserNotFound
	}

	return err
}

func (m mongoDB) ListUsers(ctx context.Context) ([]auth.User, error) {
	return m.listUserByQuery(ctx, bson.M{"is_deleted": false})
}

func (m mongoDB) ListUsersByIsActivated(ctx context.Context, isActivated bool) ([]auth.User, error) {
	return m.listUserByQuery(ctx, bson.M{"is_deleted": false, "is_activated": isActivated})
}

func (m mongoDB) CreateToken(ctx context.Context, token string, user auth.User, expiresBy time.Time) error {
	_, err := m.tokenCol.InsertOne(ctx, tokenMongoDB{
		Token:     token,
		Username:  user.Username,
		ExpiresBy: expiresBy,
	})

	if mongo.IsDuplicateKeyError(err) {
		return ErrTokenCannotBeCreated
	}

	return err
}

func (m mongoDB) ActivateUser(ctx context.Context, username string) error {
	res, err := m.userCol.UpdateOne(
		ctx,
		bson.M{
			"username":     username,
			"is_deleted":   false,
			"is_activated": false,
		},
		bson.M{
			"$set": bson.M{
				"is_activated": true,
			},
		},
	)

	if res.ModifiedCount != 1 || err == mongo.ErrNoDocuments {
		return ErrUserNotFound
	}

	return err
}

func (m mongoDB) GetUserByToken(ctx context.Context, token string) (auth.User, error) {
	var t tokenMongoDB

	doc := m.tokenCol.FindOne(
		ctx,
		bson.M{
			"token": token,
			"expires_by": bson.M{
				"$gte": time.Now(),
			},
		},
	)

	err := doc.Decode(&t)
	if err != nil {
		return auth.User{}, err
	}

	return m.GetUserByUsername(ctx, t.Username)
}

func (m mongoDB) listUserByQuery(ctx context.Context, query interface{}) ([]auth.User, error) {
	var us []auth.User
	cur, err := m.userCol.Find(ctx, query)
	if err != nil {
		return us, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var u userMongoDB
		if err = cur.Decode(&u); err != nil {
			return us, err
		}

		us = append(us, u.toUser())
	}

	return us, nil
}

func NewMongo(mongoUri, database string) (Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	// Check connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &mongoDB{
		userCol:  client.Database(database).Collection("user"),
		tokenCol: client.Database(database).Collection("token"),
	}, nil
}

func newUserMongoDB(u auth.User) userMongoDB {
	return userMongoDB{
		Username:    u.Username,
		Credentials: u.Credentials,
		Roles:       u.Roles,
		IsActivated: u.IsActivated,
		CreatedAt:   u.CreatedAt,
		IsDeleted:   false,
	}
}
