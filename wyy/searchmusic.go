package wyy

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"wyy/model"
)

type SearchMusicController struct {
	beego.Controller
}

//根据歌曲名称查询
func (c *SearchMusicController) Post() {

	content := c.GetString("XXX") //接收查询的参数 根据歌曲名称 默认返回前五条 可以自己做分页

	music := model.SearchMusic{}

	str := url.QueryEscape(content)

	req, _ := http.NewRequest("GET", "https://music.163.com/api/search/get?type=1&offset=0&limit=5&s="+str, nil) //offset limit 控制每页多少条 str要查询的字段
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(body, &music)

	c.Data["musics"] = c.getMusicList(music)

	c.TplName = "xxx.tpl"

}

func (c *SearchMusicController) getMusicList(music model.SearchMusic) []model.BackMusic {

	backMusic := make([]model.BackMusic, 0)
	lastbackMusic := make([]model.BackMusic, 0)
	sd := make(chan model.SongsDetail, 5)
	s := make(chan model.Song, 5)
	songsDetail := make([]model.SongsDetail, 0)
	songs := make([]model.Song, 0)

	defer func() {
		close(s)
		close(sd)
	}()

	for i := 0; i < len(music.Result.Songs); i++ {

		go func(ii int) {
			musics := model.SongsDetail{}
			h := strconv.Itoa(music.Result.Songs[ii].Id)
			hp := "https://api.imjad.cn/cloudmusic/?type=detail&id=" + h
			req, _ := http.NewRequest("GET", hp, nil)
			req.Header.Set("Content-type", "application/json")
			client := &http.Client{}

			response, _ := client.Do(req)

			body, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(body, &musics)

			sd <- musics

		}(i)

		go func(iii int) {
			musics := model.Song{}
			h := strconv.Itoa(music.Result.Songs[iii].Id)
			hp := "https://api.imjad.cn/cloudmusic/?type=song&id=" + h + "&br=128000"
			req, _ := http.NewRequest("GET", hp, nil)
			req.Header.Set("Content-type", "application/json")
			client := &http.Client{}

			response, _ := client.Do(req)

			body, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(body, &musics)

			s <- musics

		}(i)

	}

	for t := 0; t < len(music.Result.Songs); t++ {
		songsDetail = append(songsDetail, <-sd)
		songs = append(songs, <-s)
	}

	for n := 0; n < len(songs); n++ {
		for k := 0; k < len(songsDetail); k++ {
			if songs[n].Data[0].Id == songsDetail[k].Songs[0].Id {
				music := model.BackMusic{}
				music.Id = songsDetail[k].Songs[0].Id
				music.MusicName = songsDetail[k].Songs[0].Name
				music.Play = songs[n].Data[0].Url

				an := "" //名字很长不愿意显示的话...
				for j := 0; j < len(songsDetail[k].Songs[0].Ar); j++ {
					if len(songsDetail[k].Songs[0].Ar) >= 3 {
						an = "很多人"
						break
					}
					an += songsDetail[k].Songs[0].Ar[j].Name
					if j != len(songsDetail[k].Songs[0].Ar)-1 {
						an += "|"
					}
				}
				music.AuthorName = an

				music.PicUrl = songsDetail[k].Songs[0].Al.PicUrl

				music.PlayListName = songsDetail[k].Songs[0].Al.Name
				music.Username = "未知..." //歌单的名字 随便写了
				backMusic = append(backMusic, music)
				continue
			}
		}
	}

	for l := 0; l < len(music.Result.Songs); l++ {
		for j := 0; j < len(backMusic); j++ {
			if music.Result.Songs[l].Id == backMusic[j].Id {
				lastbackMusic = append(lastbackMusic, backMusic[j])
				continue
			}
		}
	}

	return lastbackMusic
}
