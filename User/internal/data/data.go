package data

import (
	"User/internal/conf"
	"User/internal/data/ent"
	"User/internal/data/ent/user"
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
)

// ProviderSet is data providers.

type Data struct {
	// TODO wrapped database client
	Db *ent.Client
}

func (d Data) CreateId(id string, name string) error {
	u, err := d.Db.User.Create().SetUserid(id).SetUser(name).Save(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(u)
	return nil
}
func (d Data) SelectName(id string) (string, error) {
	u, err := d.Db.User.Query().Where(user.Userid(id)).Only(context.Background())
	if err != nil {
		return "", err
	}
	return u.User, nil
}

// NewData .
func NewData(c *conf.Config) (*Data, func(), error) {
	drv, err := sql.Open("mysql", c.Mysql)
	if err != nil {
		return nil, func() {}, err
	}
	drive := entsql.OpenDB("mysql", drv)
	db := ent.NewClient(ent.Driver(drive))
	data := Data{Db: db}
	cleanup := func() {
		db.Close()
	}
	if err := db.Schema.Create(context.Background()); err != nil {
		return nil, func() {}, err
	}
	return &data, cleanup, nil
}
