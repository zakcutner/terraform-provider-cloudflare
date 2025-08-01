// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_gre_tunnel

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANGRETunnelResultEnvelope struct {
	Result MagicWANGRETunnelModel `json:"result"`
}

type MagicWANGRETunnelModel struct {
	ID                    types.String                                                      `tfsdk:"id" json:"id,computed"`
	AccountID             types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	CloudflareGREEndpoint types.String                                                      `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,required,no_refresh"`
	CustomerGREEndpoint   types.String                                                      `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,required,no_refresh"`
	InterfaceAddress      types.String                                                      `tfsdk:"interface_address" json:"interface_address,required,no_refresh"`
	Name                  types.String                                                      `tfsdk:"name" json:"name,required,no_refresh"`
	Description           types.String                                                      `tfsdk:"description" json:"description,optional,no_refresh"`
	InterfaceAddress6     types.String                                                      `tfsdk:"interface_address6" json:"interface_address6,optional,no_refresh"`
	Mtu                   types.Int64                                                       `tfsdk:"mtu" json:"mtu,computed_optional,no_refresh"`
	TTL                   types.Int64                                                       `tfsdk:"ttl" json:"ttl,computed_optional,no_refresh"`
	HealthCheck           customfield.NestedObject[MagicWANGRETunnelHealthCheckModel]       `tfsdk:"health_check" json:"health_check,computed_optional,no_refresh"`
	CreatedOn             timetypes.RFC3339                                                 `tfsdk:"created_on" json:"created_on,computed,no_refresh" format:"date-time"`
	Modified              types.Bool                                                        `tfsdk:"modified" json:"modified,computed,no_refresh"`
	ModifiedOn            timetypes.RFC3339                                                 `tfsdk:"modified_on" json:"modified_on,computed,no_refresh" format:"date-time"`
	GRETunnel             customfield.NestedObject[MagicWANGRETunnelGRETunnelModel]         `tfsdk:"gre_tunnel" json:"gre_tunnel,computed"`
	ModifiedGRETunnel     customfield.NestedObject[MagicWANGRETunnelModifiedGRETunnelModel] `tfsdk:"modified_gre_tunnel" json:"modified_gre_tunnel,computed,no_refresh"`
}

func (m MagicWANGRETunnelModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MagicWANGRETunnelModel) MarshalJSONForUpdate(state MagicWANGRETunnelModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type MagicWANGRETunnelHealthCheckModel struct {
	Direction types.String                                                      `tfsdk:"direction" json:"direction,computed_optional"`
	Enabled   types.Bool                                                        `tfsdk:"enabled" json:"enabled,computed_optional"`
	Rate      types.String                                                      `tfsdk:"rate" json:"rate,computed_optional"`
	Target    customfield.NestedObject[MagicWANGRETunnelHealthCheckTargetModel] `tfsdk:"target" json:"target,computed_optional"`
	Type      types.String                                                      `tfsdk:"type" json:"type,computed_optional"`
}

type MagicWANGRETunnelHealthCheckTargetModel struct {
	Effective types.String `tfsdk:"effective" json:"effective,computed"`
	Saved     types.String `tfsdk:"saved" json:"saved,optional"`
}

type MagicWANGRETunnelGRETunnelModel struct {
	ID                    types.String                                                         `tfsdk:"id" json:"id,computed"`
	CloudflareGREEndpoint types.String                                                         `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,computed"`
	CustomerGREEndpoint   types.String                                                         `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,computed"`
	InterfaceAddress      types.String                                                         `tfsdk:"interface_address" json:"interface_address,computed"`
	Name                  types.String                                                         `tfsdk:"name" json:"name,computed"`
	CreatedOn             timetypes.RFC3339                                                    `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description           types.String                                                         `tfsdk:"description" json:"description,computed"`
	HealthCheck           customfield.NestedObject[MagicWANGRETunnelGRETunnelHealthCheckModel] `tfsdk:"health_check" json:"health_check,computed"`
	InterfaceAddress6     types.String                                                         `tfsdk:"interface_address6" json:"interface_address6,computed"`
	ModifiedOn            timetypes.RFC3339                                                    `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Mtu                   types.Int64                                                          `tfsdk:"mtu" json:"mtu,computed"`
	TTL                   types.Int64                                                          `tfsdk:"ttl" json:"ttl,computed"`
}

type MagicWANGRETunnelGRETunnelHealthCheckModel struct {
	Direction types.String                                                               `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool                                                                 `tfsdk:"enabled" json:"enabled,computed"`
	Rate      types.String                                                               `tfsdk:"rate" json:"rate,computed"`
	Target    customfield.NestedObject[MagicWANGRETunnelGRETunnelHealthCheckTargetModel] `tfsdk:"target" json:"target,computed"`
	Type      types.String                                                               `tfsdk:"type" json:"type,computed"`
}

type MagicWANGRETunnelGRETunnelHealthCheckTargetModel struct {
	Effective types.String `tfsdk:"effective" json:"effective,computed"`
	Saved     types.String `tfsdk:"saved" json:"saved,computed"`
}

type MagicWANGRETunnelModifiedGRETunnelModel struct {
	ID                    types.String                                                                 `tfsdk:"id" json:"id,computed"`
	CloudflareGREEndpoint types.String                                                                 `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,computed"`
	CustomerGREEndpoint   types.String                                                                 `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,computed"`
	InterfaceAddress      types.String                                                                 `tfsdk:"interface_address" json:"interface_address,computed"`
	Name                  types.String                                                                 `tfsdk:"name" json:"name,computed"`
	CreatedOn             timetypes.RFC3339                                                            `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description           types.String                                                                 `tfsdk:"description" json:"description,computed"`
	HealthCheck           customfield.NestedObject[MagicWANGRETunnelModifiedGRETunnelHealthCheckModel] `tfsdk:"health_check" json:"health_check,computed"`
	InterfaceAddress6     types.String                                                                 `tfsdk:"interface_address6" json:"interface_address6,computed"`
	ModifiedOn            timetypes.RFC3339                                                            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Mtu                   types.Int64                                                                  `tfsdk:"mtu" json:"mtu,computed"`
	TTL                   types.Int64                                                                  `tfsdk:"ttl" json:"ttl,computed"`
}

type MagicWANGRETunnelModifiedGRETunnelHealthCheckModel struct {
	Direction types.String                                                                       `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool                                                                         `tfsdk:"enabled" json:"enabled,computed"`
	Rate      types.String                                                                       `tfsdk:"rate" json:"rate,computed"`
	Target    customfield.NestedObject[MagicWANGRETunnelModifiedGRETunnelHealthCheckTargetModel] `tfsdk:"target" json:"target,computed"`
	Type      types.String                                                                       `tfsdk:"type" json:"type,computed"`
}

type MagicWANGRETunnelModifiedGRETunnelHealthCheckTargetModel struct {
	Effective types.String `tfsdk:"effective" json:"effective,computed"`
	Saved     types.String `tfsdk:"saved" json:"saved,computed"`
}
