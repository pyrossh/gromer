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

type Collection struct {
	*docstore.Collection
	Type reflect.Type
}

func (c Collection) Query() *Query {
	return &Query{c.Collection.Query(), c, c.Type}
}

type Query struct {
	*docstore.Query
	Parent Collection
	Type   reflect.Type
}

func (q *Query) Where(fp docstore.FieldPath, op string, value interface{}) *Query {
	return &Query{q.Query.Where(fp, op, value), q.Parent, q.Type}
}

func (q *Query) Limit(n int) *Query {
	return &Query{q.Query.Limit(n), q.Parent, q.Type}
}

func (q *Query) OrderBy(field, direction string) *Query {
	return &Query{q.Query.OrderBy(field, direction), q.Parent, q.Type}
}

func (q *Query) One(ctx context.Context) (interface{}, int, error) {
	results, err := q.All(ctx)
	if err != nil {
		return nil, 500, err
	}
	arr := reflect.ValueOf(results)
	if arr.Len() == 0 {
		return nil, 404, fmt.Errorf("%s not found", q.Type.Name())
	}
	return arr.Index(0).Interface(), 200, nil
}

func (q *Query) All(ctx context.Context) (interface{}, error) {
	iter := q.Get(ctx)
	defer iter.Stop()
	results := reflect.New(reflect.SliceOf(reflect.PtrTo(q.Type))).Elem()
	for {
		v := reflect.New(q.Type)
		err := iter.Next(ctx, v.Interface())
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		} else {
			results.Set(reflect.Append(results, v))
		}
	}
	return results.Interface(), nil
}

func GetCollection(project, env, name string, t interface{}) *Collection {
	coll, err := docstore.OpenCollection(context.Background(), fmt.Sprintf("firestore://projects/%s/databases/(default)/documents/%s?name_field=ID", project, env+"-"+name))
	if err != nil {
		log.Fatal().Stack().Err(err).Msgf("failed to GetCollection %s", name)
	}
	return &Collection{coll, reflect.TypeOf(t)}
}
