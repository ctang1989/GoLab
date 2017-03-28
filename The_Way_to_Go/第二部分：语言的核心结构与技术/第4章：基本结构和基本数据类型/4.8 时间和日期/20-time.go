package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}

/*
	2017-03-20 18:15:14.2358191 +0800 CST
	20.03.2017
	2017-03-20 10:15:14.2448196 +0000 UTC
	2017-03-20 18:15:14.2448196 +0800 CST
	2017-03-27 10:15:14.2448196 +0000 UTC
	20 Mar 17 10:15 UTC
	Mon Mar 20 10:15:14 2017
	20 Mar 2017 10:15
	2017-03-20 10:15:14.2448196 +0000 UTC => 20170320
*/
