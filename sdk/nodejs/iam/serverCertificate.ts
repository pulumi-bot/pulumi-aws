// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ServerCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ServerCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerCertificateState, opts?: pulumi.CustomResourceOptions): ServerCertificate {
        return new ServerCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/serverCertificate:ServerCertificate';

    /**
     * Returns true if the given object is an instance of ServerCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerCertificate.__pulumiType;
    }

    public readonly arn!: pulumi.Output<string>;
    public readonly certificateBody!: pulumi.Output<string>;
    public readonly certificateChain!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    public readonly path!: pulumi.Output<string | undefined>;
    public readonly privateKey!: pulumi.Output<string>;

    /**
     * Create a ServerCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerCertificateArgs | ServerCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServerCertificateState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["certificateBody"] = state ? state.certificateBody : undefined;
            inputs["certificateChain"] = state ? state.certificateChain : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
        } else {
            const args = argsOrState as ServerCertificateArgs | undefined;
            if (!args || args.certificateBody === undefined) {
                throw new Error("Missing required property 'certificateBody'");
            }
            if (!args || args.privateKey === undefined) {
                throw new Error("Missing required property 'privateKey'");
            }
            inputs["arn"] = args ? args.arn : undefined;
            inputs["certificateBody"] = args ? args.certificateBody : undefined;
            inputs["certificateChain"] = args ? args.certificateChain : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServerCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerCertificate resources.
 */
export interface ServerCertificateState {
    readonly arn?: pulumi.Input<string>;
    readonly certificateBody?: pulumi.Input<string>;
    readonly certificateChain?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly path?: pulumi.Input<string>;
    readonly privateKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerCertificate resource.
 */
export interface ServerCertificateArgs {
    readonly arn?: pulumi.Input<string>;
    readonly certificateBody: pulumi.Input<string>;
    readonly certificateChain?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly path?: pulumi.Input<string>;
    readonly privateKey: pulumi.Input<string>;
}
