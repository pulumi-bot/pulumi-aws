// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Elastic Container Registry Policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const currentCallerIdentity = aws.getCallerIdentity({});
 * const currentRegion = aws.getRegion({});
 * const currentPartition = aws.getPartition({});
 * const example = new aws.ecr.RegistryPolicy("example", {policy: Promise.all([currentPartition, currentCallerIdentity, currentPartition, currentRegion, currentCallerIdentity]).then(([currentPartition, currentCallerIdentity, currentPartition1, currentRegion, currentCallerIdentity1]) => JSON.stringify({
 *     Version: "2012-10-17",
 *     Statement: [{
 *         Sid: "testpolicy",
 *         Effect: "Allow",
 *         Principal: {
 *             AWS: `arn:${currentPartition.partition}:iam::${currentCallerIdentity.accountId}:root`,
 *         },
 *         Action: ["ecr:ReplicateImage"],
 *         Resource: [`arn:${currentPartition1.partition}:ecr:${currentRegion.name}:${currentCallerIdentity1.accountId}:repository/*`],
 *     }],
 * }))});
 * ```
 *
 * ## Import
 *
 * ECR Registry Policy can be imported using the registry id, e.g.
 *
 * ```sh
 *  $ pulumi import aws:ecr/registryPolicy:RegistryPolicy example 123456789012
 * ```
 */
export class RegistryPolicy extends pulumi.CustomResource {
    /**
     * Get an existing RegistryPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryPolicyState, opts?: pulumi.CustomResourceOptions): RegistryPolicy {
        return new RegistryPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ecr/registryPolicy:RegistryPolicy';

    /**
     * Returns true if the given object is an instance of RegistryPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistryPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistryPolicy.__pulumiType;
    }

    public readonly policy!: pulumi.Output<string>;
    /**
     * The registry ID where the registry was created.
     */
    public /*out*/ readonly registryId!: pulumi.Output<string>;

    /**
     * Create a RegistryPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryPolicyArgs | RegistryPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegistryPolicyState | undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["registryId"] = state ? state.registryId : undefined;
        } else {
            const args = argsOrState as RegistryPolicyArgs | undefined;
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            inputs["policy"] = args ? args.policy : undefined;
            inputs["registryId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegistryPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegistryPolicy resources.
 */
export interface RegistryPolicyState {
    policy?: pulumi.Input<string>;
    /**
     * The registry ID where the registry was created.
     */
    registryId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RegistryPolicy resource.
 */
export interface RegistryPolicyArgs {
    policy: pulumi.Input<string>;
}
