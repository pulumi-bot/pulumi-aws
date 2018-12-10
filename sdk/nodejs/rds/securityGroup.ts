// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {Tags} from "../index";

/**
 * Provides an RDS security group resource. This is only for DB instances in the
 * EC2-Classic Platform. For instances inside a VPC, use the
 * [`aws_db_instance.vpc_security_group_ids`](https://www.terraform.io/docs/providers/aws/r/db_instance.html#vpc_security_group_ids)
 * attribute instead.
 */
export class SecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupState, opts?: pulumi.CustomResourceOptions): SecurityGroup {
        return new SecurityGroup(name, <any>state, { ...opts, id: id });
    }

    /**
     * The arn of the DB security group.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The description of the DB security group. Defaults to "Managed by Terraform".
     */
    public readonly description: pulumi.Output<string>;
    /**
     * A list of ingress rules.
     */
    public readonly ingress: pulumi.Output<{ cidr?: string, securityGroupId: string, securityGroupName: string, securityGroupOwnerId: string }[]>;
    /**
     * The name of the DB security group.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<Tags | undefined>;

    /**
     * Create a SecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupArgs | SecurityGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SecurityGroupState = argsOrState as SecurityGroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ingress"] = state ? state.ingress : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SecurityGroupArgs | undefined;
            if (!args || args.ingress === undefined) {
                throw new Error("Missing required property 'ingress'");
            }
            inputs["description"] = (args ? args.description : undefined) || "Managed by Pulumi";
            inputs["ingress"] = args ? args.ingress : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:rds/securityGroup:SecurityGroup", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroup resources.
 */
export interface SecurityGroupState {
    /**
     * The arn of the DB security group.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The description of the DB security group. Defaults to "Managed by Terraform".
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of ingress rules.
     */
    readonly ingress?: pulumi.Input<pulumi.Input<{ cidr?: pulumi.Input<string>, securityGroupId?: pulumi.Input<string>, securityGroupName?: pulumi.Input<string>, securityGroupOwnerId?: pulumi.Input<string> }>[]>;
    /**
     * The name of the DB security group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<Tags>;
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    /**
     * The description of the DB security group. Defaults to "Managed by Terraform".
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of ingress rules.
     */
    readonly ingress: pulumi.Input<pulumi.Input<{ cidr?: pulumi.Input<string>, securityGroupId?: pulumi.Input<string>, securityGroupName?: pulumi.Input<string>, securityGroupOwnerId?: pulumi.Input<string> }>[]>;
    /**
     * The name of the DB security group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<Tags>;
}
