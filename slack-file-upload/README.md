# slack-file-upload

A Go project used to upload local files to a specific Slack channel.

## Features
-   Allows for easy uploading of local files to Slack.
-   Built with the official [slack-go/slack](https://github.com/slack-go/slack) library.

## Configuration
Before running the bot, you'll need the following information:
1.  **SLACK_BOT_TOKEN**: Your Slack Bot Token.
2.  **CHANNEL_ID**: The ID of the channel where you'd like to upload the files.

## Local Configuration
Update the following variables in `main.go` for your specific file:
-   `fileArr`: Array of file names to upload.

## Installation and Run

1.  Navigate to the `slack-file-upload` folder:
    ```bash
    cd slack-file-upload
    ```
2.  Install dependencies:
    ```bash
    go mod tidy
    ```
3.  Set the environment variables and run the application:
    ```bash
    export SLACK_BOT_TOKEN="your-bot-token"
    export CHANNEL_ID="your-channel-id"
    go run main.go
    ```
