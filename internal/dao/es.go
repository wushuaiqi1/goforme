package dao

import (
	"context"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// TODO Es客户端支持多索引查询

type EsConfig struct {
	Host       string
	UserName   string
	Password   string
	BeiBoIndex string
}

func NewEsClient() (esClient *elastic.Client, err error) {
	esConfig := &EsConfig{
		Host:       viper.GetString("es.host"),
		UserName:   viper.GetString("es.username"),
		Password:   viper.GetString("es.password"),
		BeiBoIndex: viper.GetString("es.beiboIndex"),
	}
	log.Info("NewEsClient esConfig:", *esConfig)

	esClient, err = elastic.NewClient(
		elastic.SetURL(esConfig.Host),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(esConfig.UserName, esConfig.Password),
	)

	if err != nil {
		log.Errorf("NewEsClient NewClient err:%s", err)
		return nil, err
	}

	_, _, err = esClient.Ping(esConfig.Host).Do(context.Background())
	if err != nil {
		log.Errorf("NewEsClient Ping err:%s", err)
	}
	log.Info("NewEsClient Ping Success")

	return esClient, err
}

func EsClose(client *elastic.Client) {
	client.Stop()
}
