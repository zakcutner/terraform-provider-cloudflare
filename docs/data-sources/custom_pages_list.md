---
page_title: "cloudflare_custom_pages_list Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_custom_pages_list (Data Source)



## Example Usage

```terraform
data "cloudflare_custom_pages_list" "example_custom_pages_list" {
  account_id = "account_id"
  zone_id = "zone_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `max_items` (Number) Max items to fetch, default: 1000
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `created_on` (String)
- `description` (String)
- `id` (String)
- `modified_on` (String)
- `preview_target` (String)
- `required_tokens` (List of String)
- `state` (String) The custom page state.
Available values: "default", "customized".
- `url` (String) The URL associated with the custom page.


