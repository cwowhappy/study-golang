package studybase

import (
	"fmt"
	"time"
)

func Sample01() {
	now := time.Now()
	defer func(beginTime time.Time) {
		fmt.Printf(time.Now().Sub(beginTime).String())
	}(now)
	time.Sleep(time.Second * 1)
}
