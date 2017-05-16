# convox_syslog

The `convox_syslog` resource defines a Convox Syslog.

## Example Usage

```
resource "convox_syslog" "papertrail" {
    name     = "test-app.syslog"
    hostname = "logs10000.papertrailapp.com"
    port     = "30303"
    scheme   = "tcp|tcp+tls"
}
```

## Argument Reference

* `name` - (Required) The name for reference of the convox syslog resource.
* `hostname` - (Required) The hostname for which the convox resource syslog will be created with.
* `port` - (Required) The port for which the convox resource syslog will be created with.
* `scheme` - (Required) The scheme to be used with the convox resource syslog. This must either by tcp | tcp+tls