// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an AWS Storage Gateway SMB File Share.
 *
 * ## Example Usage
 * ### Active Directory Authentication
 *
 * > **NOTE:** The gateway must have already joined the Active Directory domain prior to SMB file share creation. e.g. via "SMB Settings" in the AWS Storage Gateway console or `smbActiveDirectorySettings` in the `aws.storagegateway.Gateway` resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.storagegateway.SmbFileShare("example", {
 *     authentication: "ActiveDirectory",
 *     gatewayArn: aws_storagegateway_gateway.example.arn,
 *     locationArn: aws_s3_bucket.example.arn,
 *     roleArn: aws_iam_role.example.arn,
 * });
 * ```
 * ### Guest Authentication
 *
 * > **NOTE:** The gateway must have already had the SMB guest password set prior to SMB file share creation. e.g. via "SMB Settings" in the AWS Storage Gateway console or `smbGuestPassword` in the `aws.storagegateway.Gateway` resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.storagegateway.SmbFileShare("example", {
 *     authentication: "GuestAccess",
 *     gatewayArn: aws_storagegateway_gateway.example.arn,
 *     locationArn: aws_s3_bucket.example.arn,
 *     roleArn: aws_iam_role.example.arn,
 * });
 * ```
 *
 * ## Import
 *
 * `aws_storagegateway_smb_file_share` can be imported by using the SMB File Share Amazon Resource Name (ARN), e.g.
 *
 * ```sh
 *  $ pulumi import aws:storagegateway/smbFileShare:SmbFileShare example arn:aws:storagegateway:us-east-1:123456789012:share/share-12345678
 * ```
 */
