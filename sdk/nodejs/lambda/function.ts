// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

import {ARN} from "..";

/**
 * Provides a Lambda Function resource. Lambda allows you to trigger execution of code in response to events in AWS, enabling serverless backend solutions. The Lambda Function itself includes source code and runtime configuration.
 *
 * For information about Lambda and how to use it, see [What is AWS Lambda?](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html)
 *
 * > **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), EC2 subnets and security groups associated with Lambda Functions can take up to 45 minutes to successfully delete.
 *
 * ## Example Usage
 * ### Lambda Layers
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleLayerVersion = new aws.lambda.LayerVersion("exampleLayerVersion", {});
 * // ... other configuration ...
 * const exampleFunction = new aws.lambda.Function("exampleFunction", {layers: [exampleLayerVersion.arn]});
 * ```
 * ### Lambda File Systems
 *
 * Lambda File Systems allow you to connect an Amazon Elastic File System (EFS) file system to a Lambda function to share data across function invocations, access existing data including large files, and save function state.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // EFS file system
 * const efsForLambda = new aws.efs.FileSystem("efsForLambda", {tags: {
 *     Name: "efs_for_lambda",
 * }});
 * // Mount target connects the file system to the subnet
 * const alpha = new aws.efs.MountTarget("alpha", {
 *     fileSystemId: efsForLambda.id,
 *     subnetId: aws_subnet.subnet_for_lambda.id,
 *     securityGroups: [aws_security_group.sg_for_lambda.id],
 * });
 * // EFS access point used by lambda file system
 * const accessPointForLambda = new aws.efs.AccessPoint("accessPointForLambda", {
 *     fileSystemId: efsForLambda.id,
 *     rootDirectory: {
 *         path: "/lambda",
 *         creationInfo: {
 *             ownerGid: 1000,
 *             ownerUid: 1000,
 *             permissions: "777",
 *         },
 *     },
 *     posixUser: {
 *         gid: 1000,
 *         uid: 1000,
 *     },
 * });
 * // A lambda function connected to an EFS file system
 * // ... other configuration ...
 * const example = new aws.lambda.Function("example", {
 *     fileSystemConfig: {
 *         arn: accessPointForLambda.arn,
 *         localMountPath: "/mnt/efs",
 *     },
 *     vpcConfig: {
 *         subnetIds: [aws_subnet.subnet_for_lambda.id],
 *         securityGroupIds: [aws_security_group.sg_for_lambda.id],
 *     },
 * }, {
 *     dependsOn: [alpha],
 * });
 * ```
 * ### CloudWatch Logging and Permissions
 *
 * For more information about CloudWatch Logs for Lambda, see the [Lambda User Guide](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-functions-logs.html).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const lambdaFunctionName = config.get("lambdaFunctionName") || "lambda_function_name";
 * // This is to optionally manage the CloudWatch Log Group for the Lambda Function.
 * // If skipping this resource configuration, also add "logs:CreateLogGroup" to the IAM policy below.
 * const example = new aws.cloudwatch.LogGroup("example", {retentionInDays: 14});
 * // See also the following AWS managed policy: AWSLambdaBasicExecutionRole
 * const lambdaLogging = new aws.iam.Policy("lambdaLogging", {
 *     path: "/",
 *     description: "IAM policy for logging from a lambda",
 *     policy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "logs:CreateLogGroup",
 *         "logs:CreateLogStream",
 *         "logs:PutLogEvents"
 *       ],
 *       "Resource": "arn:aws:logs:*:*:*",
 *       "Effect": "Allow"
 *     }
 *   ]
 * }
 * `,
 * });
 * const lambdaLogs = new aws.iam.RolePolicyAttachment("lambdaLogs", {
 *     role: aws_iam_role.iam_for_lambda.name,
 *     policyArn: lambdaLogging.arn,
 * });
 * const testLambda = new aws.lambda.Function("testLambda", {}, {
 *     dependsOn: [
 *         lambdaLogs,
 *         example,
 *     ],
 * });
 * ```
 * ## Specifying the Deployment Package
 *
 * AWS Lambda expects source code to be provided as a deployment package whose structure varies depending on which `runtime` is in use.
 * See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for the valid values of `runtime`. The expected structure of the deployment package can be found in
 * [the AWS Lambda documentation for each runtime](https://docs.aws.amazon.com/lambda/latest/dg/deployment-package-v2.html).
 *
 * Once you have created your deployment package you can specify it either directly as a local file (using the `filename` argument) or
 * indirectly via Amazon S3 (using the `s3Bucket`, `s3Key` and `s3ObjectVersion` arguments). When providing the deployment
 * package via S3 it may be useful to use the `aws.s3.BucketObject` resource to upload it.
 *
 * For larger deployment packages it is recommended by Amazon to upload via S3, since the S3 API has better support for uploading
 * large files efficiently.
 *
 * ## Import
 *
 * Lambda Functions can be imported using the `function_name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:lambda/function:Function test_lambda my_test_lambda_function
 * ```
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lambda/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
     */
    public readonly code!: pulumi.Output<pulumi.asset.Archive | undefined>;
    /**
     * Nested block to configure the function's *dead letter queue*. See details below.
     */
    public readonly deadLetterConfig!: pulumi.Output<outputs.lambda.FunctionDeadLetterConfig | undefined>;
    /**
     * Description of what your Lambda Function does.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Lambda environment's configuration settings. Fields documented below.
     */
    public readonly environment!: pulumi.Output<outputs.lambda.FunctionEnvironment | undefined>;
    /**
     * The connection settings for an EFS file system. Fields documented below. Before creating or updating Lambda functions with `fileSystemConfig`, EFS mount targets much be in available lifecycle state. Use `dependsOn` to explicitly declare this dependency. See [Using Amazon EFS with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/services-efs.html).
     */
    public readonly fileSystemConfig!: pulumi.Output<outputs.lambda.FunctionFileSystemConfig | undefined>;
    /**
     * The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
     */
    public readonly handler!: pulumi.Output<string>;
    /**
     * The ARN to be used for invoking Lambda Function from API Gateway - to be used in `aws.apigateway.Integration`'s `uri`
     */
    public /*out*/ readonly invokeArn!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key that is used to encrypt environment variables. If this configuration is not provided when environment variables are in use, AWS Lambda uses a default service key. If this configuration is provided when environment variables are not in use, the AWS Lambda API does not save this configuration and this provider will show a perpetual difference of adding the key. To fix the perpetual difference, remove this configuration.
     */
    public readonly kmsKeyArn!: pulumi.Output<string | undefined>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
     */
    public readonly layers!: pulumi.Output<string[] | undefined>;
    /**
     * Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
     */
    public readonly memorySize!: pulumi.Output<number | undefined>;
    /**
     * A unique name for your Lambda Function.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
     */
    public readonly publish!: pulumi.Output<boolean | undefined>;
    /**
     * The Amazon Resource Name (ARN) identifying your Lambda Function Version
     * (if versioning is enabled via `publish = true`).
     */
    public /*out*/ readonly qualifiedArn!: pulumi.Output<string>;
    /**
     * The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
     */
    public readonly reservedConcurrentExecutions!: pulumi.Output<number | undefined>;
    /**
     * IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html) for more details.
     */
    public readonly role!: pulumi.Output<ARN>;
    /**
     * See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
     */
    public readonly s3Bucket!: pulumi.Output<string | undefined>;
    /**
     * The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    public readonly s3Key!: pulumi.Output<string | undefined>;
    /**
     * The object version containing the function's deployment package. Conflicts with `filename`.
     */
    public readonly s3ObjectVersion!: pulumi.Output<string | undefined>;
    /**
     * Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
     */
    public readonly sourceCodeHash!: pulumi.Output<string>;
    /**
     * The size in bytes of the function .zip file.
     */
    public /*out*/ readonly sourceCodeSize!: pulumi.Output<number>;
    /**
     * A mapping of tags to assign to the object.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    public readonly tracingConfig!: pulumi.Output<outputs.lambda.FunctionTracingConfig>;
    /**
     * Latest published version of your Lambda Function.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;
    /**
     * Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC](http://docs.aws.amazon.com/lambda/latest/dg/vpc.html)
     */
    public readonly vpcConfig!: pulumi.Output<outputs.lambda.FunctionVpcConfig | undefined>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FunctionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["code"] = state ? state.code : undefined;
            inputs["deadLetterConfig"] = state ? state.deadLetterConfig : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["environment"] = state ? state.environment : undefined;
            inputs["fileSystemConfig"] = state ? state.fileSystemConfig : undefined;
            inputs["handler"] = state ? state.handler : undefined;
            inputs["invokeArn"] = state ? state.invokeArn : undefined;
            inputs["kmsKeyArn"] = state ? state.kmsKeyArn : undefined;
            inputs["lastModified"] = state ? state.lastModified : undefined;
            inputs["layers"] = state ? state.layers : undefined;
            inputs["memorySize"] = state ? state.memorySize : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["publish"] = state ? state.publish : undefined;
            inputs["qualifiedArn"] = state ? state.qualifiedArn : undefined;
            inputs["reservedConcurrentExecutions"] = state ? state.reservedConcurrentExecutions : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["runtime"] = state ? state.runtime : undefined;
            inputs["s3Bucket"] = state ? state.s3Bucket : undefined;
            inputs["s3Key"] = state ? state.s3Key : undefined;
            inputs["s3ObjectVersion"] = state ? state.s3ObjectVersion : undefined;
            inputs["sourceCodeHash"] = state ? state.sourceCodeHash : undefined;
            inputs["sourceCodeSize"] = state ? state.sourceCodeSize : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["tracingConfig"] = state ? state.tracingConfig : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["vpcConfig"] = state ? state.vpcConfig : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if (!args || args.handler === undefined) {
                throw new Error("Missing required property 'handler'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            if (!args || args.runtime === undefined) {
                throw new Error("Missing required property 'runtime'");
            }
            inputs["code"] = args ? args.code : undefined;
            inputs["deadLetterConfig"] = args ? args.deadLetterConfig : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["fileSystemConfig"] = args ? args.fileSystemConfig : undefined;
            inputs["handler"] = args ? args.handler : undefined;
            inputs["kmsKeyArn"] = args ? args.kmsKeyArn : undefined;
            inputs["layers"] = args ? args.layers : undefined;
            inputs["memorySize"] = args ? args.memorySize : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publish"] = args ? args.publish : undefined;
            inputs["reservedConcurrentExecutions"] = args ? args.reservedConcurrentExecutions : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["runtime"] = args ? args.runtime : undefined;
            inputs["s3Bucket"] = args ? args.s3Bucket : undefined;
            inputs["s3Key"] = args ? args.s3Key : undefined;
            inputs["s3ObjectVersion"] = args ? args.s3ObjectVersion : undefined;
            inputs["sourceCodeHash"] = args ? args.sourceCodeHash : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["tracingConfig"] = args ? args.tracingConfig : undefined;
            inputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["invokeArn"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["qualifiedArn"] = undefined /*out*/;
            inputs["sourceCodeSize"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Function.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * The Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
     */
    readonly code?: pulumi.Input<pulumi.asset.Archive>;
    /**
     * Nested block to configure the function's *dead letter queue*. See details below.
     */
    readonly deadLetterConfig?: pulumi.Input<inputs.lambda.FunctionDeadLetterConfig>;
    /**
     * Description of what your Lambda Function does.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Lambda environment's configuration settings. Fields documented below.
     */
    readonly environment?: pulumi.Input<inputs.lambda.FunctionEnvironment>;
    /**
     * The connection settings for an EFS file system. Fields documented below. Before creating or updating Lambda functions with `fileSystemConfig`, EFS mount targets much be in available lifecycle state. Use `dependsOn` to explicitly declare this dependency. See [Using Amazon EFS with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/services-efs.html).
     */
    readonly fileSystemConfig?: pulumi.Input<inputs.lambda.FunctionFileSystemConfig>;
    /**
     * The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
     */
    readonly handler?: pulumi.Input<string>;
    /**
     * The ARN to be used for invoking Lambda Function from API Gateway - to be used in `aws.apigateway.Integration`'s `uri`
     */
    readonly invokeArn?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key that is used to encrypt environment variables. If this configuration is not provided when environment variables are in use, AWS Lambda uses a default service key. If this configuration is provided when environment variables are not in use, the AWS Lambda API does not save this configuration and this provider will show a perpetual difference of adding the key. To fix the perpetual difference, remove this configuration.
     */
    readonly kmsKeyArn?: pulumi.Input<string>;
    /**
     * The date this resource was last modified.
     */
    readonly lastModified?: pulumi.Input<string>;
    /**
     * List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
     */
    readonly layers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
     */
    readonly memorySize?: pulumi.Input<number>;
    /**
     * A unique name for your Lambda Function.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
     */
    readonly publish?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) identifying your Lambda Function Version
     * (if versioning is enabled via `publish = true`).
     */
    readonly qualifiedArn?: pulumi.Input<string>;
    /**
     * The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
     */
    readonly reservedConcurrentExecutions?: pulumi.Input<number>;
    /**
     * IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html) for more details.
     */
    readonly role?: pulumi.Input<ARN>;
    /**
     * See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
     */
    readonly runtime?: pulumi.Input<string>;
    /**
     * The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
     */
    readonly s3Bucket?: pulumi.Input<string>;
    /**
     * The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    readonly s3Key?: pulumi.Input<string>;
    /**
     * The object version containing the function's deployment package. Conflicts with `filename`.
     */
    readonly s3ObjectVersion?: pulumi.Input<string>;
    /**
     * Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
     */
    readonly sourceCodeHash?: pulumi.Input<string>;
    /**
     * The size in bytes of the function .zip file.
     */
    readonly sourceCodeSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
     */
    readonly timeout?: pulumi.Input<number>;
    readonly tracingConfig?: pulumi.Input<inputs.lambda.FunctionTracingConfig>;
    /**
     * Latest published version of your Lambda Function.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC](http://docs.aws.amazon.com/lambda/latest/dg/vpc.html)
     */
    readonly vpcConfig?: pulumi.Input<inputs.lambda.FunctionVpcConfig>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
     */
    readonly code?: pulumi.Input<pulumi.asset.Archive>;
    /**
     * Nested block to configure the function's *dead letter queue*. See details below.
     */
    readonly deadLetterConfig?: pulumi.Input<inputs.lambda.FunctionDeadLetterConfig>;
    /**
     * Description of what your Lambda Function does.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Lambda environment's configuration settings. Fields documented below.
     */
    readonly environment?: pulumi.Input<inputs.lambda.FunctionEnvironment>;
    /**
     * The connection settings for an EFS file system. Fields documented below. Before creating or updating Lambda functions with `fileSystemConfig`, EFS mount targets much be in available lifecycle state. Use `dependsOn` to explicitly declare this dependency. See [Using Amazon EFS with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/services-efs.html).
     */
    readonly fileSystemConfig?: pulumi.Input<inputs.lambda.FunctionFileSystemConfig>;
    /**
     * The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
     */
    readonly handler: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key that is used to encrypt environment variables. If this configuration is not provided when environment variables are in use, AWS Lambda uses a default service key. If this configuration is provided when environment variables are not in use, the AWS Lambda API does not save this configuration and this provider will show a perpetual difference of adding the key. To fix the perpetual difference, remove this configuration.
     */
    readonly kmsKeyArn?: pulumi.Input<string>;
    /**
     * List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
     */
    readonly layers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
     */
    readonly memorySize?: pulumi.Input<number>;
    /**
     * A unique name for your Lambda Function.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
     */
    readonly publish?: pulumi.Input<boolean>;
    /**
     * The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
     */
    readonly reservedConcurrentExecutions?: pulumi.Input<number>;
    /**
     * IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html) for more details.
     */
    readonly role: pulumi.Input<ARN>;
    /**
     * See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
     */
    readonly runtime: pulumi.Input<string>;
    /**
     * The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
     */
    readonly s3Bucket?: pulumi.Input<string>;
    /**
     * The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    readonly s3Key?: pulumi.Input<string>;
    /**
     * The object version containing the function's deployment package. Conflicts with `filename`.
     */
    readonly s3ObjectVersion?: pulumi.Input<string>;
    /**
     * Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
     */
    readonly sourceCodeHash?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
     */
    readonly timeout?: pulumi.Input<number>;
    readonly tracingConfig?: pulumi.Input<inputs.lambda.FunctionTracingConfig>;
    /**
     * Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC](http://docs.aws.amazon.com/lambda/latest/dg/vpc.html)
     */
    readonly vpcConfig?: pulumi.Input<inputs.lambda.FunctionVpcConfig>;
}
