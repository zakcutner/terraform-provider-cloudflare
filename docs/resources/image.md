---
page_title: "cloudflare_image Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_image (Resource)



## Example Usage

```terraform
resource "cloudflare_image" "example_image" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  id = "id"
  creator = "creator"
  file = null
  metadata = {

  }
  require_signed_urls = true
  url = "https://example.com/path/to/logo.png"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Account identifier tag.
- `id` (String) An optional custom unique identifier for your image.

### Optional

- `creator` (String) Can set the creator field with an internal user ID.
- `file` (String) An image binary data. Only needed when type is uploading a file.
- `metadata` (String) User modifiable key-value store. Can use used for keeping references to another system of record for managing images.
- `require_signed_urls` (Boolean) Indicates whether the image requires a signature token for the access.
- `url` (String) A URL to fetch an image from origin. Only needed when type is uploading from a URL.

### Read-Only

- `filename` (String) Image file name.
- `meta` (String) User modifiable key-value store. Can be used for keeping references to another system of record for managing images. Metadata must not exceed 1024 bytes.
- `uploaded` (String) When the media item was uploaded.
- `variants` (List of String) Object specifying available variants for an image.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_image.example '<account_id>/<image_id>'
```
