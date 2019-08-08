// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AMI resource allows the creation and management of a completely-custom
 * *Amazon Machine Image* (AMI).
 * 
 * If you just want to duplicate an existing AMI, possibly copying it to another
 * region, it's better to use `aws.ec2.AmiCopy` instead.
 * 
 * If you just want to share an existing AMI with another AWS account,
 * it's better to use `aws.ec2.AmiLaunchPermission` instead.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * // Create an AMI that will start a machine whose root device is backed by
 * // an EBS volume populated from a snapshot. It is assumed that such a snapshot
 * // already exists with the id "snap-xxxxxxxx".
 * const example = new aws.ec2.Ami("example", {
 *     ebsBlockDevices: [{
 *         deviceName: "/dev/xvda",
 *         snapshotId: "snap-xxxxxxxx",
 *         volumeSize: 8,
 *     }],
 *     rootDeviceName: "/dev/xvda",
 *     virtualizationType: "hvm",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ami.html.markdown.
 */
export class Ami extends pulumi.CustomResource {
    /**
     * Get an existing Ami resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AmiState, opts?: pulumi.CustomResourceOptions): Ami {
        return new Ami(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/ami:Ami';

    /**
     * Returns true if the given object is an instance of Ami.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ami {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ami.__pulumiType;
    }

    /**
     * Machine architecture for created instances. Defaults to "x8664".
     */
    public readonly architecture!: pulumi.Output<string | undefined>;
    /**
     * A longer, human-readable description for the AMI.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Nested block describing an EBS block device that should be
     * attached to created instances. The structure of this block is described below.
     */
    public readonly ebsBlockDevices!: pulumi.Output<{ deleteOnTermination?: boolean, deviceName: string, encrypted?: boolean, iops?: number, snapshotId?: string, volumeSize: number, volumeType?: string }[]>;
    /**
     * Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
     */
    public readonly enaSupport!: pulumi.Output<boolean | undefined>;
    /**
     * Nested block describing an ephemeral block device that
     * should be attached to created instances. The structure of this block is described below.
     */
    public readonly ephemeralBlockDevices!: pulumi.Output<{ deviceName: string, virtualName: string }[]>;
    /**
     * Path to an S3 object containing an image manifest, e.g. created
     * by the `ec2-upload-bundle` command in the EC2 command line tools.
     */
    public readonly imageLocation!: pulumi.Output<string>;
    /**
     * The id of the kernel image (AKI) that will be used as the paravirtual
     * kernel in created instances.
     */
    public readonly kernelId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly manageEbsSnapshots!: pulumi.Output<boolean>;
    /**
     * A region-unique name for the AMI.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of an initrd image (ARI) that will be used when booting the
     * created instances.
     */
    public readonly ramdiskId!: pulumi.Output<string | undefined>;
    /**
     * The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
     */
    public readonly rootDeviceName!: pulumi.Output<string | undefined>;
    /**
     * The Snapshot ID for the root volume (for EBS-backed AMIs)
     */
    public /*out*/ readonly rootSnapshotId!: pulumi.Output<string>;
    /**
     * When set to "simple" (the default), enables enhanced networking
     * for created instances. No other value is supported at this time.
     */
    public readonly sriovNetSupport!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Keyword to choose what virtualization mode created instances
     * will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
     * changes the set of further arguments that are required, as described below.
     */
    public readonly virtualizationType!: pulumi.Output<string | undefined>;

    /**
     * Create a Ami resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AmiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AmiArgs | AmiState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AmiState | undefined;
            inputs["architecture"] = state ? state.architecture : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ebsBlockDevices"] = state ? state.ebsBlockDevices : undefined;
            inputs["enaSupport"] = state ? state.enaSupport : undefined;
            inputs["ephemeralBlockDevices"] = state ? state.ephemeralBlockDevices : undefined;
            inputs["imageLocation"] = state ? state.imageLocation : undefined;
            inputs["kernelId"] = state ? state.kernelId : undefined;
            inputs["manageEbsSnapshots"] = state ? state.manageEbsSnapshots : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ramdiskId"] = state ? state.ramdiskId : undefined;
            inputs["rootDeviceName"] = state ? state.rootDeviceName : undefined;
            inputs["rootSnapshotId"] = state ? state.rootSnapshotId : undefined;
            inputs["sriovNetSupport"] = state ? state.sriovNetSupport : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualizationType"] = state ? state.virtualizationType : undefined;
        } else {
            const args = argsOrState as AmiArgs | undefined;
            inputs["architecture"] = args ? args.architecture : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["ebsBlockDevices"] = args ? args.ebsBlockDevices : undefined;
            inputs["enaSupport"] = args ? args.enaSupport : undefined;
            inputs["ephemeralBlockDevices"] = args ? args.ephemeralBlockDevices : undefined;
            inputs["imageLocation"] = args ? args.imageLocation : undefined;
            inputs["kernelId"] = args ? args.kernelId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ramdiskId"] = args ? args.ramdiskId : undefined;
            inputs["rootDeviceName"] = args ? args.rootDeviceName : undefined;
            inputs["sriovNetSupport"] = args ? args.sriovNetSupport : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualizationType"] = args ? args.virtualizationType : undefined;
            inputs["manageEbsSnapshots"] = undefined /*out*/;
            inputs["rootSnapshotId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Ami.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ami resources.
 */
export interface AmiState {
    /**
     * Machine architecture for created instances. Defaults to "x8664".
     */
    readonly architecture?: pulumi.Input<string>;
    /**
     * A longer, human-readable description for the AMI.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Nested block describing an EBS block device that should be
     * attached to created instances. The structure of this block is described below.
     */
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<{ deleteOnTermination?: pulumi.Input<boolean>, deviceName: pulumi.Input<string>, encrypted?: pulumi.Input<boolean>, iops?: pulumi.Input<number>, snapshotId?: pulumi.Input<string>, volumeSize?: pulumi.Input<number>, volumeType?: pulumi.Input<string> }>[]>;
    /**
     * Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
     */
    readonly enaSupport?: pulumi.Input<boolean>;
    /**
     * Nested block describing an ephemeral block device that
     * should be attached to created instances. The structure of this block is described below.
     */
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<{ deviceName: pulumi.Input<string>, virtualName: pulumi.Input<string> }>[]>;
    /**
     * Path to an S3 object containing an image manifest, e.g. created
     * by the `ec2-upload-bundle` command in the EC2 command line tools.
     */
    readonly imageLocation?: pulumi.Input<string>;
    /**
     * The id of the kernel image (AKI) that will be used as the paravirtual
     * kernel in created instances.
     */
    readonly kernelId?: pulumi.Input<string>;
    readonly manageEbsSnapshots?: pulumi.Input<boolean>;
    /**
     * A region-unique name for the AMI.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The id of an initrd image (ARI) that will be used when booting the
     * created instances.
     */
    readonly ramdiskId?: pulumi.Input<string>;
    /**
     * The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
     */
    readonly rootDeviceName?: pulumi.Input<string>;
    /**
     * The Snapshot ID for the root volume (for EBS-backed AMIs)
     */
    readonly rootSnapshotId?: pulumi.Input<string>;
    /**
     * When set to "simple" (the default), enables enhanced networking
     * for created instances. No other value is supported at this time.
     */
    readonly sriovNetSupport?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Keyword to choose what virtualization mode created instances
     * will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
     * changes the set of further arguments that are required, as described below.
     */
    readonly virtualizationType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ami resource.
 */
export interface AmiArgs {
    /**
     * Machine architecture for created instances. Defaults to "x8664".
     */
    readonly architecture?: pulumi.Input<string>;
    /**
     * A longer, human-readable description for the AMI.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Nested block describing an EBS block device that should be
     * attached to created instances. The structure of this block is described below.
     */
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<{ deleteOnTermination?: pulumi.Input<boolean>, deviceName: pulumi.Input<string>, encrypted?: pulumi.Input<boolean>, iops?: pulumi.Input<number>, snapshotId?: pulumi.Input<string>, volumeSize?: pulumi.Input<number>, volumeType?: pulumi.Input<string> }>[]>;
    /**
     * Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
     */
    readonly enaSupport?: pulumi.Input<boolean>;
    /**
     * Nested block describing an ephemeral block device that
     * should be attached to created instances. The structure of this block is described below.
     */
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<{ deviceName: pulumi.Input<string>, virtualName: pulumi.Input<string> }>[]>;
    /**
     * Path to an S3 object containing an image manifest, e.g. created
     * by the `ec2-upload-bundle` command in the EC2 command line tools.
     */
    readonly imageLocation?: pulumi.Input<string>;
    /**
     * The id of the kernel image (AKI) that will be used as the paravirtual
     * kernel in created instances.
     */
    readonly kernelId?: pulumi.Input<string>;
    /**
     * A region-unique name for the AMI.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The id of an initrd image (ARI) that will be used when booting the
     * created instances.
     */
    readonly ramdiskId?: pulumi.Input<string>;
    /**
     * The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
     */
    readonly rootDeviceName?: pulumi.Input<string>;
    /**
     * When set to "simple" (the default), enables enhanced networking
     * for created instances. No other value is supported at this time.
     */
    readonly sriovNetSupport?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Keyword to choose what virtualization mode created instances
     * will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
     * changes the set of further arguments that are required, as described below.
     */
    readonly virtualizationType?: pulumi.Input<string>;
}
