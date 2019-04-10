// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to invoke custom lambda functions as data source.
// The lambda function is invoked with [RequestResponse](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax)
// invocation type.
func LookupInvocation(ctx *pulumi.Context, args *GetInvocationArgs) (*GetInvocationResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["functionName"] = args.FunctionName
		inputs["input"] = args.Input
		inputs["qualifier"] = args.Qualifier
	}
	outputs, err := ctx.Invoke("aws:lambda/getInvocation:getInvocation", inputs)
	if err != nil {
		return nil, err
	}
	return &GetInvocationResult{
		FunctionName: outputs["functionName"],
		Input: outputs["input"],
		Qualifier: outputs["qualifier"],
		Result: outputs["result"],
		ResultMap: outputs["resultMap"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getInvocation.
type GetInvocationArgs struct {
	// The name of the lambda function.
	FunctionName interface{}
	// A string in JSON format that is passed as payload to the lambda function.
	Input interface{}
	// The qualifier (a.k.a version) of the lambda function. Defaults
	// to `$LATEST`.
	Qualifier interface{}
}

// A collection of values returned by getInvocation.
type GetInvocationResult struct {
	FunctionName interface{}
	Input interface{}
	Qualifier interface{}
	// A result of the lambda function invocation.
	Result interface{}
	// This field is set only if result is a map of primitive types.
	ResultMap interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
