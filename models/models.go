package models

type Music struct {
	id int64
	name string
	singer int64
	searchWorld []string
	url string
}

type Playbill struct {
	id int64
	name string
	searchWorld []string
	music []int64
}

type Player struct {
	id int64
	name string
	searchWorld []string
	music []int64
	playbill []int64
}
