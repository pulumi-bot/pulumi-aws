// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

import {InstanceProfile} from "../iam";

/**
 * Provides a resource to create a new launch configuration, used for autoscaling groups.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ubuntu = aws.ec2.getAmi({
 *     mostRecent: true,
 *     filters: [
 *         {
 *             name: "name",
 *             values: ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
 *         },
 *         {
 *             name: "virtualization-type",
 *             values: ["hvm"],
 *         },
 *     ],
 *     owners: ["099720109477"],
 * });
 * const asConf = new aws.ec2.LaunchConfiguration("asConf", {
 *     imageId: ubuntu.then(ubuntu => ubuntu.id),
 *     instanceType: "t2.micro",
 * });
 * ```
 * ## Using with AutoScaling Groups
 *
 * Launch Configurations cannot be updated after creation with the Amazon
 * Web Service API. In order to update a Launch Configuration, this provider will
 * destroy the existing resource and create a replacement. In order to effectively
 * use a Launch Configuration resource with an [AutoScaling Group resource](https://www.terraform.io/docs/providers/aws/r/autoscaling_group.html),
 * it's recommended to specify `createBeforeDestroy` in a [lifecycle](https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html) block.
 * Either omit the Launch Configuration `name` attribute, or specify a partial name
 * with `namePrefix`.  Example:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ubuntu = aws.ec2.getAmi({
 *     mostRecent: true,
 *     filters: [
 *         {
 *             name: "name",
 *             values: ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
 *         },
 *         {
 *             name: "virtualization-type",
 *             values: ["hvm"],
 *         },
 *     ],
 *     owners: ["099720109477"],
 * });
 * const asConf = new aws.ec2.LaunchConfiguration("asConf", {
 *     namePrefix: "lc-example-",
 *     imageId: ubuntu.then(ubuntu => ubuntu.id),
 *     instanceType: "t2.micro",
 * });
 * const bar = new aws.autoscaling.Group("bar", {
 *     launchConfiguration: asConf.name,
 *     minSize: 1,
 *     maxSize: 2,
 * });
 * ```
 *
 * With this setup this provider generates a unique name for your Launch
 * Configuration and can then update the AutoScaling Group without conflict before
 * destroying the previous Launch Configuration.
 *
 * ## Using with Spot Instances
 *
 * Launch configurations can set the spot instance pricing to be used for the
 * Auto Scaling Group to reserve instances. Simply specifying the `spotPrice`
 * parameter will set the price on the Launch Configuration which will attempt to
 * reserve your instances at this price.  See the [AWS Spot Instance
 * documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances.html)
 * for more information or how to launch [Spot Instances](https://www.terraform.io/docs/providers/aws/r/spot_instance_request.html) with this provider.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ubuntu = aws.ec2.getAmi({
 *     mostRecent: true,
 *     filters: [
 *         {
 *             name: "name",
 *             values: ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
 *         },
 *         {
 *             name: "virtualization-type",
 *             values: ["hvm"],
 *         },
 *     ],
 *     owners: ["099720109477"],
 * });
 * const asConf = new aws.ec2.LaunchConfiguration("asConf", {
 *     imageId: ubuntu.then(ubuntu => ubuntu.id),
 *     instanceType: "m4.large",
 *     spotPrice: "0.001",
 * });
 * const bar = new aws.autoscaling.Group("bar", {launchConfiguration: asConf.name});
 * ```
 *
 * ## Block devices
 *
 * Each of the `*_block_device` attributes controls a portion of the AWS
 * Launch Configuration's "Block Device Mapping". It's a good idea to familiarize yourself with [AWS's Block Device
 * Mapping docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
 * to understand the implications of using these attributes.
 *
 * The `rootBlockDevice` mapping supports the following:
 *
 * * `volumeType` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`, `"gp3"`, `"st1"`, `"sc1"`
 *   or `"io1"`. (Default: `"standard"`).
 * * `volumeSize` - (Optional) The size of the volume in gigabytes.
 * * `iops` - (Optional) The amount of provisioned
 *   [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
 *   This must be set with a `volumeType` of `"io1"`.
 * * `throughput` - (Optional) The throughput (MiBps) to provision for a `gp3` volume.
 * * `deleteOnTermination` - (Optional) Whether the volume should be destroyed
 *   on instance termination (Default: `true`).
 * * `encrypted` - (Optional) Whether the volume should be encrypted or not. (Default: `false`).
 *
 * Modifying any of the `rootBlockDevice` settings requires resource
 * replacement.
 *
 * Each `ebsBlockDevice` supports the following:
 *
 * * `deviceName` - (Required) The name of the device to mount.
 * * `snapshotId` - (Optional) The Snapshot ID to mount.
 * * `volumeType` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`, `"gp3"`, `"st1"`, `"sc1"`
 *   or `"io1"`. (Default: `"standard"`).
 * * `volumeSize` - (Optional) The size of the volume in gigabytes.
 * * `iops` - (Optional) The amount of provisioned
 *   [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
 *   This must be set with a `volumeType` of `"io1"`.
 * * `throughput` - (Optional) The throughput (MiBps) to provision for a `gp3` volume.
 * * `deleteOnTermination` - (Optional) Whether the volume should be destroyed
 *   on instance termination (Default: `true`).
 * * `encrypted` - (Optional) Whether the volume should be encrypted or not. Do not use this option if you are using `snapshotId` as the encrypted flag will be determined by the snapshot. (Default: `false`).
 * * `noDevice` - (Optional) Whether the device in the block device mapping of the AMI is suppressed.
 *
 * Modifying any `ebsBlockDevice` currently requires resource replacement.
 *
 * Each `ephemeralBlockDevice` supports the following:
 *
 * * `deviceName` - The name of the block device to mount on the instance.
 * * `virtualName` - The [Instance Store Device
 *   Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
 *   (e.g. `"ephemeral0"`)
 *
 * Each AWS Instance type has a different set of Instance Store block devices
 * available for attachment. AWS [publishes a
 * list](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes)
 * of which ephemeral devices are available on each type. The devices are always
 * identified by the `virtualName` in the format `"ephemeral{0..N}"`.
 *
 * > **NOTE:** Changes to `*_block_device` configuration of _existing_ resources
 * cannot currently be detected by this provider. After updating to block device
 * configuration, resource recreation can be manually triggered by using the
 * [`up` command with the --replace argument](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
 *
 * ## Import
 *
 * Launch configurations can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:ec2/launchConfiguration:LaunchConfiguration as_conf lg-123456
 * ```
 */
