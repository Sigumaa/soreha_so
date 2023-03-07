package soreha_so

import (
	"math/rand"
	"testing"
	"time"
)

func Test_reply(t *testing.T) {
	so := [...]string{"それはそう", "それは、そう", "sorehasou", "soreha,sou", "soreha sou", "soreha so", "soreha-so", "sorehaso"}
	uz := [...]string{"うざいよ", "その言葉遣いやめたほうがいいよ", "は？"}
	for _, v := range so {
		if Reply(v) != uz[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(uz))] {
			t.Errorf("reply(%q) = %q, want %q", v, Reply(v), uz[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(uz))])
		}
	}
}
