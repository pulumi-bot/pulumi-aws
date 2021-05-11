// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Amazon API Gateway Version 2 domain name.
 * More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
 *
 * > **Note:** This resource establishes ownership of and the TLS settings for
 * a particular domain name. An API stage can be associated with the domain name using the `aws.apigatewayv2.ApiMapping` resource.
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apigatewayv2.DomainName("example", {
 *     domainName: "ws-api.example.com",
 *     domainNameConfiguration: {
 *         certificateArn: aws_acm_certificate.example.arn,
 *         endpointType: "REGIONAL",
 *         securityPolicy: "TLS_1_2",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * `aws_apigatewayv2_domain_name` can be imported by using the domain name, e.g.
 *
 * ```sh
 *  $ pulumi import aws:apigatewayv2/domainName:DomainName example ws-api.example.com
 * ```
 */
export class DomainName extends pulumi.CustomResource {
    /**
     * Get an existing DomainName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainNameState, opts?: pulumi.CustomResourceOptions): DomainName {
        return new DomainName(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigatewayv2/domainName:DomainName';

    /**
     * Returns true if the given object is an instance of DomainName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainName.__pulumiType;
    }

    /**
     * The [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
     */
    public /*out*/ readonly apiMappingSelectionExpression!: pulumi.Output<string>;
    /**
     * The ARN of the domain name.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The domain name. Must be between 1 and 512 characters in length.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The domain name configuration.
     */
    public readonly domainNameConfiguration!: pulumi.Output<outputs.apigatewayv2.DomainNameDomainNameConfiguration>;
    /**
     * The mutual TLS authentication configuration for the domain name.
     */
    public readonly mutualTlsAuthentication!: pulumi.Output<outputs.apigatewayv2.DomainNameMutualTlsAuthentication | undefined>;
    /**
     * A map of tags to assign to the domain name. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a DomainName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainNameArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainNameArgs | DomainNameState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainNameState | undefined;
            inputs["apiMappingSelectionExpression"] = state ? state.apiMappingSelectionExpression : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["domainNameConfiguration"] = state ? state.domainNameConfiguration : undefined;
            inputs["mutualTlsAuthentication"] = state ? state.mutualTlsAuthentication : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as DomainNameArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.domainNameConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainNameConfiguration'");
            }
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["domainNameConfiguration"] = args ? args.domainNameConfiguration : undefined;
            inputs["mutualTlsAuthentication"] = args ? args.mutualTlsAuthentication : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["apiMappingSelectionExpression"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DomainName.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainName resources.
 */
export interface DomainNameState {
    /**
     * The [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
     */
    apiMappingSelectionExpression?: pulumi.Input<string>;
    /**
     * The ARN of the domain name.
     */
    arn?: pulumi.Input<string>;
    /**
     * The domain name. Must be between 1 and 512 characters in length.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The domain name configuration.
     */
    domainNameConfiguration?: pulumi.Input<inputs.apigatewayv2.DomainNameDomainNameConfiguration>;
    /**
     * The mutual TLS authentication configuration for the domain name.
     */
    mutualTlsAuthentication?: pulumi.Input<inputs.apigatewayv2.DomainNameMutualTlsAuthentication>;
    /**
     * A map of tags to assign to the domain name. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a DomainName resource.
 */
export interface DomainNameArgs {
    /**
     * The domain name. Must be between 1 and 512 characters in length.
     */
    domainName: pulumi.Input<string>;
    /**
     * The domain name configuration.
     */
    domainNameConfiguration: pulumi.Input<inputs.apigatewayv2.DomainNameDomainNameConfiguration>;
    /**
     * The mutual TLS authentication configuration for the domain name.
     */
    mutualTlsAuthentication?: pulumi.Input<inputs.apigatewayv2.DomainNameMutualTlsAuthentication>;
    /**
     * A map of tags to assign to the domain name. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
