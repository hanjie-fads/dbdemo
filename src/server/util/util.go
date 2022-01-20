package util

import (
	"context"
	"os"

	"github.com/qiniu/qmgo"
	"github.com/sirupsen/logrus"
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
        mgoClient, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://root:example@localhost:27018"})
        mgoDB = mgoClient.Database(os.Getenv("DATABASE_NAME"))
        dbName := os.Getenv("DATABASE_NAME")
        collName := os.Getenv("STAFF_COLL_NAME")
        mgoDB = mgoClient.Database(dbName)
        mgoCollStaff = mgoDB.Collection(collName)
        logrus.Printf("connect mongo at %v/%v", dbName, collName)
    }
}
func GetStaffColl(ctx context.Context) *qmgo.Collection {
    initQmgo(ctx)
	return mgoCollStaff
}