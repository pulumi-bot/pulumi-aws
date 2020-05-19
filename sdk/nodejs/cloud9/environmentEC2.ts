// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cloud9 EC2 Development Environment.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cloud9.EnvironmentEC2("example", {
 *     instanceType: "t2.micro",
 * });
 * ```
 */
export class EnvironmentEC2 extends pulumi.CustomResource {
    /**
     * Get an existing EnvironmentEC2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentEC2State, opts?: pulumi.CustomResourceOptions): EnvironmentEC2 {
        return new EnvironmentEC2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloud9/environmentEC2:EnvironmentEC2';

    /**
     * Returns true if the given object is an instance of EnvironmentEC2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnvironmentEC2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnvironmentEC2.__pulumiType;
    }

    /**
     * The ARN of the environment.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The number of minutes until the running instance is shut down after the environment has last been used.
     */
    public readonly automaticStopTimeMinutes!: pulumi.Output<number | undefined>;
    /**
     * The description of the environment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The type of instance to connect to the environment, e.g. `t2.micro`.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The name of the environment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
     */
    public readonly ownerArn!: pulumi.Output<string>;
    /**
     * The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The type of the environment (e.g. `ssh` or `ec2`)
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a EnvironmentEC2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentEC2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentEC2Args | EnvironmentEC2State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EnvironmentEC2State | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["automaticStopTimeMinutes"] = state ? state.automaticStopTimeMinutes : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ownerArn"] = state ? state.ownerArn : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as EnvironmentEC2Args | undefined;
            if (!args || args.instanceType === undefined) {
                throw new Error("Missing required property 'instanceType'");
            }
            inputs["automaticStopTimeMinutes"] = args ? args.automaticStopTimeMinutes : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ownerArn"] = args ? args.ownerArn : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EnvironmentEC2.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EnvironmentEC2 resources.
 */
export interface EnvironmentEC2State {
    /**
     * The ARN of the environment.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The number of minutes until the running instance is shut down after the environment has last been used.
     */
    readonly automaticStopTimeMinutes?: pulumi.Input<number>;
    /**
     * The description of the environment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The type of instance to connect to the environment, e.g. `t2.micro`.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * The name of the environment.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
     */
    readonly ownerArn?: pulumi.Input<string>;
    /**
     * The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The type of the environment (e.g. `ssh` or `ec2`)
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EnvironmentEC2 resource.
 */
export interface EnvironmentEC2Args {
    /**
     * The number of minutes until the running instance is shut down after the environment has last been used.
     */
    readonly automaticStopTimeMinutes?: pulumi.Input<number>;
    /**
     * The description of the environment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The type of instance to connect to the environment, e.g. `t2.micro`.
     */
    readonly instanceType: pulumi.Input<string>;
    /**
     * The name of the environment.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
     */
    readonly ownerArn?: pulumi.Input<string>;
    /**
     * The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
