package model

import "time"

// BeiboRubricRecord 主观题答题记录
type BeiboRubricRecord struct {
	Id                  string     `bson:"_id"`
	Cid                 string     `bson:"cid"`             // 班级id
	CNum                string     `bson:"cnum"`            // 班级讲次序号
	TutorId             string     `bson:"tutor_id"`        // 老师id
	Stuid               string     `bson:"stuid"`           // 学生id
	StuName             string     `bson:"stuname"`         // 学生名称
	ItemContentId       string     `bson:"itemContentId"`   // 题目id
	ItemContentName     string     `bson:"itemContentName"` // 题目名称
	QuestionNum         string     `bson:"questionNum"`     // 第几道题
	StuSubmitImages     string     `bson:"stuSubmitImages"` // 学生提交图片 json数组
	InteractionId       string     `bson:"interactionId"`   // 主观题互动id
	CorrectResult       string     `bson:"correctResult"`   // 自主答题结果
	Intro               string     `bson:"intro"`
	OpenTime            *time.Time `bson:"openTime"`
	ReviseCorrectResult string     `bson:"reviseCorrectResult"` // 老师修正结果
}
