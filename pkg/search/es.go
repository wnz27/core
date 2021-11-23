package search

import (
	"context"
	"crypto/tls"
	"net/http"
	"reflect"

	"github.com/pkg/errors"
	pb "github.com/tkeel-io/core/api/core/v1"
	"github.com/tkeel-io/core/pkg/logger"

	"github.com/olivere/elastic/v7"
	"google.golang.org/protobuf/types/known/structpb"
)

var log = logger.NewLogger("core.search.es")

const EntityIndex = "entity"

type ESClient struct {
	client *elastic.Client
}

func interface2string(in interface{}) (out string) {
	if in == nil {
		return
	}
	switch inString := in.(type) {
	case string:
		out = inString
	default:
		out = ""
	}
	return
}

func (es *ESClient) Index(ctx context.Context, req *pb.IndexObject) (out *pb.IndexResponse, err error) {
	var indexID string
	out = &pb.IndexResponse{}
	out.Status = "SUCCESS"

	switch kv := req.Obj.AsInterface().(type) {
	case map[string]interface{}:
		indexID = interface2string(kv["id"])
	case nil:
		// do nothing.
		return out, nil
	default:
		return out, ErrIndexParamInvalid
	}

	objBytes, _ := req.Obj.MarshalJSON()
	_, err = es.client.Index().Index(EntityIndex).Id(indexID).BodyString(string(objBytes)).Do(context.Background())
	return out, errors.Wrap(err, "es index failed")
}

func condition2boolQuery(condition []*pb.SearchCondition, boolQuery *elastic.BoolQuery) {
	for _, condition := range condition {
		switch condition.Operator {
		case "$lt":
			boolQuery = boolQuery.Must(elastic.NewRangeQuery(condition.Field).Lt(condition.Value.AsInterface()))
		case "$lte":
			boolQuery = boolQuery.Must(elastic.NewRangeQuery(condition.Field).Lte(condition.Value.AsInterface()))
		case "$gt":
			boolQuery = boolQuery.Must(elastic.NewRangeQuery(condition.Field).Gt(condition.Value.AsInterface()))
		case "$gte":
			boolQuery = boolQuery.Must(elastic.NewRangeQuery(condition.Field).Gte(condition.Value.AsInterface()))
		case "$neq":
			boolQuery = boolQuery.MustNot(elastic.NewTermQuery(condition.Field, condition.Value.AsInterface()))
		case "$eq":
			boolQuery = boolQuery.Must(elastic.NewTermQuery(condition.Field, condition.Value.AsInterface()))
		default:
			boolQuery = boolQuery.Must(elastic.NewMatchQuery(condition.Field, condition.Value.AsInterface()))
		}
	}
}

func (es *ESClient) Search(ctx context.Context, req *pb.SearchRequest) (out *pb.SearchResponse, err error) {
	out = &pb.SearchResponse{}
	out.Items = make([]*structpb.Value, 0)
	searchQuery := es.client.Search().Index(EntityIndex)

	boolQuery := elastic.NewBoolQuery()
	if req.Condition != nil {
		condition2boolQuery(req.Condition, boolQuery)
	}
	if req.Query != "" {
		boolQuery = boolQuery.Must(elastic.NewMultiMatchQuery(req.Query))
	}
	searchResult, err := searchQuery.Query(boolQuery).Pretty(true).Do(ctx)
	if err != nil {
		return
	}
	var ttyp map[string]interface{}
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(map[string]interface{}); ok {
			tt, _ := structpb.NewValue(t)
			out.Items = append(out.Items, tt)
		}
	}
	out.Total = searchResult.TotalHits()
	if req.Page != nil {
		out.Limit = req.Page.Limit
		out.Offset = req.Page.Offset
	}
	return out, errors.Wrap(err, "search failed")
}

func NewESClient(url ...string) pb.SearchHTTPServer {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //nolint
	client, err := elastic.NewClient(elastic.SetURL(url...), elastic.SetSniff(false), elastic.SetBasicAuth("admin", "admin"))
	if err != nil {
		panic(err)
	}

	// ping.
	info, code, err := client.Ping(url[0]).Do(context.Background())
	if nil != err {
		panic(err)
	}

	log.Infof("Elasticsearch returned with code<%d>, version<%s>\n", code, info.Version.Number)

	return &ESClient{client: client}
}
