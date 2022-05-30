package mongo

import (
	"context"

	"github.com/qiniu/qmgo"
	opts "github.com/qiniu/qmgo/options"
)

type Service interface {
	Collection(model interface{}) Collection
}

type Collection interface {
	Find(ctx context.Context, filter interface{}, opts ...opts.FindOptions) qmgo.QueryI
	InsertOne(ctx context.Context, doc interface{}, opts ...opts.InsertOneOptions) (result *qmgo.InsertOneResult, err error)
	InsertMany(ctx context.Context, docs interface{}, opts ...opts.InsertManyOptions) (result *qmgo.InsertManyResult, err error)
	Upsert(ctx context.Context, filter interface{}, replacement interface{}, opts ...opts.UpsertOptions) (result *qmgo.UpdateResult, err error)
	UpsertId(ctx context.Context, id interface{}, replacement interface{}, opts ...opts.UpsertOptions) (result *qmgo.UpdateResult, err error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...opts.UpdateOptions) (err error)
	UpdateId(ctx context.Context, id interface{}, update interface{}, opts ...opts.UpdateOptions) (err error)
	UpdateAll(ctx context.Context, filter interface{}, update interface{}, opts ...opts.UpdateOptions) (result *qmgo.UpdateResult, err error)
	ReplaceOne(ctx context.Context, filter interface{}, doc interface{}, opts ...opts.ReplaceOptions) (err error)
	Remove(ctx context.Context, filter interface{}, opts ...opts.RemoveOptions) (err error)
	RemoveId(ctx context.Context, id interface{}, opts ...opts.RemoveOptions) (err error)
	RemoveAll(ctx context.Context, filter interface{}, opts ...opts.RemoveOptions) (result *qmgo.DeleteResult, err error)
	Aggregate(ctx context.Context, pipeline interface{}, opts ...opts.AggregateOptions) qmgo.AggregateI
	EnsureIndexes(ctx context.Context, uniques []string, indexes []string) (err error)
	CreateIndexes(ctx context.Context, indexes []opts.IndexModel) (err error)
	CreateOneIndex(ctx context.Context, index opts.IndexModel) error
	DropAllIndexes(ctx context.Context) (err error)
	DropIndex(ctx context.Context, indexes []string) error
}
