package main

import (
	"context"
	"fmt"
	"kintai/mock"
	"kintai/model"
	"log"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DemoQmgo() {
    var (
        ctx = context.Background()
        // client, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
        // db = client.Database("hjdb")
        // coll = db.Collection("coll-1")
        cli, err = qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017", Database: "kintai-db", Coll: "staff"})
        iResult    *qmgo.InsertOneResult
        mResult    *qmgo.InsertManyResult
        //cursor     *qmgo.Cursor
        id         primitive.ObjectID
    )
    defer func() {
        if err = cli.Close(ctx); err != nil {
            panic(err)
        }
    }()

    // make indexes
    cli.EnsureIndexes(ctx, []string{}, []string{"uid", "name,email"})

    // 4. insert one record
    if iResult, err = cli.InsertOne(ctx, mock.OneStaff); err != nil {
        fmt.Print(err)
        return
    }
    //_id: auto unique id
    id = iResult.InsertedID.(primitive.ObjectID)
    fmt.Println("unique ID:", id.Hex())

    // 4. bulk insert
    mResult, err = cli.InsertMany(context.TODO(), mock.MultiStaff)
    if err != nil{
        log.Fatal(err)
    }
    if mResult == nil {
        log.Fatal("result nil")
    }
    for _, v := range mResult.InsertedIDs {
        id = v.(primitive.ObjectID)
        fmt.Println("auto ID", id.Hex())
    }

	// find one document
    one := model.Staff{}
    err = cli.Find(ctx, bson.M{"uid": 15}).One(&one)
    if err != nil {
        fmt.Print(err)
        return
    }
    fmt.Printf("find one: %v", one)
}
