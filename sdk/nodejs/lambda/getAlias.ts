// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information about a Lambda Alias.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const production = pulumi.output(aws.lambda.getAlias({
 *     functionName: "my-lambda-func",
 *     name: "production",
 * }));
 * ```
 */
export function getAlias(args: GetAliasArgs, opts?: pulumi.InvokeOptions): Promise<GetAliasResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:lambda/getAlias:getAlias", {
        "functionName": args.functionName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getAlias.
 */
export interface GetAliasArgs {
    /**
     * Name of the aliased Lambda function.
     */
    readonly functionName: string;
    /**
     * Name of the Lambda alias.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getAlias.
 */
export interface GetAliasResult {
    /**
     * The Amazon Resource Name (ARN) identifying the Lambda function alias.
     */
    readonly arn: string;
    /**
     * Description of alias.
     */
    readonly description: string;
    readonly functionName: string;
    /**
     * Lambda function version which the alias uses.
     */
    readonly functionVersion: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ARN to be used for invoking Lambda Function from API Gateway - to be used in aws_api_gateway_integration's `uri`.
     */
    readonly invokeArn: string;
    readonly name: string;
}
