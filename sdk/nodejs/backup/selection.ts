// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages selection conditions for AWS Backup plan resources.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/backup_selection.html.markdown.
 */
export class Selection extends pulumi.CustomResource {
    /**
     * Get an existing Selection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SelectionState, opts?: pulumi.CustomResourceOptions): Selection {
        return new Selection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:backup/selection:Selection';

    /**
     * Returns true if the given object is an instance of Selection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Selection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Selection.__pulumiType;
    }

    /**
     * The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
     */
    public readonly iamRoleArn!: pulumi.Output<string>;
    /**
     * The display name of a resource selection document.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The backup plan ID to be associated with the selection of resources.
     */
    public readonly planId!: pulumi.Output<string>;
    /**
     * An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
     */
    public readonly resources!: pulumi.Output<string[] | undefined>;
    /**
     * Tag-based conditions used to specify a set of resources to assign to a backup plan.
     */
    public readonly selectionTags!: pulumi.Output<outputs.backup.SelectionSelectionTag[] | undefined>;

    /**
     * Create a Selection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SelectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SelectionArgs | SelectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SelectionState | undefined;
            inputs["iamRoleArn"] = state ? state.iamRoleArn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["planId"] = state ? state.planId : undefined;
            inputs["resources"] = state ? state.resources : undefined;
            inputs["selectionTags"] = state ? state.selectionTags : undefined;
        } else {
            const args = argsOrState as SelectionArgs | undefined;
            if (!args || args.iamRoleArn === undefined) {
                throw new Error("Missing required property 'iamRoleArn'");
            }
            if (!args || args.planId === undefined) {
                throw new Error("Missing required property 'planId'");
            }
            inputs["iamRoleArn"] = args ? args.iamRoleArn : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["planId"] = args ? args.planId : undefined;
            inputs["resources"] = args ? args.resources : undefined;
            inputs["selectionTags"] = args ? args.selectionTags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Selection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Selection resources.
 */
export interface SelectionState {
    /**
     * The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
     */
    readonly iamRoleArn?: pulumi.Input<string>;
    /**
     * The display name of a resource selection document.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The backup plan ID to be associated with the selection of resources.
     */
    readonly planId?: pulumi.Input<string>;
    /**
     * An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
     */
    readonly resources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tag-based conditions used to specify a set of resources to assign to a backup plan.
     */
    readonly selectionTags?: pulumi.Input<pulumi.Input<inputs.backup.SelectionSelectionTag>[]>;
}

/**
 * The set of arguments for constructing a Selection resource.
 */
export interface SelectionArgs {
    /**
     * The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
     */
    readonly iamRoleArn: pulumi.Input<string>;
    /**
     * The display name of a resource selection document.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The backup plan ID to be associated with the selection of resources.
     */
    readonly planId: pulumi.Input<string>;
    /**
     * An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
     */
    readonly resources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tag-based conditions used to specify a set of resources to assign to a backup plan.
     */
    readonly selectionTags?: pulumi.Input<pulumi.Input<inputs.backup.SelectionSelectionTag>[]>;
}
