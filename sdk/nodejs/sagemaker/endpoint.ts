// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a SageMaker Endpoint resource.
 * 
 * ## Example Usage
 * 
 * Basic usage:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const endpoint = new aws.sagemaker.Endpoint("e", {
 *     endpointConfigName: aws_sagemaker_endpoint_configuration_ec.name,
 *     tags: {
 *         Name: "foo",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sagemaker_endpoint.html.markdown.
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointState, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/endpoint:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this endpoint.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the endpoint configuration to use.
     */
    public readonly endpointConfigName!: pulumi.Output<string>;
    /**
     * The name of the endpoint. If omitted, this provider will assign a random, unique name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointArgs | EndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EndpointState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["endpointConfigName"] = state ? state.endpointConfigName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EndpointArgs | undefined;
            if (!args || args.endpointConfigName === undefined) {
                throw new Error("Missing required property 'endpointConfigName'");
            }
            inputs["endpointConfigName"] = args ? args.endpointConfigName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Endpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Endpoint resources.
 */
export interface EndpointState {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this endpoint.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name of the endpoint configuration to use.
     */
    readonly endpointConfigName?: pulumi.Input<string>;
    /**
     * The name of the endpoint. If omitted, this provider will assign a random, unique name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * The name of the endpoint configuration to use.
     */
    readonly endpointConfigName: pulumi.Input<string>;
    /**
     * The name of the endpoint. If omitted, this provider will assign a random, unique name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
