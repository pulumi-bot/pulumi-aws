// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Secret extends pulumi.CustomResource {
    /**
     * Get an existing Secret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretState, opts?: pulumi.CustomResourceOptions): Secret {
        return new Secret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:secretsmanager/secret:Secret';

    /**
     * Returns true if the given object is an instance of Secret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Secret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Secret.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string>;
    public readonly policy!: pulumi.Output<string | undefined>;
    public readonly recoveryWindowInDays!: pulumi.Output<number | undefined>;
    /**
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    public /*out*/ readonly rotationEnabled!: pulumi.Output<boolean>;
    /**
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    public readonly rotationLambdaArn!: pulumi.Output<string>;
    /**
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    public readonly rotationRules!: pulumi.Output<outputs.secretsmanager.SecretRotationRules>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Secret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretArgs | SecretState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["recoveryWindowInDays"] = state ? state.recoveryWindowInDays : undefined;
            inputs["rotationEnabled"] = state ? state.rotationEnabled : undefined;
            inputs["rotationLambdaArn"] = state ? state.rotationLambdaArn : undefined;
            inputs["rotationRules"] = state ? state.rotationRules : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SecretArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["recoveryWindowInDays"] = args ? args.recoveryWindowInDays : undefined;
            inputs["rotationLambdaArn"] = args ? args.rotationLambdaArn : undefined;
            inputs["rotationRules"] = args ? args.rotationRules : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["rotationEnabled"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Secret.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Secret resources.
 */
export interface SecretState {
    readonly arn?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly kmsKeyId?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly policy?: pulumi.Input<string>;
    readonly recoveryWindowInDays?: pulumi.Input<number>;
    /**
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    readonly rotationEnabled?: pulumi.Input<boolean>;
    /**
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    readonly rotationLambdaArn?: pulumi.Input<string>;
    /**
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    readonly rotationRules?: pulumi.Input<inputs.secretsmanager.SecretRotationRules>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Secret resource.
 */
export interface SecretArgs {
    readonly description?: pulumi.Input<string>;
    readonly kmsKeyId?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly policy?: pulumi.Input<string>;
    readonly recoveryWindowInDays?: pulumi.Input<number>;
    /**
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    readonly rotationLambdaArn?: pulumi.Input<string>;
    /**
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    readonly rotationRules?: pulumi.Input<inputs.secretsmanager.SecretRotationRules>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
