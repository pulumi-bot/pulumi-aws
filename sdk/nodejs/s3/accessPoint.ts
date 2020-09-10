// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class AccessPoint extends pulumi.CustomResource {
    /**
     * Get an existing AccessPoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessPointState, opts?: pulumi.CustomResourceOptions): AccessPoint {
        return new AccessPoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3/accessPoint:AccessPoint';

    /**
     * Returns true if the given object is an instance of AccessPoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPoint.__pulumiType;
    }

    public readonly accountId!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly bucket!: pulumi.Output<string>;
    public /*out*/ readonly domainName!: pulumi.Output<string>;
    public /*out*/ readonly hasPublicAccessPolicy!: pulumi.Output<boolean>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly networkOrigin!: pulumi.Output<string>;
    public readonly policy!: pulumi.Output<string | undefined>;
    public readonly publicAccessBlockConfiguration!: pulumi.Output<outputs.s3.AccessPointPublicAccessBlockConfiguration | undefined>;
    public readonly vpcConfiguration!: pulumi.Output<outputs.s3.AccessPointVpcConfiguration | undefined>;

    /**
     * Create a AccessPoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessPointArgs | AccessPointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccessPointState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["hasPublicAccessPolicy"] = state ? state.hasPublicAccessPolicy : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkOrigin"] = state ? state.networkOrigin : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["publicAccessBlockConfiguration"] = state ? state.publicAccessBlockConfiguration : undefined;
            inputs["vpcConfiguration"] = state ? state.vpcConfiguration : undefined;
        } else {
            const args = argsOrState as AccessPointArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["publicAccessBlockConfiguration"] = args ? args.publicAccessBlockConfiguration : undefined;
            inputs["vpcConfiguration"] = args ? args.vpcConfiguration : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["domainName"] = undefined /*out*/;
            inputs["hasPublicAccessPolicy"] = undefined /*out*/;
            inputs["networkOrigin"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AccessPoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessPoint resources.
 */
export interface AccessPointState {
    readonly accountId?: pulumi.Input<string>;
    readonly arn?: pulumi.Input<string>;
    readonly bucket?: pulumi.Input<string>;
    readonly domainName?: pulumi.Input<string>;
    readonly hasPublicAccessPolicy?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly networkOrigin?: pulumi.Input<string>;
    readonly policy?: pulumi.Input<string>;
    readonly publicAccessBlockConfiguration?: pulumi.Input<inputs.s3.AccessPointPublicAccessBlockConfiguration>;
    readonly vpcConfiguration?: pulumi.Input<inputs.s3.AccessPointVpcConfiguration>;
}

/**
 * The set of arguments for constructing a AccessPoint resource.
 */
export interface AccessPointArgs {
    readonly accountId?: pulumi.Input<string>;
    readonly bucket: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly policy?: pulumi.Input<string>;
    readonly publicAccessBlockConfiguration?: pulumi.Input<inputs.s3.AccessPointPublicAccessBlockConfiguration>;
    readonly vpcConfiguration?: pulumi.Input<inputs.s3.AccessPointVpcConfiguration>;
}
