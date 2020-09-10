// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:opsworks/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    public readonly agentVersion!: pulumi.Output<string | undefined>;
    public readonly amiId!: pulumi.Output<string>;
    public readonly architecture!: pulumi.Output<string | undefined>;
    public readonly autoScalingType!: pulumi.Output<string | undefined>;
    public readonly availabilityZone!: pulumi.Output<string>;
    public readonly createdAt!: pulumi.Output<string>;
    public readonly deleteEbs!: pulumi.Output<boolean | undefined>;
    public readonly deleteEip!: pulumi.Output<boolean | undefined>;
    public readonly ebsBlockDevices!: pulumi.Output<outputs.opsworks.InstanceEbsBlockDevice[]>;
    public readonly ebsOptimized!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly ec2InstanceId!: pulumi.Output<string>;
    public readonly ecsClusterArn!: pulumi.Output<string>;
    public readonly elasticIp!: pulumi.Output<string>;
    public readonly ephemeralBlockDevices!: pulumi.Output<outputs.opsworks.InstanceEphemeralBlockDevice[]>;
    public readonly hostname!: pulumi.Output<string>;
    public readonly infrastructureClass!: pulumi.Output<string>;
    public readonly installUpdatesOnBoot!: pulumi.Output<boolean | undefined>;
    public readonly instanceProfileArn!: pulumi.Output<string>;
    public readonly instanceType!: pulumi.Output<string | undefined>;
    public readonly lastServiceErrorId!: pulumi.Output<string>;
    public readonly layerIds!: pulumi.Output<string[]>;
    public readonly os!: pulumi.Output<string>;
    public readonly platform!: pulumi.Output<string>;
    public readonly privateDns!: pulumi.Output<string>;
    public readonly privateIp!: pulumi.Output<string>;
    public readonly publicDns!: pulumi.Output<string>;
    public readonly publicIp!: pulumi.Output<string>;
    public readonly registeredBy!: pulumi.Output<string>;
    public readonly reportedAgentVersion!: pulumi.Output<string>;
    public readonly reportedOsFamily!: pulumi.Output<string>;
    public readonly reportedOsName!: pulumi.Output<string>;
    public readonly reportedOsVersion!: pulumi.Output<string>;
    public readonly rootBlockDevices!: pulumi.Output<outputs.opsworks.InstanceRootBlockDevice[]>;
    public readonly rootDeviceType!: pulumi.Output<string>;
    public readonly rootDeviceVolumeId!: pulumi.Output<string>;
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    public readonly sshHostDsaKeyFingerprint!: pulumi.Output<string>;
    public readonly sshHostRsaKeyFingerprint!: pulumi.Output<string>;
    public readonly sshKeyName!: pulumi.Output<string>;
    public readonly stackId!: pulumi.Output<string>;
    public readonly state!: pulumi.Output<string | undefined>;
    public readonly status!: pulumi.Output<string>;
    public readonly subnetId!: pulumi.Output<string>;
    public readonly tenancy!: pulumi.Output<string>;
    public readonly virtualizationType!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["agentVersion"] = state ? state.agentVersion : undefined;
            inputs["amiId"] = state ? state.amiId : undefined;
            inputs["architecture"] = state ? state.architecture : undefined;
            inputs["autoScalingType"] = state ? state.autoScalingType : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["deleteEbs"] = state ? state.deleteEbs : undefined;
            inputs["deleteEip"] = state ? state.deleteEip : undefined;
            inputs["ebsBlockDevices"] = state ? state.ebsBlockDevices : undefined;
            inputs["ebsOptimized"] = state ? state.ebsOptimized : undefined;
            inputs["ec2InstanceId"] = state ? state.ec2InstanceId : undefined;
            inputs["ecsClusterArn"] = state ? state.ecsClusterArn : undefined;
            inputs["elasticIp"] = state ? state.elasticIp : undefined;
            inputs["ephemeralBlockDevices"] = state ? state.ephemeralBlockDevices : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["infrastructureClass"] = state ? state.infrastructureClass : undefined;
            inputs["installUpdatesOnBoot"] = state ? state.installUpdatesOnBoot : undefined;
            inputs["instanceProfileArn"] = state ? state.instanceProfileArn : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["lastServiceErrorId"] = state ? state.lastServiceErrorId : undefined;
            inputs["layerIds"] = state ? state.layerIds : undefined;
            inputs["os"] = state ? state.os : undefined;
            inputs["platform"] = state ? state.platform : undefined;
            inputs["privateDns"] = state ? state.privateDns : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["publicDns"] = state ? state.publicDns : undefined;
            inputs["publicIp"] = state ? state.publicIp : undefined;
            inputs["registeredBy"] = state ? state.registeredBy : undefined;
            inputs["reportedAgentVersion"] = state ? state.reportedAgentVersion : undefined;
            inputs["reportedOsFamily"] = state ? state.reportedOsFamily : undefined;
            inputs["reportedOsName"] = state ? state.reportedOsName : undefined;
            inputs["reportedOsVersion"] = state ? state.reportedOsVersion : undefined;
            inputs["rootBlockDevices"] = state ? state.rootBlockDevices : undefined;
            inputs["rootDeviceType"] = state ? state.rootDeviceType : undefined;
            inputs["rootDeviceVolumeId"] = state ? state.rootDeviceVolumeId : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["sshHostDsaKeyFingerprint"] = state ? state.sshHostDsaKeyFingerprint : undefined;
            inputs["sshHostRsaKeyFingerprint"] = state ? state.sshHostRsaKeyFingerprint : undefined;
            inputs["sshKeyName"] = state ? state.sshKeyName : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tenancy"] = state ? state.tenancy : undefined;
            inputs["virtualizationType"] = state ? state.virtualizationType : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.layerIds === undefined) {
                throw new Error("Missing required property 'layerIds'");
            }
            if (!args || args.stackId === undefined) {
                throw new Error("Missing required property 'stackId'");
            }
            inputs["agentVersion"] = args ? args.agentVersion : undefined;
            inputs["amiId"] = args ? args.amiId : undefined;
            inputs["architecture"] = args ? args.architecture : undefined;
            inputs["autoScalingType"] = args ? args.autoScalingType : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["createdAt"] = args ? args.createdAt : undefined;
            inputs["deleteEbs"] = args ? args.deleteEbs : undefined;
            inputs["deleteEip"] = args ? args.deleteEip : undefined;
            inputs["ebsBlockDevices"] = args ? args.ebsBlockDevices : undefined;
            inputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            inputs["ecsClusterArn"] = args ? args.ecsClusterArn : undefined;
            inputs["elasticIp"] = args ? args.elasticIp : undefined;
            inputs["ephemeralBlockDevices"] = args ? args.ephemeralBlockDevices : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["infrastructureClass"] = args ? args.infrastructureClass : undefined;
            inputs["installUpdatesOnBoot"] = args ? args.installUpdatesOnBoot : undefined;
            inputs["instanceProfileArn"] = args ? args.instanceProfileArn : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["lastServiceErrorId"] = args ? args.lastServiceErrorId : undefined;
            inputs["layerIds"] = args ? args.layerIds : undefined;
            inputs["os"] = args ? args.os : undefined;
            inputs["platform"] = args ? args.platform : undefined;
            inputs["privateDns"] = args ? args.privateDns : undefined;
            inputs["privateIp"] = args ? args.privateIp : undefined;
            inputs["publicDns"] = args ? args.publicDns : undefined;
            inputs["publicIp"] = args ? args.publicIp : undefined;
            inputs["registeredBy"] = args ? args.registeredBy : undefined;
            inputs["reportedAgentVersion"] = args ? args.reportedAgentVersion : undefined;
            inputs["reportedOsFamily"] = args ? args.reportedOsFamily : undefined;
            inputs["reportedOsName"] = args ? args.reportedOsName : undefined;
            inputs["reportedOsVersion"] = args ? args.reportedOsVersion : undefined;
            inputs["rootBlockDevices"] = args ? args.rootBlockDevices : undefined;
            inputs["rootDeviceType"] = args ? args.rootDeviceType : undefined;
            inputs["rootDeviceVolumeId"] = args ? args.rootDeviceVolumeId : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["sshHostDsaKeyFingerprint"] = args ? args.sshHostDsaKeyFingerprint : undefined;
            inputs["sshHostRsaKeyFingerprint"] = args ? args.sshHostRsaKeyFingerprint : undefined;
            inputs["sshKeyName"] = args ? args.sshKeyName : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tenancy"] = args ? args.tenancy : undefined;
            inputs["virtualizationType"] = args ? args.virtualizationType : undefined;
            inputs["ec2InstanceId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    readonly agentVersion?: pulumi.Input<string>;
    readonly amiId?: pulumi.Input<string>;
    readonly architecture?: pulumi.Input<string>;
    readonly autoScalingType?: pulumi.Input<string>;
    readonly availabilityZone?: pulumi.Input<string>;
    readonly createdAt?: pulumi.Input<string>;
    readonly deleteEbs?: pulumi.Input<boolean>;
    readonly deleteEip?: pulumi.Input<boolean>;
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<inputs.opsworks.InstanceEbsBlockDevice>[]>;
    readonly ebsOptimized?: pulumi.Input<boolean>;
    readonly ec2InstanceId?: pulumi.Input<string>;
    readonly ecsClusterArn?: pulumi.Input<string>;
    readonly elasticIp?: pulumi.Input<string>;
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<inputs.opsworks.InstanceEphemeralBlockDevice>[]>;
    readonly hostname?: pulumi.Input<string>;
    readonly infrastructureClass?: pulumi.Input<string>;
    readonly installUpdatesOnBoot?: pulumi.Input<boolean>;
    readonly instanceProfileArn?: pulumi.Input<string>;
    readonly instanceType?: pulumi.Input<string>;
    readonly lastServiceErrorId?: pulumi.Input<string>;
    readonly layerIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly os?: pulumi.Input<string>;
    readonly platform?: pulumi.Input<string>;
    readonly privateDns?: pulumi.Input<string>;
    readonly privateIp?: pulumi.Input<string>;
    readonly publicDns?: pulumi.Input<string>;
    readonly publicIp?: pulumi.Input<string>;
    readonly registeredBy?: pulumi.Input<string>;
    readonly reportedAgentVersion?: pulumi.Input<string>;
    readonly reportedOsFamily?: pulumi.Input<string>;
    readonly reportedOsName?: pulumi.Input<string>;
    readonly reportedOsVersion?: pulumi.Input<string>;
    readonly rootBlockDevices?: pulumi.Input<pulumi.Input<inputs.opsworks.InstanceRootBlockDevice>[]>;
    readonly rootDeviceType?: pulumi.Input<string>;
    readonly rootDeviceVolumeId?: pulumi.Input<string>;
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly sshHostDsaKeyFingerprint?: pulumi.Input<string>;
    readonly sshHostRsaKeyFingerprint?: pulumi.Input<string>;
    readonly sshKeyName?: pulumi.Input<string>;
    readonly stackId?: pulumi.Input<string>;
    readonly state?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly subnetId?: pulumi.Input<string>;
    readonly tenancy?: pulumi.Input<string>;
    readonly virtualizationType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    readonly agentVersion?: pulumi.Input<string>;
    readonly amiId?: pulumi.Input<string>;
    readonly architecture?: pulumi.Input<string>;
    readonly autoScalingType?: pulumi.Input<string>;
    readonly availabilityZone?: pulumi.Input<string>;
    readonly createdAt?: pulumi.Input<string>;
    readonly deleteEbs?: pulumi.Input<boolean>;
    readonly deleteEip?: pulumi.Input<boolean>;
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<inputs.opsworks.InstanceEbsBlockDevice>[]>;
    readonly ebsOptimized?: pulumi.Input<boolean>;
    readonly ecsClusterArn?: pulumi.Input<string>;
    readonly elasticIp?: pulumi.Input<string>;
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<inputs.opsworks.InstanceEphemeralBlockDevice>[]>;
    readonly hostname?: pulumi.Input<string>;
    readonly infrastructureClass?: pulumi.Input<string>;
    readonly installUpdatesOnBoot?: pulumi.Input<boolean>;
    readonly instanceProfileArn?: pulumi.Input<string>;
    readonly instanceType?: pulumi.Input<string>;
    readonly lastServiceErrorId?: pulumi.Input<string>;
    readonly layerIds: pulumi.Input<pulumi.Input<string>[]>;
    readonly os?: pulumi.Input<string>;
    readonly platform?: pulumi.Input<string>;
    readonly privateDns?: pulumi.Input<string>;
    readonly privateIp?: pulumi.Input<string>;
    readonly publicDns?: pulumi.Input<string>;
    readonly publicIp?: pulumi.Input<string>;
    readonly registeredBy?: pulumi.Input<string>;
    readonly reportedAgentVersion?: pulumi.Input<string>;
    readonly reportedOsFamily?: pulumi.Input<string>;
    readonly reportedOsName?: pulumi.Input<string>;
    readonly reportedOsVersion?: pulumi.Input<string>;
    readonly rootBlockDevices?: pulumi.Input<pulumi.Input<inputs.opsworks.InstanceRootBlockDevice>[]>;
    readonly rootDeviceType?: pulumi.Input<string>;
    readonly rootDeviceVolumeId?: pulumi.Input<string>;
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly sshHostDsaKeyFingerprint?: pulumi.Input<string>;
    readonly sshHostRsaKeyFingerprint?: pulumi.Input<string>;
    readonly sshKeyName?: pulumi.Input<string>;
    readonly stackId: pulumi.Input<string>;
    readonly state?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly subnetId?: pulumi.Input<string>;
    readonly tenancy?: pulumi.Input<string>;
    readonly virtualizationType?: pulumi.Input<string>;
}
