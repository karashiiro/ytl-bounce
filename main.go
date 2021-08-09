package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type failure struct {
	Error string `json:"error"`
}

var videoIdRegex = regexp.MustCompile(`https?:\/\/www\.youtube\.com\/watch\?v=(?P<videoId>[^#\&\?%"<>]*)`)

func main() {
	r := gin.Default()
	r.GET("/:channelId", func(c *gin.Context) {
		channelId := c.Param("channelId")
		livePath := fmt.Sprintf("https://www.youtube.com/channel/%s/live", channelId)
		res, err := http.Get(livePath)
		if err != nil {
			c.JSON(400, failure{Error: err.Error()})
			return
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.JSON(400, failure{Error: err.Error()})
			return
		}
		matches := videoIdRegex.FindStringSubmatch(string(body))
		if len(matches) == 0 {
			c.JSON(400, failure{Error: "No live stream found"})
			return
		}
		videoId := videoIdRegex.FindStringSubmatch(string(body))[1]
		videoPath := fmt.Sprintf("https://www.youtube.com/live_chat?is_popout=1&v=%s", videoId)
		c.Redirect(302, videoPath)
	})
	r.Run(":3799")
}
