// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage the accepter's side of a Direct Connect hosted transit virtual interface.
 * This resource accepts ownership of a transit virtual interface created by another AWS account.
 * 
 * > **NOTE:** AWS allows a Direct Connect hosted transit virtual interface to be deleted from either the allocator's or accepter's side. However, this provider only allows the Direct Connect hosted transit virtual interface to be deleted from the allocator's side by removing the corresponding `aws.directconnect.HostedTransitVirtualInterface` resource from your configuration. Removing a `aws.directconnect.HostedTransitVirtualInterfaceAcceptor` resource from your configuration will remove it from your statefile and management, **but will not delete the Direct Connect virtual interface.**
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const accepter = new aws.Provider("accepter", {});
 * const accepterCallerIdentity = pulumi.output(aws.getCallerIdentity({ provider: accepter, async: true }));
 * // Accepter's side of the VIF.
 * const example = new aws.directconnect.Gateway("example", {
 *     amazonSideAsn: "64512",
 *     name: "tf-dxg-example",
 * }, { provider: accepter });
 * // Creator's side of the VIF
 * const creator = new aws.directconnect.HostedTransitVirtualInterface("creator", {
 *     addressFamily: "ipv4",
 *     bgpAsn: 65352,
 *     connectionId: "dxcon-zzzzzzzz",
 *     name: "tf-transit-vif-example",
 *     ownerAccountId: accepterCallerIdentity.accountId,
 *     vlan: 4094,
 * }, { dependsOn: [example] });
 * const accepterHostedTransitVirtualInterfaceAcceptor = new aws.directconnect.HostedTransitVirtualInterfaceAcceptor("accepter", {
 *     dxGatewayId: example.id,
 *     tags: {
 *         Side: "Accepter",
 *     },
 *     virtualInterfaceId: creator.id,
 * }, { provider: accepter });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_hosted_transit_virtual_interface_accepter.html.markdown.
 */
export class HostedTransitVirtualInterfaceAcceptor extends pulumi.CustomResource {
    /**
     * Get an existing HostedTransitVirtualInterfaceAcceptor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostedTransitVirtualInterfaceAcceptorState, opts?: pulumi.CustomResourceOptions): HostedTransitVirtualInterfaceAcceptor {
        return new HostedTransitVirtualInterfaceAcceptor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor';

    /**
     * Returns true if the given object is an instance of HostedTransitVirtualInterfaceAcceptor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostedTransitVirtualInterfaceAcceptor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostedTransitVirtualInterfaceAcceptor.__pulumiType;
    }

    /**
     * The ARN of the virtual interface.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    public readonly dxGatewayId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the Direct Connect virtual interface to accept.
     */
    public readonly virtualInterfaceId!: pulumi.Output<string>;

    /**
     * Create a HostedTransitVirtualInterfaceAcceptor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostedTransitVirtualInterfaceAcceptorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostedTransitVirtualInterfaceAcceptorArgs | HostedTransitVirtualInterfaceAcceptorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HostedTransitVirtualInterfaceAcceptorState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["dxGatewayId"] = state ? state.dxGatewayId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualInterfaceId"] = state ? state.virtualInterfaceId : undefined;
        } else {
            const args = argsOrState as HostedTransitVirtualInterfaceAcceptorArgs | undefined;
            if (!args || args.dxGatewayId === undefined) {
                throw new Error("Missing required property 'dxGatewayId'");
            }
            if (!args || args.virtualInterfaceId === undefined) {
                throw new Error("Missing required property 'virtualInterfaceId'");
            }
            inputs["dxGatewayId"] = args ? args.dxGatewayId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualInterfaceId"] = args ? args.virtualInterfaceId : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(HostedTransitVirtualInterfaceAcceptor.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostedTransitVirtualInterfaceAcceptor resources.
 */
export interface HostedTransitVirtualInterfaceAcceptorState {
    /**
     * The ARN of the virtual interface.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    readonly dxGatewayId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the Direct Connect virtual interface to accept.
     */
    readonly virtualInterfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HostedTransitVirtualInterfaceAcceptor resource.
 */
export interface HostedTransitVirtualInterfaceAcceptorArgs {
    /**
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    readonly dxGatewayId: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the Direct Connect virtual interface to accept.
     */
    readonly virtualInterfaceId: pulumi.Input<string>;
}
