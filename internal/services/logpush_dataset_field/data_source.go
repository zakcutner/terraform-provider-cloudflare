// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_dataset_field

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type LogpushDatasetFieldDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*LogpushDatasetFieldDataSource)(nil)

func NewLogpushDatasetFieldDataSource() datasource.DataSource {
	return &LogpushDatasetFieldDataSource{}
}

func (d *LogpushDatasetFieldDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_logpush_dataset_field"
}

func (d *LogpushDatasetFieldDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *LogpushDatasetFieldDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *LogpushDatasetFieldDataSourceModel

	// resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	// params, diags := data.toReadParams(ctx)
	// resp.Diagnostics.Append(diags...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	// res := new(http.Response)
	// env := LogpushDatasetFieldResultDataSourceEnvelope{*data}
	// _, err := d.client.Logpush.Datasets.Fields.Get(
	// 	ctx,
	// 	logpush.DatasetFieldGetParamsDatasetID(data.DatasetID.ValueString()),
	// 	params,
	// 	option.WithResponseBodyInto(&res),
	// 	option.WithMiddleware(logging.Middleware(ctx)),
	// )
	// if err != nil {
	// 	resp.Diagnostics.AddError("failed to make http request", err.Error())
	// 	return
	// }
	// bytes, _ := io.ReadAll(res.Body)
	// err = apijson.UnmarshalComputed(bytes, &env)
	// if err != nil {
	// 	resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
	// 	return
	// }
	// data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
