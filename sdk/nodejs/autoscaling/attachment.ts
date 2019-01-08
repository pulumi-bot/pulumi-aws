// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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

    public readonly albTargetGroupArn: pulumi.Output<string | undefined>;
    public readonly autoscalingGroupName: pulumi.Output<string>;
    public readonly elb: pulumi.Output<string | undefined>;

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
            inputs["albTargetGroupArn"] = state ? state.albTargetGroupArn : undefined;
            inputs["autoscalingGroupName"] = state ? state.autoscalingGroupName : undefined;
            inputs["elb"] = state ? state.elb : undefined;
        } else {
            const args = argsOrState as AttachmentArgs | undefined;
            if (!args || args.autoscalingGroupName === undefined) {
                throw new Error("Missing required property 'autoscalingGroupName'");
            }
            inputs["albTargetGroupArn"] = args ? args.albTargetGroupArn : undefined;
            inputs["autoscalingGroupName"] = args ? args.autoscalingGroupName : undefined;
            inputs["elb"] = args ? args.elb : undefined;
        }
        super("aws:autoscaling/attachment:Attachment", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Attachment resources.
 */
export interface AttachmentState {
    readonly albTargetGroupArn?: pulumi.Input<string>;
    readonly autoscalingGroupName?: pulumi.Input<string>;
    readonly elb?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Attachment resource.
 */
export interface AttachmentArgs {
    readonly albTargetGroupArn?: pulumi.Input<string>;
    readonly autoscalingGroupName: pulumi.Input<string>;
    readonly elb?: pulumi.Input<string>;
}
