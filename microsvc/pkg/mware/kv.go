package mware

import (
	"context"

	// "os"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

var (
	dialTimeout = 5 * time.Second
	endpoints   []string
)

// KV is wrapper for KV Database
type KV struct {
	cli  *clientv3.Client
	conf Config
	err  error
}

// New returns singleton instance of KV
func (kv *KV) New() {
	kv.conf.New()
	endpoints = kv.conf.GetStringSlice("etcd.hosts")
	logger.New()

	// clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	kv.cli, kv.err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if kv.err != nil {
		logger.Error("Error while connecting KV store ", zap.Error(kv.err))
	}
}

// Put upserts data into KV store
func (kv *KV) Put(key, value string) {
	start := time.Now()
	_, kv.err = kv.cli.Put(context.TODO(), key, value)
	if kv.err != nil {
		logger.Error("Error writing KV store ", zap.Error(kv.err))
	} else {
		logger.Info("Successfully put to etcd - ", zap.String("key", key), zap.String("value", value))
		elapsed := time.Since(start)
		logger.Info("Put to etcd took", zap.Duration("duration", elapsed))
	}

}

// Get fetches data from KV store
func (kv *KV) Get(key string) string {
	start := time.Now()
	getResp, err := kv.cli.Get(context.TODO(), key)
	if err != nil {
		logger.Error("Error Reading KV store ", zap.Error(kv.err))
	}

	if getResp.Count >= 1 {
		logger.Info("Successfully got from etcd - ", zap.ByteString("key", getResp.Kvs[0].Key), zap.ByteString("value", getResp.Kvs[0].Value))
		elapsed := time.Since(start)
		logger.Info("Get from etcd took", zap.Duration("duration", elapsed))
		return string(getResp.Kvs[0].Value)
	}
	return "{}"

}

// GetFromKey fetches data after a time range
func (kv *KV) GetFromKey(key string) string {
	start := time.Now()
	getResp, err := kv.cli.Get(context.TODO(), key, clientv3.WithFromKey(), clientv3.WithLimit(0))
	if err != nil {
		logger.Error("Error Reading KV store ", zap.Error(kv.err))
	}

	if getResp.Count >= 1 {
		elapsed := time.Since(start)
		logger.Info("Get from etcd took", zap.Duration("duration", elapsed))
		var rb1 strings.Builder
		rb1.WriteString("[")
		logger.Info("Successfully got values for key - ", zap.String("key", key), zap.Int64("count", getResp.Count))
		for _, ev := range getResp.Kvs {
			rb1.WriteString(string(ev.Value))
			rb1.WriteString(",")
			logger.Info("Successfully got value for key - ", zap.ByteString("key", ev.Key), zap.ByteString("value", ev.Value))
		}
		rb2 := strings.TrimSuffix(rb1.String(), ",")
		rb2 = rb2 + "]"
		return rb2
	}
	logger.Info("No Value found for key - ", zap.String("key", key))
	return "{}"

}

// GetFromKeyWithLimit fetches data after a time range with limit
func (kv *KV) GetFromKeyWithLimit(key string, limit int64) string {
	start := time.Now()
	getResp, err := kv.cli.Get(context.TODO(), key, clientv3.WithFromKey(), clientv3.WithLimit(limit))
	if err != nil {
		logger.Error("Error Reading KV store ", zap.Error(kv.err))
	}

	if getResp.Count >= 1 {
		elapsed := time.Since(start)
		logger.Info("Get from etcd took", zap.Duration("duration", elapsed))
		var rb1 strings.Builder
		rb1.WriteString("[")
		logger.Info("Successfully got values for key - ", zap.String("key", key), zap.Int64("count", getResp.Count))
		for _, ev := range getResp.Kvs {
			rb1.WriteString(string(ev.Value))
			rb1.WriteString(",")
			logger.Info("Successfully got value for key - ", zap.ByteString("key", ev.Key), zap.ByteString("value", ev.Value))
		}
		rb2 := strings.TrimSuffix(rb1.String(), ",")
		rb2 = rb2 + "]"
		return rb2
	}
	logger.Info("No Value found for key - ", zap.String("key", key))
	return "{}"

}
