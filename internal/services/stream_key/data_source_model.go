// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_key

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/stream"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamKeyResultDataSourceEnvelope struct {
	Result StreamKeyDataSourceModel `json:"result,computed"`
}

type StreamKeyDataSourceModel struct {
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	ID        types.String      `tfsdk:"id" json:"id,computed"`
}

func (m *StreamKeyDataSourceModel) toReadParams(_ context.Context) (params stream.KeyGetParams, diags diag.Diagnostics) {
	params = stream.KeyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
