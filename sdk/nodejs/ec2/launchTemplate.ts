// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class LaunchTemplate extends pulumi.CustomResource {
    /**
     * Get an existing LaunchTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LaunchTemplateState, opts?: pulumi.CustomResourceOptions): LaunchTemplate {
        return new LaunchTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/launchTemplate:LaunchTemplate';

    /**
     * Returns true if the given object is an instance of LaunchTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LaunchTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LaunchTemplate.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly blockDeviceMappings!: pulumi.Output<outputs.ec2.LaunchTemplateBlockDeviceMapping[] | undefined>;
    public readonly capacityReservationSpecification!: pulumi.Output<outputs.ec2.LaunchTemplateCapacityReservationSpecification | undefined>;
    public readonly cpuOptions!: pulumi.Output<outputs.ec2.LaunchTemplateCpuOptions | undefined>;
    public readonly creditSpecification!: pulumi.Output<outputs.ec2.LaunchTemplateCreditSpecification | undefined>;
    public readonly defaultVersion!: pulumi.Output<number>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly disableApiTermination!: pulumi.Output<boolean | undefined>;
    public readonly ebsOptimized!: pulumi.Output<string | undefined>;
    public readonly elasticGpuSpecifications!: pulumi.Output<outputs.ec2.LaunchTemplateElasticGpuSpecification[] | undefined>;
    public readonly elasticInferenceAccelerator!: pulumi.Output<outputs.ec2.LaunchTemplateElasticInferenceAccelerator | undefined>;
    public readonly hibernationOptions!: pulumi.Output<outputs.ec2.LaunchTemplateHibernationOptions | undefined>;
    public readonly iamInstanceProfile!: pulumi.Output<outputs.ec2.LaunchTemplateIamInstanceProfile | undefined>;
    public readonly imageId!: pulumi.Output<string | undefined>;
    public readonly instanceInitiatedShutdownBehavior!: pulumi.Output<string | undefined>;
    public readonly instanceMarketOptions!: pulumi.Output<outputs.ec2.LaunchTemplateInstanceMarketOptions | undefined>;
    public readonly instanceType!: pulumi.Output<string | undefined>;
    public readonly kernelId!: pulumi.Output<string | undefined>;
    public readonly keyName!: pulumi.Output<string | undefined>;
    public /*out*/ readonly latestVersion!: pulumi.Output<number>;
    public readonly licenseSpecifications!: pulumi.Output<outputs.ec2.LaunchTemplateLicenseSpecification[] | undefined>;
    public readonly metadataOptions!: pulumi.Output<outputs.ec2.LaunchTemplateMetadataOptions>;
    public readonly monitoring!: pulumi.Output<outputs.ec2.LaunchTemplateMonitoring | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    public readonly networkInterfaces!: pulumi.Output<outputs.ec2.LaunchTemplateNetworkInterface[] | undefined>;
    public readonly placement!: pulumi.Output<outputs.ec2.LaunchTemplatePlacement | undefined>;
    public readonly ramDiskId!: pulumi.Output<string | undefined>;
    public readonly securityGroupNames!: pulumi.Output<string[] | undefined>;
    public readonly tagSpecifications!: pulumi.Output<outputs.ec2.LaunchTemplateTagSpecification[] | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly updateDefaultVersion!: pulumi.Output<boolean | undefined>;
    public readonly userData!: pulumi.Output<string | undefined>;
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a LaunchTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LaunchTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LaunchTemplateArgs | LaunchTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LaunchTemplateState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["blockDeviceMappings"] = state ? state.blockDeviceMappings : undefined;
            inputs["capacityReservationSpecification"] = state ? state.capacityReservationSpecification : undefined;
            inputs["cpuOptions"] = state ? state.cpuOptions : undefined;
            inputs["creditSpecification"] = state ? state.creditSpecification : undefined;
            inputs["defaultVersion"] = state ? state.defaultVersion : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disableApiTermination"] = state ? state.disableApiTermination : undefined;
            inputs["ebsOptimized"] = state ? state.ebsOptimized : undefined;
            inputs["elasticGpuSpecifications"] = state ? state.elasticGpuSpecifications : undefined;
            inputs["elasticInferenceAccelerator"] = state ? state.elasticInferenceAccelerator : undefined;
            inputs["hibernationOptions"] = state ? state.hibernationOptions : undefined;
            inputs["iamInstanceProfile"] = state ? state.iamInstanceProfile : undefined;
            inputs["imageId"] = state ? state.imageId : undefined;
            inputs["instanceInitiatedShutdownBehavior"] = state ? state.instanceInitiatedShutdownBehavior : undefined;
            inputs["instanceMarketOptions"] = state ? state.instanceMarketOptions : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["kernelId"] = state ? state.kernelId : undefined;
            inputs["keyName"] = state ? state.keyName : undefined;
            inputs["latestVersion"] = state ? state.latestVersion : undefined;
            inputs["licenseSpecifications"] = state ? state.licenseSpecifications : undefined;
            inputs["metadataOptions"] = state ? state.metadataOptions : undefined;
            inputs["monitoring"] = state ? state.monitoring : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            inputs["placement"] = state ? state.placement : undefined;
            inputs["ramDiskId"] = state ? state.ramDiskId : undefined;
            inputs["securityGroupNames"] = state ? state.securityGroupNames : undefined;
            inputs["tagSpecifications"] = state ? state.tagSpecifications : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["updateDefaultVersion"] = state ? state.updateDefaultVersion : undefined;
            inputs["userData"] = state ? state.userData : undefined;
            inputs["vpcSecurityGroupIds"] = state ? state.vpcSecurityGroupIds : undefined;
        } else {
            const args = argsOrState as LaunchTemplateArgs | undefined;
            inputs["blockDeviceMappings"] = args ? args.blockDeviceMappings : undefined;
            inputs["capacityReservationSpecification"] = args ? args.capacityReservationSpecification : undefined;
            inputs["cpuOptions"] = args ? args.cpuOptions : undefined;
            inputs["creditSpecification"] = args ? args.creditSpecification : undefined;
            inputs["defaultVersion"] = args ? args.defaultVersion : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disableApiTermination"] = args ? args.disableApiTermination : undefined;
            inputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            inputs["elasticGpuSpecifications"] = args ? args.elasticGpuSpecifications : undefined;
            inputs["elasticInferenceAccelerator"] = args ? args.elasticInferenceAccelerator : undefined;
            inputs["hibernationOptions"] = args ? args.hibernationOptions : undefined;
            inputs["iamInstanceProfile"] = args ? args.iamInstanceProfile : undefined;
            inputs["imageId"] = args ? args.imageId : undefined;
            inputs["instanceInitiatedShutdownBehavior"] = args ? args.instanceInitiatedShutdownBehavior : undefined;
            inputs["instanceMarketOptions"] = args ? args.instanceMarketOptions : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["kernelId"] = args ? args.kernelId : undefined;
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["licenseSpecifications"] = args ? args.licenseSpecifications : undefined;
            inputs["metadataOptions"] = args ? args.metadataOptions : undefined;
            inputs["monitoring"] = args ? args.monitoring : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            inputs["placement"] = args ? args.placement : undefined;
            inputs["ramDiskId"] = args ? args.ramDiskId : undefined;
            inputs["securityGroupNames"] = args ? args.securityGroupNames : undefined;
            inputs["tagSpecifications"] = args ? args.tagSpecifications : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["updateDefaultVersion"] = args ? args.updateDefaultVersion : undefined;
            inputs["userData"] = args ? args.userData : undefined;
            inputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["latestVersion"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LaunchTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LaunchTemplate resources.
 */
export interface LaunchTemplateState {
    readonly arn?: pulumi.Input<string>;
    readonly blockDeviceMappings?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchTemplateBlockDeviceMapping>[]>;
    readonly capacityReservationSpecification?: pulumi.Input<inputs.ec2.LaunchTemplateCapacityReservationSpecification>;
    readonly cpuOptions?: pulumi.Input<inputs.ec2.LaunchTemplateCpuOptions>;
    readonly creditSpecification?: pulumi.Input<inputs.ec2.LaunchTemplateCreditSpecification>;
    readonly defaultVersion?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly disableApiTermination?: pulumi.Input<boolean>;
    readonly ebsOptimized?: pulumi.Input<string>;
    readonly elasticGpuSpecifications?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchTemplateElasticGpuSpecification>[]>;
    readonly elasticInferenceAccelerator?: pulumi.Input<inputs.ec2.LaunchTemplateElasticInferenceAccelerator>;
    readonly hibernationOptions?: pulumi.Input<inputs.ec2.LaunchTemplateHibernationOptions>;
    readonly iamInstanceProfile?: pulumi.Input<inputs.ec2.LaunchTemplateIamInstanceProfile>;
    readonly imageId?: pulumi.Input<string>;
    readonly instanceInitiatedShutdownBehavior?: pulumi.Input<string>;
    readonly instanceMarketOptions?: pulumi.Input<inputs.ec2.LaunchTemplateInstanceMarketOptions>;
    readonly instanceType?: pulumi.Input<string>;
    readonly kernelId?: pulumi.Input<string>;
    readonly keyName?: pulumi.Input<string>;
    readonly latestVersion?: pulumi.Input<number>;
    readonly licenseSpecifications?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchTemplateLicenseSpecification>[]>;
    readonly metadataOptions?: pulumi.Input<inputs.ec2.LaunchTemplateMetadataOptions>;
    readonly monitoring?: pulumi.Input<inputs.ec2.LaunchTemplateMonitoring>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchTemplateNetworkInterface>[]>;
    readonly placement?: pulumi.Input<inputs.ec2.LaunchTemplatePlacement>;
    readonly ramDiskId?: pulumi.Input<string>;
    readonly securityGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tagSpecifications?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchTemplateTagSpecification>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly updateDefaultVersion?: pulumi.Input<boolean>;
    readonly userData?: pulumi.Input<string>;
    readonly vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a LaunchTemplate resource.
 */
export interface LaunchTemplateArgs {
    readonly blockDeviceMappings?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchTemplateBlockDeviceMapping>[]>;
    readonly capacityReservationSpecification?: pulumi.Input<inputs.ec2.LaunchTemplateCapacityReservationSpecification>;
    readonly cpuOptions?: pulumi.Input<inputs.ec2.LaunchTemplateCpuOptions>;
    readonly creditSpecification?: pulumi.Input<inputs.ec2.LaunchTemplateCreditSpecification>;
    readonly defaultVersion?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly disableApiTermination?: pulumi.Input<boolean>;
    readonly ebsOptimized?: pulumi.Input<string>;
    readonly elasticGpuSpecifications?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchTemplateElasticGpuSpecification>[]>;
    readonly elasticInferenceAccelerator?: pulumi.Input<inputs.ec2.LaunchTemplateElasticInferenceAccelerator>;
    readonly hibernationOptions?: pulumi.Input<inputs.ec2.LaunchTemplateHibernationOptions>;
    readonly iamInstanceProfile?: pulumi.Input<inputs.ec2.LaunchTemplateIamInstanceProfile>;
    readonly imageId?: pulumi.Input<string>;
    readonly instanceInitiatedShutdownBehavior?: pulumi.Input<string>;
    readonly instanceMarketOptions?: pulumi.Input<inputs.ec2.LaunchTemplateInstanceMarketOptions>;
    readonly instanceType?: pulumi.Input<string>;
    readonly kernelId?: pulumi.Input<string>;
    readonly keyName?: pulumi.Input<string>;
    readonly licenseSpecifications?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchTemplateLicenseSpecification>[]>;
    readonly metadataOptions?: pulumi.Input<inputs.ec2.LaunchTemplateMetadataOptions>;
    readonly monitoring?: pulumi.Input<inputs.ec2.LaunchTemplateMonitoring>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchTemplateNetworkInterface>[]>;
    readonly placement?: pulumi.Input<inputs.ec2.LaunchTemplatePlacement>;
    readonly ramDiskId?: pulumi.Input<string>;
    readonly securityGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tagSpecifications?: pulumi.Input<pulumi.Input<inputs.ec2.LaunchTemplateTagSpecification>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly updateDefaultVersion?: pulumi.Input<boolean>;
    readonly userData?: pulumi.Input<string>;
    readonly vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}
