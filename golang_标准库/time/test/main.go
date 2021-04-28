package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("time now :", t)
	fmt.Println("location:", t.Location())

	name, offset := t.Zone()
	fmt.Println("time zone;  name:", name, "time zone; offset:", offset)

	fmt.Println("Local:", t.Local())

	fmt.Println("UTC time:", t.UTC())

	fmt.Println("In:", t.In(t.Location()))

	fmt.Println("Unix:", t.Unix())

	fmt.Println("UnixNano:", t.UnixNano())

	year, month, day := t.Date()
	fmt.Println("year:", year, "month:", month, "day:", day)

	hour, min, sec := t.Clock()
	fmt.Println("hour:", hour, "min:", min, "sec:", sec)

	fmt.Println("time year:", t.Year())
	fmt.Println("time month:", t.Month())

	fmt.Println("一年的第几天:", t.YearDay())

	fmt.Println("一月的第几日:", t.Day())

	fmt.Println("一天的第几个小时：", t.Hour())
	fmt.Println("一个小时的第几分钟：", t.Minute())
	fmt.Println("一分钟的第几秒钟：", t.Second())
	fmt.Println("一秒钟的纳秒偏移量：", t.Nanosecond())

	fmt.Println("time string: ", t.String())

	res, _ := t.MarshalJSON()
	fmt.Println("json:", string(res))

	

}
