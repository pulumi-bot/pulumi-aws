// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC/Subnet/ENI Flow Log to capture IP traffic for a specific network
 * interface, subnet, or VPC. Logs are sent to a CloudWatch Log Group or a S3 Bucket.
 *
 * ## Example Usage
 * ### CloudWatch Logging
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
 * const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Sid": "",
 *       "Effect": "Allow",
 *       "Principal": {
 *         "Service": "vpc-flow-logs.amazonaws.com"
 *       },
 *       "Action": "sts:AssumeRole"
 *     }
 *   ]
 * }
 * `});
 * const exampleFlowLog = new aws.ec2.FlowLog("exampleFlowLog", {
 *     iamRoleArn: exampleRole.arn,
 *     logDestination: exampleLogGroup.arn,
 *     trafficType: "ALL",
 *     vpcId: aws_vpc.example.id,
 * });
 * const exampleRolePolicy = new aws.iam.RolePolicy("exampleRolePolicy", {
 *     role: exampleRole.id,
 *     policy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "logs:CreateLogGroup",
 *         "logs:CreateLogStream",
 *         "logs:PutLogEvents",
 *         "logs:DescribeLogGroups",
 *         "logs:DescribeLogStreams"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": "*"
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 * ### S3 Logging
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleBucket = new aws.s3.Bucket("exampleBucket", {});
 * const exampleFlowLog = new aws.ec2.FlowLog("exampleFlowLog", {
 *     logDestination: exampleBucket.arn,
 *     logDestinationType: "s3",
 *     trafficType: "ALL",
 *     vpcId: aws_vpc.example.id,
 * });
 * ```
 *
 * ## Import
 *
 * Flow Logs can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:ec2/flowLog:FlowLog test_flow_log fl-1a2b3c4d
 * ```
 */
export class FlowLog extends pulumi.CustomResource {
    /**
     * Get an existing FlowLog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlowLogState, opts?: pulumi.CustomResourceOptions): FlowLog {
        return new FlowLog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/flowLog:FlowLog';

    /**
     * Returns true if the given object is an instance of FlowLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlowLog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlowLog.__pulumiType;
    }

    /**
     * The ARN of the Flow Log.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Elastic Network Interface ID to attach to
     */
    public readonly eniId!: pulumi.Output<string | undefined>;
    /**
     * The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
     */
    public readonly iamRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the logging destination.
     */
    public readonly logDestination!: pulumi.Output<string>;
    /**
     * The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
     */
    public readonly logDestinationType!: pulumi.Output<string | undefined>;
    /**
     * The fields to include in the flow log record, in the order in which they should appear.
     */
    public readonly logFormat!: pulumi.Output<string>;
    /**
     * *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group.
     *
     * @deprecated use 'log_destination' argument instead
     */
    public readonly logGroupName!: pulumi.Output<string>;
    /**
     * The maximum interval of time
     * during which a flow of packets is captured and aggregated into a flow
     * log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
     * minutes). Default: `600`.
     */
    public readonly maxAggregationInterval!: pulumi.Output<number | undefined>;
    /**
     * Subnet ID to attach to
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
     */
    public readonly trafficType!: pulumi.Output<string>;
    /**
     * VPC ID to attach to
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a FlowLog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlowLogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlowLogArgs | FlowLogState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FlowLogState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["eniId"] = state ? state.eniId : undefined;
            inputs["iamRoleArn"] = state ? state.iamRoleArn : undefined;
            inputs["logDestination"] = state ? state.logDestination : undefined;
            inputs["logDestinationType"] = state ? state.logDestinationType : undefined;
            inputs["logFormat"] = state ? state.logFormat : undefined;
            inputs["logGroupName"] = state ? state.logGroupName : undefined;
            inputs["maxAggregationInterval"] = state ? state.maxAggregationInterval : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["trafficType"] = state ? state.trafficType : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as FlowLogArgs | undefined;
            if ((!args || args.trafficType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficType'");
            }
            inputs["eniId"] = args ? args.eniId : undefined;
            inputs["iamRoleArn"] = args ? args.iamRoleArn : undefined;
            inputs["logDestination"] = args ? args.logDestination : undefined;
            inputs["logDestinationType"] = args ? args.logDestinationType : undefined;
            inputs["logFormat"] = args ? args.logFormat : undefined;
            inputs["logGroupName"] = args ? args.logGroupName : undefined;
            inputs["maxAggregationInterval"] = args ? args.maxAggregationInterval : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["trafficType"] = args ? args.trafficType : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FlowLog.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FlowLog resources.
 */
export interface FlowLogState {
    /**
     * The ARN of the Flow Log.
     */
    arn?: pulumi.Input<string>;
    /**
     * Elastic Network Interface ID to attach to
     */
    eniId?: pulumi.Input<string>;
    /**
     * The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
     */
    iamRoleArn?: pulumi.Input<string>;
    /**
     * The ARN of the logging destination.
     */
    logDestination?: pulumi.Input<string>;
    /**
     * The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
     */
    logDestinationType?: pulumi.Input<string>;
    /**
     * The fields to include in the flow log record, in the order in which they should appear.
     */
    logFormat?: pulumi.Input<string>;
    /**
     * *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group.
     *
     * @deprecated use 'log_destination' argument instead
     */
    logGroupName?: pulumi.Input<string>;
    /**
     * The maximum interval of time
     * during which a flow of packets is captured and aggregated into a flow
     * log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
     * minutes). Default: `600`.
     */
    maxAggregationInterval?: pulumi.Input<number>;
    /**
     * Subnet ID to attach to
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
     */
    trafficType?: pulumi.Input<string>;
    /**
     * VPC ID to attach to
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FlowLog resource.
 */
export interface FlowLogArgs {
    /**
     * Elastic Network Interface ID to attach to
     */
    eniId?: pulumi.Input<string>;
    /**
     * The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
     */
    iamRoleArn?: pulumi.Input<string>;
    /**
     * The ARN of the logging destination.
     */
    logDestination?: pulumi.Input<string>;
    /**
     * The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
     */
    logDestinationType?: pulumi.Input<string>;
    /**
     * The fields to include in the flow log record, in the order in which they should appear.
     */
    logFormat?: pulumi.Input<string>;
    /**
     * *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group.
     *
     * @deprecated use 'log_destination' argument instead
     */
    logGroupName?: pulumi.Input<string>;
    /**
     * The maximum interval of time
     * during which a flow of packets is captured and aggregated into a flow
     * log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
     * minutes). Default: `600`.
     */
    maxAggregationInterval?: pulumi.Input<number>;
    /**
     * Subnet ID to attach to
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
     */
    trafficType: pulumi.Input<string>;
    /**
     * VPC ID to attach to
     */
    vpcId?: pulumi.Input<string>;
}
