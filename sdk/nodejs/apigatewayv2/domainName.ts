// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class DomainName extends pulumi.CustomResource {
    /**
     * Get an existing DomainName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainNameState, opts?: pulumi.CustomResourceOptions): DomainName {
        return new DomainName(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigatewayv2/domainName:DomainName';

    /**
     * Returns true if the given object is an instance of DomainName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainName.__pulumiType;
    }

    public /*out*/ readonly apiMappingSelectionExpression!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly domainName!: pulumi.Output<string>;
    public readonly domainNameConfiguration!: pulumi.Output<outputs.apigatewayv2.DomainNameDomainNameConfiguration>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

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
            const state = argsOrState as DomainNameState | undefined;
            inputs["apiMappingSelectionExpression"] = state ? state.apiMappingSelectionExpression : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["domainNameConfiguration"] = state ? state.domainNameConfiguration : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DomainNameArgs | undefined;
            if (!args || args.domainName === undefined) {
                throw new Error("Missing required property 'domainName'");
            }
            if (!args || args.domainNameConfiguration === undefined) {
                throw new Error("Missing required property 'domainNameConfiguration'");
            }
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["domainNameConfiguration"] = args ? args.domainNameConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["apiMappingSelectionExpression"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DomainName.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainName resources.
 */
export interface DomainNameState {
    readonly apiMappingSelectionExpression?: pulumi.Input<string>;
    readonly arn?: pulumi.Input<string>;
    readonly domainName?: pulumi.Input<string>;
    readonly domainNameConfiguration?: pulumi.Input<inputs.apigatewayv2.DomainNameDomainNameConfiguration>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a DomainName resource.
 */
export interface DomainNameArgs {
    readonly domainName: pulumi.Input<string>;
    readonly domainNameConfiguration: pulumi.Input<inputs.apigatewayv2.DomainNameDomainNameConfiguration>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
