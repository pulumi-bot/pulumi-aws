// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage the default AWS Network ACL. VPC Only.
 * 
 * Each VPC created in AWS comes with a Default Network ACL that can be managed, but not
 * destroyed. **This is an advanced resource**, and has special caveats to be aware
 * of when using it. Please read this document in its entirety before using this
 * resource.
 * 
 * The `aws.ec2.DefaultNetworkAcl` behaves differently from normal resources, in that
 * this provider does not _create_ this resource, but instead attempts to "adopt" it
 * into management. We can do this because each VPC created has a Default Network
 * ACL that cannot be destroyed, and is created with a known set of default rules.
 * 
 * When this provider first adopts the Default Network ACL, it **immediately removes all
 * rules in the ACL**. It then proceeds to create any rules specified in the
 * configuration. This step is required so that only the rules specified in the
 * configuration are created.
 * 
 * This resource treats its inline rules as absolute; only the rules defined
 * inline are created, and any additions/removals external to this resource will
 * result in diffs being shown. For these reasons, this resource is incompatible with the
 * `aws.ec2.NetworkAclRule` resource.
 * 
 * For more information about Network ACLs, see the AWS Documentation on
 * [Network ACLs][aws-network-acls].
 * 
 * ## Example config to deny all traffic to any Subnet in the Default Network ACL
 * 
 * {{% examples %}}
 * 
 * This config denies all traffic in the Default ACL. This can be useful if you
 * want a locked down default to force all resources in the VPC to assign a
 * non-default ACL.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const mainvpc = new aws.ec2.Vpc("mainvpc", {
 *     cidrBlock: "10.1.0.0/16",
 * });
 * const defaultDefaultNetworkAcl = new aws.ec2.DefaultNetworkAcl("default", {
 *     defaultNetworkAclId: mainvpc.defaultNetworkAclId,
 * });
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/default_network_acl.html.markdown.
 */
export class DefaultNetworkAcl extends pulumi.CustomResource {
    /**
     * Get an existing DefaultNetworkAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultNetworkAclState, opts?: pulumi.CustomResourceOptions): DefaultNetworkAcl {
        return new DefaultNetworkAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/defaultNetworkAcl:DefaultNetworkAcl';

    /**
     * Returns true if the given object is an instance of DefaultNetworkAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultNetworkAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultNetworkAcl.__pulumiType;
    }

    /**
     * The Network ACL ID to manage. This
     * attribute is exported from `aws.ec2.Vpc`, or manually found via the AWS Console.
     */
    public readonly defaultNetworkAclId!: pulumi.Output<string>;
    /**
     * Specifies an egress rule. Parameters defined below.
     */
    public readonly egress!: pulumi.Output<outputs.ec2.DefaultNetworkAclEgress[] | undefined>;
    /**
     * Specifies an ingress rule. Parameters defined below.
     */
    public readonly ingress!: pulumi.Output<outputs.ec2.DefaultNetworkAclIngress[] | undefined>;
    /**
     * The ID of the AWS account that owns the Default Network ACL
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * A list of Subnet IDs to apply the ACL to. See the
     * notes below on managing Subnets in the Default Network ACL
     */
    public readonly subnetIds!: pulumi.Output<string[] | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the associated VPC
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a DefaultNetworkAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultNetworkAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultNetworkAclArgs | DefaultNetworkAclState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DefaultNetworkAclState | undefined;
            inputs["defaultNetworkAclId"] = state ? state.defaultNetworkAclId : undefined;
            inputs["egress"] = state ? state.egress : undefined;
            inputs["ingress"] = state ? state.ingress : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as DefaultNetworkAclArgs | undefined;
            if (!args || args.defaultNetworkAclId === undefined) {
                throw new Error("Missing required property 'defaultNetworkAclId'");
            }
            inputs["defaultNetworkAclId"] = args ? args.defaultNetworkAclId : undefined;
            inputs["egress"] = args ? args.egress : undefined;
            inputs["ingress"] = args ? args.ingress : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["ownerId"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DefaultNetworkAcl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultNetworkAcl resources.
 */
export interface DefaultNetworkAclState {
    /**
     * The Network ACL ID to manage. This
     * attribute is exported from `aws.ec2.Vpc`, or manually found via the AWS Console.
     */
    readonly defaultNetworkAclId?: pulumi.Input<string>;
    /**
     * Specifies an egress rule. Parameters defined below.
     */
    readonly egress?: pulumi.Input<pulumi.Input<inputs.ec2.DefaultNetworkAclEgress>[]>;
    /**
     * Specifies an ingress rule. Parameters defined below.
     */
    readonly ingress?: pulumi.Input<pulumi.Input<inputs.ec2.DefaultNetworkAclIngress>[]>;
    /**
     * The ID of the AWS account that owns the Default Network ACL
     */
    readonly ownerId?: pulumi.Input<string>;
    /**
     * A list of Subnet IDs to apply the ACL to. See the
     * notes below on managing Subnets in the Default Network ACL
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the associated VPC
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DefaultNetworkAcl resource.
 */
export interface DefaultNetworkAclArgs {
    /**
     * The Network ACL ID to manage. This
     * attribute is exported from `aws.ec2.Vpc`, or manually found via the AWS Console.
     */
    readonly defaultNetworkAclId: pulumi.Input<string>;
    /**
     * Specifies an egress rule. Parameters defined below.
     */
    readonly egress?: pulumi.Input<pulumi.Input<inputs.ec2.DefaultNetworkAclEgress>[]>;
    /**
     * Specifies an ingress rule. Parameters defined below.
     */
    readonly ingress?: pulumi.Input<pulumi.Input<inputs.ec2.DefaultNetworkAclIngress>[]>;
    /**
     * A list of Subnet IDs to apply the ACL to. See the
     * notes below on managing Subnets in the Default Network ACL
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
