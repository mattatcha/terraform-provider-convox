# Convox Provider

The Convox provider is used to interact with Convox-managed apps.

## Example Usage

```
provider "convox" {
    host     = "console.convox.com"
        password = "${var.convox_password}"
}
```

## Argument Reference

* `host` -  The convox host to use. If omitted `host` will default to the first nonempty value in: `CONVOX_RACK` environment variable, `config_path/hostfile`, DefaultHost
* `password` - The password to access the host. If omitted the `password` will use the value in the `CONVOX_PASSWORD` environment variable.
* `config_path` - The convox config path. If omitted, will first look in the `CONVOX_CONFIG` environment variable, then the `$HOME/.convox` if it exists.

## Attribute Reference

Provides no atrributes.
