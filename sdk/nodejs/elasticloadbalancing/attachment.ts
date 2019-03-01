// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches an EC2 instance to an Elastic Load Balancer (ELB). For attaching resources with Application Load Balancer (ALB) or Network Load Balancer (NLB), see the [`aws_lb_target_group_attachment` resource](https://www.terraform.io/docs/providers/aws/r/lb_target_group_attachment.html).
 * 
 * > **NOTE on ELB Instances and ELB Attachments:** Terraform currently provides
 * both a standalone ELB Attachment resource (describing an instance attached to
 * an ELB), and an Elastic Load Balancer resource with
 * `instances` defined in-line. At this time you cannot use an ELB with in-line
 * instances in conjunction with an ELB Attachment resource. Doing so will cause a
 * conflict and will overwrite attachments.
 */
export class Attachment extends pulumi.CustomResource {
    /**
     * Get an existing Attachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttachmentState, opts?: pulumi.CustomResourceOptions): Attachment {
        return new Attachment(name, <any>state, { ...opts, id: id });
    }

    /**
     * The name of the ELB.
     */
    public readonly elb: pulumi.Output<string>;
    /**
     * Instance ID to place in the ELB pool.
     */
    public readonly instance: pulumi.Output<string>;

    /**
     * Create a Attachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttachmentArgs | AttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AttachmentState = argsOrState as AttachmentState | undefined;
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
        super("aws:elasticloadbalancing/attachment:Attachment", name, inputs, opts);
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
