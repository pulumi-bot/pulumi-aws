// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Config Organization Custom Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Managed Rules (those invoking an AWS managed rule), see the [`aws_config_organization_managed__rule` resource](https://www.terraform.io/docs/providers/aws/r/config_organization_managed_rule.html).
 *
 * > **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excludedAccounts` argument.
 *
 * > **NOTE:** The proper Lambda permission to allow the AWS Config service invoke the Lambda Function must be in place before the rule will successfully create or update. See also the [`aws.lambda.Permission` resource](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html).
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const examplePermission = new aws.lambda.Permission("example", {
 *     action: "lambda:InvokeFunction",
 *     function: aws_lambda_function_example.arn,
 *     principal: "config.amazonaws.com",
 * });
 * const exampleOrganization = new aws.organizations.Organization("example", {
 *     awsServiceAccessPrincipals: ["config-multiaccountsetup.amazonaws.com"],
 *     featureSet: "ALL",
 * });
 * const exampleOrganizationCustomRule = new aws.cfg.OrganizationCustomRule("example", {
 *     lambdaFunctionArn: aws_lambda_function_example.arn,
 *     triggerTypes: ["ConfigurationItemChangeNotification"],
 * }, { dependsOn: [examplePermission, exampleOrganization] });
 * ```
 */
export class OrganizationCustomRule extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationCustomRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationCustomRuleState, opts?: pulumi.CustomResourceOptions): OrganizationCustomRule {
        return new OrganizationCustomRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cfg/organizationCustomRule:OrganizationCustomRule';

    /**
     * Returns true if the given object is an instance of OrganizationCustomRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationCustomRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationCustomRule.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the rule
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Description of the rule
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of AWS account identifiers to exclude from the rule
     */
    public readonly excludedAccounts!: pulumi.Output<string[] | undefined>;
    /**
     * A string in JSON format that is passed to the AWS Config Rule Lambda Function
     */
    public readonly inputParameters!: pulumi.Output<string | undefined>;
    /**
     * Amazon Resource Name (ARN) of the rule Lambda Function
     */
    public readonly lambdaFunctionArn!: pulumi.Output<string>;
    /**
     * The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
     */
    public readonly maximumExecutionFrequency!: pulumi.Output<string | undefined>;
    /**
     * The name of the rule
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Identifier of the AWS resource to evaluate
     */
    public readonly resourceIdScope!: pulumi.Output<string | undefined>;
    /**
     * List of types of AWS resources to evaluate
     */
    public readonly resourceTypesScopes!: pulumi.Output<string[] | undefined>;
    /**
     * Tag key of AWS resources to evaluate
     */
    public readonly tagKeyScope!: pulumi.Output<string | undefined>;
    /**
     * Tag value of AWS resources to evaluate
     */
    public readonly tagValueScope!: pulumi.Output<string | undefined>;
    /**
     * List of notification types that trigger AWS Config to run an evaluation for the rule. Valid values: `ConfigurationItemChangeNotification`, `OversizedConfigurationItemChangeNotification`, and `ScheduledNotification`
     */
    public readonly triggerTypes!: pulumi.Output<string[]>;

    /**
     * Create a OrganizationCustomRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationCustomRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationCustomRuleArgs | OrganizationCustomRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OrganizationCustomRuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["excludedAccounts"] = state ? state.excludedAccounts : undefined;
            inputs["inputParameters"] = state ? state.inputParameters : undefined;
            inputs["lambdaFunctionArn"] = state ? state.lambdaFunctionArn : undefined;
            inputs["maximumExecutionFrequency"] = state ? state.maximumExecutionFrequency : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceIdScope"] = state ? state.resourceIdScope : undefined;
            inputs["resourceTypesScopes"] = state ? state.resourceTypesScopes : undefined;
            inputs["tagKeyScope"] = state ? state.tagKeyScope : undefined;
            inputs["tagValueScope"] = state ? state.tagValueScope : undefined;
            inputs["triggerTypes"] = state ? state.triggerTypes : undefined;
        } else {
            const args = argsOrState as OrganizationCustomRuleArgs | undefined;
            if (!args || args.lambdaFunctionArn === undefined) {
                throw new Error("Missing required property 'lambdaFunctionArn'");
            }
            if (!args || args.triggerTypes === undefined) {
                throw new Error("Missing required property 'triggerTypes'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["excludedAccounts"] = args ? args.excludedAccounts : undefined;
            inputs["inputParameters"] = args ? args.inputParameters : undefined;
            inputs["lambdaFunctionArn"] = args ? args.lambdaFunctionArn : undefined;
            inputs["maximumExecutionFrequency"] = args ? args.maximumExecutionFrequency : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceIdScope"] = args ? args.resourceIdScope : undefined;
            inputs["resourceTypesScopes"] = args ? args.resourceTypesScopes : undefined;
            inputs["tagKeyScope"] = args ? args.tagKeyScope : undefined;
            inputs["tagValueScope"] = args ? args.tagValueScope : undefined;
            inputs["triggerTypes"] = args ? args.triggerTypes : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OrganizationCustomRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationCustomRule resources.
 */
export interface OrganizationCustomRuleState {
    /**
     * Amazon Resource Name (ARN) of the rule
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Description of the rule
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of AWS account identifiers to exclude from the rule
     */
    readonly excludedAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A string in JSON format that is passed to the AWS Config Rule Lambda Function
     */
    readonly inputParameters?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the rule Lambda Function
     */
    readonly lambdaFunctionArn?: pulumi.Input<string>;
    /**
     * The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
     */
    readonly maximumExecutionFrequency?: pulumi.Input<string>;
    /**
     * The name of the rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Identifier of the AWS resource to evaluate
     */
    readonly resourceIdScope?: pulumi.Input<string>;
    /**
     * List of types of AWS resources to evaluate
     */
    readonly resourceTypesScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tag key of AWS resources to evaluate
     */
    readonly tagKeyScope?: pulumi.Input<string>;
    /**
     * Tag value of AWS resources to evaluate
     */
    readonly tagValueScope?: pulumi.Input<string>;
    /**
     * List of notification types that trigger AWS Config to run an evaluation for the rule. Valid values: `ConfigurationItemChangeNotification`, `OversizedConfigurationItemChangeNotification`, and `ScheduledNotification`
     */
    readonly triggerTypes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a OrganizationCustomRule resource.
 */
export interface OrganizationCustomRuleArgs {
    /**
     * Description of the rule
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of AWS account identifiers to exclude from the rule
     */
    readonly excludedAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A string in JSON format that is passed to the AWS Config Rule Lambda Function
     */
    readonly inputParameters?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the rule Lambda Function
     */
    readonly lambdaFunctionArn: pulumi.Input<string>;
    /**
     * The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
     */
    readonly maximumExecutionFrequency?: pulumi.Input<string>;
    /**
     * The name of the rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Identifier of the AWS resource to evaluate
     */
    readonly resourceIdScope?: pulumi.Input<string>;
    /**
     * List of types of AWS resources to evaluate
     */
    readonly resourceTypesScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tag key of AWS resources to evaluate
     */
    readonly tagKeyScope?: pulumi.Input<string>;
    /**
     * Tag value of AWS resources to evaluate
     */
    readonly tagValueScope?: pulumi.Input<string>;
    /**
     * List of notification types that trigger AWS Config to run an evaluation for the rule. Valid values: `ConfigurationItemChangeNotification`, `OversizedConfigurationItemChangeNotification`, and `ScheduledNotification`
     */
    readonly triggerTypes: pulumi.Input<pulumi.Input<string>[]>;
}
