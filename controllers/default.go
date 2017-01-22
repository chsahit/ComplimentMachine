package controllers

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"sort"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	clarifai "github.com/clarifai/clarifai-go"
)

var Redisptr *Client 
type MainController struct {
	beego.Controller
}

type UploadController struct {
	beego.Controller
}

type DataController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

func (c *UploadController) Post() {
	file, header, err := c.GetFile("artifact")
	if err != nil {
		fmt.Println(err)
	}
	if file != nil {
		// get the filename
		fileName := header.Filename
		// save to server
		c.SaveToFile("artifact", fileName)
		cmd := exec.Command("python", "pygur.py", fileName)
		output, _ := cmd.CombinedOutput()
		link := string(output)
		fmt.Println(link)
		c.Ctx.WriteString(link)
	}

}

func (c *DataController) Get() {

	client := clarifai.NewClient("ixACIQvGqKKcJCGLT_xnEh4_jlG7dRKXuzF4jam3",
		"-XQ5gLtB0ZljTUEmoaqV4LI8UXdlwZNEvLTkSXt-")

	// Let's get some context about these images
	urls := []string{c.GetString("image")}
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
		res.Trait = tag_data.Results[0].Result.Tag.Classes[i]
		tags = append(tags, res)
	}
	sort.Sort(tags)

	c.TplName = "data.tpl"

	c.Data["tags"] = tags
	c.Data["image"] = c.GetString("image")
        Redisptr.Set(tags[0], 0, 0)
        for i:= 0; i < 3; i++ {
            if Redisptr.Get(tags[i]) == nil {
                Redisptr.Set(tags[i], 0, 0);
            } else {
                Redisptr.Set(tags[i], Redisptr.Get(tags[i]) + 1, 0);
            }
        }
}
