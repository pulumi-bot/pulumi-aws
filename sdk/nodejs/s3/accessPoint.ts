// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage an S3 Access Point.
 *
 * ## Example Usage
 * ### AWS Partition Bucket
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleBucket = new aws.s3.Bucket("exampleBucket", {});
 * const exampleAccessPoint = new aws.s3.AccessPoint("exampleAccessPoint", {bucket: exampleBucket.id});
 * ```
 * ### S3 on Outposts Bucket
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleBucket = new aws.s3control.Bucket("exampleBucket", {bucket: "example"});
 * const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.0.0.0/16"});
 * const exampleAccessPoint = new aws.s3.AccessPoint("exampleAccessPoint", {
 *     bucket: exampleBucket.arn,
 *     vpcConfiguration: {
 *         vpcId: exampleVpc.id,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * For Access Points associated with an AWS Partition S3 Bucket, this resource can be imported using the `account_id` and `name` separated by a colon (`:`), e.g.
 *
 * ```sh
 *  $ pulumi import aws:s3/accessPoint:AccessPoint example 123456789012:example
 * ```
 *
 *  For Access Points associated with an S3 on Outposts Bucket, this resource can be imported using the Amazon Resource Name (ARN), e.g.
 *
 * ```sh
 *  $ pulumi import aws:s3/accessPoint:AccessPoint example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-1234567890123456/accesspoint/example
 * ```
 */
export class AccessPoint extends pulumi.CustomResource {
    /**
     * Get an existing AccessPoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessPointState, opts?: pulumi.CustomResourceOptions): AccessPoint {
        return new AccessPoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3/accessPoint:AccessPoint';

    /**
     * Returns true if the given object is an instance of AccessPoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPoint.__pulumiType;
    }

    /**
     * The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the S3 Access Point.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of an AWS Partition S3 Bucket or the Amazon Resource Name (ARN) of S3 on Outposts Bucket that you want to associate this access point with.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The DNS domain name of the S3 Access Point in the format _`name`_-_`accountId`_.s3-accesspoint._region_.amazonaws.com.
     * Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
     */
    public /*out*/ readonly domainName!: pulumi.Output<string>;
    /**
     * Indicates whether this access point currently has a policy that allows public access.
     */
    public /*out*/ readonly hasPublicAccessPolicy!: pulumi.Output<boolean>;
    /**
     * The name you want to assign to this access point.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
     */
    public /*out*/ readonly networkOrigin!: pulumi.Output<string>;
    /**
     * A valid JSON document that specifies the policy that you want to apply to this access point.
     */
    public readonly policy!: pulumi.Output<string | undefined>;
    /**
     * Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
     */
    public readonly publicAccessBlockConfiguration!: pulumi.Output<outputs.s3.AccessPointPublicAccessBlockConfiguration | undefined>;
    /**
     * Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
     */
    public readonly vpcConfiguration!: pulumi.Output<outputs.s3.AccessPointVpcConfiguration | undefined>;

    /**
     * Create a AccessPoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessPointArgs | AccessPointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessPointState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["hasPublicAccessPolicy"] = state ? state.hasPublicAccessPolicy : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkOrigin"] = state ? state.networkOrigin : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["publicAccessBlockConfiguration"] = state ? state.publicAccessBlockConfiguration : undefined;
            inputs["vpcConfiguration"] = state ? state.vpcConfiguration : undefined;
        } else {
            const args = argsOrState as AccessPointArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["publicAccessBlockConfiguration"] = args ? args.publicAccessBlockConfiguration : undefined;
            inputs["vpcConfiguration"] = args ? args.vpcConfiguration : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["domainName"] = undefined /*out*/;
            inputs["hasPublicAccessPolicy"] = undefined /*out*/;
            inputs["networkOrigin"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AccessPoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessPoint resources.
 */
export interface AccessPointState {
    /**
     * The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the S3 Access Point.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name of an AWS Partition S3 Bucket or the Amazon Resource Name (ARN) of S3 on Outposts Bucket that you want to associate this access point with.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * The DNS domain name of the S3 Access Point in the format _`name`_-_`accountId`_.s3-accesspoint._region_.amazonaws.com.
     * Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * Indicates whether this access point currently has a policy that allows public access.
     */
    readonly hasPublicAccessPolicy?: pulumi.Input<boolean>;
    /**
     * The name you want to assign to this access point.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
     */
    readonly networkOrigin?: pulumi.Input<string>;
    /**
     * A valid JSON document that specifies the policy that you want to apply to this access point.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
     */
    readonly publicAccessBlockConfiguration?: pulumi.Input<inputs.s3.AccessPointPublicAccessBlockConfiguration>;
    /**
     * Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
     */
    readonly vpcConfiguration?: pulumi.Input<inputs.s3.AccessPointVpcConfiguration>;
}

/**
 * The set of arguments for constructing a AccessPoint resource.
 */
export interface AccessPointArgs {
    /**
     * The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * The name of an AWS Partition S3 Bucket or the Amazon Resource Name (ARN) of S3 on Outposts Bucket that you want to associate this access point with.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * The name you want to assign to this access point.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A valid JSON document that specifies the policy that you want to apply to this access point.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
     */
    readonly publicAccessBlockConfiguration?: pulumi.Input<inputs.s3.AccessPointPublicAccessBlockConfiguration>;
    /**
     * Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
     */
    readonly vpcConfiguration?: pulumi.Input<inputs.s3.AccessPointVpcConfiguration>;
}
