// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Security Hub member resource.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleAccount = new aws.securityhub.Account("example", {});
 * const exampleMember = new aws.securityhub.Member("example", {
 *     accountId: "123456789012",
 *     email: "example@example.com",
 *     invite: true,
 * }, { dependsOn: [exampleAccount] });
 * ```
 */
export class Member extends pulumi.CustomResource {
    /**
     * Get an existing Member resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MemberState, opts?: pulumi.CustomResourceOptions): Member {
        return new Member(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:securityhub/member:Member';

    /**
     * Returns true if the given object is an instance of Member.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Member {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Member.__pulumiType;
    }

    /**
     * The ID of the member AWS account.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * The email of the member AWS account.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
     */
    public readonly invite!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the master Security Hub AWS account.
     */
    public /*out*/ readonly masterId!: pulumi.Output<string>;
    /**
     * The status of the relationship between the member account and its master account.
     */
    public /*out*/ readonly memberStatus!: pulumi.Output<string>;

    /**
     * Create a Member resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MemberArgs | MemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MemberState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["invite"] = state ? state.invite : undefined;
            inputs["masterId"] = state ? state.masterId : undefined;
            inputs["memberStatus"] = state ? state.memberStatus : undefined;
        } else {
            const args = argsOrState as MemberArgs | undefined;
            if (!args || args.accountId === undefined) {
                throw new Error("Missing required property 'accountId'");
            }
            if (!args || args.email === undefined) {
                throw new Error("Missing required property 'email'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["invite"] = args ? args.invite : undefined;
            inputs["masterId"] = undefined /*out*/;
            inputs["memberStatus"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Member.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Member resources.
 */
export interface MemberState {
    /**
     * The ID of the member AWS account.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * The email of the member AWS account.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
     */
    readonly invite?: pulumi.Input<boolean>;
    /**
     * The ID of the master Security Hub AWS account.
     */
    readonly masterId?: pulumi.Input<string>;
    /**
     * The status of the relationship between the member account and its master account.
     */
    readonly memberStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Member resource.
 */
export interface MemberArgs {
    /**
     * The ID of the member AWS account.
     */
    readonly accountId: pulumi.Input<string>;
    /**
     * The email of the member AWS account.
     */
    readonly email: pulumi.Input<string>;
    /**
     * Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
     */
    readonly invite?: pulumi.Input<boolean>;
}
