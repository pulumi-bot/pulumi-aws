// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a FSx Lustre File System. See the [FSx Lustre Guide](https://docs.aws.amazon.com/fsx/latest/LustreGuide/what-is.html) for more information.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.fsx.LustreFileSystem("example", {
 *     importPath: pulumi.interpolate`s3://${aws_s3_bucket_example.bucket}`,
 *     storageCapacity: 1200,
 *     subnetIds: aws_subnet_example.id,
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/fsx_lustre_file_system.html.markdown.
 */
export class LustreFileSystem extends pulumi.CustomResource {
    /**
     * Get an existing LustreFileSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LustreFileSystemState, opts?: pulumi.CustomResourceOptions): LustreFileSystem {
        return new LustreFileSystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:fsx/lustreFileSystem:LustreFileSystem';

    /**
     * Returns true if the given object is an instance of LustreFileSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LustreFileSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LustreFileSystem.__pulumiType;
    }

    /**
     * Amazon Resource Name of the file system.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `importPath` argument and the path must use the same Amazon S3 bucket as specified in `importPath`. Set equal to `importPath` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
     */
    public readonly exportPath!: pulumi.Output<string>;
    /**
     * S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
     */
    public readonly importPath!: pulumi.Output<string | undefined>;
    /**
     * For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `importPath` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
     */
    public readonly importedFileChunkSize!: pulumi.Output<number>;
    /**
     * Set of Elastic Network Interface identifiers from which the file system is accessible.
     */
    public /*out*/ readonly networkInterfaceIds!: pulumi.Output<string[]>;
    /**
     * AWS account identifier that created the file system.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The storage capacity (GiB) of the file system. Minimum of `1200`. Storage capacity is provisioned in increments of 3,600 GiB.
     */
    public readonly storageCapacity!: pulumi.Output<number>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
     */
    public readonly subnetIds!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the file system.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Identifier of the Virtual Private Cloud for the file system.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
     */
    public readonly weeklyMaintenanceStartTime!: pulumi.Output<string>;

    /**
     * Create a LustreFileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LustreFileSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LustreFileSystemArgs | LustreFileSystemState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LustreFileSystemState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["exportPath"] = state ? state.exportPath : undefined;
            inputs["importPath"] = state ? state.importPath : undefined;
            inputs["importedFileChunkSize"] = state ? state.importedFileChunkSize : undefined;
            inputs["networkInterfaceIds"] = state ? state.networkInterfaceIds : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["storageCapacity"] = state ? state.storageCapacity : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["weeklyMaintenanceStartTime"] = state ? state.weeklyMaintenanceStartTime : undefined;
        } else {
            const args = argsOrState as LustreFileSystemArgs | undefined;
            if (!args || args.storageCapacity === undefined) {
                throw new Error("Missing required property 'storageCapacity'");
            }
            if (!args || args.subnetIds === undefined) {
                throw new Error("Missing required property 'subnetIds'");
            }
            inputs["exportPath"] = args ? args.exportPath : undefined;
            inputs["importPath"] = args ? args.importPath : undefined;
            inputs["importedFileChunkSize"] = args ? args.importedFileChunkSize : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["storageCapacity"] = args ? args.storageCapacity : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["weeklyMaintenanceStartTime"] = args ? args.weeklyMaintenanceStartTime : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["dnsName"] = undefined /*out*/;
            inputs["networkInterfaceIds"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LustreFileSystem.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LustreFileSystem resources.
 */
export interface LustreFileSystemState {
    /**
     * Amazon Resource Name of the file system.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `importPath` argument and the path must use the same Amazon S3 bucket as specified in `importPath`. Set equal to `importPath` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
     */
    readonly exportPath?: pulumi.Input<string>;
    /**
     * S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
     */
    readonly importPath?: pulumi.Input<string>;
    /**
     * For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `importPath` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
     */
    readonly importedFileChunkSize?: pulumi.Input<number>;
    /**
     * Set of Elastic Network Interface identifiers from which the file system is accessible.
     */
    readonly networkInterfaceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * AWS account identifier that created the file system.
     */
    readonly ownerId?: pulumi.Input<string>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The storage capacity (GiB) of the file system. Minimum of `1200`. Storage capacity is provisioned in increments of 3,600 GiB.
     */
    readonly storageCapacity?: pulumi.Input<number>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
     */
    readonly subnetIds?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the file system.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Identifier of the Virtual Private Cloud for the file system.
     */
    readonly vpcId?: pulumi.Input<string>;
    /**
     * The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
     */
    readonly weeklyMaintenanceStartTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LustreFileSystem resource.
 */
export interface LustreFileSystemArgs {
    /**
     * S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `importPath` argument and the path must use the same Amazon S3 bucket as specified in `importPath`. Set equal to `importPath` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
     */
    readonly exportPath?: pulumi.Input<string>;
    /**
     * S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
     */
    readonly importPath?: pulumi.Input<string>;
    /**
     * For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `importPath` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
     */
    readonly importedFileChunkSize?: pulumi.Input<number>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The storage capacity (GiB) of the file system. Minimum of `1200`. Storage capacity is provisioned in increments of 3,600 GiB.
     */
    readonly storageCapacity: pulumi.Input<number>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
     */
    readonly subnetIds: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the file system.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
     */
    readonly weeklyMaintenanceStartTime?: pulumi.Input<string>;
}
