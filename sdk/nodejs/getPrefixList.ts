// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * `aws..getPrefixList` provides details about a specific prefix list (PL)
 * in the current region.
 * 
 * This can be used both to validate a prefix list given in a variable
 * and to obtain the CIDR blocks (IP address ranges) for the associated
 * AWS service. The latter may be useful e.g. for adding network ACL
 * rules.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const privateS3VpcEndpoint = new aws.ec2.VpcEndpoint("privateS3", {
 *     serviceName: "com.amazonaws.us-west-2.s3",
 *     vpcId: aws_vpc_foo.id,
 * });
 * const privateS3PrefixList = privateS3VpcEndpoint.prefixListId.apply(prefixListId => aws.getPrefixList({
 *     prefixListId: prefixListId,
 * }));
 * const bar = new aws.ec2.NetworkAcl("bar", {
 *     vpcId: aws_vpc_foo.id,
 * });
 * const privateS3NetworkAclRule = new aws.ec2.NetworkAclRule("privateS3", {
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
 * 
 * ### Filter
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const test = aws.getPrefixList({
 *     filters: [{
 *         name: "prefix-list-id",
 *         values: ["pl-68a54001"],
 *     }],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/prefix_list.html.markdown.
 */
export function getPrefixList(args?: GetPrefixListArgs, opts?: pulumi.InvokeOptions): Promise<GetPrefixListResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:index/getPrefixList:getPrefixList", {
        "filters": args.filters,
        "name": args.name,
        "prefixListId": args.prefixListId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrefixList.
 */
export interface GetPrefixListArgs {
    /**
     * Configuration block(s) for filtering. Detailed below.
     */
    readonly filters?: inputs.GetPrefixListFilter[];
    /**
     * The name of the filter field. Valid values can be found in the [EC2 DescribePrefixLists API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribePrefixLists.html).
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
     * The list of CIDR blocks for the AWS service associated with the prefix list.
     */
    readonly cidrBlocks: string[];
    readonly filters?: outputs.GetPrefixListFilter[];
    /**
     * The name of the selected prefix list.
     */
    readonly name: string;
    readonly prefixListId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
