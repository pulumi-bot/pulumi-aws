// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an existing autoscaling group.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const foo = aws.autoscaling.getGroup({
 *     name: "foo",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/autoscaling_group.html.markdown.
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> & GetGroupResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetGroupResult> = pulumi.runtime.invoke("aws:autoscaling/getGroup:getGroup", {
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * Specify the exact name of the desired autoscaling group.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * The Amazon Resource Name (ARN) of the Auto Scaling group.
     */
    readonly arn: string;
    /**
     * One or more Availability Zones for the group.
     */
    readonly availabilityZones: string[];
    readonly defaultCooldown: number;
    /**
     * The desired size of the group.
     */
    readonly desiredCapacity: number;
    /**
     * The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
     */
    readonly healthCheckGracePeriod: number;
    /**
     * The service to use for the health checks. The valid values are EC2 and ELB.
     */
    readonly healthCheckType: string;
    /**
     * The name of the associated launch configuration.
     */
    readonly launchConfiguration: string;
    /**
     * One or more load balancers associated with the group.
     */
    readonly loadBalancers: string[];
    /**
     * The maximum size of the group.
     */
    readonly maxSize: number;
    /**
     * The minimum size of the group.
     */
    readonly minSize: number;
    /**
     * The name of the Auto Scaling group.
     */
    readonly name: string;
    readonly newInstancesProtectedFromScaleIn: boolean;
    /**
     * The name of the placement group into which to launch your instances, if any. For more information, see Placement Groups (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the Amazon Elastic Compute Cloud User Guide.
     */
    readonly placementGroup: string;
    /**
     * The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.
     */
    readonly serviceLinkedRoleArn: string;
    /**
     * The current state of the group when DeleteAutoScalingGroup is in progress.
     */
    readonly status: string;
    /**
     * The Amazon Resource Names (ARN) of the target groups for your load balancer.
     */
    readonly targetGroupArns: string[];
    /**
     * The termination policies for the group.
     */
    readonly terminationPolicies: string[];
    /**
     * VPC ID for the group.
     */
    readonly vpcZoneIdentifier: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
