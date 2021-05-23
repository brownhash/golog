# golog

Efficient logging for Go

## Usage

### Importing

```shell
$ go get -u github.com/brownhash/golog
```

```go
import "github.com/brownhash/golog"
```

### Setting logging-level

```go
golog.SetLogLevel("logLevel")
```

or

```shell
$ export GOLOG_LOGGING_LEVEL="logLevel"
```

Available logging levels -

1. DEBUG | 10
2. INFO | 20
3. WARN | 30
4. SUCCESS | 30
5. ERROR | 40

> Note: The above logging levels are in decreasing order of their precedence. i.e. DEBUG / 10 will allow all loggers to log, while ERROR / 40 wont allow any other than itself.

### Logging

#### Debug (DEBUG / 10)

```go
golog.Debug("log message")
```

```go
golog.Debugf("log message")
```

#### Info (INFO / 20)

```go
golog.Info("log message")
```

```go
golog.Infof("log message")
```

#### Warn (WARN / 30)

```go
golog.Warn("log message")
```

```go
golog.Warnf("log message")
```

#### Success (SUCCESS / 30)

```go
golog.Success("log message")
```

```go
golog.Successf("log message")
```

#### Error (ERROR / 40)

```go
golog.Error("log message")
```

```go
golog.Errorf("log message")
```

> Exits with status 1

### Log to file

```go
golog.LogToFile("filename.log")
```

> Add this declaration before logging anything

### Request Logging

```go
http.ListenAndServe("127.0.0.1:8080", golog.LogRequest(http.DefaultServeMux))
```

> Works as log handler

### Log format

To set log format to - `<date> <time> <log>`

```go
golog.SetLogFormat()
```

To unset log format to - `<log>`

```go
golog.UnsetLogFormat()
```
