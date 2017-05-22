# convox_resource_link

The `convox_resource_link` defines the connection between a Convox Resource and a Convox App.

## Example Usage

```
resource "convox_resource_link" "myapp_papertrail" {
    rack          = "rack/name-here"
    app_name      = "myapp"
    resource_name = "papertrail-syslog"
}
```

## Argument Reference

* `rack` - (Required) The rack in which the resource exists.
* `app_name` - (Required) The name of the Convox App being linked.
* `resource_name` - (Required) The name of the Convox Resource being linked.
