package wyy

import (
	"c/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strconv"
	"wyy/model"
)

type GetMusicController struct {
	beego.Controller
	MList model.WyyMusic
}

//得到music列表
func (c *GetMusicController) Post() {

	music := model.WyyMusic{}

	req, _ := http.NewRequest("GET", "https://api.imjad.cn/cloudmusic/?type=playlist&id="+"很好找的歌单id", nil)
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(body, &music)

	c.getMusicList(music)

	c.Data["musics"] = c.getBackMusic()

	c.TplName = "xxx.tpl"

}

func (c *GetMusicController) getBackMusic() []model.BackMusic {
	backMusic := make([]model.BackMusic, 0)
	for i := 0; i < len(c.MList.Playlist.Tracks); i++ {
		music := model.BackMusic{}

		music.PlayListName = c.MList.Playlist.Name
		music.Username = c.MList.Playlist.Creator.Nickname

		music.Id = c.MList.Playlist.Tracks[i].Id

		music.MusicName = c.MList.Playlist.Tracks[i].Name

		music.PicUrl = c.MList.Playlist.Tracks[i].Al.PicUrl
		music.Play = c.MList.Playlist.Tracks[i].DownloadUrl

		an := "" //名字很长不愿意显示的话...
		for j := 0; j < len(c.MList.Playlist.Tracks[i].Ar); j++ {
			if len(c.MList.Playlist.Tracks[i].Ar) > 2 {
				an = "很多人"
				break
			}
			an += c.MList.Playlist.Tracks[i].Ar[j].Name
			if j != len(c.MList.Playlist.Tracks[i].Ar)-1 {
				an += "|"
			}
		}

		music.AuthorName = an
		backMusic = append(backMusic, music)
	}
	return backMusic
}

func (c *GetMusicController) getMusicList(music model.WyyMusic) {

	s := make(chan models.Song, 100)
	songs := make([]models.Song, 0)
	defer func() {
		close(s)
	}()
	for i := 0; i < len(music.Playlist.Tracks); i++ {

		go func(iii int) {
			musics := models.Song{}
			h := strconv.Itoa(music.Playlist.Tracks[iii].Id)
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

	for t := 0; t < len(music.Playlist.Tracks); t++ {
		songs = append(songs, <-s)
	}

	for n := 0; n < len(music.Playlist.Tracks); n++ {
		for k := 0; k < len(songs); k++ {
			if music.Playlist.Tracks[n].Id == songs[k].Data[0].Id {
				music.Playlist.Tracks[n].DownloadUrl = songs[k].Data[0].Url
			}
		}
	}

	c.MList = music
}
