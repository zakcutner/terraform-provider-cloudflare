---
page_title: "cloudflare_argo_smart_routing Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_argo_smart_routing (Data Source)



## Example Usage

```terraform
data "cloudflare_argo_smart_routing" "example_argo_smart_routing" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Specifies the zone associated with the API call.


