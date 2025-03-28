package dao

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"myproject/internal/model"
	"strconv"
)

// ListRubricRecord 获取学生某一讲主观题答题记录
func (d *Dao) ListRubricRecord(ctx context.Context, classId, lessonIndex int, stuId string) (rubricReceive []*model.BeiboRubricRecord, err error) {
	cid := strconv.Itoa(classId)
	cnum := strconv.Itoa(lessonIndex)
	query := bson.M{
		"cid":   cid,
		"cnum":  cnum,
		"stuid": stuId,
	}

	err = d.qmgoClient.Find(ctx, query).All(&rubricReceive)
	if err != nil {
		log.WithContext(ctx).Info("ListRubricRecord err:", err)
		return
	}
	return rubricReceive, err
}