export class LaunchConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing LaunchConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LaunchConfigurationState, opts?: pulumi.CustomResourceOptions): LaunchConfiguration {
        return new LaunchConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/launchConfiguration:LaunchConfiguration';

    /**
     * Returns true if the given object is an instance of LaunchConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LaunchConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LaunchConfiguration.__pulumiType;
    }

    /**
     * The Amazon Resource Name of the launch configuration.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Associate a public ip address with an instance in a VPC.
     */
    public readonly associatePublicIpAddress!: pulumi.Output<boolean | undefined>;
    /**
     * Additional EBS block devices to attach to the
     * instance.  See Block Devices below for details.
     */
    public readonly ebsBlockDevices!: pulumi.Output<outputs.ec2.LaunchConfigurationEbsBlockDevice[]>;
    /**
     * If true, the launched EC2 instance will be EBS-optimized.
     */
    public readonly ebsOptimized!: pulumi.Output<boolean>;
    /**
     * Enables/disables detailed monitoring. This is enabled by default.
     */
    public readonly enableMonitoring!: pulumi.Output<boolean | undefined>;
    /**
     * Customize Ephemeral (also known as
     * "Instance Store") volumes on the instance. See Block Devices below for details.
     */
    public readonly ephemeralBlockDevices!: pulumi.Output<outputs.ec2.LaunchConfigurationEphemeralBlockDevice[] | undefined>;
    /**
     * The name attribute of the IAM instance profile to associate
     * with launched instances.
     */
    public readonly iamInstanceProfile!: pulumi.Output<string | undefined>;
    /**
     * The EC2 image ID to launch.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The size of instance to launch.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The key name that should be used for the instance.
     */
    public readonly keyName!: pulumi.Output<string>;
    /**
     * The metadata options for the instance.
     */
    public readonly metadataOptions!: pulumi.Output<outputs.ec2.LaunchConfigurationMetadataOptions>;
    /**
     * The name of the launch configuration. If you leave
     * this blank, this provider will auto-generate a unique name. Conflicts with `namePrefix`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * The tenancy of the instance. Valid values are
     * `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
     * for more details
     */
    public readonly placementTenancy!: pulumi.Output<string | undefined>;
    /**
     * Customize details about the root block
     * device of the instance. See Block Devices below for details.
     */
    public readonly rootBlockDevice!: pulumi.Output<outputs.ec2.LaunchConfigurationRootBlockDevice>;
    /**
     * A list of associated security group IDS.
     */
    public readonly securityGroups!: pulumi.Output<string[] | undefined>;
    /**
     * The maximum price to use for reserving spot instances.
     */
    public readonly spotPrice!: pulumi.Output<string | undefined>;
    /**
     * The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
     */
    public readonly userData!: pulumi.Output<string | undefined>;
    /**
     * Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
     */
    public readonly userDataBase64!: pulumi.Output<string | undefined>;
    /**
     * The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
     */
    public readonly vpcClassicLinkId!: pulumi.Output<string | undefined>;
    /**
     * The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
     */
    public readonly vpcClassicLinkSecurityGroups!: pulumi.Output<string[] | undefined>;

