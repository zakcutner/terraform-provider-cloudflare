// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/page_shield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldPolicyResultDataSourceEnvelope struct {
	Result PageShieldPolicyDataSourceModel `json:"result,computed"`
}

type PageShieldPolicyDataSourceModel struct {
	ID          types.String `tfsdk:"id" path:"policy_id,computed"`
	PolicyID    types.String `tfsdk:"policy_id" path:"policy_id,optional"`
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Action      types.String `tfsdk:"action" json:"action,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Expression  types.String `tfsdk:"expression" json:"expression,computed"`
	Value       types.String `tfsdk:"value" json:"value,computed"`
}

func (m *PageShieldPolicyDataSourceModel) toReadParams(_ context.Context) (params page_shield.PolicyGetParams, diags diag.Diagnostics) {
	params = page_shield.PolicyGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
