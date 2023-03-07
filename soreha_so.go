package soreha_so

import (
	"math/rand"
	"strings"
	"time"
)

func reply(text string) string {
	so := [...]string{"それはそう", "それは、そう", "sorehasou", "soreha,sou", "soreha sou", "soreha so", "soreha-so", "sorehaso"}
	uz := [...]string{"うざいよ", "その言葉遣いやめたほうがいいよ", "は？"}
	for _, v := range so {
		if strings.Contains(text, v) {
			return uz[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(uz))]
		}
	}
	return ""
}
