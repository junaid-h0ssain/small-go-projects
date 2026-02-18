# slackbot

A simple Slack bot built in Go using the [slacker](https://github.com/shomali11/slacker) framework. It can calculate your birth year and provide response based on your inputs.

## Features
-   Calculates the age of the user based on birth year.
-   Provides real-time command analytics in the terminal.

## Prerequisites
Before running, ensure you have your Slack bot and app tokens ready.

## Environment Variables
The application uses the following environment variables:
- `SLACK_BOT_TOKEN`: Your Slack Bot Token.
- `SLACK_APP_TOKEN`: Your Slack App Token.

## Installation and Run

1.  Navigate to the `slackbot` folder:
    ```bash
    cd slackbot
    ```
2.  Install dependencies:
    ```bash
    go mod tidy
    ```
3.  Set the environment variables and run the application:
    ```bash
    export SLACK_BOT_TOKEN="your-bot-token"
    export SLACK_APP_TOKEN="your-app-token"
    go run main.go
    ```

## Example Usage
In Slack, type:
`My year of birth is 1990`
The bot will respond with your calculated age.
