// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides the ability to register instances and containers with an Application Load Balancer (ALB) or Network Load Balancer (NLB) target group. For attaching resources with Elastic Load Balancer (ELB), see the `aws.elb.Attachment` resource.
 *
 * > **Note:** `aws.alb.TargetGroupAttachment` is known as `aws.lb.TargetGroupAttachment`. The functionality is identical.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testTargetGroup = new aws.lb.TargetGroup("testTargetGroup", {});
 * // ... other configuration ...
 * const testInstance = new aws.ec2.Instance("testInstance", {});
 * // ... other configuration ...
 * const testTargetGroupAttachment = new aws.lb.TargetGroupAttachment("testTargetGroupAttachment", {
 *     targetGroupArn: testTargetGroup.arn,
 *     targetId: testInstance.id,
 *     port: 80,
 * });
 * ```
 * ## Usage with lambda
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testTargetGroup = new aws.lb.TargetGroup("testTargetGroup", {targetType: "lambda"});
 * const testFunction = new aws.lambda.Function("testFunction", {});
 * // ... other configuration ...
 * const withLb = new aws.lambda.Permission("withLb", {
 *     action: "lambda:InvokeFunction",
 *     "function": testFunction.arn,
 *     principal: "elasticloadbalancing.amazonaws.com",
 *     sourceArn: testTargetGroup.arn,
 * });
 * const testTargetGroupAttachment = new aws.lb.TargetGroupAttachment("testTargetGroupAttachment", {
 *     targetGroupArn: testTargetGroup.arn,
 *     targetId: testFunction.arn,
 * }, {
 *     dependsOn: [withLb],
 * });
 * ```
 *
 * ## Import
 *
 * Target Group Attachments cannot be imported.
 */
export class TargetGroupAttachment extends pulumi.CustomResource {
    /**
     * Get an existing TargetGroupAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetGroupAttachmentState, opts?: pulumi.CustomResourceOptions): TargetGroupAttachment {
        return new TargetGroupAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lb/targetGroupAttachment:TargetGroupAttachment';

    /**
     * Returns true if the given object is an instance of TargetGroupAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetGroupAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetGroupAttachment.__pulumiType;
    }

    /**
     * The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
     */
    public readonly availabilityZone!: pulumi.Output<string | undefined>;
    /**
     * The port on which targets receive traffic.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The ARN of the target group with which to register targets
     */
    public readonly targetGroupArn!: pulumi.Output<string>;
    /**
     * The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
     */
    public readonly targetId!: pulumi.Output<string>;

    /**
     * Create a TargetGroupAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetGroupAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetGroupAttachmentArgs | TargetGroupAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TargetGroupAttachmentState | undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["targetGroupArn"] = state ? state.targetGroupArn : undefined;
            inputs["targetId"] = state ? state.targetId : undefined;
        } else {
            const args = argsOrState as TargetGroupAttachmentArgs | undefined;
            if ((!args || args.targetGroupArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetGroupArn'");
            }
            if ((!args || args.targetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetId'");
            }
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["targetGroupArn"] = args ? args.targetGroupArn : undefined;
            inputs["targetId"] = args ? args.targetId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "aws:elasticloadbalancingv2/targetGroupAttachment:TargetGroupAttachment" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(TargetGroupAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetGroupAttachment resources.
 */
export interface TargetGroupAttachmentState {
    /**
     * The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * The port on which targets receive traffic.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The ARN of the target group with which to register targets
     */
    readonly targetGroupArn?: pulumi.Input<string>;
    /**
     * The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
     */
    readonly targetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetGroupAttachment resource.
 */
export interface TargetGroupAttachmentArgs {
    /**
     * The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * The port on which targets receive traffic.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The ARN of the target group with which to register targets
     */
    readonly targetGroupArn: pulumi.Input<string>;
    /**
     * The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
     */
    readonly targetId: pulumi.Input<string>;
}
