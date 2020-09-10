// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Directory extends pulumi.CustomResource {
    /**
     * Get an existing Directory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DirectoryState, opts?: pulumi.CustomResourceOptions): Directory {
        return new Directory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:workspaces/directory:Directory';

    /**
     * Returns true if the given object is an instance of Directory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Directory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Directory.__pulumiType;
    }

    public /*out*/ readonly alias!: pulumi.Output<string>;
    public /*out*/ readonly customerUserName!: pulumi.Output<string>;
    public readonly directoryId!: pulumi.Output<string>;
    public /*out*/ readonly directoryName!: pulumi.Output<string>;
    public /*out*/ readonly directoryType!: pulumi.Output<string>;
    public /*out*/ readonly dnsIpAddresses!: pulumi.Output<string[]>;
    public /*out*/ readonly iamRoleId!: pulumi.Output<string>;
    public /*out*/ readonly ipGroupIds!: pulumi.Output<string[]>;
    public /*out*/ readonly registrationCode!: pulumi.Output<string>;
    public readonly selfServicePermissions!: pulumi.Output<outputs.workspaces.DirectorySelfServicePermissions>;
    public readonly subnetIds!: pulumi.Output<string[]>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly workspaceSecurityGroupId!: pulumi.Output<string>;

    /**
     * Create a Directory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DirectoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DirectoryArgs | DirectoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DirectoryState | undefined;
            inputs["alias"] = state ? state.alias : undefined;
            inputs["customerUserName"] = state ? state.customerUserName : undefined;
            inputs["directoryId"] = state ? state.directoryId : undefined;
            inputs["directoryName"] = state ? state.directoryName : undefined;
            inputs["directoryType"] = state ? state.directoryType : undefined;
            inputs["dnsIpAddresses"] = state ? state.dnsIpAddresses : undefined;
            inputs["iamRoleId"] = state ? state.iamRoleId : undefined;
            inputs["ipGroupIds"] = state ? state.ipGroupIds : undefined;
            inputs["registrationCode"] = state ? state.registrationCode : undefined;
            inputs["selfServicePermissions"] = state ? state.selfServicePermissions : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["workspaceSecurityGroupId"] = state ? state.workspaceSecurityGroupId : undefined;
        } else {
            const args = argsOrState as DirectoryArgs | undefined;
            if (!args || args.directoryId === undefined) {
                throw new Error("Missing required property 'directoryId'");
            }
            inputs["directoryId"] = args ? args.directoryId : undefined;
            inputs["selfServicePermissions"] = args ? args.selfServicePermissions : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["alias"] = undefined /*out*/;
            inputs["customerUserName"] = undefined /*out*/;
            inputs["directoryName"] = undefined /*out*/;
            inputs["directoryType"] = undefined /*out*/;
            inputs["dnsIpAddresses"] = undefined /*out*/;
            inputs["iamRoleId"] = undefined /*out*/;
            inputs["ipGroupIds"] = undefined /*out*/;
            inputs["registrationCode"] = undefined /*out*/;
            inputs["workspaceSecurityGroupId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Directory.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Directory resources.
 */
export interface DirectoryState {
    readonly alias?: pulumi.Input<string>;
    readonly customerUserName?: pulumi.Input<string>;
    readonly directoryId?: pulumi.Input<string>;
    readonly directoryName?: pulumi.Input<string>;
    readonly directoryType?: pulumi.Input<string>;
    readonly dnsIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly iamRoleId?: pulumi.Input<string>;
    readonly ipGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly registrationCode?: pulumi.Input<string>;
    readonly selfServicePermissions?: pulumi.Input<inputs.workspaces.DirectorySelfServicePermissions>;
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly workspaceSecurityGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Directory resource.
 */
export interface DirectoryArgs {
    readonly directoryId: pulumi.Input<string>;
    readonly selfServicePermissions?: pulumi.Input<inputs.workspaces.DirectorySelfServicePermissions>;
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
