// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Image Builder Component.
 *
 * ## Example Usage
 * ### URI Document
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.imagebuilder.Component("example", {
 *     platform: "Linux",
 *     uri: pulumi.interpolate`s3://${aws_s3_bucket_object_example.bucket}/${aws_s3_bucket_object_example.key}`,
 *     version: "1.0.0",
 * });
 * ```
 *
 * ## Import
 *
 * `aws_imagebuilder_components` resources can be imported by using the Amazon Resource Name (ARN), e.g.
 *
 * ```sh
 *  $ pulumi import aws:imagebuilder/component:Component example arn:aws:imagebuilder:us-east-1:123456789012:component/example/1.0.0/1
 * ```
 *
 *  Certain resource arguments, such as `uri`, cannot be read via the API and imported into Terraform. Terraform will display a difference for these arguments the first run after import if declared in the Terraform configuration for an imported resource.
 */
export class Component extends pulumi.CustomResource {
    /**
     * Get an existing Component resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComponentState, opts?: pulumi.CustomResourceOptions): Component {
        return new Component(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:imagebuilder/component:Component';

    /**
     * Returns true if the given object is an instance of Component.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Component {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Component.__pulumiType;
    }

    /**
     * (Required) Amazon Resource Name (ARN) of the component.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Change description of the component.
     */
    public readonly changeDescription!: pulumi.Output<string | undefined>;
    public readonly data!: pulumi.Output<string>;
    /**
     * Date the component was created.
     */
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    /**
     * Description of the component.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Encryption status of the component.
     */
    public /*out*/ readonly encrypted!: pulumi.Output<boolean>;
    /**
     * Amazon Resource Name (ARN) of the Key Management Service (KMS) Key used to encrypt the component.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Name of the component.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Owner of the component.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * Platform of the component.
     */
    public readonly platform!: pulumi.Output<string>;
    /**
     * Set of Operating Systems (OS) supported by the component.
     */
    public readonly supportedOsVersions!: pulumi.Output<string[] | undefined>;
    /**
     * Key-value map of resource tags for the component.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Type of the component.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * S3 URI with data of the component. Exactly one of `data` and `uri` can be specified.
     */
    public readonly uri!: pulumi.Output<string | undefined>;
    /**
     * Version of the component.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Component resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComponentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComponentArgs | ComponentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComponentState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["changeDescription"] = state ? state.changeDescription : undefined;
            inputs["data"] = state ? state.data : undefined;
            inputs["dateCreated"] = state ? state.dateCreated : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["owner"] = state ? state.owner : undefined;
            inputs["platform"] = state ? state.platform : undefined;
            inputs["supportedOsVersions"] = state ? state.supportedOsVersions : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["uri"] = state ? state.uri : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ComponentArgs | undefined;
            if ((!args || args.platform === undefined) && !opts.urn) {
                throw new Error("Missing required property 'platform'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            inputs["changeDescription"] = args ? args.changeDescription : undefined;
            inputs["data"] = args ? args.data : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["platform"] = args ? args.platform : undefined;
            inputs["supportedOsVersions"] = args ? args.supportedOsVersions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["uri"] = args ? args.uri : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["dateCreated"] = undefined /*out*/;
            inputs["encrypted"] = undefined /*out*/;
            inputs["owner"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Component.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Component resources.
 */
export interface ComponentState {
    /**
     * (Required) Amazon Resource Name (ARN) of the component.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Change description of the component.
     */
    readonly changeDescription?: pulumi.Input<string>;
    readonly data?: pulumi.Input<string>;
    /**
     * Date the component was created.
     */
    readonly dateCreated?: pulumi.Input<string>;
    /**
     * Description of the component.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Encryption status of the component.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) of the Key Management Service (KMS) Key used to encrypt the component.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Name of the component.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Owner of the component.
     */
    readonly owner?: pulumi.Input<string>;
    /**
     * Platform of the component.
     */
    readonly platform?: pulumi.Input<string>;
    /**
     * Set of Operating Systems (OS) supported by the component.
     */
    readonly supportedOsVersions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value map of resource tags for the component.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of the component.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * S3 URI with data of the component. Exactly one of `data` and `uri` can be specified.
     */
    readonly uri?: pulumi.Input<string>;
    /**
     * Version of the component.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Component resource.
 */
export interface ComponentArgs {
    /**
     * Change description of the component.
     */
    readonly changeDescription?: pulumi.Input<string>;
    readonly data?: pulumi.Input<string>;
    /**
     * Description of the component.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the Key Management Service (KMS) Key used to encrypt the component.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Name of the component.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Platform of the component.
     */
    readonly platform: pulumi.Input<string>;
    /**
     * Set of Operating Systems (OS) supported by the component.
     */
    readonly supportedOsVersions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value map of resource tags for the component.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * S3 URI with data of the component. Exactly one of `data` and `uri` can be specified.
     */
    readonly uri?: pulumi.Input<string>;
    /**
     * Version of the component.
     */
    readonly version: pulumi.Input<string>;
}
