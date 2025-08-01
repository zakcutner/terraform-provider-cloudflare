resource "cloudflare_zero_trust_access_application" "%[1]s" {
  name       = "%[1]s"
  account_id = "%[3]s"
  domain     = "%[1]s.%[2]s"
  type       = "self_hosted"
}

resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  name       = "%[1]s"
  account_id = "%[3]s"
  decision   = "allow"
  include = [
    {
      okta = {
        name                 = "jacob-group"
        identity_provider_id = "225934dc-14e4-4f55-87be-f5d798d23f91"
      }
    }
  ]
  approval_required              = "false"
  purpose_justification_required = "false"
  isolation_required             = "false"
}
