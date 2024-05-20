package model

import "time"

type User struct {
	UserId               string
	HaveGetTodayLuckyBag bool      // 是否已经领取当日福袋
	TodayLuckyBag        time.Time // 领取福袋的时间
}
