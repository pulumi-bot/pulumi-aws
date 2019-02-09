// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `aws_prefix_list` provides details about a specific prefix list (PL)
 * in the current region.
 * 
 * This can be used both to validate a prefix list given in a variable
 * and to obtain the CIDR blocks (IP address ranges) for the associated
 * AWS service. The latter may be useful e.g. for adding network ACL
 * rules.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const bar = new aws.ec2.NetworkAcl("bar", {
 *     vpcId: aws_vpc_foo.id,
 * });
 * const privateS3VpcEndpoint = new aws.ec2.VpcEndpoint("private_s3", {
 *     serviceName: "com.amazonaws.us-west-2.s3",
 *     vpcId: aws_vpc_foo.id,
 * });
 * const privateS3PrefixList = pulumi.output(aws.getPrefixList({
 *     prefixListId: privateS3VpcEndpoint.prefixListId,
 * }));
 * const privateS3NetworkAclRule = new aws.ec2.NetworkAclRule("private_s3", {
 *     cidrBlock: privateS3PrefixList.apply(privateS3PrefixList => privateS3PrefixList.cidrBlocks[0]),
 *     egress: false,
 *     fromPort: 443,
 *     networkAclId: bar.id,
 *     protocol: "tcp",
 *     ruleAction: "allow",
 *     ruleNumber: 200,
 *     toPort: 443,
 * });
 * ```
 */
export function getPrefixList(args?: GetPrefixListArgs, opts?: pulumi.InvokeOptions): Promise<GetPrefixListResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:index/getPrefixList:getPrefixList", {
        "name": args.name,
        "prefixListId": args.prefixListId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrefixList.
 */
export interface GetPrefixListArgs {
    /**
     * The name of the prefix list to select.
     */
    readonly name?: string;
    /**
     * The ID of the prefix list to select.
     */
    readonly prefixListId?: string;
}

/**
 * A collection of values returned by getPrefixList.
 */
export interface GetPrefixListResult {
    /**
     * The list of CIDR blocks for the AWS service associated
     * with the prefix list.
     */
    readonly cidrBlocks: string[];
    /**
     * The name of the selected prefix list.
     */
    readonly name: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
