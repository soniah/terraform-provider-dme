# terraform-provider-dme

**26-Feb-2015** _This has now been merged into the Terraform development mainline; this project is kept here for reference. Until it is merged into Terraform release, build manually using the instructions below._

[![Build
Status](https://travis-ci.org/soniah/terraform-provider-dme.svg?branch=master)](https://travis-ci.org/soniah/terraform-provider-dme)
[![GoDoc](https://godoc.org/github.com/soniah/terraform-provider-dme?status.png)](http://godoc.org/github.com/soniah/terraform-provider-dme)
https://github.com/soniah/terraform-provider-dme v0.7

A [DNSMadeEasy](http://www.dnsmadeeasy.com/) provider for [Terraform](https://github.com/hashicorp/terraform).

This is the first release of the software -- it's fully featured but
more tests may be required. It provides support for adding and removing
records on existing domains; to add/remove domains or for other
DNSMadeEasy functionality you'll need to use the web interface.

Sonia Hamilton sonia@snowfrog.net http://blog.snowfrog.net

![Terraform](https://raw.githubusercontent.com/hashicorp/terraform/master/website/source/assets/images/readme.png)

## Installation

* Terraform requires Go1.4 to build

* install the [Terraform](https://github.com/hashicorp/terraform)
  development environment, build it using the **Development Environment**
  instructions. The dev environment build will place the binaries in
  `$GOPATH/src/github.com/hashicorp/terraform/bin`.  Copy these files to
  somewhere in your $PATH eg `/usr/local/bin`.

* install this project via `go get github.com/soniah/terraform-provider-dme`. This
  will place the `terraform-provider-dme` binary in `$GOPATH/bin`. Copy
  this file to somewhere in your $PATH eg `/usr/local/bin`. (These are
  the instructions for **Installing a Plugin** detailed in
  [Plugin Basics](https://www.terraform.io/docs/plugins/basics.html)

## Usage

Here is an example `test.tf` file. Note that domainid is a
string, and other integer values are integers.

```apache
# Provide your API and Secret Keys, and whether the sandbox
# is being used (defaults to false)
provider "dme" {
  akey = "aaaaaa1a-11a1-1aa1-a101-11a1a11aa1aa"
  skey = "11a0a11a-a1a1-111a-a11a-a11110a11111"
  usesandbox = true
}

# A Record
resource "dme_record" "testa" {
  domainid = "123456"
  name = "testa"
  type = "A"
  value = "1.1.1.1"
  ttl = 1000
}

# CNAME record
resource "dme_record" "testcname" {
  domainid = "123456"
  name = "testcname"
  type = "CNAME"
  value = "foo"
  ttl = 1000
}

# ANAME record
resource "dme_record" "testaname" {
  domainid = "123456"
  name = "testaname"
  type = "ANAME"
  value = "foo"
  ttl = 1000
}

# MX record
resource "dme_record" "testmx" {
  domainid = "123456"
  name = "testmx"
  type = "MX"
  value = "foo"
  mxLevel = 10
  ttl = 1000
}

# HTTPRED
resource "dme_record" "testhttpred" {
  domainid = "123456"
  name = "testhttpred"
  type = "HTTPRED"
  value = "https://github.com/soniah/terraform-provider-dme"
  hardLink = true
  redirectType = "Hidden Frame Masked"
  title = "An Example"
  keywords = "terraform example"
  description = "This is a description"
  ttl = 2000
}

# TXT record
resource "dme_record" "testtxt" {
  domainid = "123456"
  name = "testtxt"
  type = "TXT"
  value = "foo"
  ttl = 1000
}

# SPF record
resource "dme_record" "testspf" {
  domainid = "123456"
  name = "testspf"
  type = "SPF"
  value = "foo"
  ttl = 1000
}

# PTR record
resource "dme_record" "testptr" {
  domainid = "123456"
  name = "testptr"
  type = "PTR"
  value = "foo"
  ttl = 1000
}

# NS record
resource "dme_record" "testns" {
  domainid = "123456"
  name = "testns"
  type = "NS"
  value = "foo"
  ttl = 1000
}

# AAAA record
resource "dme_record" "testaaaa" {
  domainid = "123456"
  name = "testaaaa"
  type = "AAAA"
  value = "FE80::0202:B3FF:FE1E:8329"
  ttl = 1000
}

# SRV record
resource "dme_record" "testsrv" {
  domainid = "123456"
  name = "testsrv"
  type = "SRV"
  value = "foo"
  priority = 10
  weight = 20
  port = 30
  ttl = 1000
}

```

## Tests

To run the tests:

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

## Related Projects

* https://github.com/soniah/dnsmadeeasy
* https://github.com/hashicorp/terraform

## License

See LICENSE. Copyright the terraform-provider-dme authors, see AUTHORS.md.
