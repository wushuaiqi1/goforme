package dao

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"myproject/internal/model"
)

func (d *Dao) ListRubricRecord(ctx context.Context, cId, cNum, stuId string) (rubricReceive []*model.BeiboRubricRecord, err error) {
	query := bson.M{
		"cid":   cId,
		"cnum":  cNum,
		"stuid": stuId,
	}

	err = d.qmgoClient.Find(ctx, query).All(&rubricReceive)
	if err != nil {
		log.WithContext(ctx).Info("ListRubricRecord err:", err)
		return
	}
	return rubricReceive, err
}
