// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a CloudWatch Events permission to support cross-account events in the current account default event bus.
 * 
 * ## Example Usage
 * 
 * ### Account Access
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const devAccountAccess = new aws.cloudwatch.EventPermission("DevAccountAccess", {
 *     principal: "123456789012",
 *     statementId: "DevAccountAccess",
 * });
 * ```
 * {{% /examples %}}
 * 
 * ### Organization Access
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const organizationAccess = new aws.cloudwatch.EventPermission("OrganizationAccess", {
 *     condition: {
 *         key: "aws:PrincipalOrgID",
 *         type: "StringEquals",
 *         value: aws_organizations_organization_example.id,
 *     },
 *     principal: "*",
 *     statementId: "OrganizationAccess",
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_permission.html.markdown.
 */
export class EventPermission extends pulumi.CustomResource {
    /**
     * Get an existing EventPermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventPermissionState, opts?: pulumi.CustomResourceOptions): EventPermission {
        return new EventPermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/eventPermission:EventPermission';

    /**
     * Returns true if the given object is an instance of EventPermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventPermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventPermission.__pulumiType;
    }

    /**
     * The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
     */
    public readonly condition!: pulumi.Output<outputs.cloudwatch.EventPermissionCondition | undefined>;
    /**
     * The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * An identifier string for the external account that you are granting permissions to.
     */
    public readonly statementId!: pulumi.Output<string>;

    /**
     * Create a EventPermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventPermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventPermissionArgs | EventPermissionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventPermissionState | undefined;
            inputs["action"] = state ? state.action : undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["principal"] = state ? state.principal : undefined;
            inputs["statementId"] = state ? state.statementId : undefined;
        } else {
            const args = argsOrState as EventPermissionArgs | undefined;
            if (!args || args.principal === undefined) {
                throw new Error("Missing required property 'principal'");
            }
            if (!args || args.statementId === undefined) {
                throw new Error("Missing required property 'statementId'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["condition"] = args ? args.condition : undefined;
            inputs["principal"] = args ? args.principal : undefined;
            inputs["statementId"] = args ? args.statementId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventPermission.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventPermission resources.
 */
export interface EventPermissionState {
    /**
     * The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
     */
    readonly action?: pulumi.Input<string>;
    /**
     * Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
     */
    readonly condition?: pulumi.Input<inputs.cloudwatch.EventPermissionCondition>;
    /**
     * The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
     */
    readonly principal?: pulumi.Input<string>;
    /**
     * An identifier string for the external account that you are granting permissions to.
     */
    readonly statementId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventPermission resource.
 */
export interface EventPermissionArgs {
    /**
     * The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
     */
    readonly action?: pulumi.Input<string>;
    /**
     * Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
     */
    readonly condition?: pulumi.Input<inputs.cloudwatch.EventPermissionCondition>;
    /**
     * The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
     */
    readonly principal: pulumi.Input<string>;
    /**
     * An identifier string for the external account that you are granting permissions to.
     */
    readonly statementId: pulumi.Input<string>;
}
