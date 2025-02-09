package provider

import (
	"context"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	cfv2 "github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/access_mutual_tls_hostname_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/api_shield_operation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/api_token_permissions_groups"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/cloud_connector_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/content_scanning"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/content_scanning_expression"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/d1"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/dcv_delegation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/dlp_datasets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/email_routing_address"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/email_routing_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/gateway_app_types"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/gateway_categories"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/hyperdrive_config"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/infrastructure_access_target_deprecated"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/leaked_credential_check"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/leaked_credential_check_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/list"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/list_item"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/origin_ca_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/r2_bucket"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/risk_behavior"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/rulesets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/snippet_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/snippets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/turnstile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/user"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/workers_for_platforms_dispatch_namespace"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/workers_for_platforms_dispatch_namespace_deprecated"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/zero_trust_access_mtls_hostname_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/zero_trust_infrastructure_access_target"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/zero_trust_risk_behavior"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/service/zero_trust_risk_score_integration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/sdkv2provider"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
)

// Ensure CloudflareProvider satisfies various provider interfaces.
var _ provider.Provider = &CloudflareProvider{}

// CloudflareProvider defines the provider implementation.
type CloudflareProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CloudflareProviderModel describes the provider data model.
type CloudflareProviderModel struct {
	APIKey                  types.String `tfsdk:"api_key"`
	APIUserServiceKey       types.String `tfsdk:"api_user_service_key"`
	Email                   types.String `tfsdk:"email"`
	MinBackOff              types.Int64  `tfsdk:"min_backoff"`
	RPS                     types.Int64  `tfsdk:"rps"`
	APIBasePath             types.String `tfsdk:"api_base_path"`
	APIToken                types.String `tfsdk:"api_token"`
	Retries                 types.Int64  `tfsdk:"retries"`
	MaxBackoff              types.Int64  `tfsdk:"max_backoff"`
	APIClientLogging        types.Bool   `tfsdk:"api_client_logging"`
	APIHostname             types.String `tfsdk:"api_hostname"`
	UserAgentOperatorSuffix types.String `tfsdk:"user_agent_operator_suffix"`
}

func (p *CloudflareProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
	resp.Version = p.version
}

func (p *CloudflareProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			consts.EmailSchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("A registered Cloudflare email address. Alternatively, can be configured using the `%s` environment variable. Required when using `api_key`. Conflicts with `api_token`.", consts.EmailEnvVarKey),
				Validators:          []validator.String{},
			},

			consts.APIKeySchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("The API key for operations. Alternatively, can be configured using the `%s` environment variable. API keys are [now considered legacy by Cloudflare](https://developers.cloudflare.com/fundamentals/api/get-started/keys/#limitations), API tokens should be used instead. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.", consts.APIKeyEnvVarKey),
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`[0-9a-f]{37}`),
						"API key must be 37 characters long and only contain characters 0-9 and a-f (all lowercased)",
					),
					stringvalidator.AlsoRequires(path.Expressions{
						path.MatchRoot(consts.EmailSchemaKey),
					}...),
				},
			},

			consts.APITokenSchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("The API Token for operations. Alternatively, can be configured using the `%s` environment variable. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.", consts.APITokenEnvVarKey),
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`[A-Za-z0-9-_]{40}`),
						"API tokens must be 40 characters long and only contain characters a-z, A-Z, 0-9, hyphens and underscores",
					),
				},
			},

			consts.APIUserServiceKeySchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("A special Cloudflare API key good for a restricted set of endpoints. Alternatively, can be configured using the `%s` environment variable. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.", consts.APIUserServiceKeyEnvVarKey),
			},

			consts.RPSSchemaKey: schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("RPS limit to apply when making calls to the API. Alternatively, can be configured using the `%s` environment variable.", consts.RPSEnvVarKey),
			},

			consts.RetriesSchemaKey: schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("Maximum number of retries to perform when an API request fails. Alternatively, can be configured using the `%s` environment variable.", consts.RetriesEnvVarKey),
			},

			consts.MinimumBackoffSchemaKey: schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("Minimum backoff period in seconds after failed API calls. Alternatively, can be configured using the `%s` environment variable.", consts.MinimumBackoffEnvVar),
			},

			consts.MaximumBackoffSchemaKey: schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("Maximum backoff period in seconds after failed API calls. Alternatively, can be configured using the `%s` environment variable.", consts.MaximumBackoffEnvVarKey),
			},

			consts.APIClientLoggingSchemaKey: schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("Whether to print logs from the API client (using the default log library logger). Alternatively, can be configured using the `%s` environment variable.", consts.APIClientLoggingEnvVarKey),
			},

			consts.APIHostnameSchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("Configure the hostname used by the API client. Alternatively, can be configured using the `%s` environment variable.", consts.APIHostnameEnvVarKey),
			},

			consts.APIBasePathSchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("Configure the base path used by the API client. Alternatively, can be configured using the `%s` environment variable.", consts.APIBasePathEnvVarKey),
			},

			consts.UserAgentOperatorSuffixSchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("A value to append to the HTTP User Agent for all API calls. This value is not something most users need to modify however, if you are using a non-standard provider or operator configuration, this is recommended to assist in uniquely identifying your traffic. **Setting this value will remove the Terraform version from the HTTP User Agent string and may have unintended consequences**. Alternatively, can be configured using the `%s` environment variable.", consts.UserAgentOperatorSuffixEnvVarKey),
			},
		},
	}
}

