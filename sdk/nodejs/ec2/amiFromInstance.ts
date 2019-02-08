// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The "AMI from instance" resource allows the creation of an Amazon Machine
 * Image (AMI) modelled after an existing EBS-backed EC2 instance.
 * 
 * The created AMI will refer to implicitly-created snapshots of the instance's
 * EBS volumes and mimick its assigned block device configuration at the time
 * the resource is created.
 * 
 * This resource is best applied to an instance that is stopped when this instance
 * is created, so that the contents of the created image are predictable. When
 * applied to an instance that is running, *the instance will be stopped before taking
 * the snapshots and then started back up again*, resulting in a period of
 * downtime.
 * 
 * Note that the source instance is inspected only at the initial creation of this
 * resource. Ongoing updates to the referenced instance will not be propagated into
 * the generated AMI. Users may taint or otherwise recreate the resource in order
 * to produce a fresh snapshot.
 */
export class AmiFromInstance extends pulumi.CustomResource {
    /**
     * Get an existing AmiFromInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AmiFromInstanceState, opts?: pulumi.CustomResourceOptions): AmiFromInstance {
        return new AmiFromInstance(name, <any>state, { ...opts, id: id });
    }

    /**
     * Machine architecture for created instances. Defaults to "x86_64".
     */
    public /*out*/ readonly architecture: pulumi.Output<string>;
    /**
     * A longer, human-readable description for the AMI.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * Nested block describing an EBS block device that should be
     * attached to created instances. The structure of this block is described below.
     */
    public readonly ebsBlockDevices: pulumi.Output<{ deleteOnTermination: boolean, deviceName: string, encrypted: boolean, iops: number, snapshotId: string, volumeSize: number, volumeType: string }[]>;
    /**
     * Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
     */
    public /*out*/ readonly enaSupport: pulumi.Output<boolean>;
    /**
     * Nested block describing an ephemeral block device that
     * should be attached to created instances. The structure of this block is described below.
     */
    public readonly ephemeralBlockDevices: pulumi.Output<{ deviceName: string, virtualName: string }[]>;
    /**
     * Path to an S3 object containing an image manifest, e.g. created
     * by the `ec2-upload-bundle` command in the EC2 command line tools.
     */
    public /*out*/ readonly imageLocation: pulumi.Output<string>;
    /**
     * The id of the kernel image (AKI) that will be used as the paravirtual
     * kernel in created instances.
     */
    public /*out*/ readonly kernelId: pulumi.Output<string>;
    public /*out*/ readonly manageEbsSnapshots: pulumi.Output<boolean>;
    /**
     * A region-unique name for the AMI.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The id of an initrd image (ARI) that will be used when booting the
     * created instances.
     */
    public /*out*/ readonly ramdiskId: pulumi.Output<string>;
    /**
     * The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
     */
    public /*out*/ readonly rootDeviceName: pulumi.Output<string>;
    public /*out*/ readonly rootSnapshotId: pulumi.Output<string>;
    /**
     * Boolean that overrides the behavior of stopping
     * the instance before snapshotting. This is risky since it may cause a snapshot of an
     * inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
     * guarantees that no filesystem writes will be underway at the time of snapshot.
     */
    public readonly snapshotWithoutReboot: pulumi.Output<boolean | undefined>;
    /**
     * The id of the instance to use as the basis of the AMI.
     */
    public readonly sourceInstanceId: pulumi.Output<string>;
    /**
     * When set to "simple" (the default), enables enhanced networking
     * for created instances. No other value is supported at this time.
     */
    public /*out*/ readonly sriovNetSupport: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Keyword to choose what virtualization mode created instances
     * will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
     * changes the set of further arguments that are required, as described below.
     */
    public /*out*/ readonly virtualizationType: pulumi.Output<string>;

