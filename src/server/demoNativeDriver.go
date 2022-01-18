package main

import (
	"context"
	"fmt"
	"kintai/mock"
	"kintai/model"
	"kintai/util"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DemoNativeDrive() {
    var (
        // ctx = context.Background()
        // 1.init mgo
        client = util.GetMgoCli()
        err        error
        db         *mongo.Database
        collection *mongo.Collection
        iResult    *mongo.InsertOneResult
        mResult    *mongo.InsertManyResult
        cursor     *mongo.Cursor
        id         primitive.ObjectID
    )
    // 2. 使用するDBを選択
    db = client.Database("demo-db")

    // 3. collectionを選択する
    collection = db.Collection("coll-1")


    // 4. insert one record
    if iResult, err = collection.InsertOne(context.TODO(), mock.OneStaff); err != nil {
        fmt.Print(err)
        return
    }
    //_id: auto unique id
    id = iResult.InsertedID.(primitive.ObjectID)
    fmt.Println("unique ID:", id.Hex())

    // 4. bulk insert
    mResult, err = collection.InsertMany(context.TODO(), mock.MultiStaff)
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
    joken := model.FindByStaffName{StaffName: "hanjie"}

    //按照jobName字段进行过滤jobName="job10",翻页参数0-2
    if cursor, err = collection.Find(context.TODO(), joken, options.Find().SetLimit(2)); err != nil {
        fmt.Println(err)
        return
    }
    // カーソルのDelay閉じる
    defer func() {
        if err = cursor.Close(context.TODO()); err != nil {
            log.Fatal(err)
        }
    }()

    // recur cursor
    for cursor.Next(context.TODO()) {
        var sr model.Staff
        // parse Bson to object
        if cursor.Decode(&sr) != nil {
            fmt.Print(err)
            return
        }
        fmt.Println(sr)
    }

    // other recur
    var results []model.Staff
    if err = cursor.All(context.TODO(), &results); err != nil {
        log.Fatal(err)
    }
    for _, rst := range results {
        fmt.Println(rst)
    }
}
