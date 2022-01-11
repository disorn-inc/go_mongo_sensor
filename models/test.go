package models

import "time"

type Test struct {
	ID    string `json:"id"`
	Model []int  `json:"model"`
	Speed int    `json:"speed"`
	Beep  bool   `json:"beep"`
	Bat   []struct {
		Bat           []int     `json:"bat"`
		Capacity      int       `json:"capacity"`
		Current       int       `json:"current"`
		Temp          []int     `json:"temp"`
		CycleCapacity int       `json:"cycleCapacity"`
		Time          time.Time `json:"time"`
	} `json:"bat"`
	Gps []struct {
		Lat  string    `json:"lat"`
		Long string    `json:"long"`
		Time time.Time `json:"time"`
	} `json:"gps"`
	Distance int       `json:"distance"`
	Time     time.Time `json:"time"`
}