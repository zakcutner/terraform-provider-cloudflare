// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ResourceGroupsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "ID of the resource group to be fetched.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the resource group to be fetched.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ResourceGroupsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier of the resource group.",
							Computed:    true,
						},
						"scope": schema.ListNestedAttribute{
							Description: "The scope associated to the resource group",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ResourceGroupsScopeDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description: "This is a combination of pre-defined resource name and identifier (like Account ID etc.)",
										Computed:    true,
									},
									"objects": schema.ListNestedAttribute{
										Description: "A list of scope objects for additional context.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[ResourceGroupsScopeObjectsDataSourceModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description: "This is a combination of pre-defined resource name and identifier (like Zone ID etc.)",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"meta": schema.SingleNestedAttribute{
							Description: "Attributes associated to the resource group.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ResourceGroupsMetaDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Computed: true,
								},
								"value": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "Name of the resource group.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ResourceGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ResourceGroupsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
