// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

import {Tags} from "../index";

/**
 * Provides an SSM Document resource
 * 
 * ~> **NOTE on updating SSM documents:** Only documents with a schema version of 2.0
 * or greater can update their content once created, see [SSM Schema Features][1]. To update a document with an older
 * schema version you must recreate the resource.
 */
export class Document extends pulumi.CustomResource {
    /**
     * Get an existing Document resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DocumentState): Document {
        return new Document(name, <any>state, { id });
    }

    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The JSON or YAML content of the document.
     */
    public readonly content: pulumi.Output<string>;
    /**
     * The date the document was created.
     */
    public /*out*/ readonly createdDate: pulumi.Output<string>;
    /**
     * The default version of the document.
     */
    public /*out*/ readonly defaultVersion: pulumi.Output<string>;
    /**
     * The description of the document.
     */
    public /*out*/ readonly description: pulumi.Output<string>;
    /**
     * The format of the document. Valid document types include: `JSON` and `YAML`
     */
    public readonly documentFormat: pulumi.Output<string | undefined>;
    /**
     * The type of the document. Valid document types include: `Command`, `Policy` and `Automation`
     */
    public readonly documentType: pulumi.Output<string>;
    /**
     * The sha1 or sha256 of the document content
     */
    public /*out*/ readonly hash: pulumi.Output<string>;
    /**
     * "Sha1" "Sha256". The hashing algorithm used when hashing the content.
     */
    public /*out*/ readonly hashType: pulumi.Output<string>;
    /**
     * The latest version of the document.
     */
    public /*out*/ readonly latestVersion: pulumi.Output<string>;
    /**
     * The name of the document.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The AWS user account of the person who created the document.
     */
    public /*out*/ readonly owner: pulumi.Output<string>;
    /**
     * The parameters that are available to this document.
     */
    public /*out*/ readonly parameters: pulumi.Output<{ defaultValue?: string, description?: string, name?: string, type?: string }[]>;
    /**
     * Additional Permissions to attach to the document. See Permissions below for details.
     */
    public readonly permissions: pulumi.Output<{ accountIds: string, type: string } | undefined>;
    /**
     * A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
     */
    public /*out*/ readonly platformTypes: pulumi.Output<string[]>;
    /**
     * The schema version of the document.
     */
    public /*out*/ readonly schemaVersion: pulumi.Output<string>;
    /**
     * "Creating", "Active" or "Deleting". The current status of the document.
     */
    public /*out*/ readonly status: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the object.
     */
    public readonly tags: pulumi.Output<Tags | undefined>;

    /**
     * Create a Document resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DocumentArgs | DocumentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DocumentState = argsOrState as DocumentState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["content"] = state ? state.content : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["defaultVersion"] = state ? state.defaultVersion : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["documentFormat"] = state ? state.documentFormat : undefined;
            inputs["documentType"] = state ? state.documentType : undefined;
            inputs["hash"] = state ? state.hash : undefined;
            inputs["hashType"] = state ? state.hashType : undefined;
            inputs["latestVersion"] = state ? state.latestVersion : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["owner"] = state ? state.owner : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["permissions"] = state ? state.permissions : undefined;
            inputs["platformTypes"] = state ? state.platformTypes : undefined;
            inputs["schemaVersion"] = state ? state.schemaVersion : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DocumentArgs | undefined;
            if (!args || args.content === undefined) {
                throw new Error("Missing required property 'content'");
            }
            if (!args || args.documentType === undefined) {
                throw new Error("Missing required property 'documentType'");
            }
            inputs["content"] = args ? args.content : undefined;
            inputs["documentFormat"] = args ? args.documentFormat : undefined;
            inputs["documentType"] = args ? args.documentType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["defaultVersion"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["hash"] = undefined /*out*/;
            inputs["hashType"] = undefined /*out*/;
            inputs["latestVersion"] = undefined /*out*/;
            inputs["owner"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["platformTypes"] = undefined /*out*/;
            inputs["schemaVersion"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        super("aws:ssm/document:Document", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Document resources.
 */
export interface DocumentState {
    readonly arn?: pulumi.Input<string>;
    /**
     * The JSON or YAML content of the document.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * The date the document was created.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * The default version of the document.
     */
    readonly defaultVersion?: pulumi.Input<string>;
    /**
     * The description of the document.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The format of the document. Valid document types include: `JSON` and `YAML`
     */
    readonly documentFormat?: pulumi.Input<string>;
    /**
     * The type of the document. Valid document types include: `Command`, `Policy` and `Automation`
     */
    readonly documentType?: pulumi.Input<string>;
    /**
     * The sha1 or sha256 of the document content
     */
    readonly hash?: pulumi.Input<string>;
    /**
     * "Sha1" "Sha256". The hashing algorithm used when hashing the content.
     */
    readonly hashType?: pulumi.Input<string>;
    /**
     * The latest version of the document.
     */
    readonly latestVersion?: pulumi.Input<string>;
    /**
     * The name of the document.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AWS user account of the person who created the document.
     */
    readonly owner?: pulumi.Input<string>;
    /**
     * The parameters that are available to this document.
     */
    readonly parameters?: pulumi.Input<pulumi.Input<{ defaultValue?: pulumi.Input<string>, description?: pulumi.Input<string>, name?: pulumi.Input<string>, type?: pulumi.Input<string> }>[]>;
    /**
     * Additional Permissions to attach to the document. See Permissions below for details.
     */
    readonly permissions?: pulumi.Input<{ accountIds: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
     */
    readonly platformTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The schema version of the document.
     */
    readonly schemaVersion?: pulumi.Input<string>;
    /**
     * "Creating", "Active" or "Deleting". The current status of the document.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<Tags>;
}

/**
 * The set of arguments for constructing a Document resource.
 */
export interface DocumentArgs {
    /**
     * The JSON or YAML content of the document.
     */
    readonly content: pulumi.Input<string>;
    /**
     * The format of the document. Valid document types include: `JSON` and `YAML`
     */
    readonly documentFormat?: pulumi.Input<string>;
    /**
     * The type of the document. Valid document types include: `Command`, `Policy` and `Automation`
     */
    readonly documentType: pulumi.Input<string>;
    /**
     * The name of the document.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Additional Permissions to attach to the document. See Permissions below for details.
     */
    readonly permissions?: pulumi.Input<{ accountIds: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<Tags>;
}
