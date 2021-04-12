# `yag-status`

A set of Go scripts to monitor YAGPDB status by making `GET` requests to the YAGPDB status endpoint.

## What's included

-   Script to find how many shards are down
    -   Not super useful, but perhaps it's a slight improvement over spam refreshing the status page.
-   Script to find the status of YAGPDB on a guild
    -   You should probably use `cshard`.
-   Monitor YAGPDB status by posting to a Discord webhook periodically
    -   Keeps track of how many shards have gone up/down since the last check
    -   Probably the only actual useful one to be honest

## Usage

_Prerequisites:_<br>

-   Go 1.16
-   Git

_Installation:_<br>

1. `git clone https://github.com/jo3-l/yag-status`
2. `cd` into the directory of the script you want to use.
    - e.g. `cd downshards`
3. (Only if using `monitor`) Copy `monitor/config.example.json` to `monitor/config.json` and fill it out.
4. Run `go run .`

## License

[MIT](./LICENSE.md)
