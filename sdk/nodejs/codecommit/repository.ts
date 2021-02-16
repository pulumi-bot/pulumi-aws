// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CodeCommit Repository Resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.codecommit.Repository("test", {
 *     description: "This is the Sample App Repository",
 *     repositoryName: "MyTestRepository",
 * });
 * ```
 *
 * ## Import
 *
 * Codecommit repository can be imported using repository name, e.g.
 *
 * ```sh
 *  $ pulumi import aws:codecommit/repository:Repository imported ExistingRepo
 * ```
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryState, opts?: pulumi.CustomResourceOptions): Repository {
        return new Repository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:codecommit/repository:Repository';

    /**
     * Returns true if the given object is an instance of Repository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repository.__pulumiType;
    }

    /**
     * The ARN of the repository
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The URL to use for cloning the repository over HTTPS.
     */
    public /*out*/ readonly cloneUrlHttp!: pulumi.Output<string>;
    /**
     * The URL to use for cloning the repository over SSH.
     */
    public /*out*/ readonly cloneUrlSsh!: pulumi.Output<string>;
    /**
     * The default branch of the repository. The branch specified here needs to exist.
     */
    public readonly defaultBranch!: pulumi.Output<string | undefined>;
    /**
     * The description of the repository. This needs to be less than 1000 characters
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the repository
     */
    public /*out*/ readonly repositoryId!: pulumi.Output<string>;
    /**
     * The name for the repository. This needs to be less than 100 characters.
     */
    public readonly repositoryName!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryArgs | RepositoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["cloneUrlHttp"] = state ? state.cloneUrlHttp : undefined;
            inputs["cloneUrlSsh"] = state ? state.cloneUrlSsh : undefined;
            inputs["defaultBranch"] = state ? state.defaultBranch : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["repositoryId"] = state ? state.repositoryId : undefined;
            inputs["repositoryName"] = state ? state.repositoryName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RepositoryArgs | undefined;
            if ((!args || args.repositoryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryName'");
            }
            inputs["defaultBranch"] = args ? args.defaultBranch : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["repositoryName"] = args ? args.repositoryName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["cloneUrlHttp"] = undefined /*out*/;
            inputs["cloneUrlSsh"] = undefined /*out*/;
            inputs["repositoryId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Repository.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Repository resources.
 */
export interface RepositoryState {
    /**
     * The ARN of the repository
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The URL to use for cloning the repository over HTTPS.
     */
    readonly cloneUrlHttp?: pulumi.Input<string>;
    /**
     * The URL to use for cloning the repository over SSH.
     */
    readonly cloneUrlSsh?: pulumi.Input<string>;
    /**
     * The default branch of the repository. The branch specified here needs to exist.
     */
    readonly defaultBranch?: pulumi.Input<string>;
    /**
     * The description of the repository. This needs to be less than 1000 characters
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID of the repository
     */
    readonly repositoryId?: pulumi.Input<string>;
    /**
     * The name for the repository. This needs to be less than 100 characters.
     */
    readonly repositoryName?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * The default branch of the repository. The branch specified here needs to exist.
     */
    readonly defaultBranch?: pulumi.Input<string>;
    /**
     * The description of the repository. This needs to be less than 1000 characters
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name for the repository. This needs to be less than 100 characters.
     */
    readonly repositoryName: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
