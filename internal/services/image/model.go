// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image

import (
	"bytes"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ImageResultEnvelope struct {
	Result ImageModel `json:"result"`
}

type ImageModel struct {
	ID                types.String                   `tfsdk:"id" json:"id,required"`
	AccountID         types.String                   `tfsdk:"account_id" path:"account_id,required"`
	File              types.String                   `tfsdk:"file" json:"file,optional,no_refresh"`
	URL               types.String                   `tfsdk:"url" json:"url,optional,no_refresh"`
	Creator           types.String                   `tfsdk:"creator" json:"creator,optional"`
	Metadata          jsontypes.Normalized           `tfsdk:"metadata" json:"metadata,optional,no_refresh"`
	RequireSignedURLs types.Bool                     `tfsdk:"require_signed_urls" json:"requireSignedURLs,computed_optional"`
	Filename          types.String                   `tfsdk:"filename" json:"filename,computed"`
	Uploaded          timetypes.RFC3339              `tfsdk:"uploaded" json:"uploaded,computed" format:"date-time"`
	Variants          customfield.List[types.String] `tfsdk:"variants" json:"variants,computed"`
	Meta              jsontypes.Normalized           `tfsdk:"meta" json:"meta,computed"`
}

func (r ImageModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
