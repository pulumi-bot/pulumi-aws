// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC/Subnet/ENI Flow Log to capture IP traffic for a specific network
 * interface, subnet, or VPC. Logs are sent to a CloudWatch Log Group or a S3 Bucket.
 * 
 * ## Example Usage
 * 
 * ### CloudWatch Logging
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("example", {});
 * const testRole = new aws.iam.Role("test_role", {
 *     assumeRolePolicy: `{
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
 * `,
 * });
 * const exampleRolePolicy = new aws.iam.RolePolicy("example", {
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
 *     role: aws_iam_role_example.id,
 * });
 * const exampleFlowLog = new aws.ec2.FlowLog("example", {
 *     iamRoleArn: aws_iam_role_example.arn,
 *     logDestination: exampleLogGroup.arn,
 *     trafficType: "ALL",
 *     vpcId: aws_vpc_example.id,
 * });
 * ```
 * 
 * ### S3 Logging
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleBucket = new aws.s3.Bucket("example", {
 *     name: "example",
 * });
 * const exampleFlowLog = new aws.ec2.FlowLog("example", {
 *     logDestination: exampleBucket.arn,
 *     logDestinationType: "s3",
 *     trafficType: "ALL",
 *     vpcId: aws_vpc_example.id,
 * });
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
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlowLogState, opts?: pulumi.CustomResourceOptions): FlowLog {
        return new FlowLog(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'aws:ec2/flowLog:FlowLog';

    /**
     * Returns true if the given object is an instance of FlowLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlowLog {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === FlowLog.__pulumiType;
    }

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
     * *Deprecated:* Use `log_destination` instead. The name of the CloudWatch log group.
     */
    public readonly logGroupName!: pulumi.Output<string>;
    /**
     * Subnet ID to attach to
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
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
        if (opts && opts.id) {
            const state = argsOrState as FlowLogState | undefined;
            inputs["eniId"] = state ? state.eniId : undefined;
            inputs["iamRoleArn"] = state ? state.iamRoleArn : undefined;
            inputs["logDestination"] = state ? state.logDestination : undefined;
            inputs["logDestinationType"] = state ? state.logDestinationType : undefined;
            inputs["logGroupName"] = state ? state.logGroupName : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["trafficType"] = state ? state.trafficType : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as FlowLogArgs | undefined;
            if (!args || args.trafficType === undefined) {
                throw new Error("Missing required property 'trafficType'");
            }
            inputs["eniId"] = args ? args.eniId : undefined;
            inputs["iamRoleArn"] = args ? args.iamRoleArn : undefined;
            inputs["logDestination"] = args ? args.logDestination : undefined;
            inputs["logDestinationType"] = args ? args.logDestinationType : undefined;
            inputs["logGroupName"] = args ? args.logGroupName : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["trafficType"] = args ? args.trafficType : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
        }
        super(FlowLog.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FlowLog resources.
 */
export interface FlowLogState {
    /**
     * Elastic Network Interface ID to attach to
     */
    readonly eniId?: pulumi.Input<string>;
    /**
     * The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
     */
    readonly iamRoleArn?: pulumi.Input<string>;
    /**
     * The ARN of the logging destination.
     */
    readonly logDestination?: pulumi.Input<string>;
    /**
     * The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
     */
    readonly logDestinationType?: pulumi.Input<string>;
    /**
     * *Deprecated:* Use `log_destination` instead. The name of the CloudWatch log group.
     */
    readonly logGroupName?: pulumi.Input<string>;
    /**
     * Subnet ID to attach to
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
     */
    readonly trafficType?: pulumi.Input<string>;
    /**
     * VPC ID to attach to
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FlowLog resource.
 */
export interface FlowLogArgs {
    /**
     * Elastic Network Interface ID to attach to
     */
    readonly eniId?: pulumi.Input<string>;
    /**
     * The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
     */
    readonly iamRoleArn?: pulumi.Input<string>;
    /**
     * The ARN of the logging destination.
     */
    readonly logDestination?: pulumi.Input<string>;
    /**
     * The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
     */
    readonly logDestinationType?: pulumi.Input<string>;
    /**
     * *Deprecated:* Use `log_destination` instead. The name of the CloudWatch log group.
     */
    readonly logGroupName?: pulumi.Input<string>;
    /**
     * Subnet ID to attach to
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
     */
    readonly trafficType: pulumi.Input<string>;
    /**
     * VPC ID to attach to
     */
    readonly vpcId?: pulumi.Input<string>;
}
