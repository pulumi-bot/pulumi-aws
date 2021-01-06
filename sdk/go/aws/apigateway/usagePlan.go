// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway Usage Plan.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myapi, err := apigateway.NewRestApi(ctx, "myapi", nil)
// 		if err != nil {
// 			return err
// 		}
// 		dev, err := apigateway.NewDeployment(ctx, "dev", &apigateway.DeploymentArgs{
// 			RestApi:   myapi.ID(),
// 			StageName: pulumi.String("dev"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		prod, err := apigateway.NewDeployment(ctx, "prod", &apigateway.DeploymentArgs{
// 			RestApi:   myapi.ID(),
// 			StageName: pulumi.String("prod"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewUsagePlan(ctx, "myUsagePlan", &apigateway.UsagePlanArgs{
// 			Description: pulumi.String("my description"),
// 			ProductCode: pulumi.String("MYCODE"),
// 			ApiStages: apigateway.UsagePlanApiStageArray{
// 				&apigateway.UsagePlanApiStageArgs{
// 					ApiId: myapi.ID(),
// 					Stage: dev.StageName,
// 				},
// 				&apigateway.UsagePlanApiStageArgs{
// 					ApiId: myapi.ID(),
// 					Stage: prod.StageName,
// 				},
// 			},
// 			QuotaSettings: &apigateway.UsagePlanQuotaSettingsArgs{
// 				Limit:  pulumi.Int(20),
// 				Offset: pulumi.Int(2),
// 				Period: pulumi.String("WEEK"),
// 			},
// 			ThrottleSettings: &apigateway.UsagePlanThrottleSettingsArgs{
// 				BurstLimit: pulumi.Int(5),
// 				RateLimit:  pulumi.Float64(10),
// 			},
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
// AWS API Gateway Usage Plan can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/usagePlan:UsagePlan myusageplan <usage_plan_id>
// ```
type UsagePlan struct {
	pulumi.CustomResourceState

	// The associated API stages of the usage plan.
	ApiStages UsagePlanApiStageArrayOutput `pulumi:"apiStages"`
	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of a usage plan.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the usage plan.
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode pulumi.StringPtrOutput `pulumi:"productCode"`
	// The quota settings of the usage plan.
	QuotaSettings UsagePlanQuotaSettingsPtrOutput `pulumi:"quotaSettings"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The throttling limits of the usage plan.
	ThrottleSettings UsagePlanThrottleSettingsPtrOutput `pulumi:"throttleSettings"`
}

// NewUsagePlan registers a new resource with the given unique name, arguments, and options.
func NewUsagePlan(ctx *pulumi.Context,
	name string, args *UsagePlanArgs, opts ...pulumi.ResourceOption) (*UsagePlan, error) {
	if args == nil {
		args = &UsagePlanArgs{}
	}

	var resource UsagePlan
	err := ctx.RegisterResource("aws:apigateway/usagePlan:UsagePlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUsagePlan gets an existing UsagePlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUsagePlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UsagePlanState, opts ...pulumi.ResourceOption) (*UsagePlan, error) {
	var resource UsagePlan
	err := ctx.ReadResource("aws:apigateway/usagePlan:UsagePlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UsagePlan resources.
type usagePlanState struct {
	// The associated API stages of the usage plan.
	ApiStages []UsagePlanApiStage `pulumi:"apiStages"`
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The description of a usage plan.
	Description *string `pulumi:"description"`
	// The name of the usage plan.
	Name *string `pulumi:"name"`
	// The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode *string `pulumi:"productCode"`
	// The quota settings of the usage plan.
	QuotaSettings *UsagePlanQuotaSettings `pulumi:"quotaSettings"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The throttling limits of the usage plan.
	ThrottleSettings *UsagePlanThrottleSettings `pulumi:"throttleSettings"`
}

