// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an SSM Parameter resource.
 *
 * ## Example Usage
 *
 * To store a basic string parameter:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.ssm.Parameter("foo", {
 *     type: "String",
 *     value: "bar",
 * });
 * ```
 *
 * To store an encrypted string using the default SSM KMS key:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const _default = new aws.rds.Instance("default", {
 *     allocatedStorage: 10,
 *     storageType: "gp2",
 *     engine: "mysql",
 *     engineVersion: "5.7.16",
 *     instanceClass: "db.t2.micro",
 *     name: "mydb",
 *     username: "foo",
 *     password: _var.database_master_password,
 *     dbSubnetGroupName: "my_database_subnet_group",
 *     parameterGroupName: "default.mysql5.7",
 * });
 * const secret = new aws.ssm.Parameter("secret", {
 *     description: "The parameter description",
 *     type: "SecureString",
 *     value: _var.database_master_password,
 *     tags: {
 *         environment: "production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * SSM Parameters can be imported using the `parameter store name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:ssm/parameter:Parameter my_param /my_path/my_paramname
 * ```
 */
export class Parameter extends pulumi.CustomResource {
    /**
     * Get an existing Parameter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ParameterState, opts?: pulumi.CustomResourceOptions): Parameter {
        return new Parameter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ssm/parameter:Parameter';

    /**
     * Returns true if the given object is an instance of Parameter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Parameter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Parameter.__pulumiType;
    }

    /**
     * A regular expression used to validate the parameter value.
     */
    public readonly allowedPattern!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the parameter.
     */
    public readonly arn!: pulumi.Output<string>;
    /**
     * The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
     * ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
     */
    public readonly dataType!: pulumi.Output<string>;
    /**
     * The description of the parameter.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The KMS key id or arn for encrypting a SecureString.
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
     */
    public readonly overwrite!: pulumi.Output<boolean | undefined>;
    /**
     * A map of tags to assign to the object. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
     */
    public readonly tier!: pulumi.Output<string | undefined>;
    /**
     * The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * (Required) The value of the parameter.
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * The version of the parameter.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a Parameter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ParameterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ParameterArgs | ParameterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ParameterState | undefined;
            inputs["allowedPattern"] = state ? state.allowedPattern : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["dataType"] = state ? state.dataType : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["keyId"] = state ? state.keyId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["overwrite"] = state ? state.overwrite : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["value"] = state ? state.value : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ParameterArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            inputs["allowedPattern"] = args ? args.allowedPattern : undefined;
            inputs["arn"] = args ? args.arn : undefined;
            inputs["dataType"] = args ? args.dataType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["keyId"] = args ? args.keyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["overwrite"] = args ? args.overwrite : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Parameter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Parameter resources.
 */
export interface ParameterState {
    /**
     * A regular expression used to validate the parameter value.
     */
    allowedPattern?: pulumi.Input<string>;
    /**
     * The ARN of the parameter.
     */
    arn?: pulumi.Input<string>;
    /**
     * The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
     * ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
     */
    dataType?: pulumi.Input<string>;
    /**
     * The description of the parameter.
     */
    description?: pulumi.Input<string>;
    /**
     * The KMS key id or arn for encrypting a SecureString.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
     */
    name?: pulumi.Input<string>;
    /**
     * Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
     */
    overwrite?: pulumi.Input<boolean>;
    /**
     * A map of tags to assign to the object. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
     */
    tier?: pulumi.Input<string>;
    /**
     * The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
     */
    type?: pulumi.Input<string | enums.ssm.ParameterType>;
    /**
     * (Required) The value of the parameter.
     */
    value?: pulumi.Input<string>;
    /**
     * The version of the parameter.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Parameter resource.
 */
export interface ParameterArgs {
    /**
     * A regular expression used to validate the parameter value.
     */
    allowedPattern?: pulumi.Input<string>;
    /**
     * The ARN of the parameter.
     */
    arn?: pulumi.Input<string>;
    /**
     * The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
     * ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
     */
    dataType?: pulumi.Input<string>;
    /**
     * The description of the parameter.
     */
    description?: pulumi.Input<string>;
    /**
     * The KMS key id or arn for encrypting a SecureString.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
     */
    name?: pulumi.Input<string>;
    /**
     * Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
     */
    overwrite?: pulumi.Input<boolean>;
    /**
     * A map of tags to assign to the object. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
     */
    tier?: pulumi.Input<string>;
    /**
     * The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
     */
    type: pulumi.Input<string | enums.ssm.ParameterType>;
    /**
     * (Required) The value of the parameter.
     */
    value: pulumi.Input<string>;
}
