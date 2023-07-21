# nxs-support-bot-migrate

This tool helps you to migrate the DB from latest version of [nxs-chat-srv](https://github.com/nixys/nxs-chat-srv) to [nxs-support-bot](github.com/nixys/nxs-support-bot) v1.0.0.

In text below lets use following conventions:
- [nxs-chat-srv](https://github.com/nixys/nxs-chat-srv) DB it's an `old` DB
- [nxs-support-bot](github.com/nixys/nxs-support-bot) DB it's a `new` DB

What this tool do:
- Cleans up a data in `new` DB
- Migrates necessary data from `old` MySQL to `new`
- Migrates  necessary data from `old` Redis to `new`
- While migration all incorrect data from `old` to `new` will be skipped

## Quickstart

### Install

This app is a helper for [nxs-support-bot](github.com/nixys/nxs-support-bot). Please see nxs-support-bot readme for installation.

### Settings

#### General settings

| Option         | Type   | Required | Default value | Description                                                      |
|---             | :---:  | :---:    | :---:         |---                                                               |
| `logfile`      | String | No       | `stdout`      | Log file path. Also you may use `stdout` and `stderr`                                                                                                                |
| `loglevel`     | String | No       | `info`        | Log level. Available values: `debug`, `warn`, `error` and `info` |
| `pidfile`      | String | No       | -             | Pid file path. If `pidfile` is not set it will not be created                                                                                                                 |
| `src`        | [Src](#src-settings) | Yes       | -             | Source (`old`) databases settings                   |
| `dst`        | [Dst](#dst-settings) | Yes       | -             | Destination (`new`) databases settings              |

##### Src settings

| Option         | Type   | Required | Default value | Description |
|---             | :---:  | :---:    | :---:         |---          |
| `mysql` | [MySQL](#mysql-settings) | Yes | - | MySQL settings |
| `redis` | [Redis](#redis-settings) | Yes | - | Redis settings |

##### Dst settings

| Option         | Type   | Required | Default value | Description |
|---             | :---:  | :---:    | :---:         |---          |
| `mysql` | [MySQL](#mysql-settings) | Yes | - | MySQL settings |

##### MySQL settings

| Option     | Type   | Required | Default value | Description               |
|---         | :---:  | :---:    | :---:         |---                        |
| `host`     | String | Yes      | -             | Host to connect     |
| `port`     | Int    | Yes      | -             | Port to connect     |
| `db`       | String | Yes      | -             | DB name to connect  |
| `user`     | String | Yes      | -             | User to connect     |
| `password` | String | Yes      | -             | Password to connect |

##### Redis settings

| Option | Type   | Required | Default value | Description           |
|---     | :---:  | :---:    | :---:         |---                    |
| `host` | String | Yes      | -             | Host to connect |
| `port` | Int    | Yes      | -             | Port to connect |

## Feedback

For support and feedback please contact me:
- telegram: [@borisershov](https://t.me/borisershov)
- e-mail: b.ershov@nixys.ru

## License

nxs-support-bot-migrate is released under the [Apache License 2.0](LICENSE).