type UsagePlanState struct {
	// The associated API stages of the usage plan.
	ApiStages UsagePlanApiStageArrayInput
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The description of a usage plan.
	Description pulumi.StringPtrInput
	// The name of the usage plan.
	Name pulumi.StringPtrInput
	// The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode pulumi.StringPtrInput
	// The quota settings of the usage plan.
	QuotaSettings UsagePlanQuotaSettingsPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// The throttling limits of the usage plan.
	ThrottleSettings UsagePlanThrottleSettingsPtrInput
}

func (UsagePlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*usagePlanState)(nil)).Elem()
}

type usagePlanArgs struct {
	// The associated API stages of the usage plan.
	ApiStages []UsagePlanApiStage `pulumi:"apiStages"`
	// The description of a usage plan.
	Description *string `pulumi:"description"`
	// The name of the usage plan.
	Name *string `pulumi:"name"`
	// The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode *string `pulumi:"productCode"`
	// The quota settings of the usage plan.
	QuotaSettings *UsagePlanQuotaSettings `pulumi:"quotaSettings"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The throttling limits of the usage plan.
	ThrottleSettings *UsagePlanThrottleSettings `pulumi:"throttleSettings"`
}

// The set of arguments for constructing a UsagePlan resource.
type UsagePlanArgs struct {
	// The associated API stages of the usage plan.
	ApiStages UsagePlanApiStageArrayInput
	// The description of a usage plan.
	Description pulumi.StringPtrInput
	// The name of the usage plan.
	Name pulumi.StringPtrInput
	// The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode pulumi.StringPtrInput
	// The quota settings of the usage plan.
	QuotaSettings UsagePlanQuotaSettingsPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// The throttling limits of the usage plan.
	ThrottleSettings UsagePlanThrottleSettingsPtrInput
}

func (UsagePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*usagePlanArgs)(nil)).Elem()
}

type UsagePlanInput interface {
	pulumi.Input

	ToUsagePlanOutput() UsagePlanOutput
	ToUsagePlanOutputWithContext(ctx context.Context) UsagePlanOutput
}

func (*UsagePlan) ElementType() reflect.Type {
	return reflect.TypeOf((*UsagePlan)(nil))
}

func (i *UsagePlan) ToUsagePlanOutput() UsagePlanOutput {
	return i.ToUsagePlanOutputWithContext(context.Background())
}

func (i *UsagePlan) ToUsagePlanOutputWithContext(ctx context.Context) UsagePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanOutput)
}

func (i *UsagePlan) ToUsagePlanPtrOutput() UsagePlanPtrOutput {
	return i.ToUsagePlanPtrOutputWithContext(context.Background())
}

func (i *UsagePlan) ToUsagePlanPtrOutputWithContext(ctx context.Context) UsagePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanPtrOutput)
}

type UsagePlanPtrInput interface {
	pulumi.Input

	ToUsagePlanPtrOutput() UsagePlanPtrOutput
	ToUsagePlanPtrOutputWithContext(ctx context.Context) UsagePlanPtrOutput
}

type usagePlanPtrType UsagePlanArgs

func (*usagePlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UsagePlan)(nil))
}

func (i *usagePlanPtrType) ToUsagePlanPtrOutput() UsagePlanPtrOutput {
	return i.ToUsagePlanPtrOutputWithContext(context.Background())
}

func (i *usagePlanPtrType) ToUsagePlanPtrOutputWithContext(ctx context.Context) UsagePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanOutput).ToUsagePlanPtrOutput()
}

// UsagePlanArrayInput is an input type that accepts UsagePlanArray and UsagePlanArrayOutput values.
// You can construct a concrete instance of `UsagePlanArrayInput` via:
//
//          UsagePlanArray{ UsagePlanArgs{...} }
type UsagePlanArrayInput interface {
	pulumi.Input

	ToUsagePlanArrayOutput() UsagePlanArrayOutput
	ToUsagePlanArrayOutputWithContext(context.Context) UsagePlanArrayOutput
}

type UsagePlanArray []UsagePlanInput

func (UsagePlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UsagePlan)(nil))
}

func (i UsagePlanArray) ToUsagePlanArrayOutput() UsagePlanArrayOutput {
	return i.ToUsagePlanArrayOutputWithContext(context.Background())
}

func (i UsagePlanArray) ToUsagePlanArrayOutputWithContext(ctx context.Context) UsagePlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanArrayOutput)
}

// UsagePlanMapInput is an input type that accepts UsagePlanMap and UsagePlanMapOutput values.
// You can construct a concrete instance of `UsagePlanMapInput` via:
//
//          UsagePlanMap{ "key": UsagePlanArgs{...} }
type UsagePlanMapInput interface {
	pulumi.Input

	ToUsagePlanMapOutput() UsagePlanMapOutput
	ToUsagePlanMapOutputWithContext(context.Context) UsagePlanMapOutput
}

type UsagePlanMap map[string]UsagePlanInput

func (UsagePlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UsagePlan)(nil))
}

func (i UsagePlanMap) ToUsagePlanMapOutput() UsagePlanMapOutput {
	return i.ToUsagePlanMapOutputWithContext(context.Background())
}

func (i UsagePlanMap) ToUsagePlanMapOutputWithContext(ctx context.Context) UsagePlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanMapOutput)
}

type UsagePlanOutput struct {
	*pulumi.OutputState
}

func (UsagePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UsagePlan)(nil))
}

func (o UsagePlanOutput) ToUsagePlanOutput() UsagePlanOutput {
	return o
}

func (o UsagePlanOutput) ToUsagePlanOutputWithContext(ctx context.Context) UsagePlanOutput {
	return o
}

func (o UsagePlanOutput) ToUsagePlanPtrOutput() UsagePlanPtrOutput {
	return o.ToUsagePlanPtrOutputWithContext(context.Background())
}

func (o UsagePlanOutput) ToUsagePlanPtrOutputWithContext(ctx context.Context) UsagePlanPtrOutput {
	return o.ApplyT(func(v UsagePlan) *UsagePlan {
		return &v
	}).(UsagePlanPtrOutput)
}

type UsagePlanPtrOutput struct {
	*pulumi.OutputState
}

func (UsagePlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UsagePlan)(nil))
}

func (o UsagePlanPtrOutput) ToUsagePlanPtrOutput() UsagePlanPtrOutput {
	return o
}

func (o UsagePlanPtrOutput) ToUsagePlanPtrOutputWithContext(ctx context.Context) UsagePlanPtrOutput {
	return o
}

type UsagePlanArrayOutput struct{ *pulumi.OutputState }

func (UsagePlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UsagePlan)(nil))
}

func (o UsagePlanArrayOutput) ToUsagePlanArrayOutput() UsagePlanArrayOutput {
	return o
}

func (o UsagePlanArrayOutput) ToUsagePlanArrayOutputWithContext(ctx context.Context) UsagePlanArrayOutput {
	return o
}

func (o UsagePlanArrayOutput) Index(i pulumi.IntInput) UsagePlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UsagePlan {
		return vs[0].([]UsagePlan)[vs[1].(int)]
	}).(UsagePlanOutput)
}

type UsagePlanMapOutput struct{ *pulumi.OutputState }

func (UsagePlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UsagePlan)(nil))
}

func (o UsagePlanMapOutput) ToUsagePlanMapOutput() UsagePlanMapOutput {
	return o
}

func (o UsagePlanMapOutput) ToUsagePlanMapOutputWithContext(ctx context.Context) UsagePlanMapOutput {
	return o
}

func (o UsagePlanMapOutput) MapIndex(k pulumi.StringInput) UsagePlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UsagePlan {
		return vs[0].(map[string]UsagePlan)[vs[1].(string)]
	}).(UsagePlanOutput)
}

func init() {
	pulumi.RegisterOutputType(UsagePlanOutput{})
	pulumi.RegisterOutputType(UsagePlanPtrOutput{})
	pulumi.RegisterOutputType(UsagePlanArrayOutput{})
	pulumi.RegisterOutputType(UsagePlanMapOutput{})
}
