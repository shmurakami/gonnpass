package gonnpass

var Url = "https://connpass.com/api/v1/event/"

type Option struct {
	Id        uint
	Keyword   string
	Month     string
	Date      string
	Name      string
	Owner     string
	Group     Group
	GroupFlag uint
	Offset    uint
	Limit     uint
	OrderFlag string
	Order     uint
}

type Response struct {
	//results_returned	整数	含まれる検索結果の件数	1
	//results_available	整数	検索結果の総件数	191
	//results_start	整数	検索の開始位置	1
	//events	配列(複数要素)	検索結果のイベントリスト
	//event_id	整数	イベントID	364
	//title	文字列(UTF-8)	タイトル	BPStudy#56
	//catch	文字列(UTF-8)	キャッチ	株式会社ビープラウドが主催するWeb系技術討論の会
	//description	文字列(UTF-8)	概要(HTML形式)	今回は「Python プロフェッショナル　プログラミング」執筆プロジェクトの継続的ビルドについて、お話しして頂きます。
	//event_url	文字列(UTF-8)	connpass.com 上のURL	https://connpass.com/event/364/
	//hash_tag	文字列(UTF-8)	Twitterのハッシュタグ	bpstudy
	//started_at	文字列(UTF-8)	イベント開催日時 (ISO-8601形式)	2012-04-17T18:30:00+09:00
	//ended_at	文字列(UTF-8)	イベント終了日時 (ISO-8601形式)	2012-04-17T20:30:00+09:00
	//limit	整数	定員	80
	//event_type	文字列(UTF-8)	イベント参加タイプ	participation: connpassで参加受付あり
	//advertisement: 告知のみ
	//series	オブジェクト	グループ
	//id	整数	グループID	1
	//title	文字列(UTF-8)	グループタイトル	BPStudy
	//url	文字列(UTF-8)	グループのconnpass.com 上のURL	https://connpass.com/series/1/
	//address	文字列(UTF-8)	開催場所	東京都渋谷区千駄ヶ谷5-32-7
	//place	文字列(UTF-8)	開催会場	BPオフィス (NOF南新宿ビル4F)
	//lat	浮動小数点数	開催会場の緯度	35.680236100000
	//lon	浮動小数点数	開催会場の経度	139.701308500000
	//owner_id	整数	管理者のID	8
	//owner_nickname	文字列(UTF-8)	管理者のニックネーム	haru860
	//owner_display_name	文字列(UTF-8)	管理者の表示名	佐藤 治夫
	//accepted	整数	参加者数	80
	//waiting	整数	補欠者数	15
	//updated_at	文字列(UTF-8)	更新日時 (ISO-8601形式)	2012-03-20T12:07:32+09:00
}
