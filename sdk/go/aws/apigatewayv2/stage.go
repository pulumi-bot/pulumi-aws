// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Stage struct {
	pulumi.CustomResourceState

	AccessLogSettings    StageAccessLogSettingsPtrOutput    `pulumi:"accessLogSettings"`
	ApiId                pulumi.StringOutput                `pulumi:"apiId"`
	Arn                  pulumi.StringOutput                `pulumi:"arn"`
	AutoDeploy           pulumi.BoolPtrOutput               `pulumi:"autoDeploy"`
	ClientCertificateId  pulumi.StringPtrOutput             `pulumi:"clientCertificateId"`
	DefaultRouteSettings StageDefaultRouteSettingsPtrOutput `pulumi:"defaultRouteSettings"`
	DeploymentId         pulumi.StringOutput                `pulumi:"deploymentId"`
	Description          pulumi.StringPtrOutput             `pulumi:"description"`
	ExecutionArn         pulumi.StringOutput                `pulumi:"executionArn"`
	InvokeUrl            pulumi.StringOutput                `pulumi:"invokeUrl"`
	Name                 pulumi.StringOutput                `pulumi:"name"`
	RouteSettings        StageRouteSettingArrayOutput       `pulumi:"routeSettings"`
	StageVariables       pulumi.StringMapOutput             `pulumi:"stageVariables"`
	Tags                 pulumi.StringMapOutput             `pulumi:"tags"`
}

// NewStage registers a new resource with the given unique name, arguments, and options.
func NewStage(ctx *pulumi.Context,
	name string, args *StageArgs, opts ...pulumi.ResourceOption) (*Stage, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil {
		args = &StageArgs{}
	}
	var resource Stage
	err := ctx.RegisterResource("aws:apigatewayv2/stage:Stage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStage gets an existing Stage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageState, opts ...pulumi.ResourceOption) (*Stage, error) {
	var resource Stage
	err := ctx.ReadResource("aws:apigatewayv2/stage:Stage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stage resources.
type stageState struct {
	AccessLogSettings    *StageAccessLogSettings    `pulumi:"accessLogSettings"`
	ApiId                *string                    `pulumi:"apiId"`
	Arn                  *string                    `pulumi:"arn"`
	AutoDeploy           *bool                      `pulumi:"autoDeploy"`
	ClientCertificateId  *string                    `pulumi:"clientCertificateId"`
	DefaultRouteSettings *StageDefaultRouteSettings `pulumi:"defaultRouteSettings"`
	DeploymentId         *string                    `pulumi:"deploymentId"`
	Description          *string                    `pulumi:"description"`
	ExecutionArn         *string                    `pulumi:"executionArn"`
	InvokeUrl            *string                    `pulumi:"invokeUrl"`
	Name                 *string                    `pulumi:"name"`
	RouteSettings        []StageRouteSetting        `pulumi:"routeSettings"`
	StageVariables       map[string]string          `pulumi:"stageVariables"`
	Tags                 map[string]string          `pulumi:"tags"`
}

type StageState struct {
	AccessLogSettings    StageAccessLogSettingsPtrInput
	ApiId                pulumi.StringPtrInput
	Arn                  pulumi.StringPtrInput
	AutoDeploy           pulumi.BoolPtrInput
	ClientCertificateId  pulumi.StringPtrInput
	DefaultRouteSettings StageDefaultRouteSettingsPtrInput
	DeploymentId         pulumi.StringPtrInput
	Description          pulumi.StringPtrInput
	ExecutionArn         pulumi.StringPtrInput
	InvokeUrl            pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	RouteSettings        StageRouteSettingArrayInput
	StageVariables       pulumi.StringMapInput
	Tags                 pulumi.StringMapInput
}

func (StageState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageState)(nil)).Elem()
}

type stageArgs struct {
	AccessLogSettings    *StageAccessLogSettings    `pulumi:"accessLogSettings"`
	ApiId                string                     `pulumi:"apiId"`
	AutoDeploy           *bool                      `pulumi:"autoDeploy"`
	ClientCertificateId  *string                    `pulumi:"clientCertificateId"`
	DefaultRouteSettings *StageDefaultRouteSettings `pulumi:"defaultRouteSettings"`
	DeploymentId         *string                    `pulumi:"deploymentId"`
	Description          *string                    `pulumi:"description"`
	Name                 *string                    `pulumi:"name"`
	RouteSettings        []StageRouteSetting        `pulumi:"routeSettings"`
	StageVariables       map[string]string          `pulumi:"stageVariables"`
	Tags                 map[string]string          `pulumi:"tags"`
}

// The set of arguments for constructing a Stage resource.
type StageArgs struct {
	AccessLogSettings    StageAccessLogSettingsPtrInput
	ApiId                pulumi.StringInput
	AutoDeploy           pulumi.BoolPtrInput
	ClientCertificateId  pulumi.StringPtrInput
	DefaultRouteSettings StageDefaultRouteSettingsPtrInput
	DeploymentId         pulumi.StringPtrInput
	Description          pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	RouteSettings        StageRouteSettingArrayInput
	StageVariables       pulumi.StringMapInput
	Tags                 pulumi.StringMapInput
}

func (StageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageArgs)(nil)).Elem()
}
