// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a AWS Batch compute environment. Compute environments contain the Amazon ECS container instances that are used to run containerized batch jobs.
 * 
 * For information about AWS Batch, see [What is AWS Batch?](http://docs.aws.amazon.com/batch/latest/userguide/what-is-batch.html) .
 * For information about compute environment, see [Compute Environments](http://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) .
 * 
 * > **Note:** To prevent a race condition during environment deletion, make sure to set `dependsOn` to the related `aws.iam.RolePolicyAttachment`;
 * otherwise, the policy may be destroyed too soon and the compute environment will then get stuck in the `DELETING` state, see [Troubleshooting AWS Batch](http://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html) .
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ecsInstanceRoleRole = new aws.iam.Role("ecsInstanceRole", {
 *     assumeRolePolicy: `{
 *     "Version": "2012-10-17",
 *     "Statement": [
 * 	{
 * 	    "Action": "sts:AssumeRole",
 * 	    "Effect": "Allow",
 * 	    "Principal": {
 * 		"Service": "ec2.amazonaws.com"
 * 	    }
 * 	}
 *     ]
 * }
 * `,
 *     name: "ecsInstanceRole",
 * });
 * const ecsInstanceRoleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("ecsInstanceRole", {
 *     policyArn: "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role",
 *     role: ecsInstanceRoleRole.name,
 * });
 * const ecsInstanceRoleInstanceProfile = new aws.iam.InstanceProfile("ecsInstanceRole", {
 *     name: "ecsInstanceRole",
 *     role: ecsInstanceRoleRole.name,
 * });
 * const awsBatchServiceRoleRole = new aws.iam.Role("awsBatchServiceRole", {
 *     assumeRolePolicy: `{
 *     "Version": "2012-10-17",
 *     "Statement": [
 * 	{
 * 	    "Action": "sts:AssumeRole",
 * 	    "Effect": "Allow",
 * 	    "Principal": {
 * 		"Service": "batch.amazonaws.com"
 * 	    }
 * 	}
 *     ]
 * }
 * `,
 *     name: "awsBatchServiceRole",
 * });
 * const awsBatchServiceRoleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("awsBatchServiceRole", {
 *     policyArn: "arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole",
 *     role: awsBatchServiceRoleRole.name,
 * });
 * const sampleSecurityGroup = new aws.ec2.SecurityGroup("sample", {
 *     egress: [{
 *         cidrBlocks: ["0.0.0.0/0"],
 *         fromPort: 0,
 *         protocol: "-1",
 *         toPort: 0,
 *     }],
 *     name: "awsBatchComputeEnvironmentSecurityGroup",
 * });
 * const sampleVpc = new aws.ec2.Vpc("sample", {
 *     cidrBlock: "10.1.0.0/16",
 * });
 * const sampleSubnet = new aws.ec2.Subnet("sample", {
 *     cidrBlock: "10.1.1.0/24",
 *     vpcId: sampleVpc.id,
 * });
 * const sampleComputeEnvironment = new aws.batch.ComputeEnvironment("sample", {
 *     computeEnvironmentName: "sample",
 *     computeResources: {
 *         instanceRole: ecsInstanceRoleInstanceProfile.arn,
 *         instanceTypes: ["c4.large"],
 *         maxVcpus: 16,
 *         minVcpus: 0,
 *         securityGroupIds: [sampleSecurityGroup.id],
 *         subnets: [sampleSubnet.id],
 *         type: "EC2",
 *     },
 *     serviceRole: awsBatchServiceRoleRole.arn,
 *     type: "MANAGED",
 * }, { dependsOn: [awsBatchServiceRoleRolePolicyAttachment] });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/batch_compute_environment.html.markdown.
 */
export class ComputeEnvironment extends pulumi.CustomResource {
    /**
     * Get an existing ComputeEnvironment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeEnvironmentState, opts?: pulumi.CustomResourceOptions): ComputeEnvironment {
        return new ComputeEnvironment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:batch/computeEnvironment:ComputeEnvironment';

    /**
     * Returns true if the given object is an instance of ComputeEnvironment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeEnvironment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeEnvironment.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the compute environment.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed. If omitted, this provider will assign a random, unique name.
     */
    public readonly computeEnvironmentName!: pulumi.Output<string>;
    /**
     * Creates a unique compute environment name beginning with the specified prefix. Conflicts with `computeEnvironmentName`.
     */
    public readonly computeEnvironmentNamePrefix!: pulumi.Output<string | undefined>;
    /**
     * Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
     */
    public readonly computeResources!: pulumi.Output<outputs.batch.ComputeEnvironmentComputeResources | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the underlying Amazon ECS cluster used by the compute environment.
     */
    public /*out*/ readonly ecsClusterArn!: pulumi.Output<string>;
    /**
     * The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
     */
    public readonly serviceRole!: pulumi.Output<string>;
    /**
     * The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * The current status of the compute environment (for example, CREATING or VALID).
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A short, human-readable string to provide additional details about the current status of the compute environment.
     */
    public /*out*/ readonly statusReason!: pulumi.Output<string>;
    /**
     * The type of compute environment. Valid items are `EC2` or `SPOT`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ComputeEnvironment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeEnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeEnvironmentArgs | ComputeEnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ComputeEnvironmentState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["computeEnvironmentName"] = state ? state.computeEnvironmentName : undefined;
            inputs["computeEnvironmentNamePrefix"] = state ? state.computeEnvironmentNamePrefix : undefined;
            inputs["computeResources"] = state ? state.computeResources : undefined;
            inputs["ecsClusterArn"] = state ? state.ecsClusterArn : undefined;
            inputs["serviceRole"] = state ? state.serviceRole : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["statusReason"] = state ? state.statusReason : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ComputeEnvironmentArgs | undefined;
            if (!args || args.serviceRole === undefined) {
                throw new Error("Missing required property 'serviceRole'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["computeEnvironmentName"] = args ? args.computeEnvironmentName : undefined;
            inputs["computeEnvironmentNamePrefix"] = args ? args.computeEnvironmentNamePrefix : undefined;
            inputs["computeResources"] = args ? args.computeResources : undefined;
            inputs["serviceRole"] = args ? args.serviceRole : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["ecsClusterArn"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusReason"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ComputeEnvironment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeEnvironment resources.
 */
export interface ComputeEnvironmentState {
    /**
     * The Amazon Resource Name (ARN) of the compute environment.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed. If omitted, this provider will assign a random, unique name.
     */
    readonly computeEnvironmentName?: pulumi.Input<string>;
    /**
     * Creates a unique compute environment name beginning with the specified prefix. Conflicts with `computeEnvironmentName`.
     */
    readonly computeEnvironmentNamePrefix?: pulumi.Input<string>;
    /**
     * Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
     */
    readonly computeResources?: pulumi.Input<inputs.batch.ComputeEnvironmentComputeResources>;
    /**
     * The Amazon Resource Name (ARN) of the underlying Amazon ECS cluster used by the compute environment.
     */
    readonly ecsClusterArn?: pulumi.Input<string>;
    /**
     * The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
     */
    readonly serviceRole?: pulumi.Input<string>;
    /**
     * The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The current status of the compute environment (for example, CREATING or VALID).
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A short, human-readable string to provide additional details about the current status of the compute environment.
     */
    readonly statusReason?: pulumi.Input<string>;
    /**
     * The type of compute environment. Valid items are `EC2` or `SPOT`.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeEnvironment resource.
 */
export interface ComputeEnvironmentArgs {
    /**
     * The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed. If omitted, this provider will assign a random, unique name.
     */
    readonly computeEnvironmentName?: pulumi.Input<string>;
    /**
     * Creates a unique compute environment name beginning with the specified prefix. Conflicts with `computeEnvironmentName`.
     */
    readonly computeEnvironmentNamePrefix?: pulumi.Input<string>;
    /**
     * Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
     */
    readonly computeResources?: pulumi.Input<inputs.batch.ComputeEnvironmentComputeResources>;
    /**
     * The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
     */
    readonly serviceRole: pulumi.Input<string>;
    /**
     * The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The type of compute environment. Valid items are `EC2` or `SPOT`.
     */
    readonly type: pulumi.Input<string>;
}
