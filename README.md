# tsdns

tsdns is a tool to lookup [TeamSpeak 3](https://teamspeak.com/en/) server IPs and ports (if available) using dns resolution, srv dns (ts3 & tsdns) resolution and TeamSpeak nickname resolution provided by the [myTeamspeak](https://www.myteamspeak.com/) service.

## Why use this?

Usually dns resoltion using "normal" dns requests as well as ts3 and tsdns requests does not run into any issues for most users. However since implementing custom nicknames (available at [myTeamspeak](https://www.myteamspeak.com/)) for [TeamSpeak](https://teamspeak.com/en/) servers some users I know ran into issues resolving those nicknames to hosts and ports.

Therefore I created this tool serving its sole purpose of easing up the process for non-technical users, so they can keep connecting to their favourite [TeamSpeak 3](https://teamspeak.com/en/) server(s).

## Usage

Build the program yourself (see below) or [download a pre-compiled binary](https://github.com/devusSs/tsdns/releases).

Run the tool using:

```bash
tsdns -host <servers_host_or_nickname> -service <dns_or_ts3_or_tsdns_or_nick> -protocol <udp_or_tcp>
```

If you are unsure which service to use stick to the defaults and do not provide any service for the start. If you do not get the results you are looking for try going through the services until you get a result. [TeamSpeak](https://teamspeak.com/en/) unfortunately does not have the best documentation regarding these things.

The protocol flag is not really relevant for most users and can be left untouched, the program will take care of it. For more experienced users the protocol flag will set whether we query ts3 and tsdns requests using udp or tcp.

## Building yourself

Simply make sure you have [Go](https://go.dev) installed and run `go build` in the directory.

**NOTE**: This will not write any build variables to the binary. It is therefor recommended to use a precompiled binary (for now).
