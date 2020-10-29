// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a CodeArtifact Repository Resource.
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
    public static readonly __pulumiType = 'aws:codeartifact/repository:Repository';

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
     * The account number of the AWS account that manages the repository.
     */
    public /*out*/ readonly administratorAccount!: pulumi.Output<string>;
    /**
     * The ARN of the repository.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of the repository.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The domain that contains the created repository.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The account number of the AWS account that owns the domain.
     */
    public readonly domainOwner!: pulumi.Output<string>;
    /**
     * An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
     */
    public readonly externalConnections!: pulumi.Output<outputs.codeartifact.RepositoryExternalConnections | undefined>;
    /**
     * The name of the repository to create.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
     */
    public readonly upstreams!: pulumi.Output<outputs.codeartifact.RepositoryUpstream[] | undefined>;

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
        if (opts && opts.id) {
            const state = argsOrState as RepositoryState | undefined;
            inputs["administratorAccount"] = state ? state.administratorAccount : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["domainOwner"] = state ? state.domainOwner : undefined;
            inputs["externalConnections"] = state ? state.externalConnections : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["upstreams"] = state ? state.upstreams : undefined;
        } else {
            const args = argsOrState as RepositoryArgs | undefined;
            if (!args || args.domain === undefined) {
                throw new Error("Missing required property 'domain'");
            }
            if (!args || args.repository === undefined) {
                throw new Error("Missing required property 'repository'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["domainOwner"] = args ? args.domainOwner : undefined;
            inputs["externalConnections"] = args ? args.externalConnections : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["upstreams"] = args ? args.upstreams : undefined;
            inputs["administratorAccount"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
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
     * The account number of the AWS account that manages the repository.
     */
    readonly administratorAccount?: pulumi.Input<string>;
    /**
     * The ARN of the repository.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The description of the repository.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The domain that contains the created repository.
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * The account number of the AWS account that owns the domain.
     */
    readonly domainOwner?: pulumi.Input<string>;
    /**
     * An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
     */
    readonly externalConnections?: pulumi.Input<inputs.codeartifact.RepositoryExternalConnections>;
    /**
     * The name of the repository to create.
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
     */
    readonly upstreams?: pulumi.Input<pulumi.Input<inputs.codeartifact.RepositoryUpstream>[]>;
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * The description of the repository.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The domain that contains the created repository.
     */
    readonly domain: pulumi.Input<string>;
    /**
     * The account number of the AWS account that owns the domain.
     */
    readonly domainOwner?: pulumi.Input<string>;
    /**
     * An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
     */
    readonly externalConnections?: pulumi.Input<inputs.codeartifact.RepositoryExternalConnections>;
    /**
     * The name of the repository to create.
     */
    readonly repository: pulumi.Input<string>;
    /**
     * A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
     */
    readonly upstreams?: pulumi.Input<pulumi.Input<inputs.codeartifact.RepositoryUpstream>[]>;
}
