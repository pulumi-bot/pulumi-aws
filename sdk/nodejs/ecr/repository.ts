// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an Elastic Container Registry Repository.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const foo = new aws.ecr.Repository("foo", {
 *     imageScanningConfiguration: {
 *         scanOnPush: true,
 *     },
 *     imageTagMutability: "MUTABLE",
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ecr_repository.html.markdown.
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryState, opts?: pulumi.CustomResourceOptions): Repository {
        return new Repository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ecr/repository:Repository';

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
     * Full ARN of the repository.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
     */
    public readonly imageScanningConfiguration!: pulumi.Output<outputs.ecr.RepositoryImageScanningConfiguration | undefined>;
    /**
     * The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
     */
    public readonly imageTagMutability!: pulumi.Output<string | undefined>;
    /**
     * Name of the repository.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The registry ID where the repository was created.
     */
    public /*out*/ readonly registryId!: pulumi.Output<string>;
    /**
     * The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
     */
    public /*out*/ readonly repositoryUrl!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryArgs | RepositoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RepositoryState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["imageScanningConfiguration"] = state ? state.imageScanningConfiguration : undefined;
            inputs["imageTagMutability"] = state ? state.imageTagMutability : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["registryId"] = state ? state.registryId : undefined;
            inputs["repositoryUrl"] = state ? state.repositoryUrl : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RepositoryArgs | undefined;
            inputs["imageScanningConfiguration"] = args ? args.imageScanningConfiguration : undefined;
            inputs["imageTagMutability"] = args ? args.imageTagMutability : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["registryId"] = undefined /*out*/;
            inputs["repositoryUrl"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Repository.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Repository resources.
 */
export interface RepositoryState {
    /**
     * Full ARN of the repository.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
     */
    readonly imageScanningConfiguration?: pulumi.Input<inputs.ecr.RepositoryImageScanningConfiguration>;
    /**
     * The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
     */
    readonly imageTagMutability?: pulumi.Input<string>;
    /**
     * Name of the repository.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The registry ID where the repository was created.
     */
    readonly registryId?: pulumi.Input<string>;
    /**
     * The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
     */
    readonly repositoryUrl?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
     */
    readonly imageScanningConfiguration?: pulumi.Input<inputs.ecr.RepositoryImageScanningConfiguration>;
    /**
     * The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
     */
    readonly imageTagMutability?: pulumi.Input<string>;
    /**
     * Name of the repository.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
