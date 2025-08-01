// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/images"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ImagesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[ImagesResultDataSourceModel] `json:"items,computed"`
}

type ImagesDataSourceModel struct {
	AccountID types.String                                              `tfsdk:"account_id" path:"account_id,required"`
	Creator   types.String                                              `tfsdk:"creator" query:"creator,optional"`
	MaxItems  types.Int64                                               `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ImagesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ImagesDataSourceModel) toListParams(_ context.Context) (params images.V1ListParams, diags diag.Diagnostics) {
	params = images.V1ListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Creator.IsNull() {
		params.Creator = cloudflare.F(m.Creator.ValueString())
	}

	return
}

type ImagesResultDataSourceModel struct {
	Images customfield.NestedObjectList[ImagesImagesDataSourceModel] `tfsdk:"images" json:"images,computed"`
}

type ImagesImagesDataSourceModel struct {
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	Creator           types.String                   `tfsdk:"creator" json:"creator,computed"`
	Filename          types.String                   `tfsdk:"filename" json:"filename,computed"`
	Meta              jsontypes.Normalized           `tfsdk:"meta" json:"meta,computed"`
	RequireSignedURLs types.Bool                     `tfsdk:"require_signed_urls" json:"requireSignedURLs,computed"`
	Uploaded          timetypes.RFC3339              `tfsdk:"uploaded" json:"uploaded,computed" format:"date-time"`
	Variants          customfield.List[types.String] `tfsdk:"variants" json:"variants,computed"`
}
