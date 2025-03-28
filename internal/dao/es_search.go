package dao

import (
	"context"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// QueryInteractionTotal 查询互动总数
func (d *Dao) QueryInteractionTotal(classId string, lessonIndex int, eventName []string) (total int, err error) {
	index := viper.GetString("es.beiboIndex")
	// keyword不走分词，精确匹配。ABC=ABC才行
	var eventNameSli []interface{}
	for _, va := range eventName {
		eventNameSli = append(eventNameSli, va)
	}
	query := elastic.NewBoolQuery().
		Must(
			elastic.NewTermQuery("class_id.keyword", classId),
			elastic.NewTermQuery("lesson_num", lessonIndex),
			elastic.NewTermsQuery("event_name.keyword", eventNameSli...),
		)
	searchRes, err := d.esClient.Search().Index(index).Query(query).Do(context.Background())
	if err != nil {
		log.Error("[ES]QueryInteractionTotal err:", err)
	}
	return int(searchRes.TotalHits()), nil
}
