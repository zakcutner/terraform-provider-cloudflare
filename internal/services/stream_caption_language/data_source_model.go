// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_caption_language

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/stream"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamCaptionLanguageResultDataSourceEnvelope struct {
	Result StreamCaptionLanguageDataSourceModel `json:"result,computed"`
}

type StreamCaptionLanguageDataSourceModel struct {
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	Identifier types.String `tfsdk:"identifier" path:"identifier,required"`
	Language   types.String `tfsdk:"language" path:"language,required"`
	Generated  types.Bool   `tfsdk:"generated" json:"generated,computed"`
	Label      types.String `tfsdk:"label" json:"label,computed"`
	Status     types.String `tfsdk:"status" json:"status,computed"`
}

func (m *StreamCaptionLanguageDataSourceModel) toReadParams(_ context.Context) (params stream.CaptionLanguageGetParams, diags diag.Diagnostics) {
	params = stream.CaptionLanguageGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
