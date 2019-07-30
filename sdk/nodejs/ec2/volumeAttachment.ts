// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/volume_attachment.html.markdown.
 */
export class VolumeAttachment extends pulumi.CustomResource {
    /**
     * Get an existing VolumeAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeAttachmentState, opts?: pulumi.CustomResourceOptions): VolumeAttachment {
        return new VolumeAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/volumeAttachment:VolumeAttachment';

    /**
     * Returns true if the given object is an instance of VolumeAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VolumeAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VolumeAttachment.__pulumiType;
    }

    /**
     * The device name to expose to the instance (for
     * example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances][1] and [Device Naming on Windows Instances][2] for more information.
     */
    public readonly deviceName!: pulumi.Output<string>;
    /**
     * Set to `true` if you want to force the
     * volume to detach. Useful if previous attempts failed, but use this option only
     * as a last resort, as this can result in **data loss**. See
     * [Detaching an Amazon EBS Volume from an Instance][3] for more information.
     */
    public readonly forceDetach!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the Instance to attach to
     */
    public readonly instanceId!: pulumi.Output<string>;
    public readonly skipDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the Volume to be attached
     */
    public readonly volumeId!: pulumi.Output<string>;

    /**
     * Create a VolumeAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeAttachmentArgs | VolumeAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VolumeAttachmentState | undefined;
            inputs["deviceName"] = state ? state.deviceName : undefined;
            inputs["forceDetach"] = state ? state.forceDetach : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["skipDestroy"] = state ? state.skipDestroy : undefined;
            inputs["volumeId"] = state ? state.volumeId : undefined;
        } else {
            const args = argsOrState as VolumeAttachmentArgs | undefined;
            if (!args || args.deviceName === undefined) {
                throw new Error("Missing required property 'deviceName'");
            }
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            if (!args || args.volumeId === undefined) {
                throw new Error("Missing required property 'volumeId'");
            }
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["forceDetach"] = args ? args.forceDetach : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["skipDestroy"] = args ? args.skipDestroy : undefined;
            inputs["volumeId"] = args ? args.volumeId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VolumeAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VolumeAttachment resources.
 */
export interface VolumeAttachmentState {
    /**
     * The device name to expose to the instance (for
     * example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances][1] and [Device Naming on Windows Instances][2] for more information.
     */
    readonly deviceName?: pulumi.Input<string>;
    /**
     * Set to `true` if you want to force the
     * volume to detach. Useful if previous attempts failed, but use this option only
     * as a last resort, as this can result in **data loss**. See
     * [Detaching an Amazon EBS Volume from an Instance][3] for more information.
     */
    readonly forceDetach?: pulumi.Input<boolean>;
    /**
     * ID of the Instance to attach to
     */
    readonly instanceId?: pulumi.Input<string>;
    readonly skipDestroy?: pulumi.Input<boolean>;
    /**
     * ID of the Volume to be attached
     */
    readonly volumeId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VolumeAttachment resource.
 */
export interface VolumeAttachmentArgs {
    /**
     * The device name to expose to the instance (for
     * example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances][1] and [Device Naming on Windows Instances][2] for more information.
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * Set to `true` if you want to force the
     * volume to detach. Useful if previous attempts failed, but use this option only
     * as a last resort, as this can result in **data loss**. See
     * [Detaching an Amazon EBS Volume from an Instance][3] for more information.
     */
    readonly forceDetach?: pulumi.Input<boolean>;
    /**
     * ID of the Instance to attach to
     */
    readonly instanceId: pulumi.Input<string>;
    readonly skipDestroy?: pulumi.Input<boolean>;
    /**
     * ID of the Volume to be attached
     */
    readonly volumeId: pulumi.Input<string>;
}