    /**
     * Create a LaunchConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LaunchConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LaunchConfigurationArgs | LaunchConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LaunchConfigurationState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["associatePublicIpAddress"] = state ? state.associatePublicIpAddress : undefined;
            inputs["ebsBlockDevices"] = state ? state.ebsBlockDevices : undefined;
            inputs["ebsOptimized"] = state ? state.ebsOptimized : undefined;
            inputs["enableMonitoring"] = state ? state.enableMonitoring : undefined;
            inputs["ephemeralBlockDevices"] = state ? state.ephemeralBlockDevices : undefined;
            inputs["iamInstanceProfile"] = state ? state.iamInstanceProfile : undefined;
            inputs["imageId"] = state ? state.imageId : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["keyName"] = state ? state.keyName : undefined;
            inputs["metadataOptions"] = state ? state.metadataOptions : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["placementTenancy"] = state ? state.placementTenancy : undefined;
            inputs["rootBlockDevice"] = state ? state.rootBlockDevice : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["spotPrice"] = state ? state.spotPrice : undefined;
            inputs["userData"] = state ? state.userData : undefined;
            inputs["userDataBase64"] = state ? state.userDataBase64 : undefined;
            inputs["vpcClassicLinkId"] = state ? state.vpcClassicLinkId : undefined;
            inputs["vpcClassicLinkSecurityGroups"] = state ? state.vpcClassicLinkSecurityGroups : undefined;
        } else {
            const args = argsOrState as LaunchConfigurationArgs | undefined;
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            inputs["associatePublicIpAddress"] = args ? args.associatePublicIpAddress : undefined;
            inputs["ebsBlockDevices"] = args ? args.ebsBlockDevices : undefined;
            inputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            inputs["enableMonitoring"] = args ? args.enableMonitoring : undefined;
            inputs["ephemeralBlockDevices"] = args ? args.ephemeralBlockDevices : undefined;
            inputs["iamInstanceProfile"] = args ? args.iamInstanceProfile : undefined;
            inputs["imageId"] = args ? args.imageId : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["metadataOptions"] = args ? args.metadataOptions : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["placementTenancy"] = args ? args.placementTenancy : undefined;
            inputs["rootBlockDevice"] = args ? args.rootBlockDevice : undefined;
            inputs["securityGroups"] = args ? args.securityGroups : undefined;
            inputs["spotPrice"] = args ? args.spotPrice : undefined;
            inputs["userData"] = args ? args.userData : undefined;
            inputs["userDataBase64"] = args ? args.userDataBase64 : undefined;
            inputs["vpcClassicLinkId"] = args ? args.vpcClassicLinkId : undefined;
            inputs["vpcClassicLinkSecurityGroups"] = args ? args.vpcClassicLinkSecurityGroups : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LaunchConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LaunchConfiguration resources.
 */
export interface LaunchConfigurationState {
    /**
     * The Amazon Resource Name of the launch configuration.
     */
    arn?: pulumi.Input<string>;
    /**
     * Associate a public ip address with an instance in a VPC.
     */
    associatePublicIpAddress?: pulumi.Input<boolean>;
    /**
     * Additional EBS block devices to attach to the
     * instance.  See Block Devices below for details.
     */
    ebsBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchConfigurationEbsBlockDevice>[]>;
    /**
     * If true, the launched EC2 instance will be EBS-optimized.
     */
    ebsOptimized?: pulumi.Input<boolean>;
    /**
     * Enables/disables detailed monitoring. This is enabled by default.
     */
    enableMonitoring?: pulumi.Input<boolean>;
    /**
     * Customize Ephemeral (also known as
     * "Instance Store") volumes on the instance. See Block Devices below for details.
     */
    ephemeralBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchConfigurationEphemeralBlockDevice>[]>;
    /**
     * The name attribute of the IAM instance profile to associate
     * with launched instances.
     */
    iamInstanceProfile?: pulumi.Input<string | InstanceProfile>;
    /**
     * The EC2 image ID to launch.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The size of instance to launch.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The key name that should be used for the instance.
     */
    keyName?: pulumi.Input<string>;
    /**
     * The metadata options for the instance.
     */
    metadataOptions?: pulumi.Input<inputs.ec2.LaunchConfigurationMetadataOptions>;
    /**
     * The name of the launch configuration. If you leave
     * this blank, this provider will auto-generate a unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The tenancy of the instance. Valid values are
     * `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
     * for more details
     */
    placementTenancy?: pulumi.Input<string>;
    /**
     * Customize details about the root block
     * device of the instance. See Block Devices below for details.
     */
    rootBlockDevice?: pulumi.Input<inputs.ec2.LaunchConfigurationRootBlockDevice>;
    /**
     * A list of associated security group IDS.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum price to use for reserving spot instances.
     */
    spotPrice?: pulumi.Input<string>;
    /**
     * The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
     */
    userData?: pulumi.Input<string>;
    /**
     * Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
     */
    userDataBase64?: pulumi.Input<string>;
    /**
     * The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
     */
    vpcClassicLinkId?: pulumi.Input<string>;
    /**
     * The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
     */
    vpcClassicLinkSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a LaunchConfiguration resource.
 */
export interface LaunchConfigurationArgs {
    /**
     * Associate a public ip address with an instance in a VPC.
     */
    associatePublicIpAddress?: pulumi.Input<boolean>;
    /**
     * Additional EBS block devices to attach to the
     * instance.  See Block Devices below for details.
     */
    ebsBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchConfigurationEbsBlockDevice>[]>;
    /**
     * If true, the launched EC2 instance will be EBS-optimized.
     */
    ebsOptimized?: pulumi.Input<boolean>;
    /**
     * Enables/disables detailed monitoring. This is enabled by default.
     */
    enableMonitoring?: pulumi.Input<boolean>;
    /**
     * Customize Ephemeral (also known as
     * "Instance Store") volumes on the instance. See Block Devices below for details.
     */
    ephemeralBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchConfigurationEphemeralBlockDevice>[]>;
    /**
     * The name attribute of the IAM instance profile to associate
     * with launched instances.
     */
    iamInstanceProfile?: pulumi.Input<string | InstanceProfile>;
    /**
     * The EC2 image ID to launch.
     */
    imageId: pulumi.Input<string>;
    /**
     * The size of instance to launch.
     */
    instanceType: pulumi.Input<string>;
    /**
     * The key name that should be used for the instance.
     */
    keyName?: pulumi.Input<string>;
    /**
     * The metadata options for the instance.
     */
    metadataOptions?: pulumi.Input<inputs.ec2.LaunchConfigurationMetadataOptions>;
    /**
     * The name of the launch configuration. If you leave
     * this blank, this provider will auto-generate a unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The tenancy of the instance. Valid values are
     * `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
     * for more details
     */
    placementTenancy?: pulumi.Input<string>;
    /**
     * Customize details about the root block
     * device of the instance. See Block Devices below for details.
     */
    rootBlockDevice?: pulumi.Input<inputs.ec2.LaunchConfigurationRootBlockDevice>;
    /**
     * A list of associated security group IDS.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum price to use for reserving spot instances.
     */
    spotPrice?: pulumi.Input<string>;
    /**
     * The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
     */
    userData?: pulumi.Input<string>;
    /**
     * Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
     */
    userDataBase64?: pulumi.Input<string>;
    /**
     * The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
     */
    vpcClassicLinkId?: pulumi.Input<string>;
    /**
     * The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
     */
    vpcClassicLinkSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
}
