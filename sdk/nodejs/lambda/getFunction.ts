// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides information about a Lambda Function.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const functionName = config.require("functionName");
 * 
 * const existing = pulumi.output(aws.lambda.getFunction({
 *     functionName: functionName,
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lambda_function.html.markdown.
 */
export function getFunction(args: GetFunctionArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionResult> & GetFunctionResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetFunctionResult> = pulumi.runtime.invoke("aws:lambda/getFunction:getFunction", {
        "functionName": args.functionName,
        "qualifier": args.qualifier,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getFunction.
 */
export interface GetFunctionArgs {
    /**
     * Name of the lambda function.
     */
    readonly functionName: string;
    /**
     * Alias name or version number of the lambda function. e.g. `$LATEST`, `my-alias`, or `1`
     */
    readonly qualifier?: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getFunction.
 */
export interface GetFunctionResult {
    /**
     * Unqualified (no `:QUALIFIER` or `:VERSION` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also `qualified_arn`.
     */
    readonly arn: string;
    /**
     * Configure the function's *dead letter queue*.
     */
    readonly deadLetterConfig: { targetArn: string };
    /**
     * Description of what your Lambda Function does.
     */
    readonly description: string;
    /**
     * The Lambda environment's configuration settings.
     */
    readonly environment: { variables: {[key: string]: string} };
    readonly functionName: string;
    /**
     * The function entrypoint in your code.
     */
    readonly handler: string;
    /**
     * The ARN to be used for invoking Lambda Function from API Gateway.
     */
    readonly invokeArn: string;
    /**
     * The ARN for the KMS encryption key.
     */
    readonly kmsKeyArn: string;
    /**
     * The date this resource was last modified.
     */
    readonly lastModified: string;
    /**
     * A list of Lambda Layer ARNs attached to your Lambda Function.
     */
    readonly layers: string[];
    /**
     * Amount of memory in MB your Lambda Function can use at runtime.
     */
    readonly memorySize: number;
    /**
     * Qualified (`:QUALIFIER` or `:VERSION` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also `arn`.
     */
    readonly qualifiedArn: string;
    readonly qualifier?: string;
    /**
     * The amount of reserved concurrent executions for this lambda function or `-1` if unreserved.
     */
    readonly reservedConcurrentExecutions: number;
    /**
     * IAM role attached to the Lambda Function.
     */
    readonly role: string;
    /**
     * The runtime environment for the Lambda function..
     */
    readonly runtime: string;
    /**
     * Base64-encoded representation of raw SHA-256 sum of the zip file.
     */
    readonly sourceCodeHash: string;
    /**
     * The size in bytes of the function .zip file.
     */
    readonly sourceCodeSize: number;
    readonly tags: {[key: string]: any};
    /**
     * The function execution time at which Lambda should terminate the function.
     */
    readonly timeout: number;
    /**
     * Tracing settings of the function.
     */
    readonly tracingConfig: { mode: string };
    /**
     * The version of the Lambda function.
     */
    readonly version: string;
    /**
     * VPC configuration associated with your Lambda function.
     */
    readonly vpcConfig: { securityGroupIds: string[], subnetIds: string[], vpcId: string };
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
