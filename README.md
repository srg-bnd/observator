# Observator

Metrics collection and alerting service

[![go report card](https://goreportcard.com/badge/github.com/srg-bnd/observator?style=flat-square)](https://goreportcard.com/report/github.com/srg-bnd/observator)
[![test status](https://github.com/srg-bnd/observator/workflows/mertricstest/badge.svg?branch=main "test status")](https://github.com/srg-bnd/observator/actions)

## Getting Started

Dependencies:

* Go `1.23`
* PostgreSQL
* Linux or macOS platform

### Startup

#### Server

To build a server, run in the terminal:

```bash
cd cmd/server && go run . # in the root directory of the project
```

To startup the server, run in the terminal:

```bash
./cmd/server/server # in the root directory of the project
```

##### Envs & flags

* `ADDRESS` | `-a` – address and port to run server. Default `:8080`
* `DATABASE_DSN` | `-d` – DB connection address
* `LOG_LEVEL` | `-l` – log level. Default `info`
* `FILE_STORAGE_PATH` | `-f` – file storage path. Default `./temp.storage.db`
* `RESTORE` | `-r` – load data from storage. Default `true`
* `STORE_INTERVAL` | `-i` – store interval in seconds (zero for sync). Default `300`

#### Agent

To build a agent, run in the terminal:

```bash
cd cmd/agent && go run . # in the root directory of the project
```

To startup the agent, run in the terminal:

```bash
./cmd/server/agent # in the root directory of the project
```

##### Envs & flags

* `ADDRESS` | `-a` – address and port to run agent. Default `localhost:8080`
* `KEY` | `-k` – encryption key
* `POLL_INTERVAL` | `-p` – frequency (seconds) of metric polling. Default `2`
* `RATE_LIMIT` | `-l` – number of simultaneous outgoing requests to the server. Default `1`
* `REPORT_INTERVAL` | `-r` – frequency (seconds) of sending values to the server. Default `10`

## Development

### Updating the template

To be able to receive updates to autotests and other parts of the template, run the command:

```bash
git remote add -m main template https://github.com/Yandex-Practicum/go-musthave-metrics-tpl.git
```

To update the autotest code, run the command:

```bash
git fetch template && git checkout template/main .github
```

Then add the received changes to your repository.

### Launching autotests

To successfully run autotests, name the branches `iter<number>`, where `<number>` is the sequence number of the increment. For example, in the branch named `iter4`, autotests will be launched for increments from the first to the fourth.

When you merge a branch with an increment into the main branch `main`, all autotests will be run.

Read more about local and automatic startup in [README autotests](https://github.com/Yandex-Practicum/go-autotests).
