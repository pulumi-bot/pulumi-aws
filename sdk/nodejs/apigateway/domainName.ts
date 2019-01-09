// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class DomainName extends pulumi.CustomResource {
    /**
     * Get an existing DomainName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainNameState, opts?: pulumi.CustomResourceOptions): DomainName {
        return new DomainName(name, <any>state, { ...opts, id: id });
    }

    public readonly certificateArn: pulumi.Output<string | undefined>;
    public readonly certificateBody: pulumi.Output<string | undefined>;
    public readonly certificateChain: pulumi.Output<string | undefined>;
    public readonly certificateName: pulumi.Output<string | undefined>;
    public readonly certificatePrivateKey: pulumi.Output<string | undefined>;
    public /*out*/ readonly certificateUploadDate: pulumi.Output<string>;
    public /*out*/ readonly cloudfrontDomainName: pulumi.Output<string>;
    public /*out*/ readonly cloudfrontZoneId: pulumi.Output<string>;
    public readonly domainName: pulumi.Output<string>;
    public readonly endpointConfiguration: pulumi.Output<{ types: string }>;
    public readonly regionalCertificateArn: pulumi.Output<string | undefined>;
    public readonly regionalCertificateName: pulumi.Output<string | undefined>;
    public /*out*/ readonly regionalDomainName: pulumi.Output<string>;
    public /*out*/ readonly regionalZoneId: pulumi.Output<string>;

    /**
     * Create a DomainName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainNameArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainNameArgs | DomainNameState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DomainNameState = argsOrState as DomainNameState | undefined;
            inputs["certificateArn"] = state ? state.certificateArn : undefined;
            inputs["certificateBody"] = state ? state.certificateBody : undefined;
            inputs["certificateChain"] = state ? state.certificateChain : undefined;
            inputs["certificateName"] = state ? state.certificateName : undefined;
            inputs["certificatePrivateKey"] = state ? state.certificatePrivateKey : undefined;
            inputs["certificateUploadDate"] = state ? state.certificateUploadDate : undefined;
            inputs["cloudfrontDomainName"] = state ? state.cloudfrontDomainName : undefined;
            inputs["cloudfrontZoneId"] = state ? state.cloudfrontZoneId : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["endpointConfiguration"] = state ? state.endpointConfiguration : undefined;
            inputs["regionalCertificateArn"] = state ? state.regionalCertificateArn : undefined;
            inputs["regionalCertificateName"] = state ? state.regionalCertificateName : undefined;
            inputs["regionalDomainName"] = state ? state.regionalDomainName : undefined;
            inputs["regionalZoneId"] = state ? state.regionalZoneId : undefined;
        } else {
            const args = argsOrState as DomainNameArgs | undefined;
            if (!args || args.domainName === undefined) {
                throw new Error("Missing required property 'domainName'");
            }
            inputs["certificateArn"] = args ? args.certificateArn : undefined;
            inputs["certificateBody"] = args ? args.certificateBody : undefined;
            inputs["certificateChain"] = args ? args.certificateChain : undefined;
            inputs["certificateName"] = args ? args.certificateName : undefined;
            inputs["certificatePrivateKey"] = args ? args.certificatePrivateKey : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["endpointConfiguration"] = args ? args.endpointConfiguration : undefined;
            inputs["regionalCertificateArn"] = args ? args.regionalCertificateArn : undefined;
            inputs["regionalCertificateName"] = args ? args.regionalCertificateName : undefined;
            inputs["certificateUploadDate"] = undefined /*out*/;
            inputs["cloudfrontDomainName"] = undefined /*out*/;
            inputs["cloudfrontZoneId"] = undefined /*out*/;
            inputs["regionalDomainName"] = undefined /*out*/;
            inputs["regionalZoneId"] = undefined /*out*/;
        }
        super("aws:apigateway/domainName:DomainName", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainName resources.
 */
export interface DomainNameState {
    readonly certificateArn?: pulumi.Input<string>;
    readonly certificateBody?: pulumi.Input<string>;
    readonly certificateChain?: pulumi.Input<string>;
    readonly certificateName?: pulumi.Input<string>;
    readonly certificatePrivateKey?: pulumi.Input<string>;
    readonly certificateUploadDate?: pulumi.Input<string>;
    readonly cloudfrontDomainName?: pulumi.Input<string>;
    readonly cloudfrontZoneId?: pulumi.Input<string>;
    readonly domainName?: pulumi.Input<string>;
    readonly endpointConfiguration?: pulumi.Input<{ types: pulumi.Input<string> }>;
    readonly regionalCertificateArn?: pulumi.Input<string>;
    readonly regionalCertificateName?: pulumi.Input<string>;
    readonly regionalDomainName?: pulumi.Input<string>;
    readonly regionalZoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DomainName resource.
 */
export interface DomainNameArgs {
    readonly certificateArn?: pulumi.Input<string>;
    readonly certificateBody?: pulumi.Input<string>;
    readonly certificateChain?: pulumi.Input<string>;
    readonly certificateName?: pulumi.Input<string>;
    readonly certificatePrivateKey?: pulumi.Input<string>;
    readonly domainName: pulumi.Input<string>;
    readonly endpointConfiguration?: pulumi.Input<{ types: pulumi.Input<string> }>;
    readonly regionalCertificateArn?: pulumi.Input<string>;
    readonly regionalCertificateName?: pulumi.Input<string>;
}
