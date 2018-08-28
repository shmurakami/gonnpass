package gonnpass

import (
	"testing"
	"encoding/json"
	"bytes"
)

func TestResponseUnmarshal(t *testing.T) {

	response := `
{
  "ResultsReturned": 2,
  "events": [
    {
      "event_url": "https://mgt-college.connpass.com/event/94102/",
      "event_type": "participation",
      "owner_nickname": "KanriMega",
      "series": {
        "url": "https://mgt-college.connpass.com/",
        "id": 1890,
        "title": "MGTカレッジ"
      },
      "updated_at": "2018-08-22T11:13:19+09:00",
      "lat": "35.696006400000",
      "started_at": "2018-08-24T19:30:00+09:00",
      "hash_tag": "Python,AI,ディープラーニング,Django,機械学習,クローラ,スクレイピング",
      "title": "基礎からのPythonハンズオンセミナー-1回目Pythonの基礎　新宿校",
      "event_id": 94102,
      "lon": "139.698533200000",
      "waiting": 0,
      "limit": 14,
      "owner_id": 74539,
      "owner_display_name": "ＭＧＴカレッジ管理",
      "description": "<h1>東京新宿校での開催です。</h1>",
      "address": "東京都新宿区西新宿７丁目４番４号(武蔵ビル2Ｆ)",
      "catch": "Pythonセミナー１回目",
      "accepted": 11,
      "ended_at": "2018-08-24T22:00:00+09:00",
      "place": "ＭＧＴカレッジ新宿校"
    },
    {
      "event_url": "https://yamatosecurity.connpass.com/event/88767/",
      "event_type": "participation",
      "owner_nickname": "yamatosecurity",
      "series": {
        "url": "https://yamatosecurity.connpass.com/",
        "id": 5001,
        "title": "大和セキュリティ勉強会"
      },
      "updated_at": "2018-08-22T10:42:26+09:00",
      "lat": "34.689908300000",
      "started_at": "2018-08-26T11:30:00+09:00",
      "hash_tag": "yamasec",
      "title": "大和セキュリティ勉強会：Pythonでパケット解析 (8月26日)(日)",
      "event_id": 88767,
      "lon": "135.193253600000",
      "waiting": 10,
      "limit": 40,
      "owner_id": 137055,
      "owner_display_name": "yamatosecurity",
      "description": "<h1>概要</h1>\n<p>パケット解析はWireSharkでやることが多いと思いますが、</p>",
      "address": "神戸市中央区京町72番地 新クレセントビル10F",
      "catch": "Python 4 Hackerz",
      "accepted": 40,
      "ended_at": "2018-08-26T20:00:00+09:00",
      "place": "神戸デジタル・ラボ"
    }
  ],
  "results_start": 1,
  "results_available": 1000
}
`
	var res Response
	err := json.Unmarshal(bytes.NewBufferString(response).Bytes(), &res)
	if err != nil {
		t.Fatal("failed to parse to response");
	}

	if res.Events[0].EventId != 94102 {
		t.Fatalf("failed to parse. EventId %d is not identical to 94102 ", res.Events[0].EventId)
	}

}
