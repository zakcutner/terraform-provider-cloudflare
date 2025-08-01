// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_integration

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDevicePostureIntegrationResultDataSourceEnvelope struct {
	Result ZeroTrustDevicePostureIntegrationDataSourceModel `json:"result,computed"`
}

type ZeroTrustDevicePostureIntegrationDataSourceModel struct {
	ID            types.String                                                                     `tfsdk:"id" path:"integration_id,computed"`
	IntegrationID types.String                                                                     `tfsdk:"integration_id" path:"integration_id,optional"`
	AccountID     types.String                                                                     `tfsdk:"account_id" path:"account_id,required"`
	Interval      types.String                                                                     `tfsdk:"interval" json:"interval,computed"`
	Name          types.String                                                                     `tfsdk:"name" json:"name,computed"`
	Type          types.String                                                                     `tfsdk:"type" json:"type,computed"`
	Config        customfield.NestedObject[ZeroTrustDevicePostureIntegrationConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
}

func (m *ZeroTrustDevicePostureIntegrationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DevicePostureIntegrationGetParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePostureIntegrationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDevicePostureIntegrationConfigDataSourceModel struct {
	APIURL   types.String `tfsdk:"api_url" json:"api_url,computed"`
	AuthURL  types.String `tfsdk:"auth_url" json:"auth_url,computed"`
	ClientID types.String `tfsdk:"client_id" json:"client_id,computed"`
}
