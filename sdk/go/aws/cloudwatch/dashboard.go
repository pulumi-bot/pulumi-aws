// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CloudWatch Dashboard resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_dashboard.html.markdown.
type Dashboard struct {
	s *pulumi.ResourceState
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOpt) (*Dashboard, error) {
	if args == nil || args.DashboardBody == nil {
		return nil, errors.New("missing required argument 'DashboardBody'")
	}
	if args == nil || args.DashboardName == nil {
		return nil, errors.New("missing required argument 'DashboardName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dashboardBody"] = nil
		inputs["dashboardName"] = nil
	} else {
		inputs["dashboardBody"] = args.DashboardBody
		inputs["dashboardName"] = args.DashboardName
	}
	inputs["dashboardArn"] = nil
	s, err := ctx.RegisterResource("aws:cloudwatch/dashboard:Dashboard", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Dashboard{s: s}, nil
}

// GetDashboard gets an existing Dashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DashboardState, opts ...pulumi.ResourceOpt) (*Dashboard, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["dashboardArn"] = state.DashboardArn
		inputs["dashboardBody"] = state.DashboardBody
		inputs["dashboardName"] = state.DashboardName
	}
	s, err := ctx.ReadResource("aws:cloudwatch/dashboard:Dashboard", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Dashboard{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Dashboard) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Dashboard) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The Amazon Resource Name (ARN) of the dashboard.
func (r *Dashboard) DashboardArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dashboardArn"])
}

// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
func (r *Dashboard) DashboardBody() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dashboardBody"])
}

// The name of the dashboard.
func (r *Dashboard) DashboardName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dashboardName"])
}

// Input properties used for looking up and filtering Dashboard resources.
type DashboardState struct {
	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn interface{}
	// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	DashboardBody interface{}
	// The name of the dashboard.
	DashboardName interface{}
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	DashboardBody interface{}
	// The name of the dashboard.
	DashboardName interface{}
}
