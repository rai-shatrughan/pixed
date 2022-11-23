package util

import (
	"context"
	"strings"

	// "os"

	"github.com/spf13/viper"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
	"go.uber.org/zap"
)

// KV is wrapper for KV Database
type kvStrore struct {
	cli    gohbase.Client
	conf   *viper.Viper
	logger *zap.Logger
	table  string
}

// New returns instance of KV Store
func NewKV(conf *viper.Viper, logger *zap.Logger) *kvStrore {
	kv := new(kvStrore)
	kv.conf = conf
	kv.logger = logger
	kv.cli = gohbase.NewClient(conf.GetString("hbase.zookeeper"))
	kv.table = conf.GetString("hbase.table")
	return kv
}

// Put upserts data into KV store
func (kv *kvStrore) Put(table, key, value string) error {
	kv.logger.Sugar().Info("Writing : ", key)
	values := map[string]map[string][]byte{"cf": {"a": []byte(value)}}
	putRequest, _ := hrpc.NewPutStr(context.Background(), table, key, values)
	_, err := kv.cli.Put(putRequest)
	return err
}

// GetFromKeyWithLimit fetches data after a time range with limit
func (kv *kvStrore) Get(key string) (string, error) {
	var rb1 strings.Builder
	rb1.WriteString("[")
	pFilter := filter.NewPrefixFilter([]byte(key))
	scanRequest, err := hrpc.NewScanStr(context.Background(), kv.table,
		hrpc.Filters(pFilter))
	if err != nil {
		kv.logger.Sugar().Errorf("error scanRequest : %s ", err)
	}
	scanner := kv.cli.Scan(scanRequest)
	kv.logger.Sugar().Debugf("into query : %s / %s ", kv.table, key)
	for {
		r, err := scanner.Next()
		if r != nil {
			for _, v := range r.Cells {
				rb1.WriteString(string(v.Value))
				rb1.WriteString(",")
				kv.logger.Sugar().Debugf("response : %s ", v.String())
			}
		}
		if err != nil {
			// kv.logger.Sugar().Errorf("error in scanner : %s ", err)
			break
		}
	}
	rb2 := strings.TrimSuffix(rb1.String(), ",")
	rb2 = rb2 + "]"
	kv.logger.Sugar().Debugf("final response : %s ", rb2)
	return rb2, nil
}
