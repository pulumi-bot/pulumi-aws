// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * > **Note:** `aws.alb.TargetGroup` is known as `aws.lb.TargetGroup`. The functionality is identical.
 *
 * Provides information about a Load Balancer Target Group.
 *
 * This data source can prove useful when a module accepts an LB Target Group as an
 * input variable and needs to know its attributes. It can also be used to get the ARN of
 * an LB Target Group for use in other resources, given LB Target Group name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const lbTgArn = config.get("lbTgArn") || "";
 * const lbTgName = config.get("lbTgName") || "";
 * const test = aws.lb.getTargetGroup({
 *     arn: lbTgArn,
 *     name: lbTgName,
 * });
 * ```
 */
/** @deprecated aws.elasticloadbalancingv2.getTargetGroup has been deprecated in favor of aws.lb.getTargetGroup */
export function getTargetGroup(args?: GetTargetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetTargetGroupResult> {
    pulumi.log.warn("getTargetGroup is deprecated: aws.elasticloadbalancingv2.getTargetGroup has been deprecated in favor of aws.lb.getTargetGroup")
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:elasticloadbalancingv2/getTargetGroup:getTargetGroup", {
        "arn": args.arn,
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getTargetGroup.
 */
export interface GetTargetGroupArgs {
    /**
     * The full ARN of the target group.
     */
    arn?: string;
    /**
     * The unique name of the target group.
     */
    name?: string;
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getTargetGroup.
 */
export interface GetTargetGroupResult {
    readonly arn: string;
    readonly arnSuffix: string;
    readonly deregistrationDelay: number;
    readonly healthCheck: outputs.elasticloadbalancingv2.GetTargetGroupHealthCheck;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lambdaMultiValueHeadersEnabled: boolean;
    readonly loadBalancingAlgorithmType: string;
    readonly name: string;
    readonly port: number;
    readonly preserveClientIp: string;
    readonly protocol: string;
    readonly protocolVersion: string;
    readonly proxyProtocolV2: boolean;
    readonly slowStart: number;
    readonly stickiness: outputs.elasticloadbalancingv2.GetTargetGroupStickiness;
    readonly tags: {[key: string]: string};
    readonly targetType: string;
    readonly vpcId: string;
}
