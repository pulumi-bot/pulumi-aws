// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Subscribes to a Security Hub product.
 * 
 * > **NOTE:** This AWS service is in Preview and may change before General Availability release. Backwards compatibility is not guaranteed between AWS Provider releases.
 * 
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleAccount = new aws.securityhub.Account("example", {});
 * const current = aws.getRegion();
 * const exampleProductSubscription = new aws.securityhub.ProductSubscription("example", {
 *     productArn: `arn:aws:securityhub:${current.name!}:733251395267:product/alertlogic/althreatmanagement`,
 * }, {dependsOn: [exampleAccount]});
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/securityhub_product_subscription.markdown.
 */
export class ProductSubscription extends pulumi.CustomResource {
    /**
     * Get an existing ProductSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProductSubscriptionState, opts?: pulumi.CustomResourceOptions): ProductSubscription {
        return new ProductSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:securityhub/productSubscription:ProductSubscription';

    /**
     * Returns true if the given object is an instance of ProductSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProductSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProductSubscription.__pulumiType;
    }

    /**
     * The ARN of a resource that represents your subscription to the product that generates the findings that you want to import into Security Hub.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ARN of the product that generates findings that you want to import into Security Hub - see below.
     */
    public readonly productArn!: pulumi.Output<string>;

    /**
     * Create a ProductSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProductSubscriptionArgs | ProductSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProductSubscriptionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["productArn"] = state ? state.productArn : undefined;
        } else {
            const args = argsOrState as ProductSubscriptionArgs | undefined;
            if (!args || args.productArn === undefined) {
                throw new Error("Missing required property 'productArn'");
            }
            inputs["productArn"] = args ? args.productArn : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProductSubscription.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProductSubscription resources.
 */
export interface ProductSubscriptionState {
    /**
     * The ARN of a resource that represents your subscription to the product that generates the findings that you want to import into Security Hub.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The ARN of the product that generates findings that you want to import into Security Hub - see below.
     */
    readonly productArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProductSubscription resource.
 */
export interface ProductSubscriptionArgs {
    /**
     * The ARN of the product that generates findings that you want to import into Security Hub - see below.
     */
    readonly productArn: pulumi.Input<string>;
}
