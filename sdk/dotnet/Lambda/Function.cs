// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda
{
    /// <summary>
    /// Provides a Lambda Function resource. Lambda allows you to trigger execution of code in response to events in AWS, enabling serverless backend solutions. The Lambda Function itself includes source code and runtime configuration.
    /// 
    /// For information about Lambda and how to use it, see [What is AWS Lambda?](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html)
    /// 
    /// &gt; **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), EC2 subnets and security groups associated with Lambda Functions can take up to 45 minutes to successfully delete.
    /// 
    /// 
    /// ## Specifying the Deployment Package
    /// 
    /// AWS Lambda expects source code to be provided as a deployment package whose structure varies depending on which `runtime` is in use.
    /// See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for the valid values of `runtime`. The expected structure of the deployment package can be found in
    /// [the AWS Lambda documentation for each runtime](https://docs.aws.amazon.com/lambda/latest/dg/deployment-package-v2.html).
    /// 
    /// Once you have created your deployment package you can specify it either directly as a local file (using the `filename` argument) or
    /// indirectly via Amazon S3 (using the `s3_bucket`, `s3_key` and `s3_object_version` arguments). When providing the deployment
    /// package via S3 it may be useful to use the `aws.s3.BucketObject` resource to upload it.
    /// 
    /// For larger deployment packages it is recommended by Amazon to upload via S3, since the S3 API has better support for uploading
    /// large files efficiently.
    /// </summary>
    public partial class Function : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) identifying your Lambda Function.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        /// </summary>
        [Output("code")]
        public Output<Archive?> Code { get; private set; } = null!;

        /// <summary>
        /// Nested block to configure the function's *dead letter queue*. See details below.
        /// </summary>
        [Output("deadLetterConfig")]
        public Output<Outputs.FunctionDeadLetterConfig?> DeadLetterConfig { get; private set; } = null!;

        /// <summary>
        /// Description of what your Lambda Function does.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Lambda environment's configuration settings. Fields documented below.
        /// </summary>
        [Output("environment")]
        public Output<Outputs.FunctionEnvironment?> Environment { get; private set; } = null!;

        /// <summary>
        /// The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
        /// </summary>
        [Output("handler")]
        public Output<string> Handler { get; private set; } = null!;

        /// <summary>
        /// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`aws.apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
        /// </summary>
        [Output("invokeArn")]
        public Output<string> InvokeArn { get; private set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
        /// </summary>
        [Output("layers")]
        public Output<ImmutableArray<string>> Layers { get; private set; } = null!;

        /// <summary>
        /// Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        /// </summary>
        [Output("memorySize")]
        public Output<int?> MemorySize { get; private set; } = null!;

        /// <summary>
        /// A unique name for your Lambda Function.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
        /// </summary>
        [Output("publish")]
        public Output<bool?> Publish { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) identifying your Lambda Function Version
        /// (if versioning is enabled via `publish = true`).
        /// </summary>
        [Output("qualifiedArn")]
        public Output<string> QualifiedArn { get; private set; } = null!;

        /// <summary>
        /// The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
        /// </summary>
        [Output("reservedConcurrentExecutions")]
        public Output<int?> ReservedConcurrentExecutions { get; private set; } = null!;

        /// <summary>
        /// IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html) for more details.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
        /// </summary>
        [Output("runtime")]
        public Output<string> Runtime { get; private set; } = null!;

        /// <summary>
        /// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        /// </summary>
        [Output("s3Bucket")]
        public Output<string?> S3Bucket { get; private set; } = null!;

        /// <summary>
        /// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Output("s3Key")]
        public Output<string?> S3Key { get; private set; } = null!;

        /// <summary>
        /// The object version containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Output("s3ObjectVersion")]
        public Output<string?> S3ObjectVersion { get; private set; } = null!;

        /// <summary>
        /// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
        /// </summary>
        [Output("sourceCodeHash")]
        public Output<string> SourceCodeHash { get; private set; } = null!;

        /// <summary>
        /// The size in bytes of the function .zip file.
        /// </summary>
        [Output("sourceCodeSize")]
        public Output<int> SourceCodeSize { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the object.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        [Output("tracingConfig")]
        public Output<Outputs.FunctionTracingConfig> TracingConfig { get; private set; } = null!;

        /// <summary>
        /// Latest published version of your Lambda Function.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC](http://docs.aws.amazon.com/lambda/latest/dg/vpc.html)
        /// </summary>
        [Output("vpcConfig")]
        public Output<Outputs.FunctionVpcConfig?> VpcConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Function resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Function(string name, FunctionArgs args, CustomResourceOptions? options = null)
            : base("aws:lambda/function:Function", name, args ?? new FunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Function(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
            : base("aws:lambda/function:Function", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Function resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Function Get(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
        {
            return new Function(name, id, state, options);
        }
    }

    public sealed class FunctionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        /// </summary>
        [Input("code")]
        public Input<Archive>? Code { get; set; }

        /// <summary>
        /// Nested block to configure the function's *dead letter queue*. See details below.
        /// </summary>
        [Input("deadLetterConfig")]
        public Input<Inputs.FunctionDeadLetterConfigArgs>? DeadLetterConfig { get; set; }

        /// <summary>
        /// Description of what your Lambda Function does.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Lambda environment's configuration settings. Fields documented below.
        /// </summary>
        [Input("environment")]
        public Input<Inputs.FunctionEnvironmentArgs>? Environment { get; set; }

        /// <summary>
        /// The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
        /// </summary>
        [Input("handler", required: true)]
        public Input<string> Handler { get; set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        [Input("layers")]
        private InputList<string>? _layers;

        /// <summary>
        /// List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
        /// </summary>
        public InputList<string> Layers
        {
            get => _layers ?? (_layers = new InputList<string>());
            set => _layers = value;
        }

        /// <summary>
        /// Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        /// </summary>
        [Input("memorySize")]
        public Input<int>? MemorySize { get; set; }

        /// <summary>
        /// A unique name for your Lambda Function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
        /// </summary>
        [Input("publish")]
        public Input<bool>? Publish { get; set; }

        /// <summary>
        /// The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
        /// </summary>
        [Input("reservedConcurrentExecutions")]
        public Input<int>? ReservedConcurrentExecutions { get; set; }

        /// <summary>
        /// IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html) for more details.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
        /// </summary>
        [Input("runtime", required: true)]
        public Input<string> Runtime { get; set; } = null!;

        /// <summary>
        /// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Input("s3Key")]
        public Input<string>? S3Key { get; set; }

        /// <summary>
        /// The object version containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Input("s3ObjectVersion")]
        public Input<string>? S3ObjectVersion { get; set; }

        /// <summary>
        /// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
        /// </summary>
        [Input("sourceCodeHash")]
        public Input<string>? SourceCodeHash { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the object.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("tracingConfig")]
        public Input<Inputs.FunctionTracingConfigArgs>? TracingConfig { get; set; }

        /// <summary>
        /// Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC](http://docs.aws.amazon.com/lambda/latest/dg/vpc.html)
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.FunctionVpcConfigArgs>? VpcConfig { get; set; }

        public FunctionArgs()
        {
        }
    }

    public sealed class FunctionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) identifying your Lambda Function.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        /// </summary>
        [Input("code")]
        public Input<Archive>? Code { get; set; }

        /// <summary>
        /// Nested block to configure the function's *dead letter queue*. See details below.
        /// </summary>
        [Input("deadLetterConfig")]
        public Input<Inputs.FunctionDeadLetterConfigGetArgs>? DeadLetterConfig { get; set; }

        /// <summary>
        /// Description of what your Lambda Function does.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Lambda environment's configuration settings. Fields documented below.
        /// </summary>
        [Input("environment")]
        public Input<Inputs.FunctionEnvironmentGetArgs>? Environment { get; set; }

        /// <summary>
        /// The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
        /// </summary>
        [Input("handler")]
        public Input<string>? Handler { get; set; }

        /// <summary>
        /// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`aws.apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
        /// </summary>
        [Input("invokeArn")]
        public Input<string>? InvokeArn { get; set; }

        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Input("lastModified")]
        public Input<string>? LastModified { get; set; }

        [Input("layers")]
        private InputList<string>? _layers;

        /// <summary>
        /// List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
        /// </summary>
        public InputList<string> Layers
        {
            get => _layers ?? (_layers = new InputList<string>());
            set => _layers = value;
        }

        /// <summary>
        /// Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        /// </summary>
        [Input("memorySize")]
        public Input<int>? MemorySize { get; set; }

        /// <summary>
        /// A unique name for your Lambda Function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
        /// </summary>
        [Input("publish")]
        public Input<bool>? Publish { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) identifying your Lambda Function Version
        /// (if versioning is enabled via `publish = true`).
        /// </summary>
        [Input("qualifiedArn")]
        public Input<string>? QualifiedArn { get; set; }

        /// <summary>
        /// The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
        /// </summary>
        [Input("reservedConcurrentExecutions")]
        public Input<int>? ReservedConcurrentExecutions { get; set; }

        /// <summary>
        /// IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html) for more details.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
        /// </summary>
        [Input("runtime")]
        public Input<string>? Runtime { get; set; }

        /// <summary>
        /// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Input("s3Key")]
        public Input<string>? S3Key { get; set; }

        /// <summary>
        /// The object version containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Input("s3ObjectVersion")]
        public Input<string>? S3ObjectVersion { get; set; }

        /// <summary>
        /// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
        /// </summary>
        [Input("sourceCodeHash")]
        public Input<string>? SourceCodeHash { get; set; }

        /// <summary>
        /// The size in bytes of the function .zip file.
        /// </summary>
        [Input("sourceCodeSize")]
        public Input<int>? SourceCodeSize { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the object.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("tracingConfig")]
        public Input<Inputs.FunctionTracingConfigGetArgs>? TracingConfig { get; set; }

        /// <summary>
        /// Latest published version of your Lambda Function.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC](http://docs.aws.amazon.com/lambda/latest/dg/vpc.html)
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.FunctionVpcConfigGetArgs>? VpcConfig { get; set; }

        public FunctionState()
        {
        }
    }
}
