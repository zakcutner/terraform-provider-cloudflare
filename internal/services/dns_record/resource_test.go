package dns_record_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"

	cfold "github.com/cloudflare/cloudflare-go"
	cloudflare "github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func init() {
	resource.AddTestSweepers("cloudflare_dns_record", &resource.Sweeper{
		Name: "cloudflare_dns_record",
		F:    testSweepCloudflareRecord,
	})
}

func testSweepCloudflareRecord(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	// Clean up the account level rulesets
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	records, _, err := client.ListDNSRecords(context.Background(), cfold.ZoneIdentifier(zoneID), cfold.ListDNSRecordsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare DNS records: %s", err))
	}

	if len(records) == 0 {
		log.Print("[DEBUG] No Cloudflare DNS records to sweep")
		return nil
	}

	for _, record := range records {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare DNS record ID: %s", record.ID))
		//nolint:errcheck
		client.DeleteDNSRecord(context.Background(), cfold.ZoneIdentifier(zoneID), record.ID)
	}

	return nil
}

func TestAccCloudflareRecord_Basic(t *testing.T) {
	//t.Parallel()
	var record cfold.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, "tf-acctest-basic", rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("tf-acctest-basic.%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "content", "192.168.0.10"),
					resource.TestMatchResourceAttr(resourceName, consts.ZoneIDSchemaKey, regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "tag1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1", "tag2"),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_Apex(t *testing.T) {
	//t.Parallel()
	var record cfold.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigApex(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
					resource.TestCheckResourceAttr(resourceName, "name", domain),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "content", "192.168.0.10"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_LOC(t *testing.T) {
	//t.Parallel()
	var record cfold.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigLOC(zoneID, "tf-acctest-loc."+domain, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "content", "37 46 46.000 N 122 23 35.000 W 0.00 100.00 0.00 0.00"),
					resource.TestCheckResourceAttr(resourceName, "proxiable", "false"),
					resource.TestCheckResourceAttr(resourceName, "data.lat_degrees", "37"),
					resource.TestCheckResourceAttr(resourceName, "data.lat_degrees", "37"),
					resource.TestCheckResourceAttr(resourceName, "data.lat_minutes", "46"),
					resource.TestCheckResourceAttr(resourceName, "data.lat_seconds", "46"),
					resource.TestCheckResourceAttr(resourceName, "data.lat_direction", "N"),
					resource.TestCheckResourceAttr(resourceName, "data.long_degrees", "122"),
					resource.TestCheckResourceAttr(resourceName, "data.long_minutes", "23"),
					resource.TestCheckResourceAttr(resourceName, "data.long_seconds", "35"),
					resource.TestCheckResourceAttr(resourceName, "data.long_direction", "W"),
					resource.TestCheckResourceAttr(resourceName, "data.altitude", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.size", "100"),
					resource.TestCheckResourceAttr(resourceName, "data.precision_horz", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.precision_vert", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_SRV(t *testing.T) {
	//t.Parallel()
	var record cfold.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigSRV(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("_xmpp-client._tcp.%s.%s", rnd, domain)),
					resource.TestCheckResourceAttr(resourceName, "content", "0 5222 talk.l.google.com"),
					resource.TestCheckResourceAttr(resourceName, "proxiable", "false"),
					resource.TestCheckResourceAttr(resourceName, "data.priority", "5"),
					resource.TestCheckResourceAttr(resourceName, "data.weight", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.port", "5222"),
					resource.TestCheckResourceAttr(resourceName, "data.target", "talk.l.google.com"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_CAA(t *testing.T) {
	//t.Parallel()
	var record cfold.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigCAA(rnd, zoneID, fmt.Sprintf("tf-acctest-caa.%s", domain), 600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "data.flags", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.tag", "issue"),
					resource.TestCheckResourceAttr(resourceName, "data.value", "letsencrypt.org"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigCAA(rnd, zoneID, fmt.Sprintf("tf-acctest-caa.%s", domain), 300),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "data.flags", "0"),
					resource.TestCheckResourceAttr(resourceName, "data.tag", "issue"),
					resource.TestCheckResourceAttr(resourceName, "data.value", "letsencrypt.org"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_Proxied(t *testing.T) {
	//t.Parallel()
	var record cfold.DNSRecord
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigProxied(zoneID, domain, "tf-acctest-proxied", rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(resourceName, "proxiable", "true"),
					resource.TestCheckResourceAttr(resourceName, "proxied", "true"),
					resource.TestCheckResourceAttr(resourceName, "type", "CNAME"),
					resource.TestCheckResourceAttr(resourceName, "content", domain),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_Updated(t *testing.T) {
	//t.Parallel()
	var record cfold.DNSRecord
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	recordName := "tf-acctest-update"
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, recordName, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributes(&record),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigNewValue(zoneID, recordName, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &record),
					testAccCheckCloudflareRecordAttributesUpdated(&record),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_typeForceNewRecord(t *testing.T) {
	//t.Parallel()
	var afterCreate, afterUpdate cfold.DNSRecord
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	recordName := "tf-acctest-type-force-new"
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, recordName, rnd, zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &afterCreate),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigChangeType(zoneID, recordName, zoneName, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRecordExists(resourceName, &afterUpdate),
					testAccCheckCloudflareRecordRecreated(&afterCreate, &afterUpdate),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_TtlValidation(t *testing.T) {
	//t.Parallel()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	recordName := "tf-acctest-ttl-validation"
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareRecordConfigTtlValidation(zoneID, recordName, zoneName, rnd),
				ExpectError: regexp.MustCompile("ttl must be set to 1 when `proxied` is true"),
			},
		},
	})
}

func TestAccCloudflareRecord_ExplicitProxiedFalse(t *testing.T) {
	//t.Parallel()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_dns_record." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigExplicitProxied(zoneID, rnd, zoneName, "false", "300"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "proxied", "false"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "300"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigExplicitProxied(zoneID, rnd, zoneName, "true", "1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "proxied", "true"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "1"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigExplicitProxied(zoneID, rnd, zoneName, "false", "300"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "proxied", "false"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "300"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_MXWithPriorityZero(t *testing.T) {
	//t.Parallel()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_dns_record." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigMXWithPriorityZero(zoneID, rnd, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "priority", "0"),
					resource.TestCheckResourceAttr(resourceName, "content", "mail.terraform.cfapi.net"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_HTTPS(t *testing.T) {
	//t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigHTTPS(zoneID, rnd, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "data.priority", "1"),
					resource.TestCheckResourceAttr(name, "data.target", "."),
					resource.TestCheckResourceAttr(name, "data.value", `alpn="h2"`),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_SVCB(t *testing.T) {
	//t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigSVCB(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "data.priority", "2"),
					resource.TestCheckResourceAttr(name, "data.target", "foo."),
					resource.TestCheckResourceAttr(name, "data.value", `alpn="h3,h2"`),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_MXNull(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordNullMX(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "content", "."),
					resource.TestCheckResourceAttr(name, "priority", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_DNSKEY(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending automating setup from https://developers.cloudflare.com/dns/dnssec/multi-signer-dnssec/.")

	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordDNSKEY(zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", domain),
					resource.TestCheckResourceAttr(name, "type", "DNSKEY"),
					resource.TestCheckResourceAttr(name, "data.flags", "257"),
					resource.TestCheckResourceAttr(name, "data.protocol", "13"),
					resource.TestCheckResourceAttr(name, "data.algorithm", "2"),
					resource.TestCheckResourceAttr(name, "data.public_key", "mdsswUyr3DPW132mOi8V9xESWE8jTo0dxCjjnopKl+GqJxpVXckHAeF+KkxLbxILfDLUT0rAK9iUzy1L53eKGQ=="),
				),
			},
		},
	})
}

func TestAccCloudflareRecord_ClearTags(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending investigation into clearing tags.")

	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_dns_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigMultipleTags(zoneID, rnd, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "tags.#", "2"),
					resource.TestCheckResourceAttr(name, "comment", "this is a comment"),
				),
			},
			{
				Config: testAccCheckCloudflareRecordConfigNoTags(zoneID, rnd, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "tags.#", "0"),
					resource.TestCheckNoResourceAttr(name, "comment"),
				),
			},
		},
	})
}

func TestSuppressTrailingDots(t *testing.T) {
	t.Parallel()

	cases := []struct {
		old      string
		new      string
		expected bool
	}{
		{"", "", true},
		{"", "example.com", false},
		{"", "example.com.", false},
		{"", ".", false}, // single dot is used for Null MX record
		{"example.com", "example.com", true},
		{"example.com", "example.com.", true},
		{"example.com", "sub.example.com", false},
		{"sub.example.com", "sub.example.com.", true},
		{".", ".", true},
	}

	for _, c := range cases {
		got := suppressTrailingDots("", c.old, c.new, nil)
		assert.Equal(t, c.expected, got)
	}
}

func testAccCheckCloudflareRecordRecreated(before, after *cfold.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("expected change of Record Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func testAccCheckCloudflareRecordDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_dns_record" {
			continue
		}

		_, err := client.DNS.Records.Get(context.Background(), rs.Primary.ID, dns.RecordGetParams{
			ZoneID: cloudflare.F(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
		})
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			if apierr.StatusCode != 404 {
				return fmt.Errorf("Record still exists")
			}
		}
	}

	return nil
}

func testAccManuallyDeleteRecord(record *cfold.DNSRecord, zoneID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acctest.SharedClient()
		_, err := client.DNS.Records.Delete(context.Background(), record.ID, dns.RecordDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareRecordAttributes(record *cfold.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if record.Content != "192.168.0.10" {
			return fmt.Errorf("bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudflareRecordAttributesUpdated(record *cfold.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if record.Content != "192.168.0.11" {
			return fmt.Errorf("bad content: %s", record.Content)
		}

		return nil
	}
}

func testAccCheckCloudflareRecordExists(n string, record *cfold.DNSRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundRecord, err := client.GetDNSRecord(context.Background(), cfold.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundRecord.ID != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*record = foundRecord

		return nil
	}
}

func testAccCheckCloudflareRecordConfigBasic(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfigbasic.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordConfigApex(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfigapex.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareRecordConfigLOC(zoneID, name, rnd string) string {
	return acctest.LoadTestCase("recordconfigloc.tf", zoneID, name, rnd)
}

func testAccCheckCloudflareRecordConfigSRV(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfigsrv.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareRecordConfigCAA(resourceName, zoneID, name string, ttl int) string {
	return acctest.LoadTestCase("recordconfigcaa.tf", resourceName, zoneID, name, ttl)
}

func testAccCheckCloudflareRecordConfigProxied(zoneID, domain, name, rnd string) string {
	return acctest.LoadTestCase("recordconfigproxied.tf", zoneID, domain, name, rnd)
}

func testAccCheckCloudflareRecordConfigNewValue(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfignewvalue.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordConfigChangeType(zoneID, name, zoneName, rnd string) string {
	return acctest.LoadTestCase("recordconfigchangetype.tf", zoneID, name, zoneName, rnd)
}

func testAccCheckCloudflareRecordConfigChangeHostname(zoneID, name, rnd string) string {
	return acctest.LoadTestCase("recordconfigchangehostname.tf", zoneID, name, rnd)
}

func testAccCheckCloudflareRecordConfigTtlValidation(zoneID, name, zoneName, rnd string) string {
	return acctest.LoadTestCase("recordconfigttlvalidation.tf", zoneID, name, zoneName, rnd)
}

func testAccCheckCloudflareRecordConfigExplicitProxied(zoneID, name, zoneName, proxied, ttl string) string {
	return acctest.LoadTestCase("recordconfigexplicitproxied.tf", zoneID, name, zoneName, proxied, ttl)
}

func testAccCheckCloudflareRecordConfigMXWithPriorityZero(zoneID, name, zoneName string) string {
	return acctest.LoadTestCase("recordconfigmxwithpriorityzero.tf", zoneID, name, zoneName)
}

func testAccCheckCloudflareRecordConfigHTTPS(zoneID, rnd, zoneName string) string {
	return acctest.LoadTestCase("recordconfighttps.tf", zoneID, rnd, zoneName)
}

func testAccCheckCloudflareRecordConfigSVCB(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfigsvcb.tf", zoneID, rnd, domain)
}

func testAccCheckCloudflareRecordNullMX(zoneID, rnd, domain string) string {
	return acctest.LoadTestCase("recordnullmx.tf", rnd, zoneID, domain)
}

func testAccCheckCloudflareRecordConfigMultipleTags(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfigmultipletags.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordConfigNoTags(zoneID, name, rnd, domain string) string {
	return acctest.LoadTestCase("recordconfignotags.tf", zoneID, name, rnd, domain)
}

func testAccCheckCloudflareRecordDNSKEY(zoneID, name string) string {
	return acctest.LoadTestCase("recorddnskey.tf", zoneID, name)
}

func suppressTrailingDots(k, old, new string, d *schema.ResourceData) bool {
	newTrimmed := strings.TrimSuffix(new, ".")

	// Ensure to distinguish values consists of dots only.
	if newTrimmed == "" {
		return old == new
	}

	return strings.TrimSuffix(old, ".") == newTrimmed
}
