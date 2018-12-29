// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Directory extends pulumi.CustomResource {
    /**
     * Get an existing Directory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DirectoryState, opts?: pulumi.CustomResourceOptions): Directory {
        return new Directory(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly accessUrl: pulumi.Output<string>;
    public readonly alias: pulumi.Output<string>;
    public readonly connectSettings: pulumi.Output<{ customerDnsIps: string[], customerUsername: string, subnetIds: string[], vpcId: string } | undefined>;
    public readonly description: pulumi.Output<string | undefined>;
    public /*out*/ readonly dnsIpAddresses: pulumi.Output<string[]>;
    public readonly edition: pulumi.Output<string>;
    public readonly enableSso: pulumi.Output<boolean | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly password: pulumi.Output<string>;
    public /*out*/ readonly securityGroupId: pulumi.Output<string>;
    public readonly shortName: pulumi.Output<string>;
    public readonly size: pulumi.Output<string>;
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly type: pulumi.Output<string | undefined>;
    public readonly vpcSettings: pulumi.Output<{ subnetIds: string[], vpcId: string } | undefined>;

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
            const state: DirectoryState = argsOrState as DirectoryState | undefined;
            inputs["accessUrl"] = state ? state.accessUrl : undefined;
            inputs["alias"] = state ? state.alias : undefined;
            inputs["connectSettings"] = state ? state.connectSettings : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dnsIpAddresses"] = state ? state.dnsIpAddresses : undefined;
            inputs["edition"] = state ? state.edition : undefined;
            inputs["enableSso"] = state ? state.enableSso : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["shortName"] = state ? state.shortName : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["vpcSettings"] = state ? state.vpcSettings : undefined;
        } else {
            const args = argsOrState as DirectoryArgs | undefined;
            if (!args || args.password === undefined) {
                throw new Error("Missing required property 'password'");
            }
            inputs["alias"] = args ? args.alias : undefined;
            inputs["connectSettings"] = args ? args.connectSettings : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["edition"] = args ? args.edition : undefined;
            inputs["enableSso"] = args ? args.enableSso : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["shortName"] = args ? args.shortName : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["vpcSettings"] = args ? args.vpcSettings : undefined;
            inputs["accessUrl"] = undefined /*out*/;
            inputs["dnsIpAddresses"] = undefined /*out*/;
            inputs["securityGroupId"] = undefined /*out*/;
        }
        super("aws:directoryservice/directory:Directory", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Directory resources.
 */
export interface DirectoryState {
    readonly accessUrl?: pulumi.Input<string>;
    readonly alias?: pulumi.Input<string>;
    readonly connectSettings?: pulumi.Input<{ customerDnsIps: pulumi.Input<pulumi.Input<string>[]>, customerUsername: pulumi.Input<string>, subnetIds: pulumi.Input<pulumi.Input<string>[]>, vpcId: pulumi.Input<string> }>;
    readonly description?: pulumi.Input<string>;
    readonly dnsIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly edition?: pulumi.Input<string>;
    readonly enableSso?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly password?: pulumi.Input<string>;
    readonly securityGroupId?: pulumi.Input<string>;
    readonly shortName?: pulumi.Input<string>;
    readonly size?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly type?: pulumi.Input<string>;
    readonly vpcSettings?: pulumi.Input<{ subnetIds: pulumi.Input<pulumi.Input<string>[]>, vpcId: pulumi.Input<string> }>;
}

/**
 * The set of arguments for constructing a Directory resource.
 */
export interface DirectoryArgs {
    readonly alias?: pulumi.Input<string>;
    readonly connectSettings?: pulumi.Input<{ customerDnsIps: pulumi.Input<pulumi.Input<string>[]>, customerUsername: pulumi.Input<string>, subnetIds: pulumi.Input<pulumi.Input<string>[]>, vpcId: pulumi.Input<string> }>;
    readonly description?: pulumi.Input<string>;
    readonly edition?: pulumi.Input<string>;
    readonly enableSso?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly password: pulumi.Input<string>;
    readonly shortName?: pulumi.Input<string>;
    readonly size?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly type?: pulumi.Input<string>;
    readonly vpcSettings?: pulumi.Input<{ subnetIds: pulumi.Input<pulumi.Input<string>[]>, vpcId: pulumi.Input<string> }>;
}
