// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a SageMaker model resource.
 * 
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * 
 * Basic usage:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const model = new aws.sagemaker.Model("m", {
 *     executionRoleArn: aws_iam_role_foo.arn,
 *     primaryContainer: {
 *         image: "174872318107.dkr.ecr.us-west-2.amazonaws.com/kmeans:1",
 *     },
 * });
 * const assumeRole = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [{
 *             identifiers: ["sagemaker.amazonaws.com"],
 *             type: "Service",
 *         }],
 *     }],
 * });
 * const role = new aws.iam.Role("r", {
 *     assumeRolePolicy: assumeRole.json,
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sagemaker_model.html.markdown.
 */
export class Model extends pulumi.CustomResource {
    /**
     * Get an existing Model resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ModelState, opts?: pulumi.CustomResourceOptions): Model {
        return new Model(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/model:Model';

    /**
     * Returns true if the given object is an instance of Model.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Model {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Model.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this model.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
     */
    public readonly containers!: pulumi.Output<outputs.sagemaker.ModelContainer[] | undefined>;
    /**
     * Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
     */
    public readonly enableNetworkIsolation!: pulumi.Output<boolean | undefined>;
    /**
     * A role that SageMaker can assume to access model artifacts and docker images for deployment.
     */
    public readonly executionRoleArn!: pulumi.Output<string>;
    /**
     * The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
     */
    public readonly primaryContainer!: pulumi.Output<outputs.sagemaker.ModelPrimaryContainer | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
     */
    public readonly vpcConfig!: pulumi.Output<outputs.sagemaker.ModelVpcConfig | undefined>;

    /**
     * Create a Model resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ModelArgs | ModelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ModelState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["containers"] = state ? state.containers : undefined;
            inputs["enableNetworkIsolation"] = state ? state.enableNetworkIsolation : undefined;
            inputs["executionRoleArn"] = state ? state.executionRoleArn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["primaryContainer"] = state ? state.primaryContainer : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcConfig"] = state ? state.vpcConfig : undefined;
        } else {
            const args = argsOrState as ModelArgs | undefined;
            if (!args || args.executionRoleArn === undefined) {
                throw new Error("Missing required property 'executionRoleArn'");
            }
            inputs["containers"] = args ? args.containers : undefined;
            inputs["enableNetworkIsolation"] = args ? args.enableNetworkIsolation : undefined;
            inputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["primaryContainer"] = args ? args.primaryContainer : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Model.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Model resources.
 */
export interface ModelState {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this model.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
     */
    readonly containers?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelContainer>[]>;
    /**
     * Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
     */
    readonly enableNetworkIsolation?: pulumi.Input<boolean>;
    /**
     * A role that SageMaker can assume to access model artifacts and docker images for deployment.
     */
    readonly executionRoleArn?: pulumi.Input<string>;
    /**
     * The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
     */
    readonly primaryContainer?: pulumi.Input<inputs.sagemaker.ModelPrimaryContainer>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
     */
    readonly vpcConfig?: pulumi.Input<inputs.sagemaker.ModelVpcConfig>;
}

/**
 * The set of arguments for constructing a Model resource.
 */
export interface ModelArgs {
    /**
     * Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
     */
    readonly containers?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelContainer>[]>;
    /**
     * Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
     */
    readonly enableNetworkIsolation?: pulumi.Input<boolean>;
    /**
     * A role that SageMaker can assume to access model artifacts and docker images for deployment.
     */
    readonly executionRoleArn: pulumi.Input<string>;
    /**
     * The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
     */
    readonly primaryContainer?: pulumi.Input<inputs.sagemaker.ModelPrimaryContainer>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
     */
    readonly vpcConfig?: pulumi.Input<inputs.sagemaker.ModelVpcConfig>;
}
