// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches an EC2 instance to an Elastic Load Balancer (ELB). For attaching resources with Application Load Balancer (ALB) or Network Load Balancer (NLB), see the `aws.lb.TargetGroupAttachment` resource.
 *
 * > **NOTE on ELB Instances and ELB Attachments:** This provider currently provides
 * both a standalone ELB Attachment resource (describing an instance attached to
 * an ELB), and an Elastic Load Balancer resource with
 * `instances` defined in-line. At this time you cannot use an ELB with in-line
 * instances in conjunction with an ELB Attachment resource. Doing so will cause a
 * conflict and will overwrite attachments.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // Create a new load balancer attachment
 * const baz = new aws.elb.Attachment("baz", {
 *     elb: aws_elb_bar.id,
 *     instance: aws_instance_foo.id,
 * });
 * ```
 *
 * @deprecated aws.elasticloadbalancing.Attachment has been deprecated in favor of aws.elb.Attachment
 */
export class Attachment extends pulumi.CustomResource {
    /**
     * Get an existing Attachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttachmentState, opts?: pulumi.CustomResourceOptions): Attachment {
        pulumi.log.warn("Attachment is deprecated: aws.elasticloadbalancing.Attachment has been deprecated in favor of aws.elb.Attachment")
        return new Attachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticloadbalancing/attachment:Attachment';

    /**
     * Returns true if the given object is an instance of Attachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Attachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Attachment.__pulumiType;
    }

    /**
     * The name of the ELB.
     */
    public readonly elb!: pulumi.Output<string>;
    /**
     * Instance ID to place in the ELB pool.
     */
    public readonly instance!: pulumi.Output<string>;

    /**
     * Create a Attachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated aws.elasticloadbalancing.Attachment has been deprecated in favor of aws.elb.Attachment */
    constructor(name: string, args: AttachmentArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated aws.elasticloadbalancing.Attachment has been deprecated in favor of aws.elb.Attachment */
    constructor(name: string, argsOrState?: AttachmentArgs | AttachmentState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Attachment is deprecated: aws.elasticloadbalancing.Attachment has been deprecated in favor of aws.elb.Attachment")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AttachmentState | undefined;
            inputs["elb"] = state ? state.elb : undefined;
            inputs["instance"] = state ? state.instance : undefined;
        } else {
            const args = argsOrState as AttachmentArgs | undefined;
            if (!args || args.elb === undefined) {
                throw new Error("Missing required property 'elb'");
            }
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            inputs["elb"] = args ? args.elb : undefined;
            inputs["instance"] = args ? args.instance : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Attachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Attachment resources.
 */
export interface AttachmentState {
    /**
     * The name of the ELB.
     */
    readonly elb?: pulumi.Input<string>;
    /**
     * Instance ID to place in the ELB pool.
     */
    readonly instance?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Attachment resource.
 */
export interface AttachmentArgs {
    /**
     * The name of the ELB.
     */
    readonly elb: pulumi.Input<string>;
    /**
     * Instance ID to place in the ELB pool.
     */
    readonly instance: pulumi.Input<string>;
}
