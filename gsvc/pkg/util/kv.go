package util

import (
	"context"
	"strings"

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
	kv.table = conf.GetString("hbase.table")
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
func (kv *KV) Get(key string) (string, error) {
	var rb1 strings.Builder
	rb1.WriteString("[")
	hGet, err := hrpc.NewGetStr(context.Background(), kv.table, key)
	if err != nil {
		kv.logger.Sugar().Errorf("error getRequest : %s ", err)
	}
	result, err := kv.cli.Get(hGet)
	if err != nil {
		kv.logger.Sugar().Errorf("error querying : %s - %s - %s", kv.table, key, err)
		return "{}", nil
	}

	if result != nil {
		for _, v := range result.Cells {
			rb1.WriteString(string(v.Value))
			rb1.WriteString(",")
			kv.logger.Sugar().Infof("response : %s ", v.String())
		}
	}

	rb2 := strings.TrimSuffix(rb1.String(), ",")
	rb2 = rb2 + "]"
	kv.logger.Sugar().Infof("final response : %s ", rb2)
	return rb2, nil
}

// // GetFromKey fetches data after a time range
func (kv *KV) GetKeyDefLimit(key string) (string, error) {
	var rb1 strings.Builder
	rb1.WriteString("[")
	scanRequest, err := hrpc.NewScanRangeStr(context.Background(), kv.table,
		key, "", hrpc.NumberOfRows(100))
	if err != nil {
		kv.logger.Sugar().Errorf("error scanRequest : %s ", err)
	}
	scanner := kv.cli.Scan(scanRequest)
	kv.logger.Sugar().Infof("into query : %s / %s ", kv.table, key)
	for {
		r, err := scanner.Next()
		if r != nil {
			for _, v := range r.Cells {
				rb1.WriteString(string(v.Value))
				rb1.WriteString(",")
				kv.logger.Sugar().Infof("response : %s ", v.String())
			}
		}
		if err != nil {
			kv.logger.Sugar().Errorf("error in scanner : %s ", err)
			break
		}
	}
	rb2 := strings.TrimSuffix(rb1.String(), ",")
	rb2 = rb2 + "]"
	kv.logger.Sugar().Infof("final response : %s ", rb2)
	return rb2, nil
}

// GetFromKeyWithLimit fetches data after a time range with limit
func (kv *KV) GetKeyWithLimit(key string, limit uint32) (string, error) {
	var rb1 strings.Builder
	rb1.WriteString("[")
	scanRequest, err := hrpc.NewScanRangeStr(context.Background(), kv.table,
		key, "", hrpc.NumberOfRows(limit))
	if err != nil {
		kv.logger.Sugar().Errorf("error scanRequest : %s ", err)
	}
	scanner := kv.cli.Scan(scanRequest)
	kv.logger.Sugar().Infof("into query : %s / %s ", kv.table, key)
	for {
		r, err := scanner.Next()
		if r != nil {
			for _, v := range r.Cells {
				rb1.WriteString(string(v.Value))
				rb1.WriteString(",")
				kv.logger.Sugar().Infof("response : %s ", v.String())
			}
		}
		if err != nil {
			kv.logger.Sugar().Errorf("error in scanner : %s ", err)
			break
		}
	}
	rb2 := strings.TrimSuffix(rb1.String(), ",")
	rb2 = rb2 + "]"
	kv.logger.Sugar().Infof("final response : %s ", rb2)
	return rb2, nil
}