export class SmbFileShare extends pulumi.CustomResource {
    /**
     * Get an existing SmbFileShare resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SmbFileShareState, opts?: pulumi.CustomResourceOptions): SmbFileShare {
        return new SmbFileShare(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:storagegateway/smbFileShare:SmbFileShare';

    /**
     * Returns true if the given object is an instance of SmbFileShare.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SmbFileShare {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SmbFileShare.__pulumiType;
    }

    /**
     * The files and folders on this share will only be visible to users with read access. Default value is `false`.
     */
    public readonly accessBasedEnumeration!: pulumi.Output<boolean | undefined>;
    /**
     * A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
     */
    public readonly adminUserLists!: pulumi.Output<string[] | undefined>;
    /**
     * Amazon Resource Name (ARN) of the SMB File Share.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
     */
    public readonly auditDestinationArn!: pulumi.Output<string | undefined>;
    /**
     * The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
     */
    public readonly authentication!: pulumi.Output<string | undefined>;
    /**
     * Refresh cache information. see Cache Attributes for more details.
     */
    public readonly cacheAttributes!: pulumi.Output<outputs.storagegateway.SmbFileShareCacheAttributes | undefined>;
    /**
     * The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
     */
    public readonly caseSensitivity!: pulumi.Output<string | undefined>;
    /**
     * The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
     */
    public readonly defaultStorageClass!: pulumi.Output<string | undefined>;
    /**
     * The name of the file share. Must be set if an S3 prefix name is set in `locationArn`.
     */
    public readonly fileShareName!: pulumi.Output<string>;
    /**
     * ID of the SMB File Share.
     */
    public /*out*/ readonly fileshareId!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the file gateway.
     */
    public readonly gatewayArn!: pulumi.Output<string>;
    /**
     * Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
     */
    public readonly guessMimeTypeEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
     */
    public readonly invalidUserLists!: pulumi.Output<string[] | undefined>;
    /**
     * Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
     */
    public readonly kmsEncrypted!: pulumi.Output<boolean | undefined>;
    /**
     * Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
     */
    public readonly kmsKeyArn!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the backed storage used for storing file data.
     */
    public readonly locationArn!: pulumi.Output<string>;
    /**
     * The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
     */
    public readonly notificationPolicy!: pulumi.Output<string | undefined>;
    /**
     * Access Control List permission for S3 bucket objects. Defaults to `private`.
     */
    public readonly objectAcl!: pulumi.Output<string | undefined>;
    /**
     * File share path used by the NFS client to identify the mount point.
     */
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
     */
    public readonly readOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
     */
    public readonly requesterPays!: pulumi.Output<boolean | undefined>;
    /**
     * The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
     */
    public readonly smbAclEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
     */
    public readonly validUserLists!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SmbFileShare resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SmbFileShareArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SmbFileShareArgs | SmbFileShareState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SmbFileShareState | undefined;
            inputs["accessBasedEnumeration"] = state ? state.accessBasedEnumeration : undefined;
            inputs["adminUserLists"] = state ? state.adminUserLists : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["auditDestinationArn"] = state ? state.auditDestinationArn : undefined;
            inputs["authentication"] = state ? state.authentication : undefined;
            inputs["cacheAttributes"] = state ? state.cacheAttributes : undefined;
            inputs["caseSensitivity"] = state ? state.caseSensitivity : undefined;
            inputs["defaultStorageClass"] = state ? state.defaultStorageClass : undefined;
            inputs["fileShareName"] = state ? state.fileShareName : undefined;
            inputs["fileshareId"] = state ? state.fileshareId : undefined;
            inputs["gatewayArn"] = state ? state.gatewayArn : undefined;
            inputs["guessMimeTypeEnabled"] = state ? state.guessMimeTypeEnabled : undefined;
            inputs["invalidUserLists"] = state ? state.invalidUserLists : undefined;
            inputs["kmsEncrypted"] = state ? state.kmsEncrypted : undefined;
            inputs["kmsKeyArn"] = state ? state.kmsKeyArn : undefined;
            inputs["locationArn"] = state ? state.locationArn : undefined;
            inputs["notificationPolicy"] = state ? state.notificationPolicy : undefined;
            inputs["objectAcl"] = state ? state.objectAcl : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["readOnly"] = state ? state.readOnly : undefined;
            inputs["requesterPays"] = state ? state.requesterPays : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["smbAclEnabled"] = state ? state.smbAclEnabled : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["validUserLists"] = state ? state.validUserLists : undefined;
        } else {
            const args = argsOrState as SmbFileShareArgs | undefined;
            if ((!args || args.gatewayArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayArn'");
            }
            if ((!args || args.locationArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locationArn'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["accessBasedEnumeration"] = args ? args.accessBasedEnumeration : undefined;
            inputs["adminUserLists"] = args ? args.adminUserLists : undefined;
            inputs["auditDestinationArn"] = args ? args.auditDestinationArn : undefined;
            inputs["authentication"] = args ? args.authentication : undefined;
            inputs["cacheAttributes"] = args ? args.cacheAttributes : undefined;
            inputs["caseSensitivity"] = args ? args.caseSensitivity : undefined;
            inputs["defaultStorageClass"] = args ? args.defaultStorageClass : undefined;
            inputs["fileShareName"] = args ? args.fileShareName : undefined;
            inputs["gatewayArn"] = args ? args.gatewayArn : undefined;
            inputs["guessMimeTypeEnabled"] = args ? args.guessMimeTypeEnabled : undefined;
            inputs["invalidUserLists"] = args ? args.invalidUserLists : undefined;
            inputs["kmsEncrypted"] = args ? args.kmsEncrypted : undefined;
            inputs["kmsKeyArn"] = args ? args.kmsKeyArn : undefined;
            inputs["locationArn"] = args ? args.locationArn : undefined;
            inputs["notificationPolicy"] = args ? args.notificationPolicy : undefined;
            inputs["objectAcl"] = args ? args.objectAcl : undefined;
            inputs["readOnly"] = args ? args.readOnly : undefined;
            inputs["requesterPays"] = args ? args.requesterPays : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["smbAclEnabled"] = args ? args.smbAclEnabled : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["validUserLists"] = args ? args.validUserLists : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["fileshareId"] = undefined /*out*/;
            inputs["path"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SmbFileShare.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SmbFileShare resources.
 */
export interface SmbFileShareState {
    /**
     * The files and folders on this share will only be visible to users with read access. Default value is `false`.
     */
    accessBasedEnumeration?: pulumi.Input<boolean>;
    /**
     * A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
     */
    adminUserLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Amazon Resource Name (ARN) of the SMB File Share.
     */
    arn?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
     */
    auditDestinationArn?: pulumi.Input<string>;
    /**
     * The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Refresh cache information. see Cache Attributes for more details.
     */
    cacheAttributes?: pulumi.Input<inputs.storagegateway.SmbFileShareCacheAttributes>;
    /**
     * The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
     */
    caseSensitivity?: pulumi.Input<string>;
    /**
     * The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
     */
    defaultStorageClass?: pulumi.Input<string>;
    /**
     * The name of the file share. Must be set if an S3 prefix name is set in `locationArn`.
     */
    fileShareName?: pulumi.Input<string>;
    /**
     * ID of the SMB File Share.
     */
    fileshareId?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the file gateway.
     */
    gatewayArn?: pulumi.Input<string>;
    /**
     * Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
     */
    guessMimeTypeEnabled?: pulumi.Input<boolean>;
    /**
     * A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
     */
    invalidUserLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
     */
    kmsEncrypted?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
     */
    kmsKeyArn?: pulumi.Input<string>;
    /**
     * The ARN of the backed storage used for storing file data.
     */
    locationArn?: pulumi.Input<string>;
    /**
     * The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
     */
    notificationPolicy?: pulumi.Input<string>;
    /**
     * Access Control List permission for S3 bucket objects. Defaults to `private`.
     */
    objectAcl?: pulumi.Input<string>;
    /**
     * File share path used by the NFS client to identify the mount point.
     */
    path?: pulumi.Input<string>;
    /**
     * Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
     */
    requesterPays?: pulumi.Input<boolean>;
    /**
     * The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
     */
    smbAclEnabled?: pulumi.Input<boolean>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
     */
    validUserLists?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SmbFileShare resource.
 */
export interface SmbFileShareArgs {
    /**
     * The files and folders on this share will only be visible to users with read access. Default value is `false`.
     */
    accessBasedEnumeration?: pulumi.Input<boolean>;
    /**
     * A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
     */
    adminUserLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
     */
    auditDestinationArn?: pulumi.Input<string>;
    /**
     * The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Refresh cache information. see Cache Attributes for more details.
     */
    cacheAttributes?: pulumi.Input<inputs.storagegateway.SmbFileShareCacheAttributes>;
    /**
     * The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
     */
    caseSensitivity?: pulumi.Input<string>;
    /**
     * The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
     */
    defaultStorageClass?: pulumi.Input<string>;
    /**
     * The name of the file share. Must be set if an S3 prefix name is set in `locationArn`.
     */
    fileShareName?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the file gateway.
     */
    gatewayArn: pulumi.Input<string>;
    /**
     * Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
     */
    guessMimeTypeEnabled?: pulumi.Input<boolean>;
    /**
     * A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
     */
    invalidUserLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
     */
    kmsEncrypted?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
     */
    kmsKeyArn?: pulumi.Input<string>;
    /**
     * The ARN of the backed storage used for storing file data.
     */
    locationArn: pulumi.Input<string>;
    /**
     * The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
     */
    notificationPolicy?: pulumi.Input<string>;
    /**
     * Access Control List permission for S3 bucket objects. Defaults to `private`.
     */
    objectAcl?: pulumi.Input<string>;
    /**
     * Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
     */
    requesterPays?: pulumi.Input<boolean>;
    /**
     * The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
     */
    roleArn: pulumi.Input<string>;
    /**
     * Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
     */
    smbAclEnabled?: pulumi.Input<boolean>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
     */
    validUserLists?: pulumi.Input<pulumi.Input<string>[]>;
}
