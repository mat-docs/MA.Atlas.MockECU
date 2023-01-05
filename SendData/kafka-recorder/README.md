[comment]: <> ( Generated 2023-01-05 17:28:57 by go-framework v2.1.4 )

# Kafka recorder

record and playback kafka messages

# Introduction

TODO: Give a short introduction of your project. Let this section explain the objectives, or the motivation behind this
project.

# Getting Started

TODO: Guide users through getting your code up and running on their own system. In this section you can talk about:

1. Installation process
2. Software dependencies
3. Latest releases
4. API references

# Build and Test

TODO: Describe and show how to build your code and run the tests.

# Configuration

[comment]: <> ( SECTION-START: ConfigTable )

| Setting                | Config file          | Cmdline          | Env Var     | Default Val               | Applies to | Description                                                               |
|:-----------------------|:---------------------|:-----------------|:------------|:--------------------------|:-----------|:--------------------------------------------------------------------------|
| File (string)          | File                 | -f, --file       | DATAFILE    |                           | all        | Record file                                                               |
| Brokers ([]string)     | Brokers              | -b, --broker     | BROKERS     |                           | all        | Kafka brokers                                                             |
| Topic (string)         | Topic                | --topic          | TOPIC       |                           | all        | Kafka topic                                                               |
| Timeout (int)          | Timeout              | -t, --timeout    | TIMEOUT     | 30000                     | all        | Timeout (ms)                                                              |
| ConfigFile (string)    | ConfigFile           | --config-file    |             |                           | all        | Alternate config file                                                     |
| Verbose (bool)         | Verbose              | -v, --verbose    | VERBOSE     |                           | all        | Generate more detailed output                                             |
| LogLevel (string)      | Log.Level            | --log-level      | GO_LOGLEVEL | info                      | all        | Level of logging to generate (panic,fatal,error,warning,info,debug,trace) |
| LogWriter (string)     | Log.Writer           | --log-writer     |             | stdout                    | all        | Writer for logger (stdout[:json], stderr[:json], file[:json]=path)        |
| LogTimeFormat (string) | Log.TimeFormat       | --log-timeformat |             | 2006-01-02T15:04:05Z07:00 | all        | Format for timestamp of log entries                                       |
| LogStackTrace (bool)   | Log.EnableStackTrace | --log-stacktrace |             |                           | all        | Log stack trace for errors                                                |

<details>
  <summary>Further details</summary>
Config file ids with a dot can be "scoped"<br>
e.g. "a.b.c" can be added to config file as:<br>

### yaml

```yaml
a:
  b:
    c: "value"
```

### json

```json
{
  "a": {
    "b": {
      "c": "value"
    }
  }
}
```

</details>

[comment]: <> ( SECTION-END )

