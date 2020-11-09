// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information about a Launch Template.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const defaultLaunchTemplate = pulumi.output(aws.ec2.getLaunchTemplate({
 *     name: "my-launch-template",
 * }, { async: true }));
 * ```
 * ### Filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.ec2.getLaunchTemplate({
 *     filters: [{
 *         name: "launch-template-name",
 *         values: ["some-template"],
 *     }],
 * }, { async: true }));
 * ```
 */
export function getLaunchTemplate(args?: GetLaunchTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetLaunchTemplateResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getLaunchTemplate:getLaunchTemplate", {
        "filters": args.filters,
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getLaunchTemplate.
 */
export interface GetLaunchTemplateArgs {
    /**
     * Configuration block(s) for filtering. Detailed below.
     */
    readonly filters?: inputs.ec2.GetLaunchTemplateFilter[];
    /**
     * The name of the filter field. Valid values can be found in the [EC2 DescribeLaunchTemplates API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLaunchTemplates.html).
     */
    readonly name?: string;
    /**
     * A map of tags, each pair of which must exactly match a pair on the desired Launch Template.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLaunchTemplate.
 */
export interface GetLaunchTemplateResult {
    /**
     * Amazon Resource Name (ARN) of the launch template.
     */
    readonly arn: string;
    /**
     * Specify volumes to attach to the instance besides the volumes specified by the AMI.
     */
    readonly blockDeviceMappings: outputs.ec2.GetLaunchTemplateBlockDeviceMapping[];
    /**
     * Customize the credit specification of the instance. See Credit
     * Specification below for more details.
     */
    readonly creditSpecifications: outputs.ec2.GetLaunchTemplateCreditSpecification[];
    /**
     * The default version of the launch template.
     */
    readonly defaultVersion: number;
    /**
     * Description of the launch template.
     */
    readonly description: string;
    /**
     * If `true`, enables [EC2 Instance
     * Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
     */
    readonly disableApiTermination: boolean;
    /**
     * If `true`, the launched EC2 instance will be EBS-optimized.
     */
    readonly ebsOptimized: string;
    /**
     * The elastic GPU to attach to the instance. See Elastic GPU
     * below for more details.
     */
    readonly elasticGpuSpecifications: outputs.ec2.GetLaunchTemplateElasticGpuSpecification[];
    readonly filters?: outputs.ec2.GetLaunchTemplateFilter[];
    /**
     * The hibernation options for the instance.
     */
    readonly hibernationOptions: outputs.ec2.GetLaunchTemplateHibernationOption[];
    /**
     * The IAM Instance Profile to launch the instance with. See Instance Profile
     * below for more details.
     */
    readonly iamInstanceProfiles: outputs.ec2.GetLaunchTemplateIamInstanceProfile[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The AMI from which to launch the instance.
     */
    readonly imageId: string;
    /**
     * Shutdown behavior for the instance. Can be `stop` or `terminate`.
     * (Default: `stop`).
     */
    readonly instanceInitiatedShutdownBehavior: string;
    /**
     * The market (purchasing) option for the instance.
     * below for details.
     */
    readonly instanceMarketOptions: outputs.ec2.GetLaunchTemplateInstanceMarketOption[];
    /**
     * The type of the instance.
     */
    readonly instanceType: string;
    /**
     * The kernel ID.
     */
    readonly kernelId: string;
    /**
     * The key name to use for the instance.
     */
    readonly keyName: string;
    /**
     * The latest version of the launch template.
     */
    readonly latestVersion: number;
    /**
     * The metadata options for the instance.
     */
    readonly metadataOptions: outputs.ec2.GetLaunchTemplateMetadataOption[];
    /**
     * The monitoring option for the instance.
     */
    readonly monitorings: outputs.ec2.GetLaunchTemplateMonitoring[];
    readonly name?: string;
    /**
     * Customize network interfaces to be attached at instance boot time. See Network
     * Interfaces below for more details.
     */
    readonly networkInterfaces: outputs.ec2.GetLaunchTemplateNetworkInterface[];
    /**
     * The placement of the instance.
     */
    readonly placements: outputs.ec2.GetLaunchTemplatePlacement[];
    /**
     * The ID of the RAM disk.
     */
    readonly ramDiskId: string;
    /**
     * A list of security group names to associate with. If you are creating Instances in a VPC, use
     * `vpcSecurityGroupIds` instead.
     */
    readonly securityGroupNames: string[];
    /**
     * The tags to apply to the resources during launch.
     */
    readonly tagSpecifications: outputs.ec2.GetLaunchTemplateTagSpecification[];
    /**
     * (Optional) A map of tags to assign to the launch template.
     */
    readonly tags: {[key: string]: string};
    /**
     * The Base64-encoded user data to provide when launching the instance.
     */
    readonly userData: string;
    /**
     * A list of security group IDs to associate with.
     */
    readonly vpcSecurityGroupIds: string[];
}
