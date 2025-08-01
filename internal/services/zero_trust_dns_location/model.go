// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDNSLocationResultEnvelope struct {
	Result ZeroTrustDNSLocationModel `json:"result"`
}

type ZeroTrustDNSLocationModel struct {
	ID                        types.String                                                    `tfsdk:"id" json:"id,computed"`
	AccountID                 types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	Name                      types.String                                                    `tfsdk:"name" json:"name,required"`
	Endpoints                 *ZeroTrustDNSLocationEndpointsModel                             `tfsdk:"endpoints" json:"endpoints,optional"`
	ClientDefault             types.Bool                                                      `tfsdk:"client_default" json:"client_default,computed_optional"`
	DNSDestinationIPsID       types.String                                                    `tfsdk:"dns_destination_ips_id" json:"dns_destination_ips_id,computed_optional"`
	ECSSupport                types.Bool                                                      `tfsdk:"ecs_support" json:"ecs_support,computed_optional"`
	Networks                  customfield.NestedObjectList[ZeroTrustDNSLocationNetworksModel] `tfsdk:"networks" json:"networks,computed_optional"`
	CreatedAt                 timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DNSDestinationIPV6BlockID types.String                                                    `tfsdk:"dns_destination_ipv6_block_id" json:"dns_destination_ipv6_block_id,computed"`
	DOHSubdomain              types.String                                                    `tfsdk:"doh_subdomain" json:"doh_subdomain,computed"`
	IP                        types.String                                                    `tfsdk:"ip" json:"ip,computed"`
	IPV4Destination           types.String                                                    `tfsdk:"ipv4_destination" json:"ipv4_destination,computed"`
	IPV4DestinationBackup     types.String                                                    `tfsdk:"ipv4_destination_backup" json:"ipv4_destination_backup,computed"`
	UpdatedAt                 timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustDNSLocationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDNSLocationModel) MarshalJSONForUpdate(state ZeroTrustDNSLocationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDNSLocationEndpointsModel struct {
	DOH  *ZeroTrustDNSLocationEndpointsDOHModel  `tfsdk:"doh" json:"doh,required"`
	DOT  *ZeroTrustDNSLocationEndpointsDOTModel  `tfsdk:"dot" json:"dot,required"`
	IPV4 *ZeroTrustDNSLocationEndpointsIPV4Model `tfsdk:"ipv4" json:"ipv4,required"`
	IPV6 *ZeroTrustDNSLocationEndpointsIPV6Model `tfsdk:"ipv6" json:"ipv6,required"`
}

type ZeroTrustDNSLocationEndpointsDOHModel struct {
	Enabled      types.Bool                                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Networks     customfield.NestedObjectList[ZeroTrustDNSLocationEndpointsDOHNetworksModel] `tfsdk:"networks" json:"networks,computed_optional"`
	RequireToken types.Bool                                                                  `tfsdk:"require_token" json:"require_token,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsDOHNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network,required"`
}

type ZeroTrustDNSLocationEndpointsDOTModel struct {
	Enabled  types.Bool                                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Networks customfield.NestedObjectList[ZeroTrustDNSLocationEndpointsDOTNetworksModel] `tfsdk:"networks" json:"networks,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsDOTNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network,required"`
}

type ZeroTrustDNSLocationEndpointsIPV4Model struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsIPV6Model struct {
	Enabled  types.Bool                                                                   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Networks customfield.NestedObjectList[ZeroTrustDNSLocationEndpointsIPV6NetworksModel] `tfsdk:"networks" json:"networks,computed_optional"`
}

type ZeroTrustDNSLocationEndpointsIPV6NetworksModel struct {
	Network types.String `tfsdk:"network" json:"network,required"`
}

type ZeroTrustDNSLocationNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network,required"`
}