    /**
     * Create a AmiFromInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AmiFromInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AmiFromInstanceArgs | AmiFromInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AmiFromInstanceState = argsOrState as AmiFromInstanceState | undefined;
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
            inputs["snapshotWithoutReboot"] = state ? state.snapshotWithoutReboot : undefined;
            inputs["sourceInstanceId"] = state ? state.sourceInstanceId : undefined;
            inputs["sriovNetSupport"] = state ? state.sriovNetSupport : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualizationType"] = state ? state.virtualizationType : undefined;
        } else {
            const args = argsOrState as AmiFromInstanceArgs | undefined;
            if (!args || args.sourceInstanceId === undefined) {
                throw new Error("Missing required property 'sourceInstanceId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["ebsBlockDevices"] = args ? args.ebsBlockDevices : undefined;
            inputs["ephemeralBlockDevices"] = args ? args.ephemeralBlockDevices : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["snapshotWithoutReboot"] = args ? args.snapshotWithoutReboot : undefined;
            inputs["sourceInstanceId"] = args ? args.sourceInstanceId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["architecture"] = undefined /*out*/;
            inputs["enaSupport"] = undefined /*out*/;
            inputs["imageLocation"] = undefined /*out*/;
            inputs["kernelId"] = undefined /*out*/;
            inputs["manageEbsSnapshots"] = undefined /*out*/;
            inputs["ramdiskId"] = undefined /*out*/;
            inputs["rootDeviceName"] = undefined /*out*/;
            inputs["rootSnapshotId"] = undefined /*out*/;
            inputs["sriovNetSupport"] = undefined /*out*/;
            inputs["virtualizationType"] = undefined /*out*/;
        }
        super("aws:ec2/amiFromInstance:AmiFromInstance", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AmiFromInstance resources.
 */
export interface AmiFromInstanceState {
    /**
     * Machine architecture for created instances. Defaults to "x86_64".
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
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<{ deleteOnTermination?: pulumi.Input<boolean>, deviceName?: pulumi.Input<string>, encrypted?: pulumi.Input<boolean>, iops?: pulumi.Input<number>, snapshotId?: pulumi.Input<string>, volumeSize?: pulumi.Input<number>, volumeType?: pulumi.Input<string> }>[]>;
    /**
     * Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
     */
    readonly enaSupport?: pulumi.Input<boolean>;
    /**
     * Nested block describing an ephemeral block device that
     * should be attached to created instances. The structure of this block is described below.
     */
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<{ deviceName?: pulumi.Input<string>, virtualName?: pulumi.Input<string> }>[]>;
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
    readonly rootSnapshotId?: pulumi.Input<string>;
    /**
     * Boolean that overrides the behavior of stopping
     * the instance before snapshotting. This is risky since it may cause a snapshot of an
     * inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
     * guarantees that no filesystem writes will be underway at the time of snapshot.
     */
    readonly snapshotWithoutReboot?: pulumi.Input<boolean>;
    /**
     * The id of the instance to use as the basis of the AMI.
     */
    readonly sourceInstanceId?: pulumi.Input<string>;
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
 * The set of arguments for constructing a AmiFromInstance resource.
 */
export interface AmiFromInstanceArgs {
    /**
     * A longer, human-readable description for the AMI.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Nested block describing an EBS block device that should be
     * attached to created instances. The structure of this block is described below.
     */
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<{ deleteOnTermination?: pulumi.Input<boolean>, deviceName?: pulumi.Input<string>, encrypted?: pulumi.Input<boolean>, iops?: pulumi.Input<number>, snapshotId?: pulumi.Input<string>, volumeSize?: pulumi.Input<number>, volumeType?: pulumi.Input<string> }>[]>;
    /**
     * Nested block describing an ephemeral block device that
     * should be attached to created instances. The structure of this block is described below.
     */
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<{ deviceName?: pulumi.Input<string>, virtualName?: pulumi.Input<string> }>[]>;
    /**
     * A region-unique name for the AMI.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Boolean that overrides the behavior of stopping
     * the instance before snapshotting. This is risky since it may cause a snapshot of an
     * inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
     * guarantees that no filesystem writes will be underway at the time of snapshot.
     */
    readonly snapshotWithoutReboot?: pulumi.Input<boolean>;
    /**
     * The id of the instance to use as the basis of the AMI.
     */
    readonly sourceInstanceId: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
