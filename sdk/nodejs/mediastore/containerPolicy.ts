// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a MediaStore Container Policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const currentRegion = pulumi.output(aws.getRegion({ async: true }));
 * const currentCallerIdentity = pulumi.output(aws.getCallerIdentity({ async: true }));
 * const exampleContainer = new aws.mediastore.Container("example", {});
 * const exampleContainerPolicy = new aws.mediastore.ContainerPolicy("example", {
 *     containerName: exampleContainer.name,
 *     policy: pulumi.interpolate`{
 * 	"Version": "2012-10-17",
 * 	"Statement": [{
 * 		"Sid": "MediaStoreFullAccess",
 * 		"Action": [ "mediastore:*" ],
 * 		"Principal": {"AWS" : "arn:aws:iam::${currentCallerIdentity.accountId}:root"},
 * 		"Effect": "Allow",
 * 		"Resource": "arn:aws:mediastore:${currentCallerIdentity.accountId}:${currentRegion.name!}:container/${exampleContainer.name}/*",
 * 		"Condition": {
 * 			"Bool": { "aws:SecureTransport": "true" }
 * 		}
 * 	}]
 * }
 * `,
 * });
 * ```
 */
export class ContainerPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ContainerPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerPolicyState, opts?: pulumi.CustomResourceOptions): ContainerPolicy {
        return new ContainerPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:mediastore/containerPolicy:ContainerPolicy';

    /**
     * Returns true if the given object is an instance of ContainerPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerPolicy.__pulumiType;
    }

    /**
     * The name of the container.
     */
    public readonly containerName!: pulumi.Output<string>;
    /**
     * The contents of the policy.
     */
    public readonly policy!: pulumi.Output<string>;

    /**
     * Create a ContainerPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerPolicyArgs | ContainerPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ContainerPolicyState | undefined;
            inputs["containerName"] = state ? state.containerName : undefined;
            inputs["policy"] = state ? state.policy : undefined;
        } else {
            const args = argsOrState as ContainerPolicyArgs | undefined;
            if (!args || args.containerName === undefined) {
                throw new Error("Missing required property 'containerName'");
            }
            if (!args || args.policy === undefined) {
                throw new Error("Missing required property 'policy'");
            }
            inputs["containerName"] = args ? args.containerName : undefined;
            inputs["policy"] = args ? args.policy : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ContainerPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerPolicy resources.
 */
export interface ContainerPolicyState {
    /**
     * The name of the container.
     */
    readonly containerName?: pulumi.Input<string>;
    /**
     * The contents of the policy.
     */
    readonly policy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContainerPolicy resource.
 */
export interface ContainerPolicyArgs {
    /**
     * The name of the container.
     */
    readonly containerName: pulumi.Input<string>;
    /**
     * The contents of the policy.
     */
    readonly policy: pulumi.Input<string>;
}
