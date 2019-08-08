// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about a Lambda Function.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lambda_function.html.markdown.
func LookupFunction(ctx *pulumi.Context, args *GetFunctionArgs) (*GetFunctionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["functionName"] = args.FunctionName
		inputs["qualifier"] = args.Qualifier
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:lambda/getFunction:getFunction", inputs)
	if err != nil {
		return nil, err
	}
	return &GetFunctionResult{
		Arn: outputs["arn"],
		DeadLetterConfig: outputs["deadLetterConfig"],
		Description: outputs["description"],
		Environment: outputs["environment"],
		FunctionName: outputs["functionName"],
		Handler: outputs["handler"],
		InvokeArn: outputs["invokeArn"],
		KmsKeyArn: outputs["kmsKeyArn"],
		LastModified: outputs["lastModified"],
		Layers: outputs["layers"],
		MemorySize: outputs["memorySize"],
		QualifiedArn: outputs["qualifiedArn"],
		Qualifier: outputs["qualifier"],
		ReservedConcurrentExecutions: outputs["reservedConcurrentExecutions"],
		Role: outputs["role"],
		Runtime: outputs["runtime"],
		SourceCodeHash: outputs["sourceCodeHash"],
		SourceCodeSize: outputs["sourceCodeSize"],
		Tags: outputs["tags"],
		Timeout: outputs["timeout"],
		TracingConfig: outputs["tracingConfig"],
		Version: outputs["version"],
		VpcConfig: outputs["vpcConfig"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getFunction.
type GetFunctionArgs struct {
	// Name of the lambda function.
	FunctionName interface{}
	// Alias name or version number of the lambda function. e.g. `$LATEST`, `my-alias`, or `1`
	Qualifier interface{}
	Tags interface{}
}

// A collection of values returned by getFunction.
type GetFunctionResult struct {
	// Unqualified (no `:QUALIFIER` or `:VERSION` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also `qualifiedArn`.
	Arn interface{}
	// Configure the function's *dead letter queue*.
	DeadLetterConfig interface{}
	// Description of what your Lambda Function does.
	Description interface{}
	// The Lambda environment's configuration settings.
	Environment interface{}
	FunctionName interface{}
	// The function entrypoint in your code.
	Handler interface{}
	// The ARN to be used for invoking Lambda Function from API Gateway.
	InvokeArn interface{}
	// The ARN for the KMS encryption key.
	KmsKeyArn interface{}
	// The date this resource was last modified.
	LastModified interface{}
	// A list of Lambda Layer ARNs attached to your Lambda Function.
	Layers interface{}
	// Amount of memory in MB your Lambda Function can use at runtime.
	MemorySize interface{}
	// Qualified (`:QUALIFIER` or `:VERSION` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also `arn`.
	QualifiedArn interface{}
	Qualifier interface{}
	// The amount of reserved concurrent executions for this lambda function or `-1` if unreserved.
	ReservedConcurrentExecutions interface{}
	// IAM role attached to the Lambda Function.
	Role interface{}
	// The runtime environment for the Lambda function..
	Runtime interface{}
	// Base64-encoded representation of raw SHA-256 sum of the zip file.
	SourceCodeHash interface{}
	// The size in bytes of the function .zip file.
	SourceCodeSize interface{}
	Tags interface{}
	// The function execution time at which Lambda should terminate the function.
	Timeout interface{}
	// Tracing settings of the function.
	TracingConfig interface{}
	// The version of the Lambda function.
	Version interface{}
	// VPC configuration associated with your Lambda function.
	VpcConfig interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
