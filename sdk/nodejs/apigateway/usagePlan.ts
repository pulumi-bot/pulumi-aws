// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
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
 * const aws_api_gateway_rest_api_myapi = new aws.apigateway.RestApi("myapi", {
 *     name: "MyDemoAPI",
 * });
 * const aws_api_gateway_deployment_dev = new aws.apigateway.Deployment("dev", {
 *     restApi: aws_api_gateway_rest_api_myapi.id,
 *     stageName: "dev",
 * });
 * const aws_api_gateway_deployment_prod = new aws.apigateway.Deployment("prod", {
 *     restApi: aws_api_gateway_rest_api_myapi.id,
 *     stageName: "prod",
 * });
 * const aws_api_gateway_usage_plan_MyUsagePlan = new aws.apigateway.UsagePlan("MyUsagePlan", {
 *     apiStages: [
 *         {
 *             apiId: aws_api_gateway_rest_api_myapi.id,
 *             stage: aws_api_gateway_deployment_dev.stageName,
 *         },
 *         {
 *             apiId: aws_api_gateway_rest_api_myapi.id,
 *             stage: aws_api_gateway_deployment_prod.stageName,
 *         },
 *     ],
 *     description: "my description",
 *     name: "my-usage-plan",
 *     productCode: "MYCODE",
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
 */
export class UsagePlan extends pulumi.CustomResource {
    /**
     * Get an existing UsagePlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UsagePlanState, opts?: pulumi.CustomResourceOptions): UsagePlan {
        return new UsagePlan(name, <any>state, { ...opts, id: id });
    }

    /**
     * The associated API stages of the usage plan.
     */
    public readonly apiStages: pulumi.Output<{ apiId: string, stage: string }[] | undefined>;
    /**
     * The description of a usage plan.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * The name of the usage plan.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
     */
    public readonly productCode: pulumi.Output<string | undefined>;
    /**
     * The quota settings of the usage plan.
     */
    public readonly quotaSettings: pulumi.Output<{ limit: number, offset?: number, period: string } | undefined>;
    /**
     * The throttling limits of the usage plan.
     */
    public readonly throttleSettings: pulumi.Output<{ burstLimit?: number, rateLimit?: number } | undefined>;

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
            const state: UsagePlanState = argsOrState as UsagePlanState | undefined;
            inputs["apiStages"] = state ? state.apiStages : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["productCode"] = state ? state.productCode : undefined;
            inputs["quotaSettings"] = state ? state.quotaSettings : undefined;
            inputs["throttleSettings"] = state ? state.throttleSettings : undefined;
        } else {
            const args = argsOrState as UsagePlanArgs | undefined;
            inputs["apiStages"] = args ? args.apiStages : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["productCode"] = args ? args.productCode : undefined;
            inputs["quotaSettings"] = args ? args.quotaSettings : undefined;
            inputs["throttleSettings"] = args ? args.throttleSettings : undefined;
        }
        super("aws:apigateway/usagePlan:UsagePlan", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UsagePlan resources.
 */
export interface UsagePlanState {
    /**
     * The associated API stages of the usage plan.
     */
    readonly apiStages?: pulumi.Input<pulumi.Input<{ apiId: pulumi.Input<string>, stage: pulumi.Input<string> }>[]>;
    /**
     * The description of a usage plan.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the usage plan.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
     */
    readonly productCode?: pulumi.Input<string>;
    /**
     * The quota settings of the usage plan.
     */
    readonly quotaSettings?: pulumi.Input<{ limit: pulumi.Input<number>, offset?: pulumi.Input<number>, period: pulumi.Input<string> }>;
    /**
     * The throttling limits of the usage plan.
     */
    readonly throttleSettings?: pulumi.Input<{ burstLimit?: pulumi.Input<number>, rateLimit?: pulumi.Input<number> }>;
}

/**
 * The set of arguments for constructing a UsagePlan resource.
 */
export interface UsagePlanArgs {
    /**
     * The associated API stages of the usage plan.
     */
    readonly apiStages?: pulumi.Input<pulumi.Input<{ apiId: pulumi.Input<string>, stage: pulumi.Input<string> }>[]>;
    /**
     * The description of a usage plan.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the usage plan.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
     */
    readonly productCode?: pulumi.Input<string>;
    /**
     * The quota settings of the usage plan.
     */
    readonly quotaSettings?: pulumi.Input<{ limit: pulumi.Input<number>, offset?: pulumi.Input<number>, period: pulumi.Input<string> }>;
    /**
     * The throttling limits of the usage plan.
     */
    readonly throttleSettings?: pulumi.Input<{ burstLimit?: pulumi.Input<number>, rateLimit?: pulumi.Input<number> }>;
}
