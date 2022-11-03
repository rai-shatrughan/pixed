package util

import (
	"context"

	// "os"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

// KV is wrapper for KV Database
type KV struct {
	cli    gohbase.Client
	conf   *Config
	logger *Logger
	table  string
}

// New returns instance of KV Store
func (kv *KV) New(conf *Config, logger *Logger) {
	kv.conf = conf
	kv.logger = logger
	kv.cli = gohbase.NewClient(conf.GetString("hbase.zookeeper"))
}

// Put upserts data into KV store
func (kv *KV) Put(table, key, value string) error {
	kv.logger.Sugar().Info("Writing : ", key)
	values := map[string]map[string][]byte{"cf": {"a": []byte(value)}}
	putRequest, _ := hrpc.NewPutStr(context.Background(), table, key, values)
	_, err := kv.cli.Put(putRequest)
	return err
}

// Get fetches data from KV store
// func (kv *KV) Get(key string) (string, error) {
// 	start := time.Now()
// 	getResp, err := kv.cli.Get(context.TODO(), key)
// 	if err != nil {
// 		kv.logger.Error("Error Reading KV store ", zap.Error(err))
// 		return "{}", err
// 	}

// 	if getResp.Count >= 1 {
// 		kv.logger.Debug("Successfully got from etcd - ", zap.ByteString("key", getResp.Kvs[0].Key), zap.ByteString("value", getResp.Kvs[0].Value))
// 		elapsed := time.Since(start)
// 		kv.logger.Debug("Get from etcd took", zap.Duration("duration", elapsed))
// 		return string(getResp.Kvs[0].Value), nil
// 	} else {
// 		return "{}", err
// 	}

// }

// // GetFromKey fetches data after a time range
// func (kv *KV) GetFromKey(key string) (string, error) {
// 	start := time.Now()
// 	getResp, err := kv.cli.Get(context.TODO(), key, clientv3.WithFromKey(), clientv3.WithLimit(0))
// 	if err != nil {
// 		kv.logger.Error("Error Reading KV store ", zap.Error(err))
// 		return "{}", err
// 	}

// 	if getResp.Count >= 1 {
// 		elapsed := time.Since(start)
// 		kv.logger.Debug("Get from etcd took", zap.Duration("duration", elapsed))
// 		var rb1 strings.Builder
// 		rb1.WriteString("[")
// 		kv.logger.Debug("Successfully got values for key - ", zap.String("key", key), zap.Int64("count", getResp.Count))
// 		for _, ev := range getResp.Kvs {
// 			rb1.WriteString(string(ev.Value))
// 			rb1.WriteString(",")
// 			kv.logger.Debug("Successfully got value for key - ", zap.ByteString("key", ev.Key), zap.ByteString("value", ev.Value))
// 		}
// 		rb2 := strings.TrimSuffix(rb1.String(), ",")
// 		rb2 = rb2 + "]"
// 		return rb2, nil
// 	} else {
// 		kv.logger.Info("No Value found for key - ", zap.String("key", key))
// 		return "{}", nil
// 	}
// }

// GetFromKeyWithLimit fetches data after a time range with limit
func (kv *KV) GetFromKeyWithLimit(key string, limit int64) (string, error) {
	return "", nil
}
