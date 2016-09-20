

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
        internal = false
      }

      environment {
        key   = "foo"
        value = "bar"
      }

      environment {
        key   = "foofoo"
        value = "barbar"
      }
    }
