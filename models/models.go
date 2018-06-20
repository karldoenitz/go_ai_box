package models

type MusicData struct {
	Data      []Music     `json:"data"`
	Length    int         `json:"length"`
}

type PlaybillData struct {
	Data      []Playbill  `json:"data"`
	Length    int         `json:"length"`
}

type SingerData struct {
	Data      []Singer    `json:"data"`
	Length    int         `json:"length"`
}

type Music struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Singer      []int64    `json:"singer"`
	SearchWorld []string   `json:"search_world"`
	Url         string     `json:"url"`
}

type Playbill struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	SearchWorld []string   `json:"search_world"`
	Music       []int64    `json:"music"`
}

type Singer struct {
	Id int64               `json:"id"`
	Name string            `json:"name"`
	SearchWorld []string   `json:"search_world"`
	Music []int64          `json:"music"`
	Playbill []int64       `json:"playbill"`
}


