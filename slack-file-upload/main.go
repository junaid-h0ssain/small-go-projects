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
	channelArr :=os.Getenv("CHANNEL_ID")
	fileArr := []string{"ZIPL.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.UploadFileV2Parameters{
			Channel: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Uploading file: %s\n", file)
	}
}