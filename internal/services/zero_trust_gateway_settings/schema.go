// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustGatewaySettingsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"settings": schema.SingleNestedAttribute{
				Description: "Account settings",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"activity_log": schema.SingleNestedAttribute{
						Description: "Activity log settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable activity logging.",
								Optional:    true,
							},
						},
					},
					"antivirus": schema.SingleNestedAttribute{
						Description: "Anti-virus settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled_download_phase": schema.BoolAttribute{
								Description: "Enable anti-virus scanning on downloads.",
								Computed:    true,
								Optional:    true,
							},
							"enabled_upload_phase": schema.BoolAttribute{
								Description: "Enable anti-virus scanning on uploads.",
								Computed:    true,
								Optional:    true,
							},
							"fail_closed": schema.BoolAttribute{
								Description: "Block requests for files that cannot be scanned.",
								Computed:    true,
								Optional:    true,
							},
							"notification_settings": schema.SingleNestedAttribute{
								Description: "Configure a message to display on the user's device when an antivirus search is performed.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettingsModel](ctx),
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description: "Set notification on",
										Optional:    true,
									},
									"include_context": schema.BoolAttribute{
										Description: "If true, context information will be passed as query parameters",
										Optional:    true,
									},
									"msg": schema.StringAttribute{
										Description: "Customize the message shown in the notification.",
										Optional:    true,
									},
									"support_url": schema.StringAttribute{
										Description: "Optional URL to direct users to additional information. If not set, the notification will open a block page.",
										Optional:    true,
									},
								},
							},
						},
					},
					"block_page": schema.SingleNestedAttribute{
						Description: "Block page layout settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"background_color": schema.StringAttribute{
								Description: "If mode is customized_block_page: block page background color in #rrggbb format.",
								Optional:    true,
							},
							"enabled": schema.BoolAttribute{
								Description: "Enable only cipher suites and TLS versions compliant with FIPS 140-2.",
								Optional:    true,
							},
							"footer_text": schema.StringAttribute{
								Description: "If mode is customized_block_page: block page footer text.",
								Optional:    true,
							},
							"header_text": schema.StringAttribute{
								Description: "If mode is customized_block_page: block page header text.",
								Optional:    true,
							},
							"include_context": schema.BoolAttribute{
								Description: "If mode is redirect_uri: when enabled, context will be appended to target_uri as query parameters.",
								Optional:    true,
							},
							"logo_path": schema.StringAttribute{
								Description: "If mode is customized_block_page: full URL to the logo file.",
								Optional:    true,
							},
							"mailto_address": schema.StringAttribute{
								Description: "If mode is customized_block_page: admin email for users to contact.",
								Optional:    true,
							},
							"mailto_subject": schema.StringAttribute{
								Description: "If mode is customized_block_page: subject line for emails created from block page.",
								Optional:    true,
							},
							"mode": schema.StringAttribute{
								Description: "Controls whether the user is redirected to a Cloudflare-hosted block page or to a customer-provided URI.\nAvailable values: \"customized_block_page\", \"redirect_uri\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("customized_block_page", "redirect_uri"),
								},
							},
							"name": schema.StringAttribute{
								Description: "If mode is customized_block_page: block page title.",
								Optional:    true,
							},
							"read_only": schema.BoolAttribute{
								Description: "This setting was shared via the Orgs API and cannot be edited by the current account",
								Computed:    true,
								Optional:    true,
							},
							"source_account": schema.StringAttribute{
								Description: "Account tag of account that shared this setting",
								Computed:    true,
								Optional:    true,
							},
							"suppress_footer": schema.BoolAttribute{
								Description: "If mode is customized_block_page: suppress detailed info at the bottom of the block page.",
								Optional:    true,
							},
							"target_uri": schema.StringAttribute{
								Description: "If mode is redirect_uri: URI to which the user should be redirected.",
								Optional:    true,
							},
							"version": schema.Int64Attribute{
								Description: "Version number of the setting",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"body_scanning": schema.SingleNestedAttribute{
						Description: "DLP body scanning settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"inspection_mode": schema.StringAttribute{
								Description: "Set the inspection mode to either `deep` or `shallow`.\nAvailable values: \"deep\", \"shallow\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("deep", "shallow"),
								},
							},
						},
					},
					"browser_isolation": schema.SingleNestedAttribute{
						Description: "Browser isolation settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"non_identity_enabled": schema.BoolAttribute{
								Description: "Enable non-identity onramp support for Browser Isolation.",
								Optional:    true,
							},
							"url_browser_isolation_enabled": schema.BoolAttribute{
								Description: "Enable Clientless Browser Isolation.",
								Optional:    true,
							},
						},
					},
					"certificate": schema.SingleNestedAttribute{
						Description: "Certificate settings for Gateway TLS interception. If not specified, the Cloudflare Root CA will be used.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "UUID of certificate to be used for interception. Certificate must be available (previously called 'active') on the edge. A nil UUID will indicate the Cloudflare Root CA should be used.",
								Required:    true,
							},
						},
					},
					"custom_certificate": schema.SingleNestedAttribute{
						Description:        "Custom certificate settings for BYO-PKI. (deprecated and replaced by `certificate`)",
						Optional:           true,
						DeprecationMessage: "This attribute is deprecated.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable use of custom certificate authority for signing Gateway traffic.",
								Required:    true,
							},
							"id": schema.StringAttribute{
								Description: "UUID of certificate (ID from MTLS certificate store).",
								Optional:    true,
							},
							"binding_status": schema.StringAttribute{
								Description: "Certificate status (internal).",
								Computed:    true,
								Optional:    true,
							},
							"updated_at": schema.StringAttribute{
								Computed:   true,
								Optional:   true,
								CustomType: timetypes.RFC3339Type{},
							},
						},
					},
					"extended_email_matching": schema.SingleNestedAttribute{
						Description: "Extended e-mail matching settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable matching all variants of user emails (with + or . modifiers) used as criteria in Firewall policies.",
								Optional:    true,
							},
							"read_only": schema.BoolAttribute{
								Description: "This setting was shared via the Orgs API and cannot be edited by the current account",
								Optional:    true,
								Computed:    true,
							},
							"source_account": schema.StringAttribute{
								Description: "Account tag of account that shared this setting",
								Optional:    true,
								Computed:    true,
							},
							"version": schema.Int64Attribute{
								Description: "Version number of the setting",
								Optional:    true,
								Computed:    true,
							},
						},
					},
					"fips": schema.SingleNestedAttribute{
						Description: "FIPS settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"tls": schema.BoolAttribute{
								Description: "Enable only cipher suites and TLS versions compliant with FIPS 140-2.",
								Optional:    true,
							},
						},
					},
					"host_selector": schema.SingleNestedAttribute{
						Description: "Setting to enable host selector in egress policies.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable filtering via hosts for egress policies.",
								Optional:    true,
							},
						},
					},
					"inspection": schema.SingleNestedAttribute{
						Description: "Setting to define inspection settings",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description: "Defines the mode of inspection the proxy will use.\n- static: Gateway will use static inspection to inspect HTTP on TCP(80). If TLS decryption is on, Gateway will inspect HTTPS traffic on TCP(443) & UDP(443).\n- dynamic: Gateway will use protocol detection to dynamically inspect HTTP and HTTPS traffic on any port. TLS decryption must be on to inspect HTTPS traffic.\nAvailable values: \"static\", \"dynamic\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("static", "dynamic"),
								},
							},
						},
					},
					"protocol_detection": schema.SingleNestedAttribute{
						Description: "Protocol Detection settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable detecting protocol on initial bytes of client traffic.",
								Optional:    true,
							},
						},
					},
					"sandbox": schema.SingleNestedAttribute{
						Description: "Sandbox settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable sandbox.",
								Optional:    true,
							},
							"fallback_action": schema.StringAttribute{
								Description: "Action to take when the file cannot be scanned.\nAvailable values: \"allow\", \"block\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("allow", "block"),
								},
							},
						},
					},
					"tls_decrypt": schema.SingleNestedAttribute{
						Description: "TLS interception settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable inspecting encrypted HTTP traffic.",
								Optional:    true,
							},
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustGatewaySettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustGatewaySettingsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
