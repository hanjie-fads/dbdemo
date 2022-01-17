package model

type Staff struct {
	Uid    int64  `bson:"uid"` // staff id
	Name   string `bson:"name"`   // staff name
	Email  string `bson:email`

	StartTime   int64 `bson:"startTime"` // start time
	EndTime     int64 `bson:"endTime"`   // end time
	TimeUnit	int32 `bson:"timeUnit"`

	CommuteFee  int32 `bson:"commuteFee"`
	Address		Address `bson:"address"`
}
type Address struct {
	Addr1  string `bson:addr1`
	Addr2  string `bson:addr2`
	Addr3  string `bson:addr3`
}
type TimeRange struct {
   StartTime int64 `bson:"startTime"` // start time
   EndTime   int64 `bson:"endTime"`   // end time
   RestTime	 int32 `bson:"restTime"`   // rest time
}

// query entity
type FindByStaffName struct {
	StaffName string `bson:"staffName"` // Staff任务名
 }