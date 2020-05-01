// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Lightsail Instance. Amazon Lightsail is a service to provide easy virtual private servers
 * with custom software already setup. See [What is Amazon Lightsail?](https://lightsail.aws.amazon.com/ls/docs/getting-started/article/what-is-amazon-lightsail)
 * for more information.
 * 
 * > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * // Create a new GitLab Lightsail Instance
 * const gitlabTest = new aws.lightsail.Instance("gitlabTest", {
 *     availabilityZone: "us-east-1b",
 *     blueprintId: "string",
 *     bundleId: "string",
 *     keyPairName: "someKeyName",
 *     name: "customGitlab",
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * ```
 * 
 * ## Availability Zones
 * 
 * Lightsail currently supports the following Availability Zones (e.g. `us-east-1a`):
 * 
 * - `ap-northeast-1{a,c,d}`
 * - `ap-northeast-2{a,c}`
 * - `ap-south-1{a,b}`
 * - `ap-southeast-1{a,b,c}`
 * - `ap-southeast-2{a,b,c}`
 * - `ca-central-1{a,b}`
 * - `eu-central-1{a,b,c}`
 * - `eu-west-1{a,b,c}`
 * - `eu-west-2{a,b,c}`
 * - `eu-west-3{a,b,c}`
 * - `us-east-1{a,b,c,d,e,f}`
 * - `us-east-2{a,b,c}`
 * - `us-west-2{a,b,c}`
 * 
 * ## Blueprints
 * 
 * Lightsail currently supports the following Blueprint IDs:
 * 
 * ### OS Only
 * 
 * - `amazonLinux20180302`
 * - `centos7190101`
 * - `debian87`
 * - `debian95`
 * - `freebsd111`
 * - `opensuse422`
 * - `ubuntu16042`
 * - `ubuntu1804`
 * 
 * ### Apps and OS
 * 
 * - `drupal856`
 * - `gitlab11141`
 * - `joomla3811`
 * - `lamp56372`
 * - `lamp71201`
 * - `magento225`
 * - `mean401`
 * - `nginx11401`
 * - `nodejs1080`
 * - `pleskUbuntu178111`
 * - `redmine346`
 * - `wordpress498`
 * - `wordpressMultisite498`
 * 
 * ## Bundles
 * 
 * Lightsail currently supports the following Bundle IDs (e.g. an instance in `ap-northeast-1` would use `small20`):
 * 
 * ### Prefix
 * 
 * A Bundle ID starts with one of the below size prefixes:
 * 
 * - `nano_`
 * - `micro_`
 * - `small_`
 * - `medium_`
 * - `large_`
 * - `xlarge_`
 * - `2xlarge_`
 * 
 * ### Suffix
 * 
 * A Bundle ID ends with one of the following suffixes depending on Availability Zone:
 * 
 * - ap-northeast-1: `20`
 * - ap-northeast-2: `20`
 * - ap-south-1: `21`
 * - ap-southeast-1: `20`
 * - ap-southeast-2: `22`
 * - ca-central-1: `20`
 * - eu-central-1: `20`
 * - eu-west-1: `20`
 * - eu-west-2: `20`
 * - eu-west-3: `20`
 * - us-east-1: `20`
 * - us-east-2: `20`
 * - us-west-2: `20`
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lightsail_instance.html.markdown.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lightsail/instance:Instance';

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

    /**
     * The ARN of the Lightsail instance (matches `id`).
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Availability Zone in which to create your
     * instance (see list below)
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The ID for a virtual private server image
     * (see list below)
     */
    public readonly blueprintId!: pulumi.Output<string>;
    /**
     * The bundle of specification information (see list below)
     */
    public readonly bundleId!: pulumi.Output<string>;
    public /*out*/ readonly cpuCount!: pulumi.Output<number>;
    /**
     * The timestamp when the instance was created.
     * * `availabilityZone`
     * * `blueprintId`
     * * `bundleId`
     * * `keyPairName`
     * * `userData`
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public /*out*/ readonly ipv6Address!: pulumi.Output<string>;
    public /*out*/ readonly isStaticIp!: pulumi.Output<boolean>;
    /**
     * The name of your key pair. Created in the
     * Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
     */
    public readonly keyPairName!: pulumi.Output<string | undefined>;
    /**
     * The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly privateIpAddress!: pulumi.Output<string>;
    public /*out*/ readonly publicIpAddress!: pulumi.Output<string>;
    public /*out*/ readonly ramSize!: pulumi.Output<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * launch script to configure server with additional user data
     */
    public readonly userData!: pulumi.Output<string | undefined>;
    public /*out*/ readonly username!: pulumi.Output<string>;

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
            inputs["arn"] = state ? state.arn : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["blueprintId"] = state ? state.blueprintId : undefined;
            inputs["bundleId"] = state ? state.bundleId : undefined;
            inputs["cpuCount"] = state ? state.cpuCount : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["ipv6Address"] = state ? state.ipv6Address : undefined;
            inputs["isStaticIp"] = state ? state.isStaticIp : undefined;
            inputs["keyPairName"] = state ? state.keyPairName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            inputs["publicIpAddress"] = state ? state.publicIpAddress : undefined;
            inputs["ramSize"] = state ? state.ramSize : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["userData"] = state ? state.userData : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.availabilityZone === undefined) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if (!args || args.blueprintId === undefined) {
                throw new Error("Missing required property 'blueprintId'");
            }
            if (!args || args.bundleId === undefined) {
                throw new Error("Missing required property 'bundleId'");
            }
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["blueprintId"] = args ? args.blueprintId : undefined;
            inputs["bundleId"] = args ? args.bundleId : undefined;
            inputs["keyPairName"] = args ? args.keyPairName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userData"] = args ? args.userData : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["cpuCount"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["ipv6Address"] = undefined /*out*/;
            inputs["isStaticIp"] = undefined /*out*/;
            inputs["privateIpAddress"] = undefined /*out*/;
            inputs["publicIpAddress"] = undefined /*out*/;
            inputs["ramSize"] = undefined /*out*/;
            inputs["username"] = undefined /*out*/;
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
    /**
     * The ARN of the Lightsail instance (matches `id`).
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The Availability Zone in which to create your
     * instance (see list below)
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * The ID for a virtual private server image
     * (see list below)
     */
    readonly blueprintId?: pulumi.Input<string>;
    /**
     * The bundle of specification information (see list below)
     */
    readonly bundleId?: pulumi.Input<string>;
    readonly cpuCount?: pulumi.Input<number>;
    /**
     * The timestamp when the instance was created.
     * * `availabilityZone`
     * * `blueprintId`
     * * `bundleId`
     * * `keyPairName`
     * * `userData`
     */
    readonly createdAt?: pulumi.Input<string>;
    readonly ipv6Address?: pulumi.Input<string>;
    readonly isStaticIp?: pulumi.Input<boolean>;
    /**
     * The name of your key pair. Created in the
     * Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
     */
    readonly keyPairName?: pulumi.Input<string>;
    /**
     * The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
     */
    readonly name?: pulumi.Input<string>;
    readonly privateIpAddress?: pulumi.Input<string>;
    readonly publicIpAddress?: pulumi.Input<string>;
    readonly ramSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * launch script to configure server with additional user data
     */
    readonly userData?: pulumi.Input<string>;
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The Availability Zone in which to create your
     * instance (see list below)
     */
    readonly availabilityZone: pulumi.Input<string>;
    /**
     * The ID for a virtual private server image
     * (see list below)
     */
    readonly blueprintId: pulumi.Input<string>;
    /**
     * The bundle of specification information (see list below)
     */
    readonly bundleId: pulumi.Input<string>;
    /**
     * The name of your key pair. Created in the
     * Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
     */
    readonly keyPairName?: pulumi.Input<string>;
    /**
     * The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * launch script to configure server with additional user data
     */
    readonly userData?: pulumi.Input<string>;
}
