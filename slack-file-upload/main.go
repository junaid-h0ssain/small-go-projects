package main

import (
	"github.com/slack-go/slack"
	"fmt"
	"os"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN","")
	os.Setenv("CHANNEL_ID","")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"ZIPL.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			Filetype: fileArr[i],
		}
		file, err := os.Open(params.File)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Uploading file: %s\n", file.Name)
	}
}