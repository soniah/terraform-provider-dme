# terraform-provider-dme

[![Build
Status](https://travis-ci.org/soniah/terraform-provider-dme.svg?branch=master)](https://travis-ci.org/soniah/terraform-provider-dme)
[![GoDoc](https://godoc.org/github.com/soniah/terraform-provider-dme?status.png)](http://godoc.org/github.com/soniah/terraform-provider-dme)

A [DNSMadeEasy](http://www.dnsmadeeasy.com/) provider for Terraform.

This is an early release and the software should be considered 'barely
working':

* only A records are handled
* almost no tests or documentation
* code needs to be refactored and tidied

Sonia Hamilton, sonia@snowfrog.net, http://blog.snowfrog.net.

![Terraform](https://raw.githubusercontent.com/hashicorp/terraform/master/website/source/assets/images/readme.png)

## Installation

* Travis shows that this project builds on Go 1.2, 1.3, but not 1.0 or 1.4

* install the [Terraform](https://github.com/hashicorp/terraform)
  development environment, build it using the **Development Environment**
  instructions

* install this project via `go get github.com/soniah/terraform-provider-dme`

* follow the instructions for **Installing a Plugin** detailed in
  [Plugin Basics](https://www.terraform.io/docs/plugins/basics.html)

## Usage

Here is an example `test.tf` file:

```go
provider "dme" {
  akey = "aaaaaa1a-11a1-1aa1-a101-11a1a11aa1aa"
  skey = "11a0a11a-a1a1-111a-a11a-a11110a11111"
  usesandbox = true
}

resource "dme_record" "test" {
  domainid = "123456"
  name = "test1"
  type = "A"
  value = "1.1.1.1"
  ttl = 1000
}
```

## Documentation

The full documentation is available on [Godoc](http://godoc.org/github.com/soniah/dnsmadeeasy)

## Related Projects

* https://github.com/soniah/dnsmadeeasy
* https://github.com/hashicorp/terraform

## License

See LICENSE. Copyright the terraform-provider-dme authors, see AUTHORS.md.
