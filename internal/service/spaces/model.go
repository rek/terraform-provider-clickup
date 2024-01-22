package spaces

import "github.com/hashicorp/terraform-plugin-framework/types"

type ClickUpSpacesDataSourceModel struct {
	TeamId   types.String                  `tfsdk:"team_id"`
	Archived types.Bool                    `tfsdk:"archived"`
	Spaces   []ClickUpSpaceDataSourceModel `tfsdk:"spaces"`
}

type ClickUpSpaceDataSourceModel struct {
	Id                types.String                        `tfsdk:"id"`
	Name              types.String                        `tfsdk:"name"`
	Private           types.Bool                          `tfsdk:"private"`
	MultipleAssignees types.Bool                          `tfsdk:"multiple_assignees"`
	Statuses          []ClickUpSpaceStatusDataSourceModel `tfsdk:"statuses"`
}

type ClickUpSpaceStatusDataSourceModel struct {
	Status     types.String `tfsdk:"status"`
	Type       types.String `tfsdk:"type"`
	OrderIndex types.Int64  `tfsdk:"order_index"`
	Color      types.String `tfsdk:"color"`
}
