// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class OriginAccessIdentity extends pulumi.CustomResource {
    /**
     * Get an existing OriginAccessIdentity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OriginAccessIdentityState, opts?: pulumi.CustomResourceOptions): OriginAccessIdentity {
        return new OriginAccessIdentity(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly callerReference: pulumi.Output<string>;
    public /*out*/ readonly cloudfrontAccessIdentityPath: pulumi.Output<string>;
    public readonly comment: pulumi.Output<string | undefined>;
    public /*out*/ readonly etag: pulumi.Output<string>;
    public /*out*/ readonly iamArn: pulumi.Output<string>;
    public /*out*/ readonly s3CanonicalUserId: pulumi.Output<string>;

    /**
     * Create a OriginAccessIdentity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OriginAccessIdentityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OriginAccessIdentityArgs | OriginAccessIdentityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: OriginAccessIdentityState = argsOrState as OriginAccessIdentityState | undefined;
            inputs["callerReference"] = state ? state.callerReference : undefined;
            inputs["cloudfrontAccessIdentityPath"] = state ? state.cloudfrontAccessIdentityPath : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["iamArn"] = state ? state.iamArn : undefined;
            inputs["s3CanonicalUserId"] = state ? state.s3CanonicalUserId : undefined;
        } else {
            const args = argsOrState as OriginAccessIdentityArgs | undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["callerReference"] = undefined /*out*/;
            inputs["cloudfrontAccessIdentityPath"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["iamArn"] = undefined /*out*/;
            inputs["s3CanonicalUserId"] = undefined /*out*/;
        }
        super("aws:cloudfront/originAccessIdentity:OriginAccessIdentity", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OriginAccessIdentity resources.
 */
export interface OriginAccessIdentityState {
    readonly callerReference?: pulumi.Input<string>;
    readonly cloudfrontAccessIdentityPath?: pulumi.Input<string>;
    readonly comment?: pulumi.Input<string>;
    readonly etag?: pulumi.Input<string>;
    readonly iamArn?: pulumi.Input<string>;
    readonly s3CanonicalUserId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OriginAccessIdentity resource.
 */
export interface OriginAccessIdentityArgs {
    readonly comment?: pulumi.Input<string>;
}
