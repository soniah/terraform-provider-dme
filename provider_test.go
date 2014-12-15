package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"os"
	"testing"
)

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestEnvCheck(t *testing.T) {
	if v := os.Getenv("DME_SKEY"); v == "" {
		t.Fatal("DME_SKEY must be set for acceptance tests")
	}

	if v := os.Getenv("DME_AKEY"); v == "" {
		t.Fatal("DME_AKEY must be set for acceptance tests")
	}

	if v := os.Getenv("DME_USESANDBOX"); v == "" {
		t.Fatal("DME_USESANDBOX must be set for acceptance tests. Use the strings 'true' or 'false'.")
	}
}
