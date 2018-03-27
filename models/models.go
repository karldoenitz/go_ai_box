package models

type MusicData struct {
	Data []Music
	Length int
}

type PlaybillData struct {
	Data []Playbill
	Length int
}

type SingerData struct {
	Data []Singer
	Length int
}

type Music struct {
	Id int64
	Name string
	Singer []int64
	SearchWorld []string
	Url string
}

type Playbill struct {
	Id int64
	Name string
	SearchWorld []string
	Music []int64
}

type Singer struct {
	Id int64
	Name string
	SearchWorld []string
	Music []int64
	Playbill []int64
}


