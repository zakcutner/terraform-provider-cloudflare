package zero_trust_dlp_integration_entry_test

import (
	"errors"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareZeroTrustDlpIntegrationEntry_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_dlp_integration_entry." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccZeroTrustDlpIntegrationEntryConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						return errors.New("test not implemented")
					},
					resource.TestCheckResourceAttr(name, "some_string_attribute", "string_value"),
				),
			},
		},
	})
}

func testAccZeroTrustDlpIntegrationEntryConfig(rnd string) string {
	return acctest.LoadTestCase("basic.tf", rnd)
}
