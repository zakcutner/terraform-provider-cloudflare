// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustAccessApplicationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"allow_authenticate_via_warp": schema.BoolAttribute{
				Description: "When set to true, users can authenticate to this application using their WARP session.  When set to false this application will always require direct IdP authentication. This setting always overrides the organization setting for WARP authentication.",
				Optional:    true,
				Validators: []validator.Bool{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), authenticateViaWarpCompatibleAppTypes...),
				},
			},
			"allow_iframe": schema.BoolAttribute{
				Description: "Enables loading application content in an iFrame.",
				Optional:    true,
				Validators: []validator.Bool{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
				},
			},
			"app_launcher_logo_url": schema.StringAttribute{
				Description: "The image URL of the logo shown in the App Launcher header.",
				Optional:    true,
				Validators: []validator.String{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "app_launcher"),
				},
			},
			"bg_color": schema.StringAttribute{
				Description: "The background color of the App Launcher page.",
				Optional:    true,
				Validators: []validator.String{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "app_launcher"),
				},
			},
			"custom_deny_message": schema.StringAttribute{
				Description: "The custom error message shown to a user when they are denied access to the application.",
				Optional:    true,
			},
			"custom_deny_url": schema.StringAttribute{
				Description: "The custom URL a user is redirected to when they are denied access to the application when failing identity-based rules.",
				Optional:    true,
			},
			"custom_non_identity_deny_url": schema.StringAttribute{
				Description: "The custom URL a user is redirected to when they are denied access to the application when failing non-identity rules.",
				Optional:    true,
			},
			"domain": schema.StringAttribute{
				Description: "The primary hostname and path secured by Access. This domain will be displayed if the app is visible in the App Launcher.",
				Optional:    true,
				Computed:    true,
			},
			"header_bg_color": schema.StringAttribute{
				Description: "The background color of the App Launcher header.",
				Optional:    true,
				Validators: []validator.String{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "app_launcher"),
				},
			},
			"logo_url": schema.StringAttribute{
				Description: "The image URL for the logo shown in the App Launcher dashboard.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description:   "The name of the application.",
				Optional:      true,
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"options_preflight_bypass": schema.BoolAttribute{
				Description: "Allows options preflight requests to bypass Access authentication and go directly to the origin. Cannot turn on if cors_headers is set.",
				Optional:    true,
				Validators: []validator.Bool{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
				},
			},
			"read_service_tokens_from_header": schema.StringAttribute{
				Description: "Allows matching Access Service Tokens passed HTTP in a single header with this name.\nThis works as an alternative to the (CF-Access-Client-Id, CF-Access-Client-Secret) pair of headers.\nThe header value will be interpreted as a json object similar to: \n  {\n    \"cf-access-client-id\": \"88bf3b6d86161464f6509f7219099e57.access.example.com\",\n    \"cf-access-client-secret\": \"bdd31cbc4dec990953e39163fbbb194c93313ca9f0a6e420346af9d326b1d2a5\"\n  }",
				Optional:    true,
				Validators: []validator.String{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
				},
			},
			"same_site_cookie_attribute": schema.StringAttribute{
				Description: "Sets the SameSite cookie setting, which provides increased security against CSRF attacks.",
				Optional:    true,
				Validators: []validator.String{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
				},
			},
			"service_auth_401_redirect": schema.BoolAttribute{
				Description: "Returns a 401 status code when the request is blocked by a Service Auth policy.",
				Optional:    true,
				Validators: []validator.Bool{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
				},
			},
			"skip_interstitial": schema.BoolAttribute{
				Description: "Enables automatic authentication through cloudflared.",
				Optional:    true,
				Validators: []validator.Bool{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
				},
			},
			"type": schema.StringAttribute{
				Description: "The application type.\nAvailable values: \"self_hosted\", \"saas\", \"ssh\", \"vnc\", \"app_launcher\", \"warp\", \"biso\", \"bookmark\", \"dash_sso\", \"infrastructure\", \"rdp\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"self_hosted",
						"saas",
						"ssh",
						"vnc",
						"app_launcher",
						"warp",
						"biso",
						"bookmark",
						"dash_sso",
						"infrastructure",
						"rdp",
					),
				},
			},
			"allowed_idps": schema.SetAttribute{
				Description: "The identity providers your users can select when connecting to this application. Defaults to all IdPs configured in your account.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"custom_pages": schema.ListAttribute{
				Description: "The custom pages that will be displayed when applicable for this application",
				Optional:    true,
				ElementType: types.StringType,
			},
			"cors_headers": schema.SingleNestedAttribute{
				Optional: true,
				Validators: []validator.Object{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
				},
				Attributes: map[string]schema.Attribute{
					"allow_all_headers": schema.BoolAttribute{
						Description: "Allows all HTTP request headers.",
						Optional:    true,
						Validators: []validator.Bool{
							boolvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("allowed_headers")),
						},
					},
					"allow_all_methods": schema.BoolAttribute{
						Description: "Allows all HTTP request methods.",
						Optional:    true,
						Validators: []validator.Bool{
							boolvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("allowed_methods")),
						},
					},
					"allow_all_origins": schema.BoolAttribute{
						Description: "Allows all origins.",
						Optional:    true,
						Validators: []validator.Bool{
							boolvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("allowed_origins")),
						},
					},
					"allow_credentials": schema.BoolAttribute{
						Description: "When set to `true`, includes credentials (cookies, authorization headers, or TLS client certificates) with requests.",
						Optional:    true,
						Validators: []validator.Bool{
							boolvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("allow_all_origins")),
						},
					},
					"allowed_headers": schema.SetAttribute{
						Description: "Allowed HTTP request headers.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"allowed_methods": schema.SetAttribute{
						Description: "Allowed HTTP request methods.",
						Optional:    true,
						Validators: []validator.Set{
							setvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"GET",
									"POST",
									"HEAD",
									"PUT",
									"DELETE",
									"CONNECT",
									"OPTIONS",
									"TRACE",
									"PATCH",
								),
							),
						},
						ElementType: types.StringType,
					},
					"allowed_origins": schema.SetAttribute{
						Description: "Allowed origins.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"max_age": schema.Float64Attribute{
						Description: "The maximum number of seconds the results of a preflight request can be cached.",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(-1, 86400),
						},
					},
				},
			},
			"footer_links": schema.ListNestedAttribute{
				Description: "The links in the App Launcher footer.",
				Optional:    true,
				Validators: []validator.List{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "app_launcher"),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "The hypertext in the footer link.",
							Required:    true,
						},
						"url": schema.StringAttribute{
							Description: "the hyperlink in the footer link.",
							Required:    true,
						},
					},
				},
			},
			"scim_config": schema.SingleNestedAttribute{
				Description: "Configuration for provisioning to this application via SCIM. This is currently in closed beta.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"idp_uid": schema.StringAttribute{
						Description: "The UID of the IdP to use as the source for SCIM resources to provision to this application.",
						Required:    true,
					},
					"remote_uri": schema.StringAttribute{
						Description: "The base URI for the application's SCIM-compatible API.",
						Required:    true,
					},
					"authentication": schema.SingleNestedAttribute{
						Description: "Attributes for configuring HTTP Basic authentication scheme for SCIM provisioning to an application.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"password": schema.StringAttribute{
								Description: "Password used to authenticate with the remote SCIM service.",
								Optional:    true,
								Sensitive:   true,
							},
							"scheme": schema.StringAttribute{
								Description: "The authentication scheme to use when making SCIM requests to this application.\nAvailable values: \"httpbasic\", \"oauthbearertoken\", \"oauth2\", \"access_service_token\".",
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"httpbasic",
										"oauthbearertoken",
										"oauth2",
										"access_service_token",
									),
								},
							},
							"user": schema.StringAttribute{
								Description: "User name used to authenticate with the remote SCIM service.",
								Optional:    true,
							},
							"token": schema.StringAttribute{
								Description: "Token used to authenticate with the remote SCIM service.",
								Optional:    true,
								Sensitive:   true,
							},
							"authorization_url": schema.StringAttribute{
								Description: "URL used to generate the auth code used during token generation.",
								Optional:    true,
							},
							"client_id": schema.StringAttribute{
								Description: "Client ID used to authenticate when generating a token for authenticating with the remote SCIM service.",
								Optional:    true,
							},
							"client_secret": schema.StringAttribute{
								Description: "Secret used to authenticate when generating a token for authenticating with the remove SCIM service.",
								Optional:    true,
								Sensitive:   true,
							},
							"token_url": schema.StringAttribute{
								Description: "URL used to generate the token used to authenticate with the remote SCIM service.",
								Optional:    true,
							},
							"scopes": schema.ListAttribute{
								Description: "The authorization scopes to request when generating the token used to authenticate with the remove SCIM service.",
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"deactivate_on_delete": schema.BoolAttribute{
						Description: "If false, propagates DELETE requests to the target application for SCIM resources. If true, sets 'active' to false on the SCIM resource. Note: Some targets do not support DELETE operations.",
						Optional:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Whether SCIM provisioning is turned on for this application.",
						Optional:    true,
					},
					"mappings": schema.ListNestedAttribute{
						Description: "A list of mappings to apply to SCIM resources before provisioning them in this application. These can transform or filter the resources to be provisioned.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"schema": schema.StringAttribute{
									Description: "Which SCIM resource type this mapping applies to.",
									Required:    true,
								},
								"enabled": schema.BoolAttribute{
									Description: "Whether or not this mapping is enabled.",
									Optional:    true,
								},
								"filter": schema.StringAttribute{
									Description: "A [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2) that matches resources that should be provisioned to this application.",
									Optional:    true,
								},
								"operations": schema.SingleNestedAttribute{
									Description: "Whether or not this mapping applies to creates, updates, or deletes.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"create": schema.BoolAttribute{
											Description: "Whether or not this mapping applies to create (POST) operations.",
											Optional:    true,
										},
										"delete": schema.BoolAttribute{
											Description: "Whether or not this mapping applies to DELETE operations.",
											Optional:    true,
										},
										"update": schema.BoolAttribute{
											Description: "Whether or not this mapping applies to update (PATCH/PUT) operations.",
											Optional:    true,
										},
									},
								},
								"strictness": schema.StringAttribute{
									Description: "The level of adherence to outbound resource schemas when provisioning to this mapping. ‘Strict’ removes unknown values, while ‘passthrough’ passes unknown values to the target.\nAvailable values: \"strict\", \"passthrough\".",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("strict", "passthrough"),
									},
								},
								"transform_jsonata": schema.StringAttribute{
									Description: "A [JSONata](https://jsonata.org/) expression that transforms the resource before provisioning it in the application.",
									Optional:    true,
								},
							},
						},
					},
				},
			},
			"target_criteria": schema.ListNestedAttribute{
				Optional: true,
				Validators: []validator.List{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), targetCompatibleAppTypes...),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port": schema.Int64Attribute{
							Description: "The port that the targets use for the chosen communication protocol. A port cannot be assigned to multiple protocols.",
							Required:    true,
						},
						"protocol": schema.StringAttribute{
							Description: "The communication protocol your application secures.\nAvailable values: \"SSH\", \"RDP\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("SSH", "RDP"),
							},
						},
						"target_attributes": schema.MapAttribute{
							Description: "Contains a map of target attribute keys to target attribute values.",
							Required:    true,
							ElementType: types.ListType{
								ElemType: types.StringType,
							},
						},
					},
				},
			},
			"app_launcher_visible": schema.BoolAttribute{
				Description: "Displays the application in the App Launcher.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Bool{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), appLauncherVisibleAppTypes...),
				},
			},
			"auto_redirect_to_identity": schema.BoolAttribute{
				Description: "When set to `true`, users skip the identity provider selection step during login. You must specify only one identity provider in allowed_idps.",
				Optional:    true,
			},
			"enable_binding_cookie": schema.BoolAttribute{
				Description: "Enables the binding cookie, which increases security against compromised authorization tokens and CSRF attacks.",
				Optional:    true,
				Validators: []validator.Bool{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
				},
			},
			"http_only_cookie_attribute": schema.BoolAttribute{
				Description: "Enables the HttpOnly cookie attribute, which increases security against XSS attacks.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Bool{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
				},
			},
			"path_cookie_attribute": schema.BoolAttribute{
				Description: "Enables cookie paths to scope an application's JWT to the application path. If disabled, the JWT will scope to the hostname by default",
				Optional:    true,
				Validators: []validator.Bool{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
				},
			},
			"session_duration": schema.StringAttribute{
				Description: "The amount of time that tokens issued for this application will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. Note: unsupported for infrastructure type applications.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(durationRegex, `"session_duration" only supports "ns", "us" (or "µs"), "ms", "s", "m", or "h" as valid units`),
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), sessionDurationCompatibleAppTypes...),
				},
			},
			"skip_app_launcher_login_page": schema.BoolAttribute{
				Description: "Determines when to skip the App Launcher landing page.",
				Optional:    true,
				Computed:    true,
				Validators: []validator.Bool{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "app_launcher"),
				},
			},
			"self_hosted_domains": schema.ListAttribute{
				Description:        "List of public domains that Access will secure. This field is deprecated in favor of `destinations` and will be supported until **November 21, 2025.** If `destinations` are provided, then `self_hosted_domains` will be ignored.",
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
				CustomType:         customfield.NewListType[types.String](ctx),
				ElementType:        types.StringType,
				Validators: []validator.List{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
					listvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("destinations"),
					}...),
				},
			},
			"tags": schema.ListAttribute{
				Description: "The tags you want assigned to an application. Tags are used to filter applications in the App Launcher dashboard.",
				CustomType:  customfield.NewListType[types.String](ctx),
				Optional:    true,
				ElementType: types.StringType,
			},
			"destinations": schema.ListNestedAttribute{
				Description: "List of destinations secured by Access. This supersedes `self_hosted_domains` to allow for more flexibility in defining different types of domains. If `destinations` are provided, then `self_hosted_domains` will be ignored.",
				Optional:    true,
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessApplicationDestinationsModel](ctx),
				Validators: []validator.List{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), selfHostedAppTypes...),
					listvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("self_hosted_domains"),
					}...),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: `Available values: "public", "private".`,
							Optional:    true,
							Computed:    true,
							Default:     stringdefault.StaticString("public"),
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("public", "private"),
							},
						},
						"uri": schema.StringAttribute{
							Description: "The URI of the destination. Public destinations' URIs can include a domain and path with [wildcards](https://developers.cloudflare.com/cloudflare-one/policies/access/app-paths/).",
							Optional:    true,
							Validators: []validator.String{
								customvalidator.RequiresOtherStringAttributeToNullOrBeOneOf(path.MatchRelative().AtParent().AtName("type"), "public"),
							},
						},
						"cidr": schema.StringAttribute{
							Description: "The CIDR range of the destination. Single IPs will be computed as /32.",
							Optional:    true,
							Validators: []validator.String{
								customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRelative().AtParent().AtName("type"), "private"),
							},
						},
						"hostname": schema.StringAttribute{
							Description: "The hostname of the destination. Matches a valid SNI served by an HTTPS origin.",
							Optional:    true,
							Validators: []validator.String{
								customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRelative().AtParent().AtName("type"), "private"),
							},
						},
						"l4_protocol": schema.StringAttribute{
							Description: "The L4 protocol of the destination. When omitted, both UDP and TCP traffic will match.\nAvailable values: \"tcp\", \"udp\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("tcp", "udp"),
								customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRelative().AtParent().AtName("type"), "private"),
							},
						},
						"port_range": schema.StringAttribute{
							Description: "The port range of the destination. Can be a single port or a range of ports. When omitted, all ports will match.",
							Optional:    true,
							Validators: []validator.String{
								customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRelative().AtParent().AtName("type"), "private"),
							},
						},
						"vnet_id": schema.StringAttribute{
							Description: "The VNET ID to match the destination. When omitted, all VNETs will match.",
							Optional:    true,
							Validators: []validator.String{
								customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRelative().AtParent().AtName("type"), "private"),
							},
						},
					},
				},
			},
			"landing_page_design": schema.SingleNestedAttribute{
				Description: "The design of the App Launcher landing page shown to users when they log in.",
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationLandingPageDesignModel](ctx),
				Validators: []validator.Object{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), "app_launcher"),
				},
				Attributes: map[string]schema.Attribute{
					"button_color": schema.StringAttribute{
						Description: "The background color of the log in button on the landing page.",
						Optional:    true,
					},
					"button_text_color": schema.StringAttribute{
						Description: "The color of the text in the log in button on the landing page.",
						Optional:    true,
					},
					"image_url": schema.StringAttribute{
						Description: "The URL of the image shown on the landing page.",
						Optional:    true,
					},
					"message": schema.StringAttribute{
						Description: "The message shown on the landing page.",
						Optional:    true,
					},
					"title": schema.StringAttribute{
						Description: "The title shown on the landing page.",
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("Welcome!"),
					},
				},
			},
			"policies": schema.ListNestedAttribute{
				Description: "The policies that Access applies to the application, in ascending order of precedence. Items can reference existing policies or create new policies exclusive to the application.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the policy",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("include")),
							},
						},
						"precedence": schema.Int64Attribute{
							Description: "The order of execution for this policy. Must be unique for each policy within an app.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
						},
						"decision": schema.StringAttribute{
							Description: "The action Access will take if a user matches this policy. Infrastructure application policies can only use the Allow action.\nAvailable values: \"allow\", \"deny\", \"non_identity\", \"bypass\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"allow",
									"deny",
									"non_identity",
									"bypass",
								),
								stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("include")),
							},
						},
						"include": schema.ListNestedAttribute{
							Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.",
							Optional:    true,
							Validators: []validator.List{
								listvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("decision")),
							},
							CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesIncludeModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Validators: []validator.Object{
									customvalidator.ObjectSizeAtMost(1),
								},
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Required:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"auth_context": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Authentication context.",
												Required:    true,
											},
											"ac_id": schema.StringAttribute{
												Description: "The ACID of an Authentication context.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Required:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Optional:   true,
										Attributes: map[string]schema.Attribute{},
									},
									"common_name": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"common_name": schema.StringAttribute{
												Description: "The common name to match.",
												Required:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Required:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Required:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Required:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Required:    true,
											},
										},
									},
									"email": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Required:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Required:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Required:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Required:    true,
											},
											"team": schema.StringAttribute{
												Description: "The name of the team",
												Optional:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Required:    true,
											},
										},
									},
									"login_method": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an identity provider.",
												Required:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Required:    true,
											},
										},
									},
									"ip": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Required:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Required:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Required:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Required:    true,
											},
										},
									},
									"oidc": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"claim_name": schema.StringAttribute{
												Description: "The name of the OIDC claim.",
												Required:    true,
											},
											"claim_value": schema.StringAttribute{
												Description: "The OIDC claim value to look for.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your OIDC identity provider.",
												Required:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Required:    true,
											},
										},
									},
									"linked_app_token": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"app_uid": schema.StringAttribute{
												Description: "The ID of an Access OIDC SaaS application",
												Required:    true,
											},
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the Access policy.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("include")),
							},
						},
						"connection_rules": schema.SingleNestedAttribute{
							Description: "The rules that define how users may connect to the targets secured by your application.",
							Optional:    true,
							Validators: []validator.Object{
								objectvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("include")),
								customvalidator.RequiredWhenOtherStringIsOneOf(path.MatchRoot("type"), "infrastructure"),
							},
							Attributes: map[string]schema.Attribute{
								"ssh": schema.SingleNestedAttribute{
									Description: "The SSH-specific rules that define how users may connect to the targets secured by your application.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"usernames": schema.ListAttribute{
											Description: "Contains the Unix usernames that may be used when connecting over SSH.",
											Required:    true,
											ElementType: types.StringType,
										},
										"allow_email_alias": schema.BoolAttribute{
											Description: "Enables using Identity Provider email alias as SSH username.",
											Optional:    true,
										},
									},
								},
							},
						},
						"exclude": schema.ListNestedAttribute{
							Description: "Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.",
							Optional:    true,
							Validators: []validator.List{
								listvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("include")),
							},
							CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesExcludeModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Validators: []validator.Object{
									customvalidator.ObjectSizeAtMost(1),
								},
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Required:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"auth_context": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Authentication context.",
												Required:    true,
											},
											"ac_id": schema.StringAttribute{
												Description: "The ACID of an Authentication context.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Required:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Optional:   true,
										Attributes: map[string]schema.Attribute{},
									},
									"common_name": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"common_name": schema.StringAttribute{
												Description: "The common name to match.",
												Required:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Required:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Required:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Required:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Required:    true,
											},
										},
									},
									"email": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Required:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Required:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Required:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Required:    true,
											},
											"team": schema.StringAttribute{
												Description: "The name of the team",
												Optional:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Required:    true,
											},
										},
									},
									"login_method": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an identity provider.",
												Required:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Required:    true,
											},
										},
									},
									"ip": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Required:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Required:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Required:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Required:    true,
											},
										},
									},
									"oidc": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"claim_name": schema.StringAttribute{
												Description: "The name of the OIDC claim.",
												Required:    true,
											},
											"claim_value": schema.StringAttribute{
												Description: "The OIDC claim value to look for.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your OIDC identity provider.",
												Required:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Required:    true,
											},
										},
									},
									"linked_app_token": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"app_uid": schema.StringAttribute{
												Description: "The ID of an Access OIDC SaaS application",
												Required:    true,
											},
										},
									},
								},
							},
						},
						"require": schema.ListNestedAttribute{
							Description: "Rules evaluated with an AND logical operator. To match the policy, a user must meet all of the Require rules.",
							Optional:    true,
							Validators: []validator.List{
								listvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("include")),
							},
							CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessApplicationPoliciesRequireModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Validators: []validator.Object{
									customvalidator.ObjectSizeAtMost(1),
								},
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Required:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"auth_context": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Authentication context.",
												Required:    true,
											},
											"ac_id": schema.StringAttribute{
												Description: "The ACID of an Authentication context.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Required:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Required:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Optional:   true,
										Attributes: map[string]schema.Attribute{},
									},
									"common_name": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"common_name": schema.StringAttribute{
												Description: "The common name to match.",
												Required:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Required:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Required:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Required:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Required:    true,
											},
										},
									},
									"email": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Required:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Optional:    true,
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Required:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Required:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Required:    true,
											},
											"team": schema.StringAttribute{
												Description: "The name of the team",
												Optional:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Required:    true,
											},
										},
									},
									"login_method": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an identity provider.",
												Required:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Required:    true,
											},
										},
									},
									"ip": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Required:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Required:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Required:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Required:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Required:    true,
											},
										},
									},
									"oidc": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"claim_name": schema.StringAttribute{
												Description: "The name of the OIDC claim.",
												Required:    true,
											},
											"claim_value": schema.StringAttribute{
												Description: "The OIDC claim value to look for.",
												Required:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your OIDC identity provider.",
												Required:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Required:    true,
											},
										},
									},
									"linked_app_token": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"app_uid": schema.StringAttribute{
												Description: "The ID of an Access OIDC SaaS application",
												Required:    true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"saas_app": schema.SingleNestedAttribute{
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustAccessApplicationSaaSAppModel](ctx),
				Validators: []validator.Object{
					customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("type"), saasAppTypes...),
					customvalidator.RequiredWhenOtherStringIsOneOf(path.MatchRoot("type"), saasAppTypes...),
				},
				Attributes: map[string]schema.Attribute{
					"auth_type": schema.StringAttribute{
						Description: "Optional identifier indicating the authentication protocol used for the saas app. Required for OIDC. Default if unset is \"saml\"\nAvailable values: \"saml\", \"oidc\".",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("saml"),
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("saml", "oidc"),
						},
					},
					"consumer_service_url": schema.StringAttribute{
						Description: "The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToNullOrBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "saml"),
						},
					},
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"custom_attributes": schema.ListNestedAttribute{
						Optional: true,
						Validators: []validator.List{
							customvalidator.RequiresOtherStringAttributeToNullOrBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "saml"),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"friendly_name": schema.StringAttribute{
									Description: "The SAML FriendlyName of the attribute.",
									Optional:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the attribute.",
									Optional:    true,
								},
								"name_format": schema.StringAttribute{
									Description: "A globally unique name for an identity or service provider.\nAvailable values: \"urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified\", \"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\", \"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\".",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified",
											"urn:oasis:names:tc:SAML:2.0:attrname-format:basic",
											"urn:oasis:names:tc:SAML:2.0:attrname-format:uri",
										),
									},
								},
								"required": schema.BoolAttribute{
									Description: "If the attribute is required when building a SAML assertion.",
									Optional:    true,
								},
								"source": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "The name of the IdP attribute.",
											Optional:    true,
										},
										"name_by_idp": schema.ListNestedAttribute{
											Description: "A mapping from IdP ID to attribute name.",
											Optional:    true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"idp_id": schema.StringAttribute{
														Description: "The UID of the IdP.",
														Optional:    true,
													},
													"source_name": schema.StringAttribute{
														Description: "The name of the IdP provided attribute.",
														Optional:    true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"default_relay_state": schema.StringAttribute{
						Description: "The URL that the user will be redirected to after a successful login for IDP initiated logins.",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToNullOrBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "saml"),
						},
					},
					"idp_entity_id": schema.StringAttribute{
						Description: "The unique identifier for your SaaS application.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToNullOrBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "saml"),
						},
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name_id_format": schema.StringAttribute{
						Description: "The format of the name identifier sent to the SaaS application.\nAvailable values: \"id\", \"email\".",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("id", "email"),
							customvalidator.RequiresOtherStringAttributeToNullOrBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "saml"),
						},
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name_id_transform_jsonata": schema.StringAttribute{
						Description: "A [JSONata](https://jsonata.org/) expression that transforms an application's user identities into a NameID value for its SAML assertion. This expression should evaluate to a singular string. The output of this expression can override the `name_id_format` setting.",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToNullOrBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "saml"),
						},
					},
					"public_key": schema.StringAttribute{
						Description: "The Access public certificate that will be used to verify your identity.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"saml_attribute_transform_jsonata": schema.StringAttribute{
						Description: "A [JSONata] (https://jsonata.org/) expression that transforms an application's user identities into attribute assertions in the SAML response. The expression can transform id, email, name, and groups values. It can also transform fields listed in the saml_attributes or oidc_fields of the identity provider used to authenticate. The output of this expression must be a JSON object.",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToNullOrBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "saml"),
						},
					},
					"sp_entity_id": schema.StringAttribute{
						Description: "A globally unique name for an identity or service provider.",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToNullOrBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "saml"),
						},
					},
					"sso_endpoint": schema.StringAttribute{
						Description: "The endpoint where your SaaS application will send login requests.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToNullOrBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "saml"),
						},
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"updated_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"access_token_lifetime": schema.StringAttribute{
						Description: "The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must be greater than or equal to 1m and less than or equal to 24h.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "oidc"),
						},
					},
					"allow_pkce_without_client_secret": schema.BoolAttribute{
						Description: "If client secret should be required on the token endpoint when authorization_code_with_pkce grant is used.",
						Optional:    true,
						Validators: []validator.Bool{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "oidc"),
						},
					},
					"app_launcher_url": schema.StringAttribute{
						Description: "The URL where this applications tile redirects users",
						Optional:    true,
					},
					"client_id": schema.StringAttribute{
						Description: "The application client id",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"client_secret": schema.StringAttribute{
						Description: "The application client secret, only returned on POST request.",
						Computed:    true,
						Sensitive:   true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"custom_claims": schema.ListNestedAttribute{
						Optional: true,
						Validators: []validator.List{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "oidc"),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description: "The name of the claim.",
									Optional:    true,
								},
								"required": schema.BoolAttribute{
									Description: "If the claim is required when building an OIDC token.",
									Optional:    true,
								},
								"scope": schema.StringAttribute{
									Description: "The scope of the claim.\nAvailable values: \"groups\", \"profile\", \"email\", \"openid\".",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"groups",
											"profile",
											"email",
											"openid",
										),
									},
								},
								"source": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "The name of the IdP claim.",
											Optional:    true,
										},
										"name_by_idp": schema.MapAttribute{
											Description: "A mapping from IdP ID to claim name.",
											Optional:    true,
											ElementType: types.StringType,
										},
									},
								},
							},
						},
					},
					"grant_types": schema.ListAttribute{
						Description: "The OIDC flows supported by this application",
						Optional:    true,
						Validators: []validator.List{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "oidc"),
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"authorization_code",
									"authorization_code_with_pkce",
									"refresh_tokens",
									"hybrid",
									"implicit",
								),
							),
						},
						ElementType: types.StringType,
					},
					"group_filter_regex": schema.StringAttribute{
						Description: "A regex to filter Cloudflare groups returned in ID token and userinfo endpoint",
						Optional:    true,
						Validators: []validator.String{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "oidc"),
						},
					},
					"hybrid_and_implicit_options": schema.SingleNestedAttribute{
						Optional: true,
						Validators: []validator.Object{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "oidc"),
						},
						Attributes: map[string]schema.Attribute{
							"return_access_token_from_authorization_endpoint": schema.BoolAttribute{
								Description: "If an Access Token should be returned from the OIDC Authorization endpoint",
								Optional:    true,
							},
							"return_id_token_from_authorization_endpoint": schema.BoolAttribute{
								Description: "If an ID Token should be returned from the OIDC Authorization endpoint",
								Optional:    true,
							},
						},
					},
					"redirect_uris": schema.ListAttribute{
						Description: "The permitted URL's for Cloudflare to return Authorization codes and Access/ID tokens",
						Optional:    true,
						ElementType: types.StringType,
						Validators: []validator.List{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "oidc"),
						},
					},
					"refresh_token_options": schema.SingleNestedAttribute{
						Optional: true,
						Validators: []validator.Object{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "oidc"),
						},
						Attributes: map[string]schema.Attribute{
							"lifetime": schema.StringAttribute{
								Description: "How long a refresh token will be valid for after creation. Valid units are m,h,d. Must be longer than 1m.",
								Optional:    true,
							},
						},
					},
					"scopes": schema.ListAttribute{
						Description: `Define the user information shared with access, "offline_access" scope will be automatically enabled if refresh tokens are enabled`,
						Optional:    true,
						Validators: []validator.List{
							customvalidator.RequiresOtherStringAttributeToBeOneOf(path.MatchRoot("saas_app").AtName("auth_type"), "oidc"),
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"openid",
									"groups",
									"email",
									"profile",
								),
							),
						},
						ElementType: types.StringType,
					},
				},
			},
			"aud": schema.StringAttribute{
				Description: "Audience tag.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ZeroTrustAccessApplicationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustAccessApplicationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
