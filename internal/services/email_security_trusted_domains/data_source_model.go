// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_trusted_domains

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/email_security"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityTrustedDomainsResultDataSourceEnvelope struct {
	Result EmailSecurityTrustedDomainsDataSourceModel `json:"result,computed"`
}

type EmailSecurityTrustedDomainsDataSourceModel struct {
	ID              types.Int64                                          `tfsdk:"id" path:"trusted_domain_id,computed"`
	TrustedDomainID types.Int64                                          `tfsdk:"trusted_domain_id" path:"trusted_domain_id,optional"`
	AccountID       types.String                                         `tfsdk:"account_id" path:"account_id,required"`
	Comments        types.String                                         `tfsdk:"comments" json:"comments,computed"`
	CreatedAt       timetypes.RFC3339                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IsRecent        types.Bool                                           `tfsdk:"is_recent" json:"is_recent,computed"`
	IsRegex         types.Bool                                           `tfsdk:"is_regex" json:"is_regex,computed"`
	IsSimilarity    types.Bool                                           `tfsdk:"is_similarity" json:"is_similarity,computed"`
	LastModified    timetypes.RFC3339                                    `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	Pattern         types.String                                         `tfsdk:"pattern" json:"pattern,computed"`
	Filter          *EmailSecurityTrustedDomainsFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *EmailSecurityTrustedDomainsDataSourceModel) toReadParams(_ context.Context) (params email_security.SettingTrustedDomainGetParams, diags diag.Diagnostics) {
	params = email_security.SettingTrustedDomainGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *EmailSecurityTrustedDomainsDataSourceModel) toListParams(_ context.Context) (params email_security.SettingTrustedDomainListParams, diags diag.Diagnostics) {
	params = email_security.SettingTrustedDomainListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(email_security.SettingTrustedDomainListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.IsRecent.IsNull() {
		params.IsRecent = cloudflare.F(m.Filter.IsRecent.ValueBool())
	}
	if !m.Filter.IsSimilarity.IsNull() {
		params.IsSimilarity = cloudflare.F(m.Filter.IsSimilarity.ValueBool())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(email_security.SettingTrustedDomainListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Pattern.IsNull() {
		params.Pattern = cloudflare.F(m.Filter.Pattern.ValueString())
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	return
}

type EmailSecurityTrustedDomainsFindOneByDataSourceModel struct {
	Direction    types.String `tfsdk:"direction" query:"direction,optional"`
	IsRecent     types.Bool   `tfsdk:"is_recent" query:"is_recent,optional"`
	IsSimilarity types.Bool   `tfsdk:"is_similarity" query:"is_similarity,optional"`
	Order        types.String `tfsdk:"order" query:"order,optional"`
	Pattern      types.String `tfsdk:"pattern" query:"pattern,optional"`
	Search       types.String `tfsdk:"search" query:"search,optional"`
}
