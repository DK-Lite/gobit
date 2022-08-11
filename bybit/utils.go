package bybit

import (
	"fmt"
	"strconv"
	"time"
)

func GetNowSecond() string {
	intNow := int(time.Now().UTC().UnixNano() / int64(time.Second))
	return strconv.Itoa(intNow)
}

func GetNowUTC() string {
	intNow := int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	now := strconv.Itoa(intNow)
	return now
}

func Map[T, M any](datas []T, f func(x T) M) []M {
	result := []M{}
	for _, data := range datas {
		result = append(result, f(data))
	}
	return result
}

func CalculateFrom(unit, amount int) string {
	return fmt.Sprintf("%d", GetServerTime()-int64(unit*60*amount))
}
