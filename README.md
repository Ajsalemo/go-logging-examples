# go-logging-examples
Various examples of logging to file and console

This repo shows basic examples of how to write to console, files, both file and console - with and without custom loggers - as well as third-party loggers.

Each folder matches a route to it's name that invokes the logger logic on said route.

**How to run**: Clone this repo and run `go run .` in the root of the project.

- `console`: Maps to `/logger/console` and writes to console with the Go standard logger
- `customloggertoconsole`: Maps to `/logger/customloggertoconsole` and writes to console with a custom logger using the Go standard logger
- `customloggertoconsoleandfile`: Maps to `/logger/customloggertoconsoleandfile` and writes to a file and console with a custom logger using the Go standard logger
- `customloggertofile`: Maps to `/logger/customloggertofile` and writes to a file with a custom logger using the Go standard logger
- `index`: Maps to `/` - root path
- `logfile`: Maps to `/logger/logfile` and writes to a file with the Go standard logger
- `logrus`: Maps to `/logger/logrus` and writes to console with [Logrus](https://github.com/Sirupsen/logrus)
- `zap`: Maps to `/logger/zap` and writes to console with [Zap](https://github.com/uber-go/zap)
- `zerolog`: Maps to `/logger/zerolog` and writes to console with [Zerolog](https://github.com/rs/zerolog)
