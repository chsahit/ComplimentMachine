package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	clarifai "github.com/clarifai/clarifai-go"
)

type MainController struct {
	beego.Controller
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func (c *MainController) Get() {
	client := clarifai.NewClient("ixACIQvGqKKcJCGLT_xnEh4_jlG7dRKXuzF4jam3", "-XQ5gLtB0ZljTUEmoaqV4LI8UXdlwZNEvLTkSXt-")

	// Let's get some context about these images
	urls := []string{"https://avatars1.githubusercontent.com/u/3252741?v=3&s=400"}
	// Give it to Clarifai to run their magic
	tag_data, err := client.Tag(clarifai.TagRequest{URLs: urls})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tag_data.Results[0].Result.Tag.Classes[0])
	}

	tags := tag_data.Results[0].Result.Tag.Classes

	for i := 0; i < len(tags); i++ {
		c.Ctx.ResponseWriter.Write([]byte(tag_data.Results[0].Result.Tag.Classes[i] + "\n"))

	}
}
