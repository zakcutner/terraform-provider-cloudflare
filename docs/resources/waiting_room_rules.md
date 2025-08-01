---
page_title: "cloudflare_waiting_room_rules Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_waiting_room_rules (Resource)



## Example Usage

```terraform
resource "cloudflare_waiting_room_rules" "example_waiting_room_rules" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  waiting_room_id = "699d98642c564d2e855e9661899b7252"
  rules = [{
    action = "bypass_waiting_room"
    expression = "ip.src in {10.20.30.40}"
    description = "allow all traffic from 10.20.30.40"
    enabled = true
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `rules` (Attributes List) (see [below for nested schema](#nestedatt--rules))
- `waiting_room_id` (String)
- `zone_id` (String) Identifier.

### Read-Only

- `id` (String) The ID of the rule.

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Required:

- `action` (String) The action to take when the expression matches.
Available values: "bypass_waiting_room".
- `expression` (String) Criteria defining when there is a match for the current rule.

Optional:

- `description` (String) The description of the rule.
- `enabled` (Boolean) When set to true, the rule is enabled.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_waiting_room_rules.example '<zone_id>/<waiting_room_id>'
```
