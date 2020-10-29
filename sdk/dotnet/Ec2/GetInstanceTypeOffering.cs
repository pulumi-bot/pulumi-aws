// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetInstanceTypeOffering
    {
        /// <summary>
        /// Information about single EC2 Instance Type Offering.
        /// </summary>
        public static Task<GetInstanceTypeOfferingResult> InvokeAsync(GetInstanceTypeOfferingArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTypeOfferingResult>("aws:ec2/getInstanceTypeOffering:getInstanceTypeOffering", args ?? new GetInstanceTypeOfferingArgs(), options.WithVersion());
    }


    public sealed class GetInstanceTypeOfferingArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetInstanceTypeOfferingFilterArgs>? _filters;

        /// <summary>
        /// One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstanceTypeOfferings.html) for supported filters. Detailed below.
        /// </summary>
        public List<Inputs.GetInstanceTypeOfferingFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInstanceTypeOfferingFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Location type. Defaults to `region`. Valid values: `availability-zone`, `availability-zone-id`, and `region`.
        /// </summary>
        [Input("locationType")]
        public string? LocationType { get; set; }

        [Input("preferredInstanceTypes")]
        private List<string>? _preferredInstanceTypes;

        /// <summary>
        /// Ordered list of preferred EC2 Instance Types. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
        /// </summary>
        public List<string> PreferredInstanceTypes
        {
            get => _preferredInstanceTypes ?? (_preferredInstanceTypes = new List<string>());
            set => _preferredInstanceTypes = value;
        }

        public GetInstanceTypeOfferingArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceTypeOfferingResult
    {
        public readonly ImmutableArray<Outputs.GetInstanceTypeOfferingFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// EC2 Instance Type.
        /// </summary>
        public readonly string InstanceType;
        public readonly string? LocationType;
        public readonly ImmutableArray<string> PreferredInstanceTypes;

        [OutputConstructor]
        private GetInstanceTypeOfferingResult(
            ImmutableArray<Outputs.GetInstanceTypeOfferingFilterResult> filters,

            string id,

            string instanceType,

            string? locationType,

            ImmutableArray<string> preferredInstanceTypes)
        {
            Filters = filters;
            Id = id;
            InstanceType = instanceType;
            LocationType = locationType;
            PreferredInstanceTypes = preferredInstanceTypes;
        }
    }
}