func (p *CloudflareProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var (
		data CloudflareProviderModel

		email             string
		apiKey            string
		apiToken          string
		apiUserServiceKey string
		rps               int64
		retries           int64
		minBackOff        int64
		maxBackOff        int64
		baseHostname      string
		basePath          string
	)

	cfv2Options := make([]option.RequestOption, 0)

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.APIHostname.ValueString() != "" {
		baseHostname = data.APIHostname.ValueString()
	} else {
		baseHostname = utils.GetDefaultFromEnv(consts.APIHostnameEnvVarKey, consts.APIHostnameDefault)
	}

	if data.APIBasePath.ValueString() != "" {
		basePath = data.APIBasePath.ValueString()
	} else {
		basePath = utils.GetDefaultFromEnv(consts.APIBasePathEnvVarKey, consts.APIBasePathDefault)
	}

	baseURL := cfv1.BaseURL(fmt.Sprintf("https://%s%s", baseHostname, basePath))

	// Ensure there is a trailing slash for client.V2 basePath
	basePathV2 := strings.TrimSuffix(basePath, "/") + "/"
	cfv2Options = append(cfv2Options, option.WithBaseURL(fmt.Sprintf("https://%s%s", baseHostname, basePathV2)))

	if !data.RPS.IsNull() {
		rps = int64(data.RPS.ValueInt64())
	} else {
		i, _ := strconv.ParseInt(utils.GetDefaultFromEnv(consts.RPSEnvVarKey, consts.RPSDefault), 10, 64)
		rps = i
	}
	limitOpt := cfv1.UsingRateLimit(float64(rps))

	if !data.Retries.IsNull() {
		retries = int64(data.Retries.ValueInt64())
	} else {
		i, _ := strconv.ParseInt(utils.GetDefaultFromEnv(consts.RetriesEnvVarKey, consts.RetriesDefault), 10, 64)
		retries = i
	}

	if !data.MinBackOff.IsNull() {
		minBackOff = int64(data.MaxBackoff.ValueInt64())
	} else {
		i, _ := strconv.ParseInt(utils.GetDefaultFromEnv(consts.MinimumBackoffEnvVar, consts.MinimumBackoffDefault), 10, 64)
		minBackOff = i
	}

	if !data.MinBackOff.IsNull() {
		maxBackOff = int64(data.MaxBackoff.ValueInt64())
	} else {
		i, _ := strconv.ParseInt(utils.GetDefaultFromEnv(consts.MaximumBackoffEnvVarKey, consts.MaximumBackoffDefault), 10, 64)
		maxBackOff = i
	}

	if retries >= math.MaxInt32 {
		resp.Diagnostics.AddError(
			fmt.Sprintf("retries value of %d is too large, try a smaller value.", retries),
			fmt.Sprintf("retries value of %d is too large, try a smaller value.", retries),
		)
		return
	}

	if minBackOff >= math.MaxInt32 {
		resp.Diagnostics.AddError(
			fmt.Sprintf("min_backoff value of %d is too large, try a smaller value.", minBackOff),
			fmt.Sprintf("min_backoff value of %d is too large, try a smaller value.", minBackOff),
		)
		return
	}

	if maxBackOff >= math.MaxInt32 {
		resp.Diagnostics.AddError(
			fmt.Sprintf("max_backoff value of %d is too large, try a smaller value.", maxBackOff),
			fmt.Sprintf("max_backoff value of %d is too large, try a smaller value.", maxBackOff),
		)
		return
	}

	retryOpt := cfv1.UsingRetryPolicy(int(retries), int(minBackOff), int(maxBackOff))
	cfv1Options := []cfv1.Option{limitOpt, retryOpt, baseURL}
	cfv1Options = append(cfv1Options, cfv1.Debug(logging.IsDebugOrHigher()))

	pluginVersion := utils.FindGoModuleVersion("github.com/hashicorp/terraform-plugin-framework")
	userAgentParams := utils.UserAgentBuilderParams{
		ProviderVersion: &p.version,
		PluginType:      cfv1.StringPtr("terraform-plugin-framework"),
		PluginVersion:   pluginVersion,
	}
	if !data.UserAgentOperatorSuffix.IsNull() {
		userAgentParams.OperatorSuffix = cfv1.StringPtr(data.UserAgentOperatorSuffix.String())
	} else {
		userAgentParams.TerraformVersion = cfv1.StringPtr(req.TerraformVersion)
	}
	cfv1Options = append(cfv1Options, cfv1.UserAgent(userAgentParams.String()))
	cfv2Options = append(cfv2Options, option.WithHeader("user-agent", userAgentParams.String()))

	v1Config := Config{Options: cfv1Options}

	if !data.APIToken.IsNull() {
		apiToken = data.APIToken.ValueString()
	} else {
		apiToken = utils.GetDefaultFromEnv(consts.APITokenEnvVarKey, "")
	}

	if apiToken != "" {
		v1Config.APIToken = apiToken
		cfv2Options = append(cfv2Options, option.WithAPIToken(apiToken))
	}

	if !data.APIKey.IsNull() {
		apiKey = data.APIKey.ValueString()
	} else {
		apiKey = utils.GetDefaultFromEnv(consts.APIKeyEnvVarKey, "")
	}

	if apiKey != "" {
		v1Config.APIKey = apiKey
		cfv2Options = append(cfv2Options, option.WithAPIKey(apiKey))

		if !data.Email.IsNull() {
			email = data.Email.ValueString()
		} else {
			email = utils.GetDefaultFromEnv(consts.EmailEnvVarKey, "")
		}

		if email == "" {
			resp.Diagnostics.AddError(
				fmt.Sprintf("%q is not set correctly", consts.EmailSchemaKey),
				fmt.Sprintf("%q is required with %q and was not configured", consts.EmailSchemaKey, consts.APIKeySchemaKey))
			return
		}

		if email != "" {
			v1Config.Email = email
			cfv2Options = append(cfv2Options, option.WithAPIEmail(email))
		}
	}

	if !data.APIUserServiceKey.IsNull() {
		apiUserServiceKey = data.APIUserServiceKey.ValueString()
	} else {
		apiUserServiceKey = utils.GetDefaultFromEnv(consts.APIUserServiceKeyEnvVarKey, "")
	}

	if apiUserServiceKey != "" {
		v1Config.APIUserServiceKey = apiUserServiceKey
		cfv2Options = append(cfv2Options, option.WithUserServiceKey(apiUserServiceKey))
	}

	if apiKey == "" && apiToken == "" && apiUserServiceKey == "" {
		resp.Diagnostics.AddError(
			fmt.Sprintf("must provide one of %q, %q or %q.", consts.APIKeySchemaKey, consts.APITokenSchemaKey, consts.APIUserServiceKeySchemaKey),
			fmt.Sprintf("must provide one of %q, %q or %q.", consts.APIKeySchemaKey, consts.APITokenSchemaKey, consts.APIUserServiceKeySchemaKey),
		)
		return
	}

	v1Config.Options = cfv1Options
	v1Client, err := v1Config.Client(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to initialize a new v1 client",
			err.Error(),
		)
		return
	}

	v2Client := cfv2.NewClient(cfv2Options...)

	client := &muxclient.Client{
		V1: v1Client,
		V2: v2Client,
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CloudflareProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		snippet_rules.NewResource,
		snippets.NewResource,
		cloud_connector_rules.NewResource,
		d1.NewResource,
		email_routing_address.NewResource,
		email_routing_rule.NewResource,
		hyperdrive_config.NewResource,
		list_item.NewResource,
		list.NewResource,
		r2_bucket.NewResource,
		risk_behavior.NewResource,
		zero_trust_risk_behavior.NewResource,
		rulesets.NewResource,
		turnstile.NewResource,
		access_mutual_tls_hostname_settings.NewResource,
		zero_trust_access_mtls_hostname_settings.NewResource,
		workers_for_platforms_dispatch_namespace_deprecated.NewResource,
		workers_for_platforms_dispatch_namespace.NewResource,
		zero_trust_risk_score_integration.NewResource,
		infrastructure_access_target_deprecated.NewResource,
		zero_trust_infrastructure_access_target.NewResource,
		leaked_credential_check.NewResource,
		leaked_credential_check_rule.NewResource,
		content_scanning.NewResource,
		content_scanning_expression.NewResource,
		apishieldoperation.NewResource,
	}
}

func (p *CloudflareProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		api_token_permissions_groups.NewDataSource,
		origin_ca_certificate.NewDataSource,
		user.NewDataSource,
		dlp_datasets.NewDataSource,
		gateway_categories.NewDataSource,
		gateway_app_types.NewDataSource,
		dcv_delegation.NewDataSource,
		infrastructure_access_target_deprecated.NewDataSource,
		zero_trust_infrastructure_access_target.NewDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudflareProvider{
			version: version,
		}
	}
}

var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cloudflare": func() (tfprotov6.ProviderServer, error) {
		upgradedSdkProvider, err := tf5to6server.UpgradeServer(context.Background(), sdkv2provider.New("dev")().GRPCProvider)
		if err != nil {
			log.Fatal(err)
		}
		providers := []func() tfprotov6.ProviderServer{
			func() tfprotov6.ProviderServer {
				return upgradedSdkProvider
			},
			providerserver.NewProtocol6(New("dev")()),
		}

		return tf6muxserver.NewMuxServer(context.Background(), providers...)
	},
}

func TestAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}
