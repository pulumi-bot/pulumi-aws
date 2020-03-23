// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Registers an on-premises server or virtual machine with Amazon EC2 so that it can be managed using Run Command.
 * 
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const testRole = new aws.iam.Role("testRole", {
 *     assumeRolePolicy: `  {
 *     "Version": "2012-10-17",
 *     "Statement": {
 *       "Effect": "Allow",
 *       "Principal": {"Service": "ssm.amazonaws.com"},
 *       "Action": "sts:AssumeRole"
 *     }
 *   }
 * `,
 * });
 * const testAttach = new aws.iam.RolePolicyAttachment("testAttach", {
 *     policyArn: "arn:aws:iam::aws:policy/service-role/AmazonEC2RoleforSSM",
 *     role: testRole.name,
 * });
 * const foo = new aws.ssm.Activation("foo", {
 *     description: "Test",
 *     iamRole: testRole.id,
 *     registrationLimit: 5,
 * }, {dependsOn: [testAttach]});
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_activation.html.markdown.
 */
export class Activation extends pulumi.CustomResource {
    /**
     * Get an existing Activation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActivationState, opts?: pulumi.CustomResourceOptions): Activation {
        return new Activation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ssm/activation:Activation';

    /**
     * Returns true if the given object is an instance of Activation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Activation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Activation.__pulumiType;
    }

    /**
     * The code the system generates when it processes the activation.
     */
    public /*out*/ readonly activationCode!: pulumi.Output<string>;
    /**
     * The description of the resource that you want to register.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The date by which this activation request should expire. The default value is 24 hours.
     */
    public readonly expirationDate!: pulumi.Output<string>;
    /**
     * If the current activation has expired.
     */
    public /*out*/ readonly expired!: pulumi.Output<string>;
    /**
     * The IAM Role to attach to the managed instance.
     */
    public readonly iamRole!: pulumi.Output<string>;
    /**
     * The default name of the registered managed instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of managed instances that are currently registered using this activation.
     */
    public /*out*/ readonly registrationCount!: pulumi.Output<number>;
    /**
     * The maximum number of managed instances you want to register. The default value is 1 instance.
     */
    public readonly registrationLimit!: pulumi.Output<number | undefined>;
    /**
     * A mapping of tags to assign to the object.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Activation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActivationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActivationArgs | ActivationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ActivationState | undefined;
            inputs["activationCode"] = state ? state.activationCode : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["expirationDate"] = state ? state.expirationDate : undefined;
            inputs["expired"] = state ? state.expired : undefined;
            inputs["iamRole"] = state ? state.iamRole : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["registrationCount"] = state ? state.registrationCount : undefined;
            inputs["registrationLimit"] = state ? state.registrationLimit : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ActivationArgs | undefined;
            if (!args || args.iamRole === undefined) {
                throw new Error("Missing required property 'iamRole'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["expirationDate"] = args ? args.expirationDate : undefined;
            inputs["iamRole"] = args ? args.iamRole : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["registrationLimit"] = args ? args.registrationLimit : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["activationCode"] = undefined /*out*/;
            inputs["expired"] = undefined /*out*/;
            inputs["registrationCount"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Activation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Activation resources.
 */
export interface ActivationState {
    /**
     * The code the system generates when it processes the activation.
     */
    readonly activationCode?: pulumi.Input<string>;
    /**
     * The description of the resource that you want to register.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The date by which this activation request should expire. The default value is 24 hours.
     */
    readonly expirationDate?: pulumi.Input<string>;
    /**
     * If the current activation has expired.
     */
    readonly expired?: pulumi.Input<string>;
    /**
     * The IAM Role to attach to the managed instance.
     */
    readonly iamRole?: pulumi.Input<string>;
    /**
     * The default name of the registered managed instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of managed instances that are currently registered using this activation.
     */
    readonly registrationCount?: pulumi.Input<number>;
    /**
     * The maximum number of managed instances you want to register. The default value is 1 instance.
     */
    readonly registrationLimit?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Activation resource.
 */
export interface ActivationArgs {
    /**
     * The description of the resource that you want to register.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The date by which this activation request should expire. The default value is 24 hours.
     */
    readonly expirationDate?: pulumi.Input<string>;
    /**
     * The IAM Role to attach to the managed instance.
     */
    readonly iamRole: pulumi.Input<string>;
    /**
     * The default name of the registered managed instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The maximum number of managed instances you want to register. The default value is 1 instance.
     */
    readonly registrationLimit?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
