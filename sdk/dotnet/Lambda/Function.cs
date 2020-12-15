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
    /// ## Import
    /// 
    /// Lambda Functions can be imported using the `function_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:lambda/function:Function test_lambda my_test_lambda_function
    /// ```
    /// </summary>
    [AwsResourceType("aws:lambda/function:Function")]
    public partial class Function : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options and `image_uri` cannot be used.
        /// </summary>
        [Output("code")]
        public Output<Archive?> Code { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) for a Code Signing Configuration.
        /// </summary>
        [Output("codeSigningConfigArn")]
        public Output<string?> CodeSigningConfigArn { get; private set; } = null!;

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
        /// The connection settings for an EFS file system. Fields documented below. Before creating or updating Lambda functions with `file_system_config`, EFS mount targets much be in available lifecycle state. Use `depends_on` to explicitly declare this dependency. See [Using Amazon EFS with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/services-efs.html).
        /// </summary>
        [Output("fileSystemConfig")]
        public Output<Outputs.FunctionFileSystemConfig?> FileSystemConfig { get; private set; } = null!;

        /// <summary>
        /// The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
        /// </summary>
        [Output("handler")]
        public Output<string?> Handler { get; private set; } = null!;

        /// <summary>
        /// The Lambda OCI image configurations. Fields documented below. See [Using container images with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html)
        /// </summary>
        [Output("imageConfig")]
        public Output<Outputs.FunctionImageConfig?> ImageConfig { get; private set; } = null!;

        /// <summary>
        /// The ECR image URI containing the function's deployment package. Conflicts with `filename`, `s3_bucket`, `s3_key`, and `s3_object_version`.
        /// </summary>
        [Output("imageUri")]
        public Output<string?> ImageUri { get; private set; } = null!;

        /// <summary>
        /// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`aws.apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
        /// </summary>
        [Output("invokeArn")]
        public Output<string> InvokeArn { get; private set; } = null!;

        /// <summary>
        /// (Optional) The ARN for the KMS encryption key.
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
        /// The Lambda deployment package type. Valid values are `Zip` and `Image`. Defaults to `Zip`.
        /// </summary>
        [Output("packageType")]
        public Output<string?> PackageType { get; private set; } = null!;

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
        public Output<string?> Runtime { get; private set; } = null!;

        /// <summary>
        /// The S3 bucket location containing the function's deployment package. Conflicts with `filename` and `image_uri`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        /// </summary>
        [Output("s3Bucket")]
        public Output<string?> S3Bucket { get; private set; } = null!;

        /// <summary>
        /// The S3 key of an object containing the function's deployment package. Conflicts with `filename` and `image_uri`.
        /// </summary>
        [Output("s3Key")]
        public Output<string?> S3Key { get; private set; } = null!;

        /// <summary>
        /// The object version containing the function's deployment package. Conflicts with `filename` and `image_uri`.
        /// </summary>
        [Output("s3ObjectVersion")]
        public Output<string?> S3ObjectVersion { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of a signing job.
        /// </summary>
        [Output("signingJobArn")]
        public Output<string> SigningJobArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for a signing profile version.
        /// </summary>
        [Output("signingProfileVersionArn")]
        public Output<string> SigningProfileVersionArn { get; private set; } = null!;

        /// <summary>
        /// Base64-encoded representation of raw SHA-256 sum of the zip file, provided either via `filename` or `s3_*` parameters.
        /// </summary>
        [Output("sourceCodeHash")]
        public Output<string> SourceCodeHash { get; private set; } = null!;

        /// <summary>
        /// The size in bytes of the function .zip file.
        /// </summary>
        [Output("sourceCodeSize")]
        public Output<int> SourceCodeSize { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the object.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

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
        /// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options and `image_uri` cannot be used.
        /// </summary>
        [Input("code")]
        public Input<Archive>? Code { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) for a Code Signing Configuration.
        /// </summary>
        [Input("codeSigningConfigArn")]
        public Input<string>? CodeSigningConfigArn { get; set; }

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
        /// The connection settings for an EFS file system. Fields documented below. Before creating or updating Lambda functions with `file_system_config`, EFS mount targets much be in available lifecycle state. Use `depends_on` to explicitly declare this dependency. See [Using Amazon EFS with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/services-efs.html).
        /// </summary>
        [Input("fileSystemConfig")]
        public Input<Inputs.FunctionFileSystemConfigArgs>? FileSystemConfig { get; set; }

        /// <summary>
        /// The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
        /// </summary>
        [Input("handler")]
        public Input<string>? Handler { get; set; }

        /// <summary>
        /// The Lambda OCI image configurations. Fields documented below. See [Using container images with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html)
        /// </summary>
        [Input("imageConfig")]
        public Input<Inputs.FunctionImageConfigArgs>? ImageConfig { get; set; }

        /// <summary>
        /// The ECR image URI containing the function's deployment package. Conflicts with `filename`, `s3_bucket`, `s3_key`, and `s3_object_version`.
        /// </summary>
        [Input("imageUri")]
        public Input<string>? ImageUri { get; set; }

        /// <summary>
        /// (Optional) The ARN for the KMS encryption key.
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
        /// The Lambda deployment package type. Valid values are `Zip` and `Image`. Defaults to `Zip`.
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

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
        [Input("runtime")]
        public InputUnion<string, Pulumi.Aws.Lambda.Runtime>? Runtime { get; set; }

        /// <summary>
        /// The S3 bucket location containing the function's deployment package. Conflicts with `filename` and `image_uri`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// The S3 key of an object containing the function's deployment package. Conflicts with `filename` and `image_uri`.
        /// </summary>
        [Input("s3Key")]
        public Input<string>? S3Key { get; set; }

        /// <summary>
        /// The object version containing the function's deployment package. Conflicts with `filename` and `image_uri`.
        /// </summary>
        [Input("s3ObjectVersion")]
        public Input<string>? S3ObjectVersion { get; set; }

        /// <summary>
        /// Base64-encoded representation of raw SHA-256 sum of the zip file, provided either via `filename` or `s3_*` parameters.
        /// </summary>
        [Input("sourceCodeHash")]
        public Input<string>? SourceCodeHash { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the object.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
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
        /// The Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options and `image_uri` cannot be used.
        /// </summary>
        [Input("code")]
        public Input<Archive>? Code { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) for a Code Signing Configuration.
        /// </summary>
        [Input("codeSigningConfigArn")]
        public Input<string>? CodeSigningConfigArn { get; set; }

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
        /// The connection settings for an EFS file system. Fields documented below. Before creating or updating Lambda functions with `file_system_config`, EFS mount targets much be in available lifecycle state. Use `depends_on` to explicitly declare this dependency. See [Using Amazon EFS with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/services-efs.html).
        /// </summary>
        [Input("fileSystemConfig")]
        public Input<Inputs.FunctionFileSystemConfigGetArgs>? FileSystemConfig { get; set; }

        /// <summary>
        /// The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
        /// </summary>
        [Input("handler")]
        public Input<string>? Handler { get; set; }

        /// <summary>
        /// The Lambda OCI image configurations. Fields documented below. See [Using container images with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html)
        /// </summary>
        [Input("imageConfig")]
        public Input<Inputs.FunctionImageConfigGetArgs>? ImageConfig { get; set; }

        /// <summary>
        /// The ECR image URI containing the function's deployment package. Conflicts with `filename`, `s3_bucket`, `s3_key`, and `s3_object_version`.
        /// </summary>
        [Input("imageUri")]
        public Input<string>? ImageUri { get; set; }

        /// <summary>
        /// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`aws.apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
        /// </summary>
        [Input("invokeArn")]
        public Input<string>? InvokeArn { get; set; }

        /// <summary>
        /// (Optional) The ARN for the KMS encryption key.
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
        /// The Lambda deployment package type. Valid values are `Zip` and `Image`. Defaults to `Zip`.
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

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
        public InputUnion<string, Pulumi.Aws.Lambda.Runtime>? Runtime { get; set; }

        /// <summary>
        /// The S3 bucket location containing the function's deployment package. Conflicts with `filename` and `image_uri`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// The S3 key of an object containing the function's deployment package. Conflicts with `filename` and `image_uri`.
        /// </summary>
        [Input("s3Key")]
        public Input<string>? S3Key { get; set; }

        /// <summary>
        /// The object version containing the function's deployment package. Conflicts with `filename` and `image_uri`.
        /// </summary>
        [Input("s3ObjectVersion")]
        public Input<string>? S3ObjectVersion { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of a signing job.
        /// </summary>
        [Input("signingJobArn")]
        public Input<string>? SigningJobArn { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for a signing profile version.
        /// </summary>
        [Input("signingProfileVersionArn")]
        public Input<string>? SigningProfileVersionArn { get; set; }

        /// <summary>
        /// Base64-encoded representation of raw SHA-256 sum of the zip file, provided either via `filename` or `s3_*` parameters.
        /// </summary>
        [Input("sourceCodeHash")]
        public Input<string>? SourceCodeHash { get; set; }

        /// <summary>
        /// The size in bytes of the function .zip file.
        /// </summary>
        [Input("sourceCodeSize")]
        public Input<int>? SourceCodeSize { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the object.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
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
