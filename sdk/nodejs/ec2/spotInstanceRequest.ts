// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class SpotInstanceRequest extends pulumi.CustomResource {
    /**
     * Get an existing SpotInstanceRequest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SpotInstanceRequestState, opts?: pulumi.CustomResourceOptions): SpotInstanceRequest {
        return new SpotInstanceRequest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/spotInstanceRequest:SpotInstanceRequest';

    /**
     * Returns true if the given object is an instance of SpotInstanceRequest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SpotInstanceRequest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SpotInstanceRequest.__pulumiType;
    }

    public readonly ami!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly associatePublicIpAddress!: pulumi.Output<boolean>;
    public readonly availabilityZone!: pulumi.Output<string>;
    public readonly blockDurationMinutes!: pulumi.Output<number | undefined>;
    public readonly cpuCoreCount!: pulumi.Output<number>;
    public readonly cpuThreadsPerCore!: pulumi.Output<number>;
    public readonly creditSpecification!: pulumi.Output<outputs.ec2.SpotInstanceRequestCreditSpecification | undefined>;
    public readonly disableApiTermination!: pulumi.Output<boolean | undefined>;
    public readonly ebsBlockDevices!: pulumi.Output<outputs.ec2.SpotInstanceRequestEbsBlockDevice[]>;
    public readonly ebsOptimized!: pulumi.Output<boolean | undefined>;
    public readonly ephemeralBlockDevices!: pulumi.Output<outputs.ec2.SpotInstanceRequestEphemeralBlockDevice[]>;
    public readonly getPasswordData!: pulumi.Output<boolean | undefined>;
    public readonly hibernation!: pulumi.Output<boolean | undefined>;
    public readonly hostId!: pulumi.Output<string>;
    public readonly iamInstanceProfile!: pulumi.Output<string | undefined>;
    public readonly instanceInitiatedShutdownBehavior!: pulumi.Output<string | undefined>;
    public readonly instanceInterruptionBehaviour!: pulumi.Output<string | undefined>;
    public /*out*/ readonly instanceState!: pulumi.Output<string>;
    public readonly instanceType!: pulumi.Output<string>;
    public readonly ipv6AddressCount!: pulumi.Output<number>;
    public readonly ipv6Addresses!: pulumi.Output<string[]>;
    public readonly keyName!: pulumi.Output<string>;
    public readonly launchGroup!: pulumi.Output<string | undefined>;
    public readonly metadataOptions!: pulumi.Output<outputs.ec2.SpotInstanceRequestMetadataOptions>;
    public readonly monitoring!: pulumi.Output<boolean | undefined>;
    public readonly networkInterfaces!: pulumi.Output<outputs.ec2.SpotInstanceRequestNetworkInterface[]>;
    public /*out*/ readonly outpostArn!: pulumi.Output<string>;
    public /*out*/ readonly passwordData!: pulumi.Output<string>;
    public readonly placementGroup!: pulumi.Output<string>;
    public /*out*/ readonly primaryNetworkInterfaceId!: pulumi.Output<string>;
    public /*out*/ readonly privateDns!: pulumi.Output<string>;
    public readonly privateIp!: pulumi.Output<string>;
    public /*out*/ readonly publicDns!: pulumi.Output<string>;
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    public readonly rootBlockDevice!: pulumi.Output<outputs.ec2.SpotInstanceRequestRootBlockDevice>;
    public readonly secondaryPrivateIps!: pulumi.Output<string[]>;
    public readonly securityGroups!: pulumi.Output<string[]>;
    public readonly sourceDestCheck!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly spotBidStatus!: pulumi.Output<string>;
    public /*out*/ readonly spotInstanceId!: pulumi.Output<string>;
    public readonly spotPrice!: pulumi.Output<string | undefined>;
    public /*out*/ readonly spotRequestState!: pulumi.Output<string>;
    public readonly spotType!: pulumi.Output<string | undefined>;
    public readonly subnetId!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly tenancy!: pulumi.Output<string>;
    public readonly userData!: pulumi.Output<string | undefined>;
    public readonly userDataBase64!: pulumi.Output<string | undefined>;
    public readonly validFrom!: pulumi.Output<string>;
    public readonly validUntil!: pulumi.Output<string>;
    public readonly volumeTags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[]>;
    public readonly waitForFulfillment!: pulumi.Output<boolean | undefined>;

    /**
     * Create a SpotInstanceRequest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SpotInstanceRequestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SpotInstanceRequestArgs | SpotInstanceRequestState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SpotInstanceRequestState | undefined;
            inputs["ami"] = state ? state.ami : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["associatePublicIpAddress"] = state ? state.associatePublicIpAddress : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["blockDurationMinutes"] = state ? state.blockDurationMinutes : undefined;
            inputs["cpuCoreCount"] = state ? state.cpuCoreCount : undefined;
            inputs["cpuThreadsPerCore"] = state ? state.cpuThreadsPerCore : undefined;
            inputs["creditSpecification"] = state ? state.creditSpecification : undefined;
            inputs["disableApiTermination"] = state ? state.disableApiTermination : undefined;
            inputs["ebsBlockDevices"] = state ? state.ebsBlockDevices : undefined;
            inputs["ebsOptimized"] = state ? state.ebsOptimized : undefined;
            inputs["ephemeralBlockDevices"] = state ? state.ephemeralBlockDevices : undefined;
            inputs["getPasswordData"] = state ? state.getPasswordData : undefined;
            inputs["hibernation"] = state ? state.hibernation : undefined;
            inputs["hostId"] = state ? state.hostId : undefined;
            inputs["iamInstanceProfile"] = state ? state.iamInstanceProfile : undefined;
            inputs["instanceInitiatedShutdownBehavior"] = state ? state.instanceInitiatedShutdownBehavior : undefined;
            inputs["instanceInterruptionBehaviour"] = state ? state.instanceInterruptionBehaviour : undefined;
            inputs["instanceState"] = state ? state.instanceState : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["ipv6AddressCount"] = state ? state.ipv6AddressCount : undefined;
            inputs["ipv6Addresses"] = state ? state.ipv6Addresses : undefined;
            inputs["keyName"] = state ? state.keyName : undefined;
            inputs["launchGroup"] = state ? state.launchGroup : undefined;
            inputs["metadataOptions"] = state ? state.metadataOptions : undefined;
            inputs["monitoring"] = state ? state.monitoring : undefined;
            inputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            inputs["outpostArn"] = state ? state.outpostArn : undefined;
            inputs["passwordData"] = state ? state.passwordData : undefined;
            inputs["placementGroup"] = state ? state.placementGroup : undefined;
            inputs["primaryNetworkInterfaceId"] = state ? state.primaryNetworkInterfaceId : undefined;
            inputs["privateDns"] = state ? state.privateDns : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["publicDns"] = state ? state.publicDns : undefined;
            inputs["publicIp"] = state ? state.publicIp : undefined;
            inputs["rootBlockDevice"] = state ? state.rootBlockDevice : undefined;
            inputs["secondaryPrivateIps"] = state ? state.secondaryPrivateIps : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["sourceDestCheck"] = state ? state.sourceDestCheck : undefined;
            inputs["spotBidStatus"] = state ? state.spotBidStatus : undefined;
            inputs["spotInstanceId"] = state ? state.spotInstanceId : undefined;
            inputs["spotPrice"] = state ? state.spotPrice : undefined;
            inputs["spotRequestState"] = state ? state.spotRequestState : undefined;
            inputs["spotType"] = state ? state.spotType : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tenancy"] = state ? state.tenancy : undefined;
            inputs["userData"] = state ? state.userData : undefined;
            inputs["userDataBase64"] = state ? state.userDataBase64 : undefined;
            inputs["validFrom"] = state ? state.validFrom : undefined;
            inputs["validUntil"] = state ? state.validUntil : undefined;
            inputs["volumeTags"] = state ? state.volumeTags : undefined;
            inputs["vpcSecurityGroupIds"] = state ? state.vpcSecurityGroupIds : undefined;
            inputs["waitForFulfillment"] = state ? state.waitForFulfillment : undefined;
        } else {
            const args = argsOrState as SpotInstanceRequestArgs | undefined;
            if (!args || args.ami === undefined) {
                throw new Error("Missing required property 'ami'");
            }
            if (!args || args.instanceType === undefined) {
                throw new Error("Missing required property 'instanceType'");
            }
            inputs["ami"] = args ? args.ami : undefined;
            inputs["associatePublicIpAddress"] = args ? args.associatePublicIpAddress : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["blockDurationMinutes"] = args ? args.blockDurationMinutes : undefined;
            inputs["cpuCoreCount"] = args ? args.cpuCoreCount : undefined;
            inputs["cpuThreadsPerCore"] = args ? args.cpuThreadsPerCore : undefined;
            inputs["creditSpecification"] = args ? args.creditSpecification : undefined;
            inputs["disableApiTermination"] = args ? args.disableApiTermination : undefined;
            inputs["ebsBlockDevices"] = args ? args.ebsBlockDevices : undefined;
            inputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            inputs["ephemeralBlockDevices"] = args ? args.ephemeralBlockDevices : undefined;
            inputs["getPasswordData"] = args ? args.getPasswordData : undefined;
            inputs["hibernation"] = args ? args.hibernation : undefined;
            inputs["hostId"] = args ? args.hostId : undefined;
            inputs["iamInstanceProfile"] = args ? args.iamInstanceProfile : undefined;
            inputs["instanceInitiatedShutdownBehavior"] = args ? args.instanceInitiatedShutdownBehavior : undefined;
            inputs["instanceInterruptionBehaviour"] = args ? args.instanceInterruptionBehaviour : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["ipv6AddressCount"] = args ? args.ipv6AddressCount : undefined;
            inputs["ipv6Addresses"] = args ? args.ipv6Addresses : undefined;
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["launchGroup"] = args ? args.launchGroup : undefined;
            inputs["metadataOptions"] = args ? args.metadataOptions : undefined;
            inputs["monitoring"] = args ? args.monitoring : undefined;
            inputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            inputs["placementGroup"] = args ? args.placementGroup : undefined;
            inputs["privateIp"] = args ? args.privateIp : undefined;
            inputs["rootBlockDevice"] = args ? args.rootBlockDevice : undefined;
            inputs["secondaryPrivateIps"] = args ? args.secondaryPrivateIps : undefined;
            inputs["securityGroups"] = args ? args.securityGroups : undefined;
            inputs["sourceDestCheck"] = args ? args.sourceDestCheck : undefined;
            inputs["spotPrice"] = args ? args.spotPrice : undefined;
            inputs["spotType"] = args ? args.spotType : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenancy"] = args ? args.tenancy : undefined;
            inputs["userData"] = args ? args.userData : undefined;
            inputs["userDataBase64"] = args ? args.userDataBase64 : undefined;
            inputs["validFrom"] = args ? args.validFrom : undefined;
            inputs["validUntil"] = args ? args.validUntil : undefined;
            inputs["volumeTags"] = args ? args.volumeTags : undefined;
            inputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            inputs["waitForFulfillment"] = args ? args.waitForFulfillment : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["instanceState"] = undefined /*out*/;
            inputs["outpostArn"] = undefined /*out*/;
            inputs["passwordData"] = undefined /*out*/;
            inputs["primaryNetworkInterfaceId"] = undefined /*out*/;
            inputs["privateDns"] = undefined /*out*/;
            inputs["publicDns"] = undefined /*out*/;
            inputs["publicIp"] = undefined /*out*/;
            inputs["spotBidStatus"] = undefined /*out*/;
            inputs["spotInstanceId"] = undefined /*out*/;
            inputs["spotRequestState"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SpotInstanceRequest.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SpotInstanceRequest resources.
 */
export interface SpotInstanceRequestState {
    readonly ami?: pulumi.Input<string>;
    readonly arn?: pulumi.Input<string>;
    readonly associatePublicIpAddress?: pulumi.Input<boolean>;
    readonly availabilityZone?: pulumi.Input<string>;
    readonly blockDurationMinutes?: pulumi.Input<number>;
    readonly cpuCoreCount?: pulumi.Input<number>;
    readonly cpuThreadsPerCore?: pulumi.Input<number>;
    readonly creditSpecification?: pulumi.Input<inputs.ec2.SpotInstanceRequestCreditSpecification>;
    readonly disableApiTermination?: pulumi.Input<boolean>;
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.SpotInstanceRequestEbsBlockDevice>[]>;
    readonly ebsOptimized?: pulumi.Input<boolean>;
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.SpotInstanceRequestEphemeralBlockDevice>[]>;
    readonly getPasswordData?: pulumi.Input<boolean>;
    readonly hibernation?: pulumi.Input<boolean>;
    readonly hostId?: pulumi.Input<string>;
    readonly iamInstanceProfile?: pulumi.Input<string>;
    readonly instanceInitiatedShutdownBehavior?: pulumi.Input<string>;
    readonly instanceInterruptionBehaviour?: pulumi.Input<string>;
    readonly instanceState?: pulumi.Input<string>;
    readonly instanceType?: pulumi.Input<string>;
    readonly ipv6AddressCount?: pulumi.Input<number>;
    readonly ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly keyName?: pulumi.Input<string>;
    readonly launchGroup?: pulumi.Input<string>;
    readonly metadataOptions?: pulumi.Input<inputs.ec2.SpotInstanceRequestMetadataOptions>;
    readonly monitoring?: pulumi.Input<boolean>;
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.ec2.SpotInstanceRequestNetworkInterface>[]>;
    readonly outpostArn?: pulumi.Input<string>;
    readonly passwordData?: pulumi.Input<string>;
    readonly placementGroup?: pulumi.Input<string>;
    readonly primaryNetworkInterfaceId?: pulumi.Input<string>;
    readonly privateDns?: pulumi.Input<string>;
    readonly privateIp?: pulumi.Input<string>;
    readonly publicDns?: pulumi.Input<string>;
    readonly publicIp?: pulumi.Input<string>;
    readonly rootBlockDevice?: pulumi.Input<inputs.ec2.SpotInstanceRequestRootBlockDevice>;
    readonly secondaryPrivateIps?: pulumi.Input<pulumi.Input<string>[]>;
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    readonly sourceDestCheck?: pulumi.Input<boolean>;
    readonly spotBidStatus?: pulumi.Input<string>;
    readonly spotInstanceId?: pulumi.Input<string>;
    readonly spotPrice?: pulumi.Input<string>;
    readonly spotRequestState?: pulumi.Input<string>;
    readonly spotType?: pulumi.Input<string>;
    readonly subnetId?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly tenancy?: pulumi.Input<string>;
    readonly userData?: pulumi.Input<string>;
    readonly userDataBase64?: pulumi.Input<string>;
    readonly validFrom?: pulumi.Input<string>;
    readonly validUntil?: pulumi.Input<string>;
    readonly volumeTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly waitForFulfillment?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a SpotInstanceRequest resource.
 */
export interface SpotInstanceRequestArgs {
    readonly ami: pulumi.Input<string>;
    readonly associatePublicIpAddress?: pulumi.Input<boolean>;
    readonly availabilityZone?: pulumi.Input<string>;
    readonly blockDurationMinutes?: pulumi.Input<number>;
    readonly cpuCoreCount?: pulumi.Input<number>;
    readonly cpuThreadsPerCore?: pulumi.Input<number>;
    readonly creditSpecification?: pulumi.Input<inputs.ec2.SpotInstanceRequestCreditSpecification>;
    readonly disableApiTermination?: pulumi.Input<boolean>;
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.SpotInstanceRequestEbsBlockDevice>[]>;
    readonly ebsOptimized?: pulumi.Input<boolean>;
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.SpotInstanceRequestEphemeralBlockDevice>[]>;
    readonly getPasswordData?: pulumi.Input<boolean>;
    readonly hibernation?: pulumi.Input<boolean>;
    readonly hostId?: pulumi.Input<string>;
    readonly iamInstanceProfile?: pulumi.Input<string>;
    readonly instanceInitiatedShutdownBehavior?: pulumi.Input<string>;
    readonly instanceInterruptionBehaviour?: pulumi.Input<string>;
    readonly instanceType: pulumi.Input<string>;
    readonly ipv6AddressCount?: pulumi.Input<number>;
    readonly ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly keyName?: pulumi.Input<string>;
    readonly launchGroup?: pulumi.Input<string>;
    readonly metadataOptions?: pulumi.Input<inputs.ec2.SpotInstanceRequestMetadataOptions>;
    readonly monitoring?: pulumi.Input<boolean>;
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.ec2.SpotInstanceRequestNetworkInterface>[]>;
    readonly placementGroup?: pulumi.Input<string>;
    readonly privateIp?: pulumi.Input<string>;
    readonly rootBlockDevice?: pulumi.Input<inputs.ec2.SpotInstanceRequestRootBlockDevice>;
    readonly secondaryPrivateIps?: pulumi.Input<pulumi.Input<string>[]>;
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    readonly sourceDestCheck?: pulumi.Input<boolean>;
    readonly spotPrice?: pulumi.Input<string>;
    readonly spotType?: pulumi.Input<string>;
    readonly subnetId?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly tenancy?: pulumi.Input<string>;
    readonly userData?: pulumi.Input<string>;
    readonly userDataBase64?: pulumi.Input<string>;
    readonly validFrom?: pulumi.Input<string>;
    readonly validUntil?: pulumi.Input<string>;
    readonly volumeTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly waitForFulfillment?: pulumi.Input<boolean>;
}
