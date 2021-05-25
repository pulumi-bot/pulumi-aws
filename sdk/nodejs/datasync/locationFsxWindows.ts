// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AWS DataSync FSx Windows Location.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.datasync.LocationFsxWindows("example", {
 *     fsxFilesystemArn: aws_fsx_windows_file_system.example.arn,
 *     user: "SomeUser",
 *     password: "SuperSecretPassw0rd",
 *     securityGroupArns: [aws_security_group.example.arn],
 * });
 * ```
 *
 * ## Import
 *
 * `aws_datasync_location_fsx_windows_file_system` can be imported by using the `DataSync-ARN#FSx-Windows-ARN`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:datasync/locationFsxWindows:LocationFsxWindows example arn:aws:datasync:us-west-2:123456789012:location/loc-12345678901234567#arn:aws:fsx:us-west-2:476956259333:file-system/fs-08e04cd442c1bb94a
 * ```
 */
export class LocationFsxWindows extends pulumi.CustomResource {
    /**
     * Get an existing LocationFsxWindows resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocationFsxWindowsState, opts?: pulumi.CustomResourceOptions): LocationFsxWindows {
        return new LocationFsxWindows(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:datasync/locationFsxWindows:LocationFsxWindows';

    /**
     * Returns true if the given object is an instance of LocationFsxWindows.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocationFsxWindows {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocationFsxWindows.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The time that the FSx for Windows location was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The name of the Windows domain that the FSx for Windows server belongs to.
     */
    public readonly domain!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) for the FSx for Windows file system.
     */
    public readonly fsxFilesystemArn!: pulumi.Output<string>;
    /**
     * The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
     */
    public readonly securityGroupArns!: pulumi.Output<string[]>;
    /**
     * Subdirectory to perform actions as source or destination.
     */
    public readonly subdirectory!: pulumi.Output<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The URL of the FSx for Windows location that was described.
     */
    public /*out*/ readonly uri!: pulumi.Output<string>;
    /**
     * The user who has the permissions to access files and folders in the FSx for Windows file system.
     */
    public readonly user!: pulumi.Output<string>;

    /**
     * Create a LocationFsxWindows resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocationFsxWindowsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocationFsxWindowsArgs | LocationFsxWindowsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocationFsxWindowsState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["fsxFilesystemArn"] = state ? state.fsxFilesystemArn : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["securityGroupArns"] = state ? state.securityGroupArns : undefined;
            inputs["subdirectory"] = state ? state.subdirectory : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["uri"] = state ? state.uri : undefined;
            inputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as LocationFsxWindowsArgs | undefined;
            if ((!args || args.fsxFilesystemArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fsxFilesystemArn'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.securityGroupArns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupArns'");
            }
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            inputs["domain"] = args ? args.domain : undefined;
            inputs["fsxFilesystemArn"] = args ? args.fsxFilesystemArn : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["securityGroupArns"] = args ? args.securityGroupArns : undefined;
            inputs["subdirectory"] = args ? args.subdirectory : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["user"] = args ? args.user : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LocationFsxWindows.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LocationFsxWindows resources.
 */
export interface LocationFsxWindowsState {
    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     */
    arn?: pulumi.Input<string>;
    /**
     * The time that the FSx for Windows location was created.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The name of the Windows domain that the FSx for Windows server belongs to.
     */
    domain?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) for the FSx for Windows file system.
     */
    fsxFilesystemArn?: pulumi.Input<string>;
    /**
     * The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
     */
    password?: pulumi.Input<string>;
    /**
     * The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
     */
    securityGroupArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Subdirectory to perform actions as source or destination.
     */
    subdirectory?: pulumi.Input<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The URL of the FSx for Windows location that was described.
     */
    uri?: pulumi.Input<string>;
    /**
     * The user who has the permissions to access files and folders in the FSx for Windows file system.
     */
    user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LocationFsxWindows resource.
 */
export interface LocationFsxWindowsArgs {
    /**
     * The name of the Windows domain that the FSx for Windows server belongs to.
     */
    domain?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) for the FSx for Windows file system.
     */
    fsxFilesystemArn: pulumi.Input<string>;
    /**
     * The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
     */
    password: pulumi.Input<string>;
    /**
     * The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
     */
    securityGroupArns: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Subdirectory to perform actions as source or destination.
     */
    subdirectory?: pulumi.Input<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The user who has the permissions to access files and folders in the FSx for Windows file system.
     */
    user: pulumi.Input<string>;
}
