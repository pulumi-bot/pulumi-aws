// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pricing
{
    public static partial class GetProduct
    {
        /// <summary>
        /// Use this data source to get the pricing information of all products in AWS.
        /// This data source is only available in a us-east-1 or ap-south-1 provider.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/pricing_product.html.markdown.
        /// </summary>
        public static Task<GetProductResult> InvokeAsync(GetProductArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProductResult>("aws:pricing/getProduct:getProduct", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetProductArgs : Pulumi.InvokeArgs
    {
        [Input("filters", required: true)]
        private List<Inputs.GetProductFiltersArgs>? _filters;

        /// <summary>
        /// A list of filters. Passed directly to the API (see GetProducts API reference). These filters must describe a single product, this resource will fail if more than one product is returned by the API.
        /// </summary>
        public List<Inputs.GetProductFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetProductFiltersArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The code of the service. Available service codes can be fetched using the DescribeServices pricing API call.
        /// </summary>
        [Input("serviceCode", required: true)]
        public string ServiceCode { get; set; } = null!;

        public GetProductArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetProductResult
    {
        public readonly ImmutableArray<Outputs.GetProductFiltersResult> Filters;
        /// <summary>
        /// Set to the product returned from the API.
        /// </summary>
        public readonly string Result;
        public readonly string ServiceCode;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetProductResult(
            ImmutableArray<Outputs.GetProductFiltersResult> filters,
            string result,
            string serviceCode,
            string id)
        {
            Filters = filters;
            Result = result;
            ServiceCode = serviceCode;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetProductFiltersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The product attribute name that you want to filter on.
        /// </summary>
        [Input("field", required: true)]
        public string Field { get; set; } = null!;

        /// <summary>
        /// The product attribute value that you want to filter on.
        /// </summary>
        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public GetProductFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetProductFiltersResult
    {
        /// <summary>
        /// The product attribute name that you want to filter on.
        /// </summary>
        public readonly string Field;
        /// <summary>
        /// The product attribute value that you want to filter on.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetProductFiltersResult(
            string field,
            string value)
        {
            Field = field;
            Value = value;
        }
    }
    }
}
