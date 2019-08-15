// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a CodePipeline.
 * 
 * > **NOTE on `aws.codepipeline.Pipeline`:** - the `GITHUB_TOKEN` environment variable must be set if the GitHub provider is specified.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const codepipelineRole = new aws.iam.Role("codepipelineRole", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Effect": "Allow",
 *       "Principal": {
 *         "Service": "codepipeline.amazonaws.com"
 *       },
 *       "Action": "sts:AssumeRole"
 *     }
 *   ]
 * }
 * `,
 * });
 * const codepipelineBucket = new aws.s3.Bucket("codepipelineBucket", {
 *     acl: "private",
 * });
 * const s3kmskey = aws.kms.getAlias({
 *     name: "alias/myKmsKey",
 * });
 * const codepipeline = new aws.codepipeline.Pipeline("codepipeline", {
 *     artifactStore: {
 *         encryptionKey: {
 *             id: s3kmskey.arn,
 *             type: "KMS",
 *         },
 *         location: codepipelineBucket.bucket,
 *         type: "S3",
 *     },
 *     roleArn: codepipelineRole.arn,
 *     stages: [
 *         {
 *             actions: [{
 *                 category: "Source",
 *                 configuration: {
 *                     Branch: "master",
 *                     Owner: "my-organization",
 *                     Repo: "test",
 *                 },
 *                 name: "Source",
 *                 outputArtifacts: ["sourceOutput"],
 *                 owner: "ThirdParty",
 *                 provider: "GitHub",
 *                 version: "1",
 *             }],
 *             name: "Source",
 *         },
 *         {
 *             actions: [{
 *                 category: "Build",
 *                 configuration: {
 *                     ProjectName: "test",
 *                 },
 *                 inputArtifacts: ["sourceOutput"],
 *                 name: "Build",
 *                 outputArtifacts: ["buildOutput"],
 *                 owner: "AWS",
 *                 provider: "CodeBuild",
 *                 version: "1",
 *             }],
 *             name: "Build",
 *         },
 *         {
 *             actions: [{
 *                 category: "Deploy",
 *                 configuration: {
 *                     ActionMode: "REPLACE_ON_FAILURE",
 *                     Capabilities: "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM",
 *                     OutputFileName: "CreateStackOutput.json",
 *                     StackName: "MyStack",
 *                     TemplatePath: "build_output::sam-templated.yaml",
 *                 },
 *                 inputArtifacts: ["buildOutput"],
 *                 name: "Deploy",
 *                 owner: "AWS",
 *                 provider: "CloudFormation",
 *                 version: "1",
 *             }],
 *             name: "Deploy",
 *         },
 *     ],
 * });
 * const codepipelinePolicy = new aws.iam.RolePolicy("codepipelinePolicy", {
 *     policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Effect":"Allow",
 *       "Action": [
 *         "s3:GetObject",
 *         "s3:GetObjectVersion",
 *         "s3:GetBucketVersioning",
 *         "s3:PutObject"
 *       ],
 *       "Resource": [
 *         "${codepipelineBucket.arn}",
 *         "${codepipelineBucket.arn}/*"
 *       ]
 *     },
 *     {
 *       "Effect": "Allow",
 *       "Action": [
 *         "codebuild:BatchGetBuilds",
 *         "codebuild:StartBuild"
 *       ],
 *       "Resource": "*"
 *     }
 *   ]
 * }
 * `,
 *     role: codepipelineRole.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codepipeline.html.markdown.
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
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineState, opts?: pulumi.CustomResourceOptions): Pipeline {
        return new Pipeline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:codepipeline/pipeline:Pipeline';

    /**
     * Returns true if the given object is an instance of Pipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pipeline.__pulumiType;
    }

    /**
     * The codepipeline ARN.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * An artifactStore block. Artifact stores are documented below.
     * * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
     */
    public readonly artifactStore!: pulumi.Output<outputApi.codepipeline.PipelineArtifactStore>;
    /**
     * The name of the pipeline.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
     */
    public readonly roleArn!: pulumi.Output<string>;
    public readonly stages!: pulumi.Output<outputApi.codepipeline.PipelineStage[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineArgs | PipelineState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PipelineState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["artifactStore"] = state ? state.artifactStore : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["stages"] = state ? state.stages : undefined;
            inputs["tags"] = state ? state.tags : undefined;
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
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Pipeline.__pulumiType, name, inputs, opts);
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
     * An artifactStore block. Artifact stores are documented below.
     * * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
     */
    readonly artifactStore?: pulumi.Input<inputApi.codepipeline.PipelineArtifactStore>;
    /**
     * The name of the pipeline.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
     */
    readonly roleArn?: pulumi.Input<string>;
    readonly stages?: pulumi.Input<pulumi.Input<inputApi.codepipeline.PipelineStage>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Pipeline resource.
 */
export interface PipelineArgs {
    /**
     * An artifactStore block. Artifact stores are documented below.
     * * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
     */
    readonly artifactStore: pulumi.Input<inputApi.codepipeline.PipelineArtifactStore>;
    /**
     * The name of the pipeline.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
     */
    readonly roleArn: pulumi.Input<string>;
    readonly stages: pulumi.Input<pulumi.Input<inputApi.codepipeline.PipelineStage>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
