package main

import (
	"context"
	"fmt"
	"kintai/mock"
	"log"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DemoQmgo() {
    var (
        ctx = context.Background()
        // client, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
        // db = client.Database("hjdb")
        // coll = db.Collection("coll-1")
        cli, err = qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017", Database: "hjdb", Coll: "coll-1"})
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

    // 6. query, find
    //joken := model.FindByStaffName{StaffName: "hanjie"}

    // //按照jobName字段进行过滤jobName="job10",翻页参数0-2
    // if cursor, err = cli.Find(context.TODO(), joken, options.Find().SetLimit(2)); err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // // カーソルのDelay閉じる
    // defer func() {
    //     if err = cursor.Close(context.TODO()); err != nil {
    //         log.Fatal(err)
    //     }
    // }()

    // // recur cursor
    // for cursor.Next(context.TODO()) {
    //     var sr model.Staff
    //     // parse Bson to object
    //     if cursor.Decode(&sr) != nil {
    //         fmt.Print(err)
    //         return
    //     }
    //     fmt.Println(sr)
    // }

    // // other recur
    // var results []model.Staff
    // if err = cursor.All(context.TODO(), &results); err != nil {
    //     log.Fatal(err)
    // }
    // for _, rst := range results {
    //     fmt.Println(rst)
    // }
}
