package mongo

import (
	"context"
	"reflect"

	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/middleware"
	"github.com/qiniu/qmgo/operator"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host               string
	Scheme             string
	Username           string
	Password           string
	Database           string
	AuthDBName         string
	PoolSizeMultiplier float64
	SSL                bool
	SetSafe            bool
}

type Impl struct {
	*qmgo.Client
	DBName string
}

func isBeforeAction(opType operator.OpType) bool {
	return opType == operator.BeforeInsert || opType == operator.BeforeRemove || opType == operator.BeforeQuery ||
		opType == operator.BeforeUpdate || opType == operator.BeforeUpsert || opType == operator.BeforeReplace
}

func isAfterAction(opType operator.OpType) bool {
	return opType == operator.AfterInsert || opType == operator.AfterRemove || opType == operator.AfterQuery ||
		opType == operator.AfterUpdate || opType == operator.AfterUpsert || opType == operator.AfterReplace
}

func doWhatWeWant(ctx context.Context, doc interface{}, opType operator.OpType, opts ...interface{}) error {
	if isBeforeAction(opType) {
		logrus.Debug("before")
	}
	if isAfterAction(opType) {
		logrus.Debug("after")
	}
	return nil
}

func Init(config Config) (Service, error) {
	middleware.Register(doWhatWeWant)
	client, err := ConnectToMongo(config)
	if err != nil {
		return nil, err
	}

	return &Impl{
		Client: client,
		DBName: config.Database,
	}, nil
}

func (m *Impl) Collection(model interface{}) Collection {
	name := reflect.TypeOf(model).Name()
	if name != "" {
		// not array of model
		return m.Client.Database(m.DBName).Collection(name)
	} else {
		// if we pass array of model, we need to use Elem to get Name()
		return m.Client.Database(m.DBName).Collection(reflect.TypeOf(model).Elem().Name())
	}
}
