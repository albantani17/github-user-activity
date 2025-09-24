# GitHub User Activity CLI

A small Go command-line utility that fetches the ten most recent public GitHub events for a given user and prints a concise summary in the terminal.

This implementation follows the Roadmap.sh challenge: https://roadmap.sh/projects/github-user-activity

## Prerequisites
- Go 1.25 or newer
- Git (optional, but recommended for cloning the repository)

## Getting Started
1. Clone the repository
   ```bash
   git clone https://github.com/albantani17/github-user-activity.git
   cd github-user-activity
   ```
2. Run the program without installing (useful for quick checks)
   ```bash
   go run ./cmd/github-activity <github-username>
   ```

## Build Options
- **Local binary (run from the project folder)**
  ```bash
  go build -o github-activity ./cmd/github-activity
  ./github-activity <github-username>        # macOS / Linux
  .\github-activity.exe <github-username>   # Windows
  ```

- **Install to your PATH (run `github-activity` from anywhere)**
  ```bash
  go install ./cmd/github-activity
  ```
  The executable will be placed in your `GOBIN` (or `GOPATH/bin`). Ensure that directory is on your `PATH`, then you can call:
  ```bash
  github-activity <github-username>
  ```

## Example
```bash
github-activity octocat
```
Typical output looks like:
```
Output:
octocat started watching hello-world
octocat pushed 2 commits to spoon-knife
...
```
Event types currently handled include `WatchEvent` and `PushEvent`. Other public events are fetched but ignored by default.

## Notes
- Calls rely on the public GitHub REST API and therefore inherit its unauthenticated rate limits.
- Feel free to extend `internal/github/github.go` to support more event types or richer formatting.
