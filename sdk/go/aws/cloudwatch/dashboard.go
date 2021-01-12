// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CloudWatch Dashboard resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudwatch.NewDashboard(ctx, "main", &cloudwatch.DashboardArgs{
// 			DashboardBody: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"widgets\": [\n", "    {\n", "      \"type\": \"metric\",\n", "      \"x\": 0,\n", "      \"y\": 0,\n", "      \"width\": 12,\n", "      \"height\": 6,\n", "      \"properties\": {\n", "        \"metrics\": [\n", "          [\n", "            \"AWS/EC2\",\n", "            \"CPUUtilization\",\n", "            \"InstanceId\",\n", "            \"i-012345\"\n", "          ]\n", "        ],\n", "        \"period\": 300,\n", "        \"stat\": \"Average\",\n", "        \"region\": \"us-east-1\",\n", "        \"title\": \"EC2 Instance CPU\"\n", "      }\n", "    },\n", "    {\n", "      \"type\": \"text\",\n", "      \"x\": 0,\n", "      \"y\": 7,\n", "      \"width\": 3,\n", "      \"height\": 3,\n", "      \"properties\": {\n", "        \"markdown\": \"Hello world\"\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n")),
// 			DashboardName: pulumi.String("my-dashboard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// CloudWatch dashboards can be imported using the `dashboard_name`, e.g.
//
// ```sh
//  $ pulumi import aws:cloudwatch/dashboard:Dashboard sample <dashboard_name>
// ```
type Dashboard struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn pulumi.StringOutput `pulumi:"dashboardArn"`
	// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	DashboardBody pulumi.StringOutput `pulumi:"dashboardBody"`
	// The name of the dashboard.
	DashboardName pulumi.StringOutput `pulumi:"dashboardName"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DashboardBody == nil {
		return nil, errors.New("invalid value for required argument 'DashboardBody'")
	}
	if args.DashboardName == nil {
		return nil, errors.New("invalid value for required argument 'DashboardName'")
	}
	var resource Dashboard
	err := ctx.RegisterResource("aws:cloudwatch/dashboard:Dashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboard gets an existing Dashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardState, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	var resource Dashboard
	err := ctx.ReadResource("aws:cloudwatch/dashboard:Dashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dashboard resources.
type dashboardState struct {
	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string `pulumi:"dashboardArn"`
	// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	DashboardBody *string `pulumi:"dashboardBody"`
	// The name of the dashboard.
	DashboardName *string `pulumi:"dashboardName"`
}

type DashboardState struct {
	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn pulumi.StringPtrInput
	// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	DashboardBody pulumi.StringPtrInput
	// The name of the dashboard.
	DashboardName pulumi.StringPtrInput
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	DashboardBody string `pulumi:"dashboardBody"`
	// The name of the dashboard.
	DashboardName string `pulumi:"dashboardName"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	DashboardBody pulumi.StringInput
	// The name of the dashboard.
	DashboardName pulumi.StringInput
}

func (DashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardArgs)(nil)).Elem()
}

type DashboardInput interface {
	pulumi.Input

	ToDashboardOutput() DashboardOutput
	ToDashboardOutputWithContext(ctx context.Context) DashboardOutput
}

func (*Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((*Dashboard)(nil))
}

func (i *Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

func (i *Dashboard) ToDashboardPtrOutput() DashboardPtrOutput {
	return i.ToDashboardPtrOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardPtrOutputWithContext(ctx context.Context) DashboardPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPtrOutput)
}

type DashboardPtrInput interface {
	pulumi.Input

	ToDashboardPtrOutput() DashboardPtrOutput
	ToDashboardPtrOutputWithContext(ctx context.Context) DashboardPtrOutput
}

type dashboardPtrType DashboardArgs

func (*dashboardPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil))
}

func (i *dashboardPtrType) ToDashboardPtrOutput() DashboardPtrOutput {
	return i.ToDashboardPtrOutputWithContext(context.Background())
}

func (i *dashboardPtrType) ToDashboardPtrOutputWithContext(ctx context.Context) DashboardPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput).ToDashboardPtrOutput()
}

// DashboardArrayInput is an input type that accepts DashboardArray and DashboardArrayOutput values.
// You can construct a concrete instance of `DashboardArrayInput` via:
//
//          DashboardArray{ DashboardArgs{...} }
type DashboardArrayInput interface {
	pulumi.Input

	ToDashboardArrayOutput() DashboardArrayOutput
	ToDashboardArrayOutputWithContext(context.Context) DashboardArrayOutput
}

type DashboardArray []DashboardInput

func (DashboardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Dashboard)(nil))
}

func (i DashboardArray) ToDashboardArrayOutput() DashboardArrayOutput {
	return i.ToDashboardArrayOutputWithContext(context.Background())
}

func (i DashboardArray) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardArrayOutput)
}

// DashboardMapInput is an input type that accepts DashboardMap and DashboardMapOutput values.
// You can construct a concrete instance of `DashboardMapInput` via:
//
//          DashboardMap{ "key": DashboardArgs{...} }
type DashboardMapInput interface {
	pulumi.Input

	ToDashboardMapOutput() DashboardMapOutput
	ToDashboardMapOutputWithContext(context.Context) DashboardMapOutput
}

type DashboardMap map[string]DashboardInput

func (DashboardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Dashboard)(nil))
}

func (i DashboardMap) ToDashboardMapOutput() DashboardMapOutput {
	return i.ToDashboardMapOutputWithContext(context.Background())
}

func (i DashboardMap) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardMapOutput)
}

type DashboardOutput struct {
	*pulumi.OutputState
}

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dashboard)(nil))
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardPtrOutput() DashboardPtrOutput {
	return o.ToDashboardPtrOutputWithContext(context.Background())
}

func (o DashboardOutput) ToDashboardPtrOutputWithContext(ctx context.Context) DashboardPtrOutput {
	return o.ApplyT(func(v Dashboard) *Dashboard {
		return &v
	}).(DashboardPtrOutput)
}

type DashboardPtrOutput struct {
	*pulumi.OutputState
}

func (DashboardPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil))
}

func (o DashboardPtrOutput) ToDashboardPtrOutput() DashboardPtrOutput {
	return o
}

func (o DashboardPtrOutput) ToDashboardPtrOutputWithContext(ctx context.Context) DashboardPtrOutput {
	return o
}

type DashboardArrayOutput struct{ *pulumi.OutputState }

func (DashboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Dashboard)(nil))
}

func (o DashboardArrayOutput) ToDashboardArrayOutput() DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) Index(i pulumi.IntInput) DashboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Dashboard {
		return vs[0].([]Dashboard)[vs[1].(int)]
	}).(DashboardOutput)
}

type DashboardMapOutput struct{ *pulumi.OutputState }

func (DashboardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Dashboard)(nil))
}

func (o DashboardMapOutput) ToDashboardMapOutput() DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) MapIndex(k pulumi.StringInput) DashboardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Dashboard {
		return vs[0].(map[string]Dashboard)[vs[1].(string)]
	}).(DashboardOutput)
}

func init() {
	pulumi.RegisterOutputType(DashboardOutput{})
	pulumi.RegisterOutputType(DashboardPtrOutput{})
	pulumi.RegisterOutputType(DashboardArrayOutput{})
	pulumi.RegisterOutputType(DashboardMapOutput{})
}
