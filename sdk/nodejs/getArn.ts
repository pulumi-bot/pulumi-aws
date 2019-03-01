// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Parses an Amazon Resource Name (ARN) into its constituent parts.
 */
export function getArn(args: GetArnArgs, opts?: pulumi.InvokeOptions): Promise<GetArnResult> {
    return pulumi.runtime.invoke("aws:index/getArn:getArn", {
        "arn": args.arn,
    }, opts);
}

/**
 * A collection of arguments for invoking getArn.
 */
export interface GetArnArgs {
    /**
     * The ARN to parse.
     */
    readonly arn: string;
}

/**
 * A collection of values returned by getArn.
 */
export interface GetArnResult {
    /**
     * The [ID](https://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html) of the AWS account that owns the resource, without the hyphens.
     */
    readonly account: string;
    /**
     * The partition that the resource is in.
     */
    readonly partition: string;
    /**
     * The region the resource resides in.
     * Note that the ARNs for some resources do not require a region, so this component might be omitted.
     */
    readonly region: string;
    /**
     * The content of this part of the ARN varies by service.
     * It often includes an indicator of the type of resource—for example, an IAM user or Amazon RDS database —followed by a slash (/) or a colon (:), followed by the resource name itself.
     */
    readonly resource: string;
    /**
     * The [service namespace](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces) that identifies the AWS product.
     */
    readonly service: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
