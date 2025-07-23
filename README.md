# GitHub User Activity

A CLI tool to fetch recent activity of a GitHub User.
<br>
Built as a learning project for [roadmap.sh](https://roadmap.sh/projects/github-user-activity) using Go.

## Features

- Fresh Events from GitHub Events REST API
- Message Formatting for 10+ GitHub events
- Error handling for invalid username, internet connectivity

## Requirements

- Go 1.24 or higher

## Installation

Clone the repository and compile the program:
```bash
git clone https://github.com/architxkumar/GitHub-User-Activity.git
cd GitHub-User-Activity
go build -o ./github-activity
```

## Usage
```bash
./github-activity [github-username]
```

## Example

```text
./github-activity dmalan

Pull request on repository cs50/problems at 2025-06-26 00:06:11 +0000 UTC

Pushed 1 commit to repository  cs50/lectures at 2025-06-24 13:03:25 +0000 UTC

Pushed 2 commits to repository  cs50/lectures  at  2025-06-24 10:36:39 +0000 UTC

Issue event on repository cs50/render50 at 2025-06-24 10:34:11 +0000 UTC

Pushed 1 commit to repository  cs50/lectures at 2025-06-24 09:42:52 +0000 UTC

Pushed 1 commit to repository  cs50/lectures at 2025-06-24 04:08:40 +0000 UTC

Pushed 1 commit to repository  cs50/lectures at 2025-06-24 01:21:37 +0000 UTC

Issue event on repository cs50/jekyll-theme-cs50 at 2025-06-23 17:45:13 +0000 UTC
```

## Note

- GitHub Events has [rate limiting](https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api?apiVersion=2022-11-28) enabled 
- GitHub Events API is not built to serve real-time use cases. 

## License

This project is licensed under the [MIT License](./LICENSE)