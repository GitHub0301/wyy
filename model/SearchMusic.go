package model

//这里只是接收了一部分参数 具体看实际需求
type SearchMusic struct {
	Code   int    `json:"Code"`
	Result Result `json:"Result"`
}

type Result struct {
	SongCount int     `json:"SongCount"`
	Songs     []Songs `json:"Songs"`
}

type Songs struct {
	Id      int       `json:"Id"`
	Name    string    `json:"Name"`
	Artists []Artists `json:"Artists"`
	Album   Album     `json:"Album"`
}

type Artists struct {
	Id        int    `json:"Id"`
	Name      string `json:"Name"`
	Img1v1Url string `json:"Img1v1Url"`
}

type Album struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

//==== songs detail
type SongsDetail struct {
	Songs []MySongs `json:"Songs"`
}

type MySongs struct {
	Name string `json:"Name"`
	Id   int    `json:"Id"`
	Ar   []Ar   `json:"Ar"`
	Al   Al     `json:"Al"`
}
