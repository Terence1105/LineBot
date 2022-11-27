package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MemePictures struct {
	Success bool  `json:"success"`
	Data    Memes `json:"data"`
}

type Memes struct {
	Memes []Meme `json:"memes"`
}

type Meme struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Box_Count int    `json:"box_count"`
}

func New_MemesPicture() MemePictures {
	var memePictures MemePictures

	res, err := http.Get("https://api.imgflip.com/get_memes")
	if err != nil {
		panic(fmt.Errorf("fatal error get_memes: %w", err))
	}

	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Errorf("fatal error ReadAll: %w", err))
	}
	defer res.Body.Close()

	er := json.Unmarshal(sitemap, &memePictures)
	if er != nil {
		panic(fmt.Errorf("fatal error Unmarshal: %w", err))
	}

	return memePictures
}
