// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the id of a VPC Link in
 * API Gateway. To fetch the VPC Link you must provide a name to match against. 
 * As there is no unique name constraint on API Gateway VPC Links this data source will 
 * error if there is more than one match.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const myApiGatewayVpcLink = aws.apigateway.getVpcLink({
 *     name: "my-vpc-link",
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/api_gateway_vpc_link.html.markdown.
 */
export function getVpcLink(args: GetVpcLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcLinkResult> & GetVpcLinkResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetVpcLinkResult> = pulumi.runtime.invoke("aws:apigateway/getVpcLink:getVpcLink", {
        "name": args.name,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getVpcLink.
 */
export interface GetVpcLinkArgs {
    /**
     * The name of the API Gateway VPC Link to look up. If no API Gateway VPC Link is found with this name, an error will be returned. 
     * If multiple API Gateway VPC Links are found with this name, an error will be returned.
     */
    readonly name: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getVpcLink.
 */
export interface GetVpcLinkResult {
    /**
     * The description of the VPC link.
     */
    readonly description: string;
    /**
     * Set to the ID of the found API Gateway VPC Link.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The status of the VPC link.
     */
    readonly status: string;
    /**
     * The status message of the VPC link.
     */
    readonly statusMessage: string;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags: {[key: string]: any};
    /**
     * The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
     */
    readonly targetArns: string[];
}
