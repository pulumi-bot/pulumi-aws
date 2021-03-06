// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Provides a CodePipeline.
 * 
 * ~> **NOTE on `aws_codepipeline`:** - the `GITHUB_TOKEN` environment variable must be set if the GitHub provider is specified.
 */
export class Pipeline extends pulumi.CustomResource {
    /**
     * Get an existing Pipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineState): Pipeline {
        return new Pipeline(name, <any>state, { id });
    }

    /**
     * The codepipeline ARN.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * An artifact_store block. Artifact stores are documented below.
     */
    public readonly artifactStore: pulumi.Output<{ encryptionKey?: { id: string, type: string }, location: string, type: string }>;
    /**
     * The action declaration's name.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
     */
    public readonly roleArn: pulumi.Output<string>;
    /**
     * A stage block. Stages are documented below.
     */
    public readonly stages: pulumi.Output<{ actions: { category: string, configuration?: {[key: string]: any}, inputArtifacts?: string[], name: string, outputArtifacts?: string[], owner: string, provider: string, roleArn?: string, runOrder: number, version: string }[], name: string }[]>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: PipelineArgs | PipelineState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: PipelineState = argsOrState as PipelineState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["artifactStore"] = state ? state.artifactStore : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["stages"] = state ? state.stages : undefined;
        } else {
            const args = argsOrState as PipelineArgs | undefined;
            if (!args || args.artifactStore === undefined) {
                throw new Error("Missing required property 'artifactStore'");
            }
            if (!args || args.roleArn === undefined) {
                throw new Error("Missing required property 'roleArn'");
            }
            if (!args || args.stages === undefined) {
                throw new Error("Missing required property 'stages'");
            }
            inputs["artifactStore"] = args ? args.artifactStore : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["stages"] = args ? args.stages : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:codepipeline/pipeline:Pipeline", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pipeline resources.
 */
export interface PipelineState {
    /**
     * The codepipeline ARN.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * An artifact_store block. Artifact stores are documented below.
     */
    readonly artifactStore?: pulumi.Input<{ encryptionKey?: pulumi.Input<{ id: pulumi.Input<string>, type: pulumi.Input<string> }>, location: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * The action declaration's name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * A stage block. Stages are documented below.
     */
    readonly stages?: pulumi.Input<{ actions: pulumi.Input<{ category: pulumi.Input<string>, configuration?: pulumi.Input<{[key: string]: any}>, inputArtifacts?: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string>, outputArtifacts?: pulumi.Input<pulumi.Input<string>[]>, owner: pulumi.Input<string>, provider: pulumi.Input<string>, roleArn?: pulumi.Input<string>, runOrder?: pulumi.Input<number>, version: pulumi.Input<string> }[]>, name: pulumi.Input<string> }[]>;
}

/**
 * The set of arguments for constructing a Pipeline resource.
 */
export interface PipelineArgs {
    /**
     * An artifact_store block. Artifact stores are documented below.
     */
    readonly artifactStore: pulumi.Input<{ encryptionKey?: pulumi.Input<{ id: pulumi.Input<string>, type: pulumi.Input<string> }>, location: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * The action declaration's name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * A stage block. Stages are documented below.
     */
    readonly stages: pulumi.Input<{ actions: pulumi.Input<{ category: pulumi.Input<string>, configuration?: pulumi.Input<{[key: string]: any}>, inputArtifacts?: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string>, outputArtifacts?: pulumi.Input<pulumi.Input<string>[]>, owner: pulumi.Input<string>, provider: pulumi.Input<string>, roleArn?: pulumi.Input<string>, runOrder?: pulumi.Input<number>, version: pulumi.Input<string> }[]>, name: pulumi.Input<string> }[]>;
}
