// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an API Gateway Usage Plan.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myapi = new aws.apigateway.RestApi("myapi", {});
 * // ...
 * const dev = new aws.apigateway.Deployment("dev", {
 *     restApi: myapi.id,
 *     stageName: "dev",
 * });
 * const prod = new aws.apigateway.Deployment("prod", {
 *     restApi: myapi.id,
 *     stageName: "prod",
 * });
 * const myUsagePlan = new aws.apigateway.UsagePlan("myUsagePlan", {
 *     description: "my description",
 *     productCode: "MYCODE",
 *     apiStages: [
 *         {
 *             apiId: myapi.id,
 *             stage: dev.stageName,
 *         },
 *         {
 *             apiId: myapi.id,
 *             stage: prod.stageName,
 *         },
 *     ],
 *     quotaSettings: {
 *         limit: 20,
 *         offset: 2,
 *         period: "WEEK",
 *     },
 *     throttleSettings: {
 *         burstLimit: 5,
 *         rateLimit: 10,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * AWS API Gateway Usage Plan can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:apigateway/usagePlan:UsagePlan myusageplan <usage_plan_id>
 * ```
 */
export class UsagePlan extends pulumi.CustomResource {
    /**
     * Get an existing UsagePlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UsagePlanState, opts?: pulumi.CustomResourceOptions): UsagePlan {
        return new UsagePlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/usagePlan:UsagePlan';

    /**
     * Returns true if the given object is an instance of UsagePlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UsagePlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UsagePlan.__pulumiType;
    }

    /**
     * The associated API stages of the usage plan.
     */
    public readonly apiStages!: pulumi.Output<outputs.apigateway.UsagePlanApiStage[] | undefined>;
    /**
     * Amazon Resource Name (ARN)
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of a usage plan.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the usage plan.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
     */
    public readonly productCode!: pulumi.Output<string | undefined>;
    /**
     * The quota settings of the usage plan.
     */
    public readonly quotaSettings!: pulumi.Output<outputs.apigateway.UsagePlanQuotaSettings | undefined>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The throttling limits of the usage plan.
     */
    public readonly throttleSettings!: pulumi.Output<outputs.apigateway.UsagePlanThrottleSettings | undefined>;

    /**
     * Create a UsagePlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UsagePlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UsagePlanArgs | UsagePlanState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UsagePlanState | undefined;
            inputs["apiStages"] = state ? state.apiStages : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["productCode"] = state ? state.productCode : undefined;
            inputs["quotaSettings"] = state ? state.quotaSettings : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["throttleSettings"] = state ? state.throttleSettings : undefined;
        } else {
            const args = argsOrState as UsagePlanArgs | undefined;
            inputs["apiStages"] = args ? args.apiStages : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["productCode"] = args ? args.productCode : undefined;
            inputs["quotaSettings"] = args ? args.quotaSettings : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["throttleSettings"] = args ? args.throttleSettings : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UsagePlan.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UsagePlan resources.
 */
export interface UsagePlanState {
    /**
     * The associated API stages of the usage plan.
     */
    readonly apiStages?: pulumi.Input<pulumi.Input<inputs.apigateway.UsagePlanApiStage>[]>;
    /**
     * Amazon Resource Name (ARN)
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The description of a usage plan.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the usage plan.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
     */
    readonly productCode?: pulumi.Input<string>;
    /**
     * The quota settings of the usage plan.
     */
    readonly quotaSettings?: pulumi.Input<inputs.apigateway.UsagePlanQuotaSettings>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The throttling limits of the usage plan.
     */
    readonly throttleSettings?: pulumi.Input<inputs.apigateway.UsagePlanThrottleSettings>;
}

/**
 * The set of arguments for constructing a UsagePlan resource.
 */
export interface UsagePlanArgs {
    /**
     * The associated API stages of the usage plan.
     */
    readonly apiStages?: pulumi.Input<pulumi.Input<inputs.apigateway.UsagePlanApiStage>[]>;
    /**
     * The description of a usage plan.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the usage plan.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
     */
    readonly productCode?: pulumi.Input<string>;
    /**
     * The quota settings of the usage plan.
     */
    readonly quotaSettings?: pulumi.Input<inputs.apigateway.UsagePlanQuotaSettings>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The throttling limits of the usage plan.
     */
    readonly throttleSettings?: pulumi.Input<inputs.apigateway.UsagePlanThrottleSettings>;
}
