package gromer

import (
	"context"
	"fmt"
	"io"
	"reflect"

	"github.com/rs/zerolog/log"
	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/gcpfirestore"
)

type Collection[S any] struct {
	*docstore.Collection
	Type reflect.Type
}

func (c *Collection[S]) Query() *Query[S] {
	return &Query[S]{c.Collection.Query(), c}
}

type Query[S any] struct {
	*docstore.Query
	Parent *Collection[S]
}

func (q *Query[S]) Where(fp docstore.FieldPath, op string, value interface{}) *Query[S] {
	return &Query[S]{q.Query.Where(fp, op, value), q.Parent}
}

func (q *Query[S]) Limit(n int) *Query[S] {
	return &Query[S]{q.Query.Limit(n), q.Parent}
}

func (q *Query[S]) OrderBy(field, direction string) *Query[S] {
	return &Query[S]{q.Query.OrderBy(field, direction), q.Parent}
}

func (q *Query[S]) One(ctx context.Context) (S, int, error) {
	results, err := q.All(ctx)
	if err != nil {
		return *new(S), 500, err
	}
	arr := reflect.ValueOf(results)
	if arr.Len() == 0 {
		return *new(S), 404, fmt.Errorf("%s not found", q.Parent.Type.Name())
	}
	return arr.Index(0).Interface().(S), 200, nil
}

func (q *Query[S]) All(ctx context.Context) ([]S, error) {
	iter := q.Get(ctx)
	defer iter.Stop()
	results := reflect.New(reflect.SliceOf(reflect.PtrTo(q.Parent.Type))).Elem()
	for {
		v := reflect.New(q.Parent.Type)
		err := iter.Next(ctx, v.Interface())
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		} else {
			results.Set(reflect.Append(results, v))
		}
	}
	return results.Interface().([]S), nil
}

func NewCollection[S any](project, name string, t interface{}) *Collection[S] {
	coll, err := docstore.OpenCollection(context.Background(), fmt.Sprintf("firestore://projects/%s/databases/(default)/documents/%s-%s?name_field=ID", project, project, name))
	if err != nil {
		log.Fatal().Stack().Err(err).Msgf("failed to getTable %s", name)
	}
	return &Collection[S]{coll, reflect.TypeOf(t)}
}
