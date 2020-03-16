// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a RDS DB Instance association with an IAM Role. Example use cases:
 * 
 * * [Amazon RDS Oracle integration with Amazon S3](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html)
 * * [Importing Amazon S3 Data into an RDS PostgreSQL DB Instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.S3Import.html)
 * 
 * > To manage the RDS DB Instance IAM Role for [Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html), see the `aws.rds.Instance` resource `monitoringRoleArn` argument instead.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.rds.RoleAssociation("example", {
 *     dbInstanceIdentifier: aws_db_instance_example.id,
 *     featureName: "S3_INTEGRATION",
 *     roleArn: aws_iam_role_example.id,
 * });
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/db_instance_role_association.html.markdown.
 */
export class RoleAssociation extends pulumi.CustomResource {
    /**
     * Get an existing RoleAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleAssociationState, opts?: pulumi.CustomResourceOptions): RoleAssociation {
        return new RoleAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:rds/roleAssociation:RoleAssociation';

    /**
     * Returns true if the given object is an instance of RoleAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoleAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoleAssociation.__pulumiType;
    }

    /**
     * DB Instance Identifier to associate with the IAM Role.
     */
    public readonly dbInstanceIdentifier!: pulumi.Output<string>;
    /**
     * Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
     */
    public readonly featureName!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the IAM Role to associate with the DB Instance.
     */
    public readonly roleArn!: pulumi.Output<string>;

    /**
     * Create a RoleAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoleAssociationArgs | RoleAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RoleAssociationState | undefined;
            inputs["dbInstanceIdentifier"] = state ? state.dbInstanceIdentifier : undefined;
            inputs["featureName"] = state ? state.featureName : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
        } else {
            const args = argsOrState as RoleAssociationArgs | undefined;
            if (!args || args.dbInstanceIdentifier === undefined) {
                throw new Error("Missing required property 'dbInstanceIdentifier'");
            }
            if (!args || args.featureName === undefined) {
                throw new Error("Missing required property 'featureName'");
            }
            if (!args || args.roleArn === undefined) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["dbInstanceIdentifier"] = args ? args.dbInstanceIdentifier : undefined;
            inputs["featureName"] = args ? args.featureName : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RoleAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RoleAssociation resources.
 */
export interface RoleAssociationState {
    /**
     * DB Instance Identifier to associate with the IAM Role.
     */
    readonly dbInstanceIdentifier?: pulumi.Input<string>;
    /**
     * Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
     */
    readonly featureName?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the IAM Role to associate with the DB Instance.
     */
    readonly roleArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RoleAssociation resource.
 */
export interface RoleAssociationArgs {
    /**
     * DB Instance Identifier to associate with the IAM Role.
     */
    readonly dbInstanceIdentifier: pulumi.Input<string>;
    /**
     * Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
     */
    readonly featureName: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the IAM Role to associate with the DB Instance.
     */
    readonly roleArn: pulumi.Input<string>;
}
