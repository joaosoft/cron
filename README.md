cron
================

[![Build Status](https://travis-ci.org/joaosoft/cron.svg?branch=master)](https://travis-ci.org/joaosoft/cron) | [![codecov](https://codecov.io/gh/joaosoft/cron/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/cron) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/cron)](https://goreportcard.com/report/github.com/joaosoft/cron) | [![GoDoc](https://godoc.org/github.com/joaosoft/cron?status.svg)](https://godoc.org/github.com/joaosoft/cron)

A simple tool that allows you to cron your jobs.

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for modes
* from midnight
* recurring

## Dependency Management
>### Dependency

Project dependencies are managed using Dependency. Read more about [Dependency](https://github.com/joaosoft/dependency).
* Get dependency manager: `go get github.com/joaosoft/dependency`

###### Commands
* Install dependencies: `dependency get`
* Update dependencies: `dependency update`
* Reset dependencies: `dependency reset`
* Add dependencies: `dependency add <dependency>`
* Remove dependencies: `dependency remove <dependency>`

>### Go
```
go get github.com/joaosoft/cron
```

>### Configuration
```
{
  "cron": {
    "log": {
      "level": "debug"
    },
    "migration": {
      "path": {
        "database": "schema/db/postgres"
      },
      "db": {
        "schema": "cron",
        "driver": "postgres",
        "datasource": "postgres://user:password@localhost:7000/postgres?sslmode=disable&cron_path=cron"
      },
      "log": {
        "level": "info"
      }
    },
    "dbr": {
      "read_db": {
        "driver": "postgres",
        "datasource": "postgres://user:password@localhost:7000/postgres?sslmode=disable&cron_path=cron"
      },
      "write_db": {
        "driver": "postgres",
        "datasource": "postgres://user:password@localhost:7100/postgres?sslmode=disable&cron_path=cron"
      },
      "log": {
        "level": "debug"
      }
    }
  },
  "manager": {
    "log": {
      "level": "debug"
    }
  }
}
```

## Usage 
This examples are available in the project at [cron/examples](https://github.com/joaosoft/cron/tree/master/examples)

```go
func main() {
	// create cron
	myCron, err := cron.New()
	if err != nil {
		panic(err)
	}

	// add tasks to jobs
	myCron.
		AddJobTaskWithFuncs("teste1_key", canExecute, before, middle, after).
		AddJobTask("teste2_key", cron.NewTaskGeneric(canExecute, before, middle, after)).
		AddJobTaskWithFuncs("teste3_key", canExecute, before, middle, after)

	// wait
	myCron.Wait()
}

func canExecute() (bool, error) {
	logger.Info("executing can execute method")
	return true, nil
}

func before() error {
	logger.Info("executing before method")
	return nil
}

func middle() error {
	logger.Info("executing middle method")
	return nil
}

func after() error {
	logger.Info("executing after method")
	return nil
}
```

> ##### Result:
```
GOROOT=/usr/local/go #gosetup
GOPATH=/Users/joaoribeiro/workspace/go/personal:/Users/joaoribeiro/workspace/go/foursource:/Users/joaoribeiro/workspace/go/sonae:/Users/joaoribeiro/workspace/go/others #gosetup
/usr/local/go/bin/go build -i -o /private/var/folders/d_/ghp5zy210wbcjmwjv2xr55gh0000gn/T/___Run_cron /Users/joaoribeiro/workspace/go/personal/src/cron/examples/main.go #gosetup
/private/var/folders/d_/ghp5zy210wbcjmwjv2xr55gh0000gn/T/___Run_cron #gosetup
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"manager"},"message":"config config_app added","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"manager"},"message":"database db_postgres added","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"manager"},"message":"database db-Read added","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"manager"},"message":"database db-Write added","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"cron"},"message":"scheduling job with name: teste1","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"manager"},"message":"process teste1_key added","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"cron"},"message":"scheduling job with name: teste2","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"manager"},"message":"process teste2_key added","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"manager"},"message":"starting...","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"manager"},"message":"started [ process: teste1_key ]","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"manager"},"message":"started [ process: teste2_key ]","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"cron"},"message":"job teste2_key waiting next scheduled time 1m0s]","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"cron"},"message":"job teste1_key waiting next scheduled time 23h41m30.132184s]","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"manager"},"message":"started","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"cron"},"message":"cron service started!","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:29:19"},"tags":{"service":"cron"},"message":"waiting for quit signal...","sufixes":{"ip":"192.168.1.67"}}
{"prefixes":{"level":"info","timestamp":"2019-06-14 19:04:31:19"},"tags":{"service":"cron"},"message":"quitting cron","sufixes":{"ip":"192.168.1.67"}}
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
