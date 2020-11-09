// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage a GuardDuty IPSet.
 *
 * > **Note:** Currently in GuardDuty, users from member accounts cannot upload and further manage IPSets. IPSets that are uploaded by the primary account are imposed on GuardDuty functionality in its member accounts. See the [GuardDuty API Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/create-ip-set.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const primary = new aws.guardduty.Detector("primary", {enable: true});
 * const bucket = new aws.s3.Bucket("bucket", {acl: "private"});
 * const myIPSet = new aws.s3.BucketObject("myIPSet", {
 *     acl: "public-read",
 *     content: "10.0.0.0/8\n",
 *     bucket: bucket.id,
 *     key: "MyIPSet",
 * });
 * const example = new aws.guardduty.IPSet("example", {
 *     activate: true,
 *     detectorId: primary.id,
 *     format: "TXT",
 *     location: pulumi.interpolate`https://s3.amazonaws.com/${myIPSet.bucket}/${myIPSet.key}`,
 * });
 * ```
 *
 * ## Import
 *
 * GuardDuty IPSet can be imported using the the primary GuardDuty detector ID and IPSet ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:guardduty/iPSet:IPSet MyIPSet 00b00fd5aecc0ab60a708659477e9617:123456789012
 * ```
 */
export class IPSet extends pulumi.CustomResource {
    /**
     * Get an existing IPSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IPSetState, opts?: pulumi.CustomResourceOptions): IPSet {
        return new IPSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:guardduty/iPSet:IPSet';

    /**
     * Returns true if the given object is an instance of IPSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IPSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IPSet.__pulumiType;
    }

    /**
     * Specifies whether GuardDuty is to start using the uploaded IPSet.
     */
    public readonly activate!: pulumi.Output<boolean>;
    /**
     * Amazon Resource Name (ARN) of the GuardDuty IPSet.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The detector ID of the GuardDuty.
     */
    public readonly detectorId!: pulumi.Output<string>;
    /**
     * The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * The URI of the file that contains the IPSet.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The friendly name to identify the IPSet.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a IPSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IPSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IPSetArgs | IPSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IPSetState | undefined;
            inputs["activate"] = state ? state.activate : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["detectorId"] = state ? state.detectorId : undefined;
            inputs["format"] = state ? state.format : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as IPSetArgs | undefined;
            if (!args || args.activate === undefined) {
                throw new Error("Missing required property 'activate'");
            }
            if (!args || args.detectorId === undefined) {
                throw new Error("Missing required property 'detectorId'");
            }
            if (!args || args.format === undefined) {
                throw new Error("Missing required property 'format'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            inputs["activate"] = args ? args.activate : undefined;
            inputs["detectorId"] = args ? args.detectorId : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IPSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IPSet resources.
 */
export interface IPSetState {
    /**
     * Specifies whether GuardDuty is to start using the uploaded IPSet.
     */
    readonly activate?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) of the GuardDuty IPSet.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The detector ID of the GuardDuty.
     */
    readonly detectorId?: pulumi.Input<string>;
    /**
     * The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
     */
    readonly format?: pulumi.Input<string>;
    /**
     * The URI of the file that contains the IPSet.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The friendly name to identify the IPSet.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a IPSet resource.
 */
export interface IPSetArgs {
    /**
     * Specifies whether GuardDuty is to start using the uploaded IPSet.
     */
    readonly activate: pulumi.Input<boolean>;
    /**
     * The detector ID of the GuardDuty.
     */
    readonly detectorId: pulumi.Input<string>;
    /**
     * The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
     */
    readonly format: pulumi.Input<string>;
    /**
     * The URI of the file that contains the IPSet.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The friendly name to identify the IPSet.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
