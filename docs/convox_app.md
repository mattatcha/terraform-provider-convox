# convox_app

The `convox_app` resource defines a Convox App and provides access to its endpoints.

## Example Usage

```
resource "convox_app" "test-app" {
    name = "test-app"
    rack = "foo"

    params {
        Internal = false
    }

    environment {
        Foo    = "bar"
            FooBar = 10
    }
}
```

## Argument Reference

* `name` - (Required) The name of the app.
* `rack` - (Required) The rack in which the app exists.
* `environment` - (Optional) A map of the environment variables for the app.
* `params` - (Optional) A map of the app's parameters. See: [parameters](https://convox.com/docs/app-parameters/#setting-parameters)

## Attribute Reference

* `status` - The current status of the app.
* `balancer_endpoint` - UNKNOWN
* `formation` - List of processes in the app.

### `formation` processes

* `name` - The name of the process.
* `balancer` - The domain name of the load balancer endpoint for the process.
* `cpu` - The amount of CPU shares allocated to the process.
* `memory` - The amount of memory allocated to the process.
* `count` - The number of instances the process is scaled to.
* `ports` - List of integers. The ports exposed by the process.
