// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class VpcAssociationAuthorization extends pulumi.CustomResource {
    /**
     * Get an existing VpcAssociationAuthorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcAssociationAuthorizationState, opts?: pulumi.CustomResourceOptions): VpcAssociationAuthorization {
        return new VpcAssociationAuthorization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization';

    /**
     * Returns true if the given object is an instance of VpcAssociationAuthorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcAssociationAuthorization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcAssociationAuthorization.__pulumiType;
    }

    public readonly vpcId!: pulumi.Output<string>;
    public readonly vpcRegion!: pulumi.Output<string>;
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a VpcAssociationAuthorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcAssociationAuthorizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcAssociationAuthorizationArgs | VpcAssociationAuthorizationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcAssociationAuthorizationState | undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["vpcRegion"] = state ? state.vpcRegion : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as VpcAssociationAuthorizationArgs | undefined;
            if (!args || args.vpcId === undefined) {
                throw new Error("Missing required property 'vpcId'");
            }
            if (!args || args.zoneId === undefined) {
                throw new Error("Missing required property 'zoneId'");
            }
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["vpcRegion"] = args ? args.vpcRegion : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcAssociationAuthorization.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcAssociationAuthorization resources.
 */
export interface VpcAssociationAuthorizationState {
    readonly vpcId?: pulumi.Input<string>;
    readonly vpcRegion?: pulumi.Input<string>;
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcAssociationAuthorization resource.
 */
export interface VpcAssociationAuthorizationArgs {
    readonly vpcId: pulumi.Input<string>;
    readonly vpcRegion?: pulumi.Input<string>;
    readonly zoneId: pulumi.Input<string>;
}
