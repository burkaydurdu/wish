package couchbase

import (
	"time"

	"github.com/burkaydurdu/wish/config"
	"github.com/couchbase/gocb/v2"
	"go.uber.org/zap"
)

const (
	bucketName         = "wish"
	scopeName          = "wish"
	wishCollectionName = "wish"
)

type Client struct {
	Cluster       *gocb.Cluster
	conf          *config.DatabaseConfig
	globalTimeout time.Duration
}

func NewCouchbaseClient(conf *config.DatabaseConfig) (client *Client, err error) {
	cluster, err := gocb.Connect(
		conf.Host,
		gocb.ClusterOptions{
			Username: conf.Username,
			Password: conf.Password,
		})
	if err != nil {
		return
	}
	duration := 15
	client = &Client{Cluster: cluster, conf: conf, globalTimeout: time.Duration(duration) * time.Second}

	return
}

func CreateBucket(client *Client) *gocb.Bucket {
	return client.Cluster.Bucket(bucketName)
}

func CreateScope(bucket *gocb.Bucket) *gocb.Scope {
	return bucket.Scope(scopeName)
}

type NewCouchbaseRepositoryOpts struct {
	Client     *Client
	Bucket     *gocb.Bucket
	Scope      *gocb.Scope
	Collection *gocb.Collection
	Logger     *zap.Logger
}

type repository struct {
	logger     *zap.Logger
	collection *gocb.Collection
	scope      *gocb.Scope
	client     *Client
}
