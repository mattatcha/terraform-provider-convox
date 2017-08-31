# convox_syslog

The `convox_syslog` resource defines a Convox Syslog Resource.
(Convox uses the noun "Resource" as well. So, this is a Terraform Resource describing a Convox Resource.)

## Example Usage

```
resource "convox_syslog" "papertrail" {
    name     = "test-app.syslog"
    rack     = "rack/name-here"
    private  = true
    hostname = "logs10000.papertrailapp.com"
    port     = "30303"
    scheme   = "tcp+tls"
}
```

## Argument Reference

* `name` - (Required) The name for reference of the convox syslog resource
* `rack` - (Required) The rack in which the resource exists.
* `private` - (Optional) Create in private subnets (defaults to false)
* `hostname` - (Required) The hostname for which the convox resource syslog will be created with.
* `port` - (Required) The port for which the convox resource syslog will be created with.
* `scheme` - (Required) The scheme to be used with the convox resource syslog. This must either by `tcp` | `tcp+tls`

## Usage

To use a Resource, you need to link it to an app. Use the `convox_resource_link` resource to do so.
