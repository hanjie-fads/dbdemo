package api

import (
	"context"
	"fmt"
	"kintai/model"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

func GetStaffAll() []model.Staff {
    var (
        ctx = context.Background()
        cli, err = qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017", Database: "kintai-db", Coll: "staff"})
        ary = []model.Staff{}
    )
    defer func() {
        if err = cli.Close(ctx); err != nil {
            panic(err)
        }
    }()

    // find all 、sort and limit
    cli.Find(ctx, nil).All(&ary)

    return ary
}

func GetStaff(uid int64) model.Staff {
    var (
        ctx = context.Background()
        cli, err = qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017", Database: "kintai-db", Coll: "staff"})
        one = model.Staff{}
    )
    defer func() {
        if err = cli.Close(ctx); err != nil {
            panic(err)
        }
    }()

	// find one document
    err = cli.Find(ctx, bson.M{"uid": uid}).One(&one)
    if err != nil {
        fmt.Print(err)
        return one
    }
    return one
}