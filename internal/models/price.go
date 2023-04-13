package models

type Price struct {
	ID           int    `bson:"id"`
	JobType      string `bson:"job_type"`
	NormalPrice  int64  `bson:"normal_price"`
	HolidayPrice int64  `bson:"holiday_price"`
	WeekendPrice int64  `bson:"weekend_price"`
}
