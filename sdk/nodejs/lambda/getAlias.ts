// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides information about a Lambda Alias.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lambda_alias.html.markdown.
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
     * The ARN to be used for invoking Lambda Function from API Gateway - to be used in aws_api_gateway_integration's `uri`.
     */
    readonly invokeArn: string;
    readonly name: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
