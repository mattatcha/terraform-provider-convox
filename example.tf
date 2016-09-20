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

/*resource "convox_rack" "dev" {
  name           = "convox"   # custom rack name (default: "convox") [$STACK_NAME]
  ami            = ""         # custom AMI for rack instances
  dedicated      = false      # create EC2 instances on dedicated hardware
  instance_count = 3          # number of instances in the rack (default: 3)
  instance_type  = "t2.small" # type of instances in the rack (default: "t2.small")
  private        = false      # use private subnets and NAT gateways to shield instances

  private_cidrs = [
    "10.0.4.0/24",
    "10.0.5.0/24",
    "10.0.6.0/24",
  ] # private subnet CIDRs (default: "10.0.4.0/24,10.0.5.0/24,10.0.6.0/24")

  region = "us-east-1" # aws region (default: "us-east-1") [$AWS_REGION]

  version  = "latest"      # install a specific version (default: "latest") [$VERSION]
  vpc_cidr = "10.0.0.0/16" # custom VPC CIDR (default: "10.0.0.0/16")

  subnet_cidrs = [
    "10.0.1.0/24",
    "10.0.2.0/24",
    "10.0.3.0/24",
  ] # subnet CIDRs (default: "10.0.1.0/24,10.0.2.0/24,10.0.3.0/24")
}

resource "convox_app" "main" {
  name = "test-app"
  rack = "${convox_rack.dev.name}"

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
}*/

