// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apigateway

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an API Gateway Usage Plan.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_usage_plan.html.markdown.
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
	// The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode pulumi.StringPtrOutput `pulumi:"productCode"`
	// The quota settings of the usage plan.
	QuotaSettings UsagePlanQuotaSettingsPtrOutput `pulumi:"quotaSettings"`
	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
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
	// The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode *string `pulumi:"productCode"`
	// The quota settings of the usage plan.
	QuotaSettings *UsagePlanQuotaSettings `pulumi:"quotaSettings"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
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
	// The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode pulumi.StringPtrInput
	// The quota settings of the usage plan.
	QuotaSettings UsagePlanQuotaSettingsPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
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
	// The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode *string `pulumi:"productCode"`
	// The quota settings of the usage plan.
	QuotaSettings *UsagePlanQuotaSettings `pulumi:"quotaSettings"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
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
	// The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode pulumi.StringPtrInput
	// The quota settings of the usage plan.
	QuotaSettings UsagePlanQuotaSettingsPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
	// The throttling limits of the usage plan.
	ThrottleSettings UsagePlanThrottleSettingsPtrInput
}

func (UsagePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*usagePlanArgs)(nil)).Elem()
}

