// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a Lambda function alias. Creates an alias that points to the specified Lambda function version.
//
// For information about Lambda and how to use it, see [What is AWS Lambda?](http://docs.aws.amazon.com/lambda/latest/dg/welcome.html)
// For information about function aliases, see [CreateAlias](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateAlias.html) and [AliasRoutingConfiguration](https://docs.aws.amazon.com/lambda/latest/dg/API_AliasRoutingConfiguration.html) in the API docs.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = lambda.NewAlias(ctx, "testAlias", &lambda.AliasArgs{
// 			Description:     pulumi.String("a sample description"),
// 			FunctionName:    pulumi.String(aws_lambda_function.Lambda_function_test.Arn),
// 			FunctionVersion: pulumi.String("1"),
// 			RoutingConfig: &lambda.AliasRoutingConfigArgs{
// 				AdditionalVersionWeights: pulumi.Map{
// 					"2": pulumi.Float64(0.5),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Alias struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) identifying your Lambda function alias.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the alias.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The function ARN of the Lambda function for which you want to create an alias.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`
	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion pulumi.StringOutput `pulumi:"functionVersion"`
	// The ARN to be used for invoking Lambda Function from API Gateway - to be used in `apigateway.Integration`'s `uri`
	InvokeArn pulumi.StringOutput `pulumi:"invokeArn"`
	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name pulumi.StringOutput `pulumi:"name"`
	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig AliasRoutingConfigPtrOutput `pulumi:"routingConfig"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
	if args == nil || args.FunctionVersion == nil {
		return nil, errors.New("missing required argument 'FunctionVersion'")
	}
	if args == nil {
		args = &AliasArgs{}
	}
	var resource Alias
	err := ctx.RegisterResource("aws:lambda/alias:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("aws:lambda/alias:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
	// The Amazon Resource Name (ARN) identifying your Lambda function alias.
	Arn *string `pulumi:"arn"`
	// Description of the alias.
	Description *string `pulumi:"description"`
	// The function ARN of the Lambda function for which you want to create an alias.
	FunctionName *string `pulumi:"functionName"`
	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion *string `pulumi:"functionVersion"`
	// The ARN to be used for invoking Lambda Function from API Gateway - to be used in `apigateway.Integration`'s `uri`
	InvokeArn *string `pulumi:"invokeArn"`
	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name *string `pulumi:"name"`
	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig *AliasRoutingConfig `pulumi:"routingConfig"`
}

type AliasState struct {
	// The Amazon Resource Name (ARN) identifying your Lambda function alias.
	Arn pulumi.StringPtrInput
	// Description of the alias.
	Description pulumi.StringPtrInput
	// The function ARN of the Lambda function for which you want to create an alias.
	FunctionName pulumi.StringPtrInput
	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion pulumi.StringPtrInput
	// The ARN to be used for invoking Lambda Function from API Gateway - to be used in `apigateway.Integration`'s `uri`
	InvokeArn pulumi.StringPtrInput
	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name pulumi.StringPtrInput
	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig AliasRoutingConfigPtrInput
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// Description of the alias.
	Description *string `pulumi:"description"`
	// The function ARN of the Lambda function for which you want to create an alias.
	FunctionName string `pulumi:"functionName"`
	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion string `pulumi:"functionVersion"`
	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name *string `pulumi:"name"`
	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig *AliasRoutingConfig `pulumi:"routingConfig"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// Description of the alias.
	Description pulumi.StringPtrInput
	// The function ARN of the Lambda function for which you want to create an alias.
	FunctionName pulumi.StringInput
	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion pulumi.StringInput
	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name pulumi.StringPtrInput
	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig AliasRoutingConfigPtrInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}
