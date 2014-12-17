package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	dme "github.com/soniah/dnsmadeeasy"
)

func resourceDMERecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMERecordCreate,
		Read:   resourceDMERecordRead,
		Update: resourceDMERecordUpdate,
		Delete: resourceDMERecordDelete,

		Schema: map[string]*schema.Schema{
			// Use recordid for TF ID.
			"domainid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			/*
				"source": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"sourceid": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"dynamicdns": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"password": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"monitor": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"failover": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"failed": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"gtdlocation": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"description": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"keywords": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"title": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"hardlink": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"mxlevel": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"weight": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"priority": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"port": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
			*/
		},
	}
}

func resourceDMERecordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*dme.Client)

	domainid := d.Get("domainid").(string)
	log.Printf("[INFO] Creating record for domainid: %s", domainid)

	cr := map[string]interface{}{
		"name":  d.Get("name").(string),
		"type":  d.Get("type").(string),
		"value": d.Get("value").(string),
		"ttl":   int64(d.Get("ttl").(int)),
	}
	log.Printf("[DEBUG] record create configuration: %#v", cr)

	result, err := client.CreateRecord(domainid, cr)
	if err != nil {
		return fmt.Errorf("Failed to create record: %s", err)
	}

	d.SetId(result)
	log.Printf("[INFO] record ID: %s", d.Id())

	return resourceDMERecordRead(d, meta)
}

func resourceDMERecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*dme.Client)

	domainid := d.Get("domainid").(string)
	recordid := d.Id()
	log.Printf("[INFO] Reading record for domainid: %s recordid: %s", domainid, recordid)

	rec, err := client.ReadRecord(domainid, recordid)
	if err != nil {
		return fmt.Errorf("Couldn't find record: %s", err)
	}

	d.Set("name", rec.Name)
	d.Set("type", rec.Type)
	d.Set("value", rec.Value)
	d.Set("ttl", rec.TTL)

	return nil
}

func resourceDMERecordUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*dme.Client)

	domainid := d.Get("domainid").(string)
	recordid := d.Id()

	cr := make(map[string]interface{})
	if attr, ok := d.GetOk("name"); ok {
		cr["Name"] = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		cr["Type"] = attr.(string)
	}
	if attr, ok := d.GetOk("value"); ok {
		cr["Value"] = attr.(string)
	}
	if attr, ok := d.GetOk("ttl"); ok {
		cr["TTL"] = int64(attr.(int))
	}
	log.Printf("[DEBUG] record update configuration: %+#v", cr)

	if _, err := client.UpdateRecord(domainid, recordid, cr); err != nil {
		return fmt.Errorf("Error updating record: %s", err)
	}

	return resourceDMERecordRead(d, meta)
}

func resourceDMERecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*dme.Client)

	domainid := d.Get("domainid").(string)
	recordid := d.Id()
	log.Printf("[INFO] Deleting record for domainid: %s recordid: %s", domainid, recordid)

	if err := client.DeleteRecord(domainid, recordid); err != nil {
		return fmt.Errorf("Error deleting record: %s", err)
	}

	return nil
}
