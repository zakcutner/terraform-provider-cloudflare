// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tag

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessTagResultDataSourceEnvelope struct {
	Result ZeroTrustAccessTagDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessTagDataSourceModel struct {
	ID        types.String      `tfsdk:"id" path:"tag_name,computed"`
	TagName   types.String      `tfsdk:"tag_name" path:"tag_name,optional"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	AppCount  types.Int64       `tfsdk:"app_count" json:"app_count,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m *ZeroTrustAccessTagDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessTagGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessTagGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
