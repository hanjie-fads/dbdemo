package util

import (
	"context"
	"os"

	"github.com/qiniu/qmgo"
)

var (
    err        error
    mgoClient  *qmgo.Client
    mgoDB      *qmgo.Database
    mgoCollStaff  *qmgo.Collection
)

func initQmgo(ctx context.Context) {
    os.Setenv("DATABASE_NAME", "kintai-db")
    os.Setenv("STAFF_COLL_NAME", "coll-staff")
    if mgoClient == nil {
        mgoClient, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
        mgoDB = mgoClient.Database(os.Getenv("DATABASE_NAME"))
        //mgoDB = mgoClient.Database("aaa")
        mgoCollStaff = mgoDB.Collection(os.Getenv("STAFF_COLL_NAME"))
    }
}
func GetStaffColl(ctx context.Context) *qmgo.Collection {
    initQmgo(ctx)
	return mgoCollStaff
}