package model

//这个是最终的返回结果
type BackMusic struct {
	Id           int    `json:"Id"` //歌曲ID
	MusicName    string `json:"MusicName"`
	AuthorName   string `json:"AuthorName"`
	PicUrl       string `json:"PicUrl"`       //背景图片
	PlayListName string `json:"PlayListName"` //歌单名字
	Username     string `json:"Username"`     //用户名
	Play         string `json:"Play"`         //下载链接
}

//这里只是接收了一部分参数 具体看实际需求
type WyyMusic struct {
	Code     string   `json:"Code"`
	Playlist Playlist `json:"Playlist"`
}

type Playlist struct {
	Id          int      `json:"Id"`
	Name        string   `json:"Name"`
	CoverImgUrl string   `json:"CoverImgUrl"`
	Creator     Creator  `json:"Creator"`
	TrackCount  int      `json:"TrackCount"` //数量
	Tracks      []Tracks `json:"Tracks"`
}

type Creator struct {
	AvatarUrl string `json:"AvatarUrl"` //头像
	Nickname  string `json:"Nickname"`
}

type Tracks struct {
	Name        string `json:"Name"`
	Id          int    `json:"Id"`
	Ar          []Ar   `json:"Ar"`
	Al          Al     `json:"Al"`
	Lyric       string `json:"Lyric"`       //歌词
	DownloadUrl string `json:"DownloadUrl"` //下载链接
	PublishTime string `json:"PublishTime"`
}

type Ar struct {
	Name string `json:"Name"`
}

type Al struct {
	Name   string `json:"Name"`
	PicUrl string `json:"PicUrl"`
}

type Lyric struct {
	Lrc       Lrc       `json:"Lrc"`
	LyricUser LyricUser `json:"LyricUser"`
}

type LyricUser struct {
	Id int `json:"Id"`
}

type Lrc struct {
	Lyric string `json:"Lyric"`
}

type Song struct {
	Data []Data `json:"Data"`
}

type Data struct {
	Id  int    `json:"Id"`
	Url string `json:"Url"`
}
