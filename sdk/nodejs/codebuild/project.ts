// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a CodeBuild Project resource. See also the `aws.codebuild.Webhook` resource, which manages the webhook to the source (e.g. the "rebuild every time a code change is pushed" option in the CodeBuild web console).
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:codebuild/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * The ARN of the CodeBuild project.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Information about the project's build output artifacts. Artifact blocks are documented below.
     */
    public readonly artifacts!: pulumi.Output<outputs.codebuild.ProjectArtifacts>;
    /**
     * Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
     */
    public readonly badgeEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The URL of the build badge when `badgeEnabled` is enabled.
     */
    public /*out*/ readonly badgeUrl!: pulumi.Output<string>;
    /**
     * How long in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
     */
    public readonly buildTimeout!: pulumi.Output<number | undefined>;
    /**
     * Information about the cache storage for the project. Cache blocks are documented below.
     */
    public readonly cache!: pulumi.Output<outputs.codebuild.ProjectCache | undefined>;
    /**
     * A short description of the project.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
     */
    public readonly encryptionKey!: pulumi.Output<string>;
    /**
     * Information about the project's build environment. Environment blocks are documented below.
     */
    public readonly environment!: pulumi.Output<outputs.codebuild.ProjectEnvironment>;
    /**
     * Configuration for the builds to store log data to CloudWatch or S3.
     */
    public readonly logsConfig!: pulumi.Output<outputs.codebuild.ProjectLogsConfig | undefined>;
    /**
     * The projects name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * How long in minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
     */
    public readonly queuedTimeout!: pulumi.Output<number | undefined>;
    /**
     * A set of secondary artifacts to be used inside the build. Secondary artifacts blocks are documented below.
     */
    public readonly secondaryArtifacts!: pulumi.Output<outputs.codebuild.ProjectSecondaryArtifact[] | undefined>;
    /**
     * A set of secondary sources to be used inside the build. Secondary sources blocks are documented below.
     */
    public readonly secondarySources!: pulumi.Output<outputs.codebuild.ProjectSecondarySource[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
     */
    public readonly serviceRole!: pulumi.Output<string>;
    /**
     * Information about the project's input source code. Source blocks are documented below.
     */
    public readonly source!: pulumi.Output<outputs.codebuild.ProjectSource>;
    /**
     * A version of the build input to be built for this project. If not specified, the latest version is used.
     */
    public readonly sourceVersion!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Configuration for the builds to run inside a VPC. VPC config blocks are documented below.
     */
    public readonly vpcConfig!: pulumi.Output<outputs.codebuild.ProjectVpcConfig | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["artifacts"] = state ? state.artifacts : undefined;
            inputs["badgeEnabled"] = state ? state.badgeEnabled : undefined;
            inputs["badgeUrl"] = state ? state.badgeUrl : undefined;
            inputs["buildTimeout"] = state ? state.buildTimeout : undefined;
            inputs["cache"] = state ? state.cache : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encryptionKey"] = state ? state.encryptionKey : undefined;
            inputs["environment"] = state ? state.environment : undefined;
            inputs["logsConfig"] = state ? state.logsConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["queuedTimeout"] = state ? state.queuedTimeout : undefined;
            inputs["secondaryArtifacts"] = state ? state.secondaryArtifacts : undefined;
            inputs["secondarySources"] = state ? state.secondarySources : undefined;
            inputs["serviceRole"] = state ? state.serviceRole : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["sourceVersion"] = state ? state.sourceVersion : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcConfig"] = state ? state.vpcConfig : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if (!args || args.artifacts === undefined) {
                throw new Error("Missing required property 'artifacts'");
            }
            if (!args || args.environment === undefined) {
                throw new Error("Missing required property 'environment'");
            }
            if (!args || args.serviceRole === undefined) {
                throw new Error("Missing required property 'serviceRole'");
            }
            if (!args || args.source === undefined) {
                throw new Error("Missing required property 'source'");
            }
            inputs["artifacts"] = args ? args.artifacts : undefined;
            inputs["badgeEnabled"] = args ? args.badgeEnabled : undefined;
            inputs["buildTimeout"] = args ? args.buildTimeout : undefined;
            inputs["cache"] = args ? args.cache : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encryptionKey"] = args ? args.encryptionKey : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["logsConfig"] = args ? args.logsConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["queuedTimeout"] = args ? args.queuedTimeout : undefined;
            inputs["secondaryArtifacts"] = args ? args.secondaryArtifacts : undefined;
            inputs["secondarySources"] = args ? args.secondarySources : undefined;
            inputs["serviceRole"] = args ? args.serviceRole : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["sourceVersion"] = args ? args.sourceVersion : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["badgeUrl"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * The ARN of the CodeBuild project.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Information about the project's build output artifacts. Artifact blocks are documented below.
     */
    readonly artifacts?: pulumi.Input<inputs.codebuild.ProjectArtifacts>;
    /**
     * Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
     */
    readonly badgeEnabled?: pulumi.Input<boolean>;
    /**
     * The URL of the build badge when `badgeEnabled` is enabled.
     */
    readonly badgeUrl?: pulumi.Input<string>;
    /**
     * How long in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
     */
    readonly buildTimeout?: pulumi.Input<number>;
    /**
     * Information about the cache storage for the project. Cache blocks are documented below.
     */
    readonly cache?: pulumi.Input<inputs.codebuild.ProjectCache>;
    /**
     * A short description of the project.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
     */
    readonly encryptionKey?: pulumi.Input<string>;
    /**
     * Information about the project's build environment. Environment blocks are documented below.
     */
    readonly environment?: pulumi.Input<inputs.codebuild.ProjectEnvironment>;
    /**
     * Configuration for the builds to store log data to CloudWatch or S3.
     */
    readonly logsConfig?: pulumi.Input<inputs.codebuild.ProjectLogsConfig>;
    /**
     * The projects name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * How long in minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
     */
    readonly queuedTimeout?: pulumi.Input<number>;
    /**
     * A set of secondary artifacts to be used inside the build. Secondary artifacts blocks are documented below.
     */
    readonly secondaryArtifacts?: pulumi.Input<pulumi.Input<inputs.codebuild.ProjectSecondaryArtifact>[]>;
    /**
     * A set of secondary sources to be used inside the build. Secondary sources blocks are documented below.
     */
    readonly secondarySources?: pulumi.Input<pulumi.Input<inputs.codebuild.ProjectSecondarySource>[]>;
    /**
     * The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
     */
    readonly serviceRole?: pulumi.Input<string>;
    /**
     * Information about the project's input source code. Source blocks are documented below.
     */
    readonly source?: pulumi.Input<inputs.codebuild.ProjectSource>;
    /**
     * A version of the build input to be built for this project. If not specified, the latest version is used.
     */
    readonly sourceVersion?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration for the builds to run inside a VPC. VPC config blocks are documented below.
     */
    readonly vpcConfig?: pulumi.Input<inputs.codebuild.ProjectVpcConfig>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * Information about the project's build output artifacts. Artifact blocks are documented below.
     */
    readonly artifacts: pulumi.Input<inputs.codebuild.ProjectArtifacts>;
    /**
     * Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
     */
    readonly badgeEnabled?: pulumi.Input<boolean>;
    /**
     * How long in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
     */
    readonly buildTimeout?: pulumi.Input<number>;
    /**
     * Information about the cache storage for the project. Cache blocks are documented below.
     */
    readonly cache?: pulumi.Input<inputs.codebuild.ProjectCache>;
    /**
     * A short description of the project.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
     */
    readonly encryptionKey?: pulumi.Input<string>;
    /**
     * Information about the project's build environment. Environment blocks are documented below.
     */
    readonly environment: pulumi.Input<inputs.codebuild.ProjectEnvironment>;
    /**
     * Configuration for the builds to store log data to CloudWatch or S3.
     */
    readonly logsConfig?: pulumi.Input<inputs.codebuild.ProjectLogsConfig>;
    /**
     * The projects name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * How long in minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
     */
    readonly queuedTimeout?: pulumi.Input<number>;
    /**
     * A set of secondary artifacts to be used inside the build. Secondary artifacts blocks are documented below.
     */
    readonly secondaryArtifacts?: pulumi.Input<pulumi.Input<inputs.codebuild.ProjectSecondaryArtifact>[]>;
    /**
     * A set of secondary sources to be used inside the build. Secondary sources blocks are documented below.
     */
    readonly secondarySources?: pulumi.Input<pulumi.Input<inputs.codebuild.ProjectSecondarySource>[]>;
    /**
     * The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
     */
    readonly serviceRole: pulumi.Input<string>;
    /**
     * Information about the project's input source code. Source blocks are documented below.
     */
    readonly source: pulumi.Input<inputs.codebuild.ProjectSource>;
    /**
     * A version of the build input to be built for this project. If not specified, the latest version is used.
     */
    readonly sourceVersion?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration for the builds to run inside a VPC. VPC config blocks are documented below.
     */
    readonly vpcConfig?: pulumi.Input<inputs.codebuild.ProjectVpcConfig>;
}
