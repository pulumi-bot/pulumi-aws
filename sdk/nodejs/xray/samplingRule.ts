// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages an AWS XRay Sampling Rule.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.xray.SamplingRule("example", {
 *     attributes: {
 *         Hello: "Tris",
 *     },
 *     fixedRate: 0.05,
 *     host: "*",
 *     httpMethod: "*",
 *     priority: 10000,
 *     reservoirSize: 1,
 *     resourceArn: "*",
 *     ruleName: "example",
 *     serviceName: "*",
 *     serviceType: "*",
 *     urlPath: "*",
 *     version: 1,
 * });
 * ```
 */
export class SamplingRule extends pulumi.CustomResource {
    /**
     * Get an existing SamplingRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SamplingRuleState, opts?: pulumi.CustomResourceOptions): SamplingRule {
        return new SamplingRule(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'aws:xray/samplingRule:SamplingRule';

    /**
     * Returns true if the given object is an instance of SamplingRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SamplingRule {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === SamplingRule.__pulumiType;
    }

    /**
     * The ARN of the sampling rule.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Matches attributes derived from the request.
     */
    public readonly attributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The percentage of matching requests to instrument, after the reservoir is exhausted.
     */
    public readonly fixedRate!: pulumi.Output<number>;
    /**
     * Matches the hostname from a request URL.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * Matches the HTTP method of a request.
     */
    public readonly httpMethod!: pulumi.Output<string>;
    /**
     * The priority of the sampling rule.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
     */
    public readonly reservoirSize!: pulumi.Output<number>;
    /**
     * Matches the ARN of the AWS resource on which the service runs.
     */
    public readonly resourceArn!: pulumi.Output<string>;
    /**
     * The name of the sampling rule.
     */
    public readonly ruleName!: pulumi.Output<string | undefined>;
    /**
     * Matches the `name` that the service uses to identify itself in segments.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Matches the `origin` that the service uses to identify its type in segments.
     */
    public readonly serviceType!: pulumi.Output<string>;
    /**
     * Matches the path from a request URL.
     */
    public readonly urlPath!: pulumi.Output<string>;
    /**
     * The version of the sampling rule format (`1` )
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a SamplingRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SamplingRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SamplingRuleArgs | SamplingRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SamplingRuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["attributes"] = state ? state.attributes : undefined;
            inputs["fixedRate"] = state ? state.fixedRate : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["httpMethod"] = state ? state.httpMethod : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["reservoirSize"] = state ? state.reservoirSize : undefined;
            inputs["resourceArn"] = state ? state.resourceArn : undefined;
            inputs["ruleName"] = state ? state.ruleName : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["serviceType"] = state ? state.serviceType : undefined;
            inputs["urlPath"] = state ? state.urlPath : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as SamplingRuleArgs | undefined;
            if (!args || args.fixedRate === undefined) {
                throw new Error("Missing required property 'fixedRate'");
            }
            if (!args || args.host === undefined) {
                throw new Error("Missing required property 'host'");
            }
            if (!args || args.httpMethod === undefined) {
                throw new Error("Missing required property 'httpMethod'");
            }
            if (!args || args.priority === undefined) {
                throw new Error("Missing required property 'priority'");
            }
            if (!args || args.reservoirSize === undefined) {
                throw new Error("Missing required property 'reservoirSize'");
            }
            if (!args || args.resourceArn === undefined) {
                throw new Error("Missing required property 'resourceArn'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.serviceType === undefined) {
                throw new Error("Missing required property 'serviceType'");
            }
            if (!args || args.urlPath === undefined) {
                throw new Error("Missing required property 'urlPath'");
            }
            if (!args || args.version === undefined) {
                throw new Error("Missing required property 'version'");
            }
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["fixedRate"] = args ? args.fixedRate : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["httpMethod"] = args ? args.httpMethod : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["reservoirSize"] = args ? args.reservoirSize : undefined;
            inputs["resourceArn"] = args ? args.resourceArn : undefined;
            inputs["ruleName"] = args ? args.ruleName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["serviceType"] = args ? args.serviceType : undefined;
            inputs["urlPath"] = args ? args.urlPath : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super(SamplingRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SamplingRule resources.
 */
export interface SamplingRuleState {
    /**
     * The ARN of the sampling rule.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Matches attributes derived from the request.
     */
    readonly attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The percentage of matching requests to instrument, after the reservoir is exhausted.
     */
    readonly fixedRate?: pulumi.Input<number>;
    /**
     * Matches the hostname from a request URL.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * Matches the HTTP method of a request.
     */
    readonly httpMethod?: pulumi.Input<string>;
    /**
     * The priority of the sampling rule.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
     */
    readonly reservoirSize?: pulumi.Input<number>;
    /**
     * Matches the ARN of the AWS resource on which the service runs.
     */
    readonly resourceArn?: pulumi.Input<string>;
    /**
     * The name of the sampling rule.
     */
    readonly ruleName?: pulumi.Input<string>;
    /**
     * Matches the `name` that the service uses to identify itself in segments.
     */
    readonly serviceName?: pulumi.Input<string>;
    /**
     * Matches the `origin` that the service uses to identify its type in segments.
     */
    readonly serviceType?: pulumi.Input<string>;
    /**
     * Matches the path from a request URL.
     */
    readonly urlPath?: pulumi.Input<string>;
    /**
     * The version of the sampling rule format (`1` )
     */
    readonly version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SamplingRule resource.
 */
export interface SamplingRuleArgs {
    /**
     * Matches attributes derived from the request.
     */
    readonly attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The percentage of matching requests to instrument, after the reservoir is exhausted.
     */
    readonly fixedRate: pulumi.Input<number>;
    /**
     * Matches the hostname from a request URL.
     */
    readonly host: pulumi.Input<string>;
    /**
     * Matches the HTTP method of a request.
     */
    readonly httpMethod: pulumi.Input<string>;
    /**
     * The priority of the sampling rule.
     */
    readonly priority: pulumi.Input<number>;
    /**
     * A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
     */
    readonly reservoirSize: pulumi.Input<number>;
    /**
     * Matches the ARN of the AWS resource on which the service runs.
     */
    readonly resourceArn: pulumi.Input<string>;
    /**
     * The name of the sampling rule.
     */
    readonly ruleName?: pulumi.Input<string>;
    /**
     * Matches the `name` that the service uses to identify itself in segments.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Matches the `origin` that the service uses to identify its type in segments.
     */
    readonly serviceType: pulumi.Input<string>;
    /**
     * Matches the path from a request URL.
     */
    readonly urlPath: pulumi.Input<string>;
    /**
     * The version of the sampling rule format (`1` )
     */
    readonly version: pulumi.Input<number>;
}
