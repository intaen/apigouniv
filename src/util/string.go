package util

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomData(listData []string) string {
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())
	return fmt.Sprint(listData[rand.Intn(len(listData))])
}
