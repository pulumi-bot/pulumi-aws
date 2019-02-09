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
 */
export function getFunction(args: GetFunctionArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionResult> {
    return pulumi.runtime.invoke("aws:lambda/getFunction:getFunction", {
        "functionName": args.functionName,
        "qualifier": args.qualifier,
    }, opts);
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
     * Qualifier of the lambda function. Defaults to `$LATEST`.
     */
    readonly qualifier?: string;
}

/**
 * A collection of values returned by getFunction.
 */
export interface GetFunctionResult {
    /**
     * The Amazon Resource Name (ARN) identifying your Lambda Function.
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
     * The Amazon Resource Name (ARN) identifying your Lambda Function Version
     */
    readonly qualifiedArn: string;
    /**
     * The amount of reserved concurrent executions for this lambda function.
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
