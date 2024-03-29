package models

import "time"

type Bat struct {
	Bat           []int     `json:"bat"`
	Capacity      int       `json:"capacity"`
	Current       int       `json:"current"`
	Temp          []int     `json:"temp"`
	CycleCapacity int       `json:"cycleCapacity"`
	Time          time.Time `json:"time"`
}

type Gps struct {
	Lat  string    `json:"lat"`
	Long string    `json:"long"`
	Time time.Time `json:"time"`
}

type Sensor struct {
	ID    string `json:"id"`
	Model []int  `json:"model"`
	Speed int    `json:"speed"`
	Beep  bool   `json:"beep"`
	BatArray   []Bat  `json:"bat"`
	GpsArray   []Gps  `json:"gps"`
	Distance int       `json:"distance"`
	Time     time.Time `json:"time"`
	CreatedAt time.Time `json:"create_at"`
	UpdateAt time.Time	`json:"update_at"`
}