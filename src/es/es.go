package es

import (
	"github.com/olivere/elastic"
	"config"
	"logger"
	"context"
	"net/http"
)

var es *elastic.Client

func Init() {
	var err error

	// Connect to ElasticSearch
	es, err = connectToElasticSearch()
	if err != nil {
		logger.Log.Fatal("Connection to Elasticsearch failed with error: " + err.Error())
	}

}

func connectToElasticSearch() (es *elastic.Client, err error) {
	conf := config.GetConfig()
	esHost := conf.ElasticSearch.Host
	logger.Log.Infof("Connecting to Local Elasticsearch: %s", esHost)
	es, err = elastic.NewClient(
		elastic.SetURL(esHost),
		elastic.SetSniff(false),
		elastic.SetMaxRetries(3),
	)
	if err != nil {
		return nil, err
	}
	defer es.Stop()
	return es, nil
}

// GetESClient for getting Elasticsearch client
func GetESClient() *elastic.Client {
	return es
}

// GetESStatus returns ElasticSearch cluster health
func GetESStatus() (*elastic.ClusterHealthResponse, int) {
	response, _ := es.ClusterHealth().Do(context.TODO())
	statusCode := http.StatusOK
	if response == nil || response.TimedOut {
		statusCode = http.StatusServiceUnavailable
	}
	return response, statusCode
}
