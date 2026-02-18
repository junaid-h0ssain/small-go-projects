# Small Go Projects

A collection of small Go projects demonstrating different functionalities including web APIs and Slack integrations.

## Contents

- [fiber-app](#fiber-app)
- [slackbot](#slackbot)
- [slack-file-upload](#slack-file-upload)

---

## [fiber-app](fiber-app/)

A simple CRM-like REST API built with the [Fiber](https://gofiber.io/) web framework and [GORM](https://gorm.io/) with SQLite.

### Features
- Create, Read, and Delete leads.
- Persistent storage using SQLite.

### API Endpoints
- `GET /`: Retrieve all leads.
- `GET /:id`: Retrieve a specific lead by ID.
- `POST /`: Create a new lead (requires JSON body).
- `DELETE /:id`: Delete a lead by ID.

### Setup
1. Navigate to `fiber-app/`.
2. Run `go run main.go`.
3. The server will start on port `8080`.

---

## [slackbot](slackbot/)

A Slack bot built using the [slacker](https://github.com/shomali11/slacker) framework. It provides a simple command to calculate a user's age based on their year of birth.

### Features
- Responds to the command: `My year of birth is <year>`.
- Displays command analytics in the console.

### Configuration
Requires the following environment variables:
- `SLACK_BOT_TOKEN`
- `SLACK_APP_TOKEN`

### Setup
1. Navigate to `slackbot/`.
2. Set environment variables.
3. Run `go run main.go`.

---

## [slack-file-upload](slack-file-upload/)

A utility script to upload files to a Slack channel using the [slack-go](https://github.com/slack-go/slack) library.

### Features
- Uploads local files to a specified Slack channel.

### Configuration
Requires the following environment variables:
- `SLACK_BOT_TOKEN`
- `CHANNEL_ID`

### Setup
1. Navigate to `slack-file-upload/`.
2. Update the `fileArr` in `main.go` with the file paths you wish to upload (e.g., `ZIPL.pdf`).
3. Set environment variables.
4. Run `go run main.go`.
