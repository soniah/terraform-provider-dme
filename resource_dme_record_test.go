package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	dme "github.com/soniah/dnsmadeeasy"
)

var _ = fmt.Sprintf("dummy") // dummy
var _ = os.DevNull           // dummy

func TestAccDMERecordA(t *testing.T) {
	var record dme.Record
	domainid := os.Getenv("DME_DOMAINID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMERecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testDMERecordConfigA, domainid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMERecordExists("dme_record.test", &record),
					resource.TestCheckResourceAttr(
						"dme_record.test", "domainid", domainid),
					resource.TestCheckResourceAttr(
						"dme_record.test", "name", "testa"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "type", "A"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "value", "1.1.1.1"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "ttl", "2000"),
				),
			},
		},
	})
}

func TestAccDMERecordCName(t *testing.T) {
	var record dme.Record
	domainid := os.Getenv("DME_DOMAINID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMERecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testDMERecordConfigCName, domainid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMERecordExists("dme_record.test", &record),
					resource.TestCheckResourceAttr(
						"dme_record.test", "domainid", domainid),
					resource.TestCheckResourceAttr(
						"dme_record.test", "name", "testcname"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "type", "CNAME"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "value", "foo"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "ttl", "2000"),
				),
			},
		},
	})
}

/*
func TestAccDMERecordAName(t *testing.T) {
	var record dme.Record
	domainid := os.Getenv("DME_DOMAINID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMERecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testDMERecordConfigAName, domainid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMERecordExists("dme_record.test", &record),
					resource.TestCheckResourceAttr(
						"dme_record.test", "domainid", domainid),
					resource.TestCheckResourceAttr(
						"dme_record.test", "name", "testaname"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "type", "ANAME"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "value", "foo"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "ttl", "2000"),
				),
			},
		},
	})
}
*/

func TestAccDMERecordMX(t *testing.T) {
	var record dme.Record
	domainid := os.Getenv("DME_DOMAINID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMERecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testDMERecordConfigMX, domainid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMERecordExists("dme_record.test", &record),
					resource.TestCheckResourceAttr(
						"dme_record.test", "domainid", domainid),
					resource.TestCheckResourceAttr(
						"dme_record.test", "name", "testmx"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "type", "MX"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "value", "foo"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "mxLevel", "10"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "ttl", "2000"),
				),
			},
		},
	})
}

func TestAccDMERecordTXT(t *testing.T) {
	var record dme.Record
	domainid := os.Getenv("DME_DOMAINID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMERecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testDMERecordConfigTXT, domainid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMERecordExists("dme_record.test", &record),
					resource.TestCheckResourceAttr(
						"dme_record.test", "domainid", domainid),
					resource.TestCheckResourceAttr(
						"dme_record.test", "name", "testtxt"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "type", "TXT"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "value", "\"foo\""),
					resource.TestCheckResourceAttr(
						"dme_record.test", "ttl", "2000"),
				),
			},
		},
	})
}

func TestAccDMERecordSPF(t *testing.T) {
	var record dme.Record
	domainid := os.Getenv("DME_DOMAINID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMERecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testDMERecordConfigSPF, domainid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMERecordExists("dme_record.test", &record),
					resource.TestCheckResourceAttr(
						"dme_record.test", "domainid", domainid),
					resource.TestCheckResourceAttr(
						"dme_record.test", "name", "testspf"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "type", "SPF"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "value", "\"foo\""),
					resource.TestCheckResourceAttr(
						"dme_record.test", "ttl", "2000"),
				),
			},
		},
	})
}

func TestAccDMERecordNS(t *testing.T) {
	var record dme.Record
	domainid := os.Getenv("DME_DOMAINID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMERecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testDMERecordConfigNS, domainid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMERecordExists("dme_record.test", &record),
					resource.TestCheckResourceAttr(
						"dme_record.test", "domainid", domainid),
					resource.TestCheckResourceAttr(
						"dme_record.test", "name", "testns"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "type", "NS"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "value", "foo"),
					resource.TestCheckResourceAttr(
						"dme_record.test", "ttl", "2000"),
				),
			},
		},
	})
}

func testAccCheckDMERecordDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*dme.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dnsmadeeasy_record" {
			continue
		}

		fmt.Printf("rs: %+v")
		_, err := client.ReadRecord(rs.Primary.Attributes["domainid"], rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Record still exists")
		}
	}

	return nil
}

func testAccCheckDMERecordExists(n string, record *dme.Record) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		client := testAccProvider.Meta().(*dme.Client)

		foundRecord, err := client.ReadRecord(rs.Primary.Attributes["domainid"], rs.Primary.ID)

		if err != nil {
			return err
		}

		if foundRecord.StringRecordID() != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*record = *foundRecord

		return nil
	}
}

const testDMERecordConfigA = `
resource "dme_record" "test" {
  domainid = "%s"
  name = "testa"
  type = "A"
  value = "1.1.1.1"
  ttl = 2000
}`

const testDMERecordConfigCName = `
resource "dme_record" "test" {
  domainid = "%s"
  name = "testcname"
  type = "CNAME"
  value = "foo"
  ttl = 2000
}`

const testDMERecordConfigAName = `
resource "dme_record" "test" {
  domainid = "%s"
  name = "testaname"
  type = "ANAME"
  value = "foo"
  ttl = 2000
}`

const testDMERecordConfigMX = `
resource "dme_record" "test" {
  domainid = "%s"
  name = "testmx"
  type = "MX"
  value = "foo"
  mxLevel = 10
  ttl = 2000
}`

const testDMERecordConfigTXT = `
resource "dme_record" "test" {
  domainid = "%s"
  name = "testtxt"
  type = "TXT"
  value = "foo"
  ttl = 2000
}`

const testDMERecordConfigSPF = `
resource "dme_record" "test" {
  domainid = "%s"
  name = "testspf"
  type = "SPF"
  value = "foo"
  ttl = 2000
}`

const testDMERecordConfigNS = `
resource "dme_record" "test" {
  domainid = "%s"
  name = "testns"
  type = "NS"
  value = "foo"
  ttl = 2000
}`
