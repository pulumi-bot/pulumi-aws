// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an SSM Document resource
 *
 * > **NOTE on updating SSM documents:** Only documents with a schema version of 2.0
 * or greater can update their content once created, see [SSM Schema Features](http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-ssm-docs.html#document-schemas-features). To update a document with an older
 * schema version you must recreate the resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.ssm.Document("foo", {
 *     content: `  {
 *     "schemaVersion": "1.2",
 *     "description": "Check ip configuration of a Linux instance.",
 *     "parameters": {
 *
 *     },
 *     "runtimeConfig": {
 *       "aws:runShellScript": {
 *         "properties": [
 *           {
 *             "id": "0.aws:runShellScript",
 *             "runCommand": ["ifconfig"]
 *           }
 *         ]
 *       }
 *     }
 *   }
 * `,
 *     documentType: "Command",
 * });
 * ```
 * ## Permissions
 *
 * The permissions attribute specifies how you want to share the document. If you share a document privately,
 * you must specify the AWS user account IDs for those people who can use the document. If you share a document
 * publicly, you must specify All as the account ID.
 *
 * The permissions mapping supports the following:
 *
 * * `type` - The permission type for the document. The permission type can be `Share`.
 * * `accountIds` - The AWS user accounts that should have access to the document. The account IDs can either be a group of account IDs or `All`.
 *
 * ## Import
 *
 * SSM Documents can be imported using the name, e.g.
 *
 * ```sh
 *  $ pulumi import aws:ssm/document:Document example example
 * ```
 *
 *  The `attachments_source` argument does not have an SSM API method for reading the attachment information detail after creation. If the argument is set in the provider configuration on an imported resource, this provider will always show a difference. To workaround this behavior, either omit the argument from the configuration or use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to hide the difference, e.g. hcl resource "aws_ssm_document" "test" {
 *
 *  name
 *
 * = "test_document"
 *
 *  document_type = "Package"
 *
 *  attachments_source {
 *
 *  key
 *
 * = "SourceUrl"
 *
 *  values = ["s3://${aws_s3_bucket.object_bucket.bucket}/test.zip"]
 *
 *  }
 *
 * # There is no AWS SSM API for reading attachments_source info directly
 *
 *  lifecycle {
 *
 *  ignore_changes = [attachments_source]
 *
 *  } }
 */
export class Document extends pulumi.CustomResource {
    /**
     * Get an existing Document resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DocumentState, opts?: pulumi.CustomResourceOptions): Document {
        return new Document(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ssm/document:Document';

    /**
     * Returns true if the given object is an instance of Document.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Document {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Document.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * One or more configuration blocks describing attachments sources to a version of a document. Defined below.
     */
    public readonly attachmentsSources!: pulumi.Output<outputs.ssm.DocumentAttachmentsSource[] | undefined>;
    /**
     * The JSON or YAML content of the document.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * The date the document was created.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * The default version of the document.
     */
    public /*out*/ readonly defaultVersion!: pulumi.Output<string>;
    /**
     * The description of the document.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * The format of the document. Valid document types include: `JSON` and `YAML`
     */
    public readonly documentFormat!: pulumi.Output<string | undefined>;
    /**
     * The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
     */
    public readonly documentType!: pulumi.Output<string>;
    /**
     * The document version.
     */
    public /*out*/ readonly documentVersion!: pulumi.Output<string>;
    /**
     * The sha1 or sha256 of the document content
     */
    public /*out*/ readonly hash!: pulumi.Output<string>;
    /**
     * "Sha1" "Sha256". The hashing algorithm used when hashing the content.
     */
    public /*out*/ readonly hashType!: pulumi.Output<string>;
    /**
     * The latest version of the document.
     */
    public /*out*/ readonly latestVersion!: pulumi.Output<string>;
    /**
     * The name of the document.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The AWS user account of the person who created the document.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * The parameters that are available to this document.
     */
    public /*out*/ readonly parameters!: pulumi.Output<outputs.ssm.DocumentParameter[]>;
    /**
     * Additional Permissions to attach to the document. See Permissions below for details.
     */
    public readonly permissions!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
     */
    public /*out*/ readonly platformTypes!: pulumi.Output<string[]>;
    /**
     * The schema version of the document.
     */
    public /*out*/ readonly schemaVersion!: pulumi.Output<string>;
    /**
     * "Creating", "Active" or "Deleting". The current status of the document.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the object.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
     */
    public readonly targetType!: pulumi.Output<string | undefined>;
    /**
     * A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
     */
    public readonly versionName!: pulumi.Output<string | undefined>;

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
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DocumentState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["attachmentsSources"] = state ? state.attachmentsSources : undefined;
            inputs["content"] = state ? state.content : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["defaultVersion"] = state ? state.defaultVersion : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["documentFormat"] = state ? state.documentFormat : undefined;
            inputs["documentType"] = state ? state.documentType : undefined;
            inputs["documentVersion"] = state ? state.documentVersion : undefined;
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
            inputs["targetType"] = state ? state.targetType : undefined;
            inputs["versionName"] = state ? state.versionName : undefined;
        } else {
            const args = argsOrState as DocumentArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.documentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'documentType'");
            }
            inputs["attachmentsSources"] = args ? args.attachmentsSources : undefined;
            inputs["content"] = args ? args.content : undefined;
            inputs["documentFormat"] = args ? args.documentFormat : undefined;
            inputs["documentType"] = args ? args.documentType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetType"] = args ? args.targetType : undefined;
            inputs["versionName"] = args ? args.versionName : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["defaultVersion"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["documentVersion"] = undefined /*out*/;
            inputs["hash"] = undefined /*out*/;
            inputs["hashType"] = undefined /*out*/;
            inputs["latestVersion"] = undefined /*out*/;
            inputs["owner"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["platformTypes"] = undefined /*out*/;
            inputs["schemaVersion"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Document.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Document resources.
 */
export interface DocumentState {
    readonly arn?: pulumi.Input<string>;
    /**
     * One or more configuration blocks describing attachments sources to a version of a document. Defined below.
     */
    readonly attachmentsSources?: pulumi.Input<pulumi.Input<inputs.ssm.DocumentAttachmentsSource>[]>;
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
     * The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
     */
    readonly documentType?: pulumi.Input<string>;
    /**
     * The document version.
     */
    readonly documentVersion?: pulumi.Input<string>;
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
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.ssm.DocumentParameter>[]>;
    /**
     * Additional Permissions to attach to the document. See Permissions below for details.
     */
    readonly permissions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
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
     * A map of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
     */
    readonly targetType?: pulumi.Input<string>;
    /**
     * A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
     */
    readonly versionName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Document resource.
 */
export interface DocumentArgs {
    /**
     * One or more configuration blocks describing attachments sources to a version of a document. Defined below.
     */
    readonly attachmentsSources?: pulumi.Input<pulumi.Input<inputs.ssm.DocumentAttachmentsSource>[]>;
    /**
     * The JSON or YAML content of the document.
     */
    readonly content: pulumi.Input<string>;
    /**
     * The format of the document. Valid document types include: `JSON` and `YAML`
     */
    readonly documentFormat?: pulumi.Input<string>;
    /**
     * The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
     */
    readonly documentType: pulumi.Input<string>;
    /**
     * The name of the document.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Additional Permissions to attach to the document. See Permissions below for details.
     */
    readonly permissions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
     */
    readonly targetType?: pulumi.Input<string>;
    /**
     * A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
     */
    readonly versionName?: pulumi.Input<string>;
}
