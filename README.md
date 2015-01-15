# terraform-provider-dme

[![Build
Status](https://travis-ci.org/soniah/terraform-provider-dme.svg?branch=master)](https://travis-ci.org/soniah/terraform-provider-dme)
[![GoDoc](https://godoc.org/github.com/soniah/terraform-provider-dme?status.png)](http://godoc.org/github.com/soniah/terraform-provider-dme)
https://github.com/soniah/terraform-provider-dme v0.3

A [DNSMadeEasy](http://www.dnsmadeeasy.com/) provider for [Terraform](https://github.com/hashicorp/terraform).

This is an early release and the software should be considered 'just
working':

* only A, CName records are handled
* sparse documentation and more tests required

Sonia Hamilton sonia@snowfrog.net http://blog.snowfrog.net

![Terraform](https://raw.githubusercontent.com/hashicorp/terraform/master/website/source/assets/images/readme.png)

## Installation

* Travis shows that this project builds on Go 1.2, 1.3, but not 1.0 or 1.4

* install the [Terraform](https://github.com/hashicorp/terraform)
  development environment, build it using the **Development Environment**
  instructions

* install this project via `go get github.com/soniah/terraform-provider-dme`

* follow the instructions for **Installing a Plugin** detailed in
  [Plugin Basics](https://www.terraform.io/docs/plugins/basics.html)

* to run the tests:

```shell
% export TF_ACC=1
% export DME_AKEY=aaaaaa1a-11a1-1aa1-a101-11a1a11aa1a
% export DME_SKEY=11a0a11a-a1a1-111a-a11a-a11110a11111
% export DME_DOMAINID=123456
% export DME_USESANDBOX='true'
% go test -v
```

* wireshark filter:

```shell
host 208.94.147.116
```

## Usage

Here is an example `test.tf` file:

```go
provider "dme" {
  akey = "aaaaaa1a-11a1-1aa1-a101-11a1a11aa1aa"
  skey = "11a0a11a-a1a1-111a-a11a-a11110a11111"
  usesandbox = true
}

resource "dme_record" "testa" {
  domainid = "123456"
  name = "testa"
  type = "A"
  value = "1.1.1.1"
  ttl = 1000
}

resource "dme_record" "testcname" {
  domainid = "123456"
  name = "testcname"
  type = "CNAME"
  value = "foo"
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
