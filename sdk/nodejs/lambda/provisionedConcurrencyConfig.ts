// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Lambda Provisioned Concurrency Configuration.
 *
 * ## Example Usage
 * ### Alias Name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.ProvisionedConcurrencyConfig("example", {
 *     functionName: aws_lambda_alias.example.function_name,
 *     provisionedConcurrentExecutions: 1,
 *     qualifier: aws_lambda_alias.example.name,
 * });
 * ```
 * ### Function Version
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.ProvisionedConcurrencyConfig("example", {
 *     functionName: aws_lambda_function.example.function_name,
 *     provisionedConcurrentExecutions: 1,
 *     qualifier: aws_lambda_function.example.version,
 * });
 * ```
 *
 * ## Import
 *
 * Lambda Provisioned Concurrency Configs can be imported using the `function_name` and `qualifier` separated by a colon (`:`), e.g.
 *
 * ```sh
 *  $ pulumi import aws:lambda/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig example my_function:production
 * ```
 */
export class ProvisionedConcurrencyConfig extends pulumi.CustomResource {
    /**
     * Get an existing ProvisionedConcurrencyConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProvisionedConcurrencyConfigState, opts?: pulumi.CustomResourceOptions): ProvisionedConcurrencyConfig {
        return new ProvisionedConcurrencyConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lambda/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig';

    /**
     * Returns true if the given object is an instance of ProvisionedConcurrencyConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProvisionedConcurrencyConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProvisionedConcurrencyConfig.__pulumiType;
    }

    /**
     * Name or Amazon Resource Name (ARN) of the Lambda Function.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * Amount of capacity to allocate. Must be greater than or equal to `1`.
     */
    public readonly provisionedConcurrentExecutions!: pulumi.Output<number>;
    /**
     * Lambda Function version or Lambda Alias name.
     */
    public readonly qualifier!: pulumi.Output<string>;

    /**
     * Create a ProvisionedConcurrencyConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProvisionedConcurrencyConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProvisionedConcurrencyConfigArgs | ProvisionedConcurrencyConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProvisionedConcurrencyConfigState | undefined;
            inputs["functionName"] = state ? state.functionName : undefined;
            inputs["provisionedConcurrentExecutions"] = state ? state.provisionedConcurrentExecutions : undefined;
            inputs["qualifier"] = state ? state.qualifier : undefined;
        } else {
            const args = argsOrState as ProvisionedConcurrencyConfigArgs | undefined;
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            if ((!args || args.provisionedConcurrentExecutions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'provisionedConcurrentExecutions'");
            }
            if ((!args || args.qualifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'qualifier'");
            }
            inputs["functionName"] = args ? args.functionName : undefined;
            inputs["provisionedConcurrentExecutions"] = args ? args.provisionedConcurrentExecutions : undefined;
            inputs["qualifier"] = args ? args.qualifier : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ProvisionedConcurrencyConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProvisionedConcurrencyConfig resources.
 */
export interface ProvisionedConcurrencyConfigState {
    /**
     * Name or Amazon Resource Name (ARN) of the Lambda Function.
     */
    functionName?: pulumi.Input<string>;
    /**
     * Amount of capacity to allocate. Must be greater than or equal to `1`.
     */
    provisionedConcurrentExecutions?: pulumi.Input<number>;
    /**
     * Lambda Function version or Lambda Alias name.
     */
    qualifier?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProvisionedConcurrencyConfig resource.
 */
export interface ProvisionedConcurrencyConfigArgs {
    /**
     * Name or Amazon Resource Name (ARN) of the Lambda Function.
     */
    functionName: pulumi.Input<string>;
    /**
     * Amount of capacity to allocate. Must be greater than or equal to `1`.
     */
    provisionedConcurrentExecutions: pulumi.Input<number>;
    /**
     * Lambda Function version or Lambda Alias name.
     */
    qualifier: pulumi.Input<string>;
}
