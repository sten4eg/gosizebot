package tg

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

const (
	ApiScheme = "https"
	ApiHost   = "api.telegram.org"
)

var (
	//go:embed api.key
	ApiKey string

	DefaultUri = ApiScheme + "://" + ApiHost + "/bot" + ApiKey
	ApiClient  = &fasthttp.Client{}
)

func GetUpdates(sinceUpdateId int64) (*Updates, error) {
	rq := fasthttp.AcquireRequest()
	rs := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(rq)
	defer fasthttp.ReleaseResponse(rs)

	rq.SetRequestURI(DefaultUri + "/getUpdates")
	rq.Header.SetContentType("application/application/x-www-form-urlencoded")
	rq.Header.SetMethod("POST")

	if sinceUpdateId != 0 {
		rq.PostArgs().Add("offset", strconv.FormatInt(sinceUpdateId, 10))
	}

	rq.PostArgs().Add("timeout", "60")

	err := ApiClient.Do(rq, rs)
	if err != nil {
		return nil, err
	}

	if rs.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("STATUS: %d", rs.StatusCode())
	}

	var updates = &Updates{}

	err = json.Unmarshal(rs.Body(), &updates)
	if err != nil {
		return updates, err
	}

	return updates, nil
}

func SendInlineReplyArticle(queryId string, results []InlineQueryResultArticle) error {
	rq := fasthttp.AcquireRequest()
	rs := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(rq)
	defer fasthttp.ReleaseResponse(rs)

	rq.SetRequestURI(DefaultUri + "/answerInlineQuery")
	rq.Header.SetContentType("application/x-www-form-urlencoded")
	rq.Header.SetMethod("POST")

	res, _ := json.Marshal(results)

	rq.PostArgs().Add("inline_query_id", queryId)
	rq.PostArgs().Add("cache_time", "0")
	rq.PostArgs().AddBytesV("results", res)

	err := ApiClient.Do(rq, rs)
	if err != nil {
		return err
	}

	if rs.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf("STATUS: %d", rs.StatusCode())
	}

	return nil
}
