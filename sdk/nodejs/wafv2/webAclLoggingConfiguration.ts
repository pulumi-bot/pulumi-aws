// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Creates a WAFv2 Web ACL Logging Configuration resource.
 *
 * > **Note:** To start logging from a WAFv2 Web ACL, an Amazon Kinesis Data Firehose (e.g. `aws.kinesis.FirehoseDeliveryStream` resourc must also be created with a PUT source (not a stream) and in the region that you are operating.
 * If you are capturing logs for Amazon CloudFront, always create the firehose in US East (N. Virginia).
 * Be sure to give the data firehose a name that starts with the prefix `aws-waf-logs-`.
 *
 * ## Example Usage
 * ### With Redacted Fields
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.wafv2.WebAclLoggingConfiguration("example", {
 *     logDestinationConfigs: [aws_kinesis_firehose_delivery_stream.example.arn],
 *     resourceArn: aws_wafv2_web_acl.example.arn,
 *     redactedFields: [{
 *         singleHeader: {
 *             name: "user-agent",
 *         },
 *     }],
 * });
 * ```
 * ### With Logging Filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.wafv2.WebAclLoggingConfiguration("example", {
 *     logDestinationConfigs: [aws_kinesis_firehose_delivery_stream.example.arn],
 *     resourceArn: aws_wafv2_web_acl.example.arn,
 *     loggingFilter: {
 *         defaultBehavior: "KEEP",
 *         filters: [
 *             {
 *                 behavior: "DROP",
 *                 conditions: [
 *                     {
 *                         actionCondition: {
 *                             action: "COUNT",
 *                         },
 *                     },
 *                     {
 *                         labelNameCondition: {
 *                             labelName: "awswaf:111122223333:rulegroup:testRules:LabelNameZ",
 *                         },
 *                     },
 *                 ],
 *                 requirement: "MEETS_ALL",
 *             },
 *             {
 *                 behavior: "KEEP",
 *                 conditions: [{
 *                     actionCondition: {
 *                         action: "ALLOW",
 *                     },
 *                 }],
 *                 requirement: "MEETS_ANY",
 *             },
 *         ],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * WAFv2 Web ACL Logging Configurations can be imported using the WAFv2 Web ACL ARN e.g.
 *
 * ```sh
 *  $ pulumi import aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration example arn:aws:wafv2:us-west-2:123456789012:regional/webacl/test-logs/a1b2c3d4-5678-90ab-cdef
 * ```
 */
export class WebAclLoggingConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing WebAclLoggingConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebAclLoggingConfigurationState, opts?: pulumi.CustomResourceOptions): WebAclLoggingConfiguration {
        return new WebAclLoggingConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration';

    /**
     * Returns true if the given object is an instance of WebAclLoggingConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAclLoggingConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAclLoggingConfiguration.__pulumiType;
    }

    /**
     * The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
     */
    public readonly logDestinationConfigs!: pulumi.Output<string[]>;
    /**
     * A configuration block that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation. See Logging Filter below for more details.
     */
    public readonly loggingFilter!: pulumi.Output<outputs.wafv2.WebAclLoggingConfigurationLoggingFilter | undefined>;
    /**
     * The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported. See Redacted Fields below for more details.
     */
    public readonly redactedFields!: pulumi.Output<outputs.wafv2.WebAclLoggingConfigurationRedactedField[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
     */
    public readonly resourceArn!: pulumi.Output<string>;

    /**
     * Create a WebAclLoggingConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAclLoggingConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAclLoggingConfigurationArgs | WebAclLoggingConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebAclLoggingConfigurationState | undefined;
            inputs["logDestinationConfigs"] = state ? state.logDestinationConfigs : undefined;
            inputs["loggingFilter"] = state ? state.loggingFilter : undefined;
            inputs["redactedFields"] = state ? state.redactedFields : undefined;
            inputs["resourceArn"] = state ? state.resourceArn : undefined;
        } else {
            const args = argsOrState as WebAclLoggingConfigurationArgs | undefined;
            if ((!args || args.logDestinationConfigs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logDestinationConfigs'");
            }
            if ((!args || args.resourceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceArn'");
            }
            inputs["logDestinationConfigs"] = args ? args.logDestinationConfigs : undefined;
            inputs["loggingFilter"] = args ? args.loggingFilter : undefined;
            inputs["redactedFields"] = args ? args.redactedFields : undefined;
            inputs["resourceArn"] = args ? args.resourceArn : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebAclLoggingConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebAclLoggingConfiguration resources.
 */
export interface WebAclLoggingConfigurationState {
    /**
     * The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
     */
    logDestinationConfigs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A configuration block that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation. See Logging Filter below for more details.
     */
    loggingFilter?: pulumi.Input<inputs.wafv2.WebAclLoggingConfigurationLoggingFilter>;
    /**
     * The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported. See Redacted Fields below for more details.
     */
    redactedFields?: pulumi.Input<pulumi.Input<inputs.wafv2.WebAclLoggingConfigurationRedactedField>[]>;
    /**
     * The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
     */
    resourceArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebAclLoggingConfiguration resource.
 */
export interface WebAclLoggingConfigurationArgs {
    /**
     * The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
     */
    logDestinationConfigs: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A configuration block that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation. See Logging Filter below for more details.
     */
    loggingFilter?: pulumi.Input<inputs.wafv2.WebAclLoggingConfigurationLoggingFilter>;
    /**
     * The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported. See Redacted Fields below for more details.
     */
    redactedFields?: pulumi.Input<pulumi.Input<inputs.wafv2.WebAclLoggingConfigurationRedactedField>[]>;
    /**
     * The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
     */
    resourceArn: pulumi.Input<string>;
}
