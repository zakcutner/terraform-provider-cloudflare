// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/keyless_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KeylessCertificateResultDataSourceEnvelope struct {
	Result KeylessCertificateDataSourceModel `json:"result,computed"`
}

type KeylessCertificateDataSourceModel struct {
	ID                   types.String                                                      `tfsdk:"id" path:"keyless_certificate_id,computed"`
	KeylessCertificateID types.String                                                      `tfsdk:"keyless_certificate_id" path:"keyless_certificate_id,optional"`
	ZoneID               types.String                                                      `tfsdk:"zone_id" path:"zone_id,required"`
	CreatedOn            timetypes.RFC3339                                                 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Enabled              types.Bool                                                        `tfsdk:"enabled" json:"enabled,computed"`
	Host                 types.String                                                      `tfsdk:"host" json:"host,computed"`
	ModifiedOn           timetypes.RFC3339                                                 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name                 types.String                                                      `tfsdk:"name" json:"name,computed"`
	Port                 types.Float64                                                     `tfsdk:"port" json:"port,computed"`
	Status               types.String                                                      `tfsdk:"status" json:"status,computed"`
	Permissions          customfield.List[types.String]                                    `tfsdk:"permissions" json:"permissions,computed"`
	Tunnel               customfield.NestedObject[KeylessCertificateTunnelDataSourceModel] `tfsdk:"tunnel" json:"tunnel,computed"`
}

func (m *KeylessCertificateDataSourceModel) toReadParams(_ context.Context) (params keyless_certificates.KeylessCertificateGetParams, diags diag.Diagnostics) {
	params = keyless_certificates.KeylessCertificateGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type KeylessCertificateTunnelDataSourceModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}
