package model

import "time"

// BeiboRubricRecordData 主观题记录
type BeiboRubricRecordData struct {
	Id                  string     `bson:"_id"`
	Stuid               string     `bson:"stuid"`
	StuName             string     `bson:"stuname"`
	Uuid                string     `bson:"uuid"`
	ItemContentId       string     `bson:"itemContentId"`
	ItemContentName     string     `bson:"itemContentName"`
	StuSubmitImages     string     `bson:"stuSubmitImages"`
	InteractionId       string     `bson:"interactionId"`
	CorrectResult       string     `bson:"correctResult"` // 自主答题结果
	Intro               string     `bson:"intro"`
	QuestionNum         string     `bson:"questionNum"`
	OpenTime            *time.Time `bson:"openTime"`
	ReviseCorrectResult string     `bson:"reviseCorrectResult"` // 老师修正结果
}
