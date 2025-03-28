# honeymarker

[![OSS Lifecycle](https://img.shields.io/osslifecycle/honeycombio/honeymarker?color=success)](https://github.com/honeycombio/home/blob/main/honeycomb-oss-lifecycle-and-practices.md)

`honeymarker` provides a simple CRUD interface for dealing with per-dataset [markers](https://docs.honeycomb.io/working-with-your-data/markers/) on https://honeycomb.io/.

## Installation

```
$ go install github.com/honeycombio/honeymarker@latest
$ honeymarker    # (if $GOPATH/bin is in your path.)
```

;)

## Usage

`$ honeymarker -k <your-configuration-key> -d <dataset> [--version] COMMAND [command-specific flags]`

* `<your-configuration-key>` can be found on `https://ui.honeycomb.io/<team>/environments/<environment>/api_keys`. You can also set it as `HONEYCOMB_API_KEY`
* `<dataset>` is the name of one of the datasets associated with the team whose configuration key you're using. You can use `__all__` as dataset name to work with environment-wide markers.
* `COMMAND` see below

## Available commands:

| Command   | Description            |
| --------- | ---------------------- |
| `add`     | add a new marker       |
| `list`    | list all markers       |
| `rm`      | delete a marker        |
| `update`  | update a marker        |
| `version` | print version and exit |

## Adding markers (`add`)

| Name       | Flag                              | Description                                                      | Required |
| ---------- | --------------------------------- | ---------------------------------------------------------------- | -------- |
| Start Time | `-s <arg>` / `--start_time=<arg>` | start time for the marker in unix time (seconds since the epoch) | No       |
| End Time   | `-e <arg>` / `--end_time=<arg>`   | end time for the marker in unix time (seconds since the epoch)   | No       |
| Message    | `-m <arg>` / `--msg=<arg>`        | message describing this specific marker                          | No       |
| URL        | `-u <arg>` / `--url=<arg>`        | URL associated with this marker                                  | No       |
| Type       | `-t <arg>` / `--type=<arg>`       | identifies marker type                                           | No       |

All parameters to add are optional.

If start_time is missing, the marker will be assigned the current time.

It is highly recommended that you fill in either message or type.
All markers of the same type will be shown with the same color in the UI.
The message will be visible above an individual marker.

If a URL is specified along with a message, the message will be shown
as a link in the UI, and clicking it will take you to the URL.

Example:

```
$ ./honeymarker -k $CONFIGURATION_KEY -d myservice add -t deploy -m "build 192837"
{"created_at":"2017-06-28T18:55:11Z","updated_at":"2017-06-28T18:55:11Z","start_time":1498676111,"message":"build 192837","type":"deploy","id":"n71R3zM1Tn3"}

$
```

## Listing markers (`list`)

| Name            | Flag          | Description                                                              | Required |
| --------------- | ------------- | ------------------------------------------------------------------------ | -------- |
| JSON            | `--json`      | Output the list as json instead of in tabular form                       | No       |
| Unix Timestamps | `--unix_time` | In table mode, format times as unit timestamps (seconds since the epoch) | No       |

Example:
```
$ ./honeymarker -k $CONFIGURATION_KEY -d myservice list
| ID          |         Start Time |           End Time | Type         | Message      | URL |
+-------------+--------------------+--------------------+--------------+--------------+-----+
| n71R3zM1Tn3 | Jun 28 24 11:55:11 |                    | deploy       | build 192837 |     |

$
```

## Updating markers (`update`)

| Name       | Flag                              | Description                                                      | Required |
| ---------- | --------------------------------- | ---------------------------------------------------------------- | -------- |
| Marker ID  | `-i <arg>` / `--id=<arg>`         | ID of the marker to update                                       | Yes      |
| Start Time | `-s <arg>` / `--start_time=<arg>` | start time for the marker in unix time (seconds since the epoch) | No       |
| End Time   | `-e <arg>` / `--end_time=<arg>`   | end time for the marker in unix time (seconds since the epoch)   | No       |
| Message    | `-m <arg>` / `--msg=<arg>`        | message describing this specific marker                          | No       |
| URL        | `-u <arg>` / `--url=<arg>`        | URL associated with this marker                                  | No       |
| Type       | `-t <arg>` / `--type=<arg>`       | identifies marker type                                           | No       |

The marker ID is available from the `list` command, and is also output to the console by the `add` command.

Example:
```
$ ./honeymarker -k $CONFIGURATION_KEY -d myservice update -i n71R3zM1Tn3 -u "http://my.service.co/builds/192837"
{"created_at":"2017-06-28T18:55:11Z","updated_at":"2017-06-28T18:55:11Z","start_time":1498676111,"message":"build 192837","type":"deploy","url":"http://my.service.co/builds/192837","id":"n71R3zM1Tn3"}

$
```

## Deleting markers (`rm`)

| Name      | Flag                      | Description                | Required |
| --------- | ------------------------- | -------------------------- | -------- |
| Marker ID | `-i <arg>` / `--id=<arg>` | ID of the marker to delete | Yes      |

The marker ID is available from the `list` command, and is also output to the console by the `add` command.

Example:
```
$ ./honeymarker -k $CONFIGURATION_KEY -d myservice rm -i n71R3zM1Tn3
{"created_at":"2017-06-28T18:55:11Z","updated_at":"2017-06-28T18:57:47Z","start_time":1498676111,"message":"build 192837","type":"deploy","url":"http://my.service.co/builds/192837","id":"n71R3zM1Tn3"}

$ ./honeymarker -k $CONFIGURATION_KEY -d myservice list
| ID          |         Start Time |           End Time | Type         | Message | URL |
+-------------+--------------------+--------------------+--------------+---------+-----+
$
```

