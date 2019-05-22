// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a new Amazon Redshift security group. You use security groups to control access to non-VPC clusters
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const defaultSecurityGroup = new aws.redshift.SecurityGroup("default", {
 *     ingress: [{
 *         cidr: "10.0.0.0/24",
 *     }],
 * });
 * ```
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
     * The description of the Redshift security group. Defaults to "Managed by Terraform".
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * A list of ingress rules.
     */
    public readonly ingress!: pulumi.Output<{ cidr?: string, securityGroupName: string, securityGroupOwnerId: string }[]>;
    /**
     * The name of the Redshift security group.
     */
    public readonly name!: pulumi.Output<string>;

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
            const state = argsOrState as SecurityGroupState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ingress"] = state ? state.ingress : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as SecurityGroupArgs | undefined;
            if (!args || args.ingress === undefined) {
                throw new Error("Missing required property 'ingress'");
            }
            inputs["description"] = (args ? args.description : undefined) || "Managed by Pulumi";
            inputs["ingress"] = args ? args.ingress : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        super("aws:redshift/securityGroup:SecurityGroup", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroup resources.
 */
export interface SecurityGroupState {
    /**
     * The description of the Redshift security group. Defaults to "Managed by Terraform".
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of ingress rules.
     */
    readonly ingress?: pulumi.Input<pulumi.Input<{ cidr?: pulumi.Input<string>, securityGroupName?: pulumi.Input<string>, securityGroupOwnerId?: pulumi.Input<string> }>[]>;
    /**
     * The name of the Redshift security group.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    /**
     * The description of the Redshift security group. Defaults to "Managed by Terraform".
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of ingress rules.
     */
    readonly ingress: pulumi.Input<pulumi.Input<{ cidr?: pulumi.Input<string>, securityGroupName?: pulumi.Input<string>, securityGroupOwnerId?: pulumi.Input<string> }>[]>;
    /**
     * The name of the Redshift security group.
     */
    readonly name?: pulumi.Input<string>;
}
