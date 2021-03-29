# golog

Efficient logging for Go

## Usage

### Importing

```go
import "githib.com/sharma1612harshit/golog"
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
4. ERROR | 40

> Note: The above logging levels are in decreasing order of their precedence. i.e. DEBUG / 10 will allow all loggers to log, while ERROR / 40 wont allow any other than itself.
