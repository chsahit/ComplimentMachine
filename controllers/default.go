package controllers

import (
	"encoding/json"
	"fmt"

	"sort"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	clarifai "github.com/clarifai/clarifai-go"
)

type MainController struct {
	beego.Controller
}

type Sentiment struct {
	trait       string
	Probability Prob   `json:"probability"`
	Label       string `json:"label"`
}

type Prob struct {
	Neg     float64 `json:"neg"`
	Neutral float64 `json:"neutral"`
	Pos     float64 `json:"pos"`
}

type Sentiments []Sentiment

func (slice Sentiments) Len() int {
	return len(slice)
}

func (slice Sentiments) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice Sentiments) Less(i, j int) bool {
	return slice[i].Probability.Pos*60+
		slice[i].Probability.Neg*30+
		slice[i].Probability.Neutral*10 >
		slice[j].Probability.Pos*60+
			slice[j].Probability.Neg*30+
			slice[j].Probability.Neutral*10
}

func (c *MainController) Get() {
	client := clarifai.NewClient("ixACIQvGqKKcJCGLT_xnEh4_jlG7dRKXuzF4jam3", "-XQ5gLtB0ZljTUEmoaqV4LI8UXdlwZNEvLTkSXt-")

	// Let's get some context about these images
	urls := []string{"https://avatars1.githubusercontent.com/u/3252741?v=3&s=400"}
	// Give it to Clarifai to run their magic
	tag_data, err := client.Tag(clarifai.TagRequest{URLs: urls})

	if err != nil {
		fmt.Println(err)
	}

	var tags Sentiments
	for i := 0; i < len(tag_data.Results[0].Result.Tag.Classes); i++ {
		req := httplib.Post("http://text-processing.com/api/sentiment/")
		req.Param("text", tag_data.Results[0].Result.Tag.Classes[i])
		res := Sentiment{}
		str, _ := req.String()
		json.Unmarshal([]byte(str), &res)
		res.trait = tag_data.Results[0].Result.Tag.Classes[i]
		tags = append(tags, res)
	}
	sort.Sort(tags)

	for i := 0; i < len(tags); i++ {
		c.Ctx.WriteString(tags[i].trait + "\n")
	}
}
