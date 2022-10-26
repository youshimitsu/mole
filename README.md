# GET STARTED

```shell
go run main.go -config /path/to/cfg.json -group group_name
```
## Parameters

| Name   | Type     | Required | Example             | Description                                  |
|--------|----------|----------|---------------------|----------------------------------------------|
| config | `string` | `true`   | `/path/to/cfg.json` | path to your config file, ext (JSON, YAML)   |
| group  | `string` | `false`  | `group_name`        | default: `all`, up tunnels by concrete group |

## Config example:

### JSON

```json
{
  "sshConfig": "",
  "logger": {
    "log_level": "debug"
  },
  "hosts": [
    {
      "name": "TunnelName with credentials for host",
      "host": "LOCAL_PORT:REMOTE_IP_OR_DOMAIN:REMOTE_PORT",
      "connection": {
        "login": "user@host",
        "password": "passwd"
      }
    },
    {
      "name": "TunnelName with credentials for group of hosts",
      "host": "LOCAL_PORT:REMOTE_IP_OR_DOMAIN:REMOTE_PORT",
      "group": "group_name"
    }
  ],
  "group_connections": {
    "group_name": {
      "login": "user@host",
      "password": "passwd"
    }
  }
}
```