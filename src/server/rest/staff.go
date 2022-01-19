package rest

import (
	"context"
	"fmt"
	"kintai/mock"
	"kintai/model"
	"kintai/util"
	"log"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitStaff() error {
    var (
        err error
        ctx     = context.Background()
        coll    = util.GetStaffColl(ctx)
        iResult *qmgo.InsertOneResult
        mResult *qmgo.InsertManyResult
        id      primitive.ObjectID
    )

    // make indexes
    coll.EnsureIndexes(ctx, []string{}, []string{"uid", "name,email"})

    // 4. insert one record
    if iResult, err = coll.InsertOne(ctx, mock.OneStaff); err != nil {
        fmt.Print(err)
        return err
    }
    //_id: auto unique id
    id = iResult.InsertedID.(primitive.ObjectID)
    fmt.Println("unique ID:", id.Hex())

    // 4. bulk insert
    mResult, err = coll.InsertMany(context.TODO(), mock.MultiStaff)
    if err != nil{
        log.Fatal(err)
        return err
    }
    if mResult == nil {
        log.Fatal("result nil")
    }
    for _, v := range mResult.InsertedIDs {
        id = v.(primitive.ObjectID)
        fmt.Println("auto ID", id.Hex())
    }
    return nil
}

func CreateStaff(staff model.Staff) error {
    var (
        err error
        ctx     = context.Background()
        coll    = util.GetStaffColl(ctx)
        iResult *qmgo.InsertOneResult
        id      primitive.ObjectID
    )

    // insert one record
    if iResult, err = coll.InsertOne(ctx, mock.OneStaff); err != nil {
        fmt.Print(err)
        return err
    }
    //_id: auto unique id
    id = iResult.InsertedID.(primitive.ObjectID)
    err = coll.Find(ctx, bson.M{"_id": id}).One(&staff)
    if err != nil {
        fmt.Print(err)
        return err
    }

    return nil
}

func UpdateStaff(staff model.Staff) error {
    var (
        err error
        ctx     = context.Background()
        coll    = util.GetStaffColl(ctx)
    )

    // insert one record
    if err = coll.UpdateOne(ctx, bson.M{"uid": staff.Uid}, bson.M{"$set": staff}); err != nil {
        fmt.Print(err)
        return err
    }

    return nil
}

func GetStaffAll() []model.Staff {
    var (
        ctx = context.Background()
        coll = util.GetStaffColl(ctx)
        ary = []model.Staff{}
    )

    // find all 、sort and limit
    coll.Find(ctx, nil).All(&ary)

    return ary
}

func GetStaff(uid int64) model.Staff {
    var (
        err error
        ctx = context.Background()
        coll = util.GetStaffColl(ctx)
        one = model.Staff{}
    )

	// find one document
    err = coll.Find(ctx, bson.M{"uid": uid}).One(&one)
    if err != nil {
        fmt.Print(err)
        return one
    }
    return one
}
func DelStaff(uid int64) error {
    var (
        err error
        ctx = context.Background()
        coll = util.GetStaffColl(ctx)
    )

	// find one document
    err = coll.Remove(ctx, bson.M{"uid": uid})
    if err != nil {
        fmt.Print(err)
        return err
    }
    return nil
}