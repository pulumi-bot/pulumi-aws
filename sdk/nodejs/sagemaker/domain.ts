// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Sagemaker Domain resource.
 *
 * ## Example Usage
 * ### Basic usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleDomain = new aws.sagemaker.Domain("exampleDomain", {
 *     domainName: "example",
 *     authMode: "IAM",
 *     vpcId: aws_vpc.test.id,
 *     subnetIds: [aws_subnet.test.id],
 *     defaultUserSettings: {
 *         executionRole: aws_iam_role.test.arn,
 *     },
 * });
 * const examplePolicyDocument = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["sagemaker.amazonaws.com"],
 *         }],
 *     }],
 * });
 * const exampleRole = new aws.iam.Role("exampleRole", {
 *     path: "/",
 *     assumeRolePolicy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
 * });
 * ```
 * ### Using Custom Images
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testImage = new aws.sagemaker.Image("testImage", {
 *     imageName: "example",
 *     roleArn: aws_iam_role.test.arn,
 * });
 * const testAppImageConfig = new aws.sagemaker.AppImageConfig("testAppImageConfig", {
 *     appImageConfigName: "example",
 *     kernelGatewayImageConfig: {
 *         kernelSpec: {
 *             name: "example",
 *         },
 *     },
 * });
 * const testImageVersion = new aws.sagemaker.ImageVersion("testImageVersion", {
 *     imageName: testImage.id,
 *     baseImage: "base-image",
 * });
 * const testDomain = new aws.sagemaker.Domain("testDomain", {
 *     domainName: "example",
 *     authMode: "IAM",
 *     vpcId: aws_vpc.test.id,
 *     subnetIds: [aws_subnet.test.id],
 *     defaultUserSettings: {
 *         executionRole: aws_iam_role.test.arn,
 *         kernelGatewayAppSettings: {
 *             customImages: [{
 *                 appImageConfigName: testAppImageConfig.appImageConfigName,
 *                 imageName: testImageVersion.imageName,
 *             }],
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Sagemaker Code Domains can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:sagemaker/domain:Domain test_domain d-8jgsjtilstu8
 * ```
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
     */
    public readonly appNetworkAccessType!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this Domain.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
     */
    public readonly authMode!: pulumi.Output<string>;
    /**
     * The default user settings. See Default User Settings below.
     */
    public readonly defaultUserSettings!: pulumi.Output<outputs.sagemaker.DomainDefaultUserSettings>;
    /**
     * The domain name.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The ID of the Amazon Elastic File System (EFS) managed by this Domain.
     */
    public /*out*/ readonly homeEfsFileSystemId!: pulumi.Output<string>;
    /**
     * The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * The SSO managed application instance ID.
     */
    public /*out*/ readonly singleSignOnManagedApplicationInstanceId!: pulumi.Output<string>;
    /**
     * The VPC subnets that Studio uses for communication.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The domain's URL.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainState | undefined;
            inputs["appNetworkAccessType"] = state ? state.appNetworkAccessType : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["authMode"] = state ? state.authMode : undefined;
            inputs["defaultUserSettings"] = state ? state.defaultUserSettings : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["homeEfsFileSystemId"] = state ? state.homeEfsFileSystemId : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["singleSignOnManagedApplicationInstanceId"] = state ? state.singleSignOnManagedApplicationInstanceId : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if ((!args || args.authMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authMode'");
            }
            if ((!args || args.defaultUserSettings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultUserSettings'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["appNetworkAccessType"] = args ? args.appNetworkAccessType : undefined;
            inputs["authMode"] = args ? args.authMode : undefined;
            inputs["defaultUserSettings"] = args ? args.defaultUserSettings : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["homeEfsFileSystemId"] = undefined /*out*/;
            inputs["singleSignOnManagedApplicationInstanceId"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Domain.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
     */
    appNetworkAccessType?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this Domain.
     */
    arn?: pulumi.Input<string>;
    /**
     * The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
     */
    authMode?: pulumi.Input<string>;
    /**
     * The default user settings. See Default User Settings below.
     */
    defaultUserSettings?: pulumi.Input<inputs.sagemaker.DomainDefaultUserSettings>;
    /**
     * The domain name.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The ID of the Amazon Elastic File System (EFS) managed by this Domain.
     */
    homeEfsFileSystemId?: pulumi.Input<string>;
    /**
     * The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The SSO managed application instance ID.
     */
    singleSignOnManagedApplicationInstanceId?: pulumi.Input<string>;
    /**
     * The VPC subnets that Studio uses for communication.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The domain's URL.
     */
    url?: pulumi.Input<string>;
    /**
     * The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
     */
    appNetworkAccessType?: pulumi.Input<string>;
    /**
     * The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
     */
    authMode: pulumi.Input<string>;
    /**
     * The default user settings. See Default User Settings below.
     */
    defaultUserSettings: pulumi.Input<inputs.sagemaker.DomainDefaultUserSettings>;
    /**
     * The domain name.
     */
    domainName: pulumi.Input<string>;
    /**
     * The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The VPC subnets that Studio uses for communication.
     */
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
     */
    vpcId: pulumi.Input<string>;
}
