
# Install
Run `make install` from within the `terraform-provider-convox` directory. The convox provider should automatically be discovered by `terraform`


# Example
    provider "convox" {
      rack {
        name     = "foo"
        hostname = "foo.bar.com"
        password = "xxx"
      }
    }

    resource "convox_app" "test-app" {
      name = "test-app"
      rack = "foo"

      params {
        Internal = false
      }

      environment {
        Foo = "bar"
        FooBar = 10
      }
    }
