package wyy

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"wyy/model"
)

type DownloadLrcController struct {
	beego.Controller
}

//歌词下载
func (c *DownloadLrcController) Post() {

	c.downloadLrc(111111, "歌曲名字", "演唱者")

}

func (c *DownloadLrcController) downloadLrc(songid int, songname string, authorname string) {

	musics := model.Lyric{}
	h := strconv.Itoa(songid)
	hp := "https://api.imjad.cn/cloudmusic/?type=lyric&id=" + h
	req, _ := http.NewRequest("GET", hp, nil)
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}

	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &musics)

	fmt.Println(musics.Lrc.Lyric)

	var filename = "xxx/" + songname + ".txt"
	var f *os.File
	var err1 error
	/* 使用 io.WriteString 写入文件 */
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
		n, err := io.WriteString(f, musics.Lrc.Lyric) //写入文件(字符串)

		if err != nil || err1 != nil {
			fmt.Println("删除文件")
		}
		check(err1)
		fmt.Printf("写入 %d 个字节n", n)
	}

	c.Ctx.Output.Download(filename, authorname+"_"+songname+".txt")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
