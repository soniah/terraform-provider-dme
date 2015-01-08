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
						"dme_record.test", "name", "test"),
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
  name = "test"
  type = "A"
  value = "1.1.1.1"
  ttl = 2000
}`
