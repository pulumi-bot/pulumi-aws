// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ServiceQuotas
{
    public static class GetServiceQuota
    {
        /// <summary>
        /// Retrieve information about a Service Quota.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var byQuotaCode = Output.Create(Aws.ServiceQuotas.GetServiceQuota.InvokeAsync(new Aws.ServiceQuotas.GetServiceQuotaArgs
        ///         {
        ///             QuotaCode = "L-F678F1CE",
        ///             ServiceCode = "vpc",
        ///         }));
        ///         var byQuotaName = Output.Create(Aws.ServiceQuotas.GetServiceQuota.InvokeAsync(new Aws.ServiceQuotas.GetServiceQuotaArgs
        ///         {
        ///             QuotaName = "VPCs per Region",
        ///             ServiceCode = "vpc",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceQuotaResult> InvokeAsync(GetServiceQuotaArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceQuotaResult>("aws:servicequotas/getServiceQuota:getServiceQuota", args ?? new GetServiceQuotaArgs(), options.WithVersion());

        public static Output<GetServiceQuotaResult> Invoke(GetServiceQuotaOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.QuotaCode.Box(),
                args.QuotaName.Box(),
                args.ServiceCode.Box()
            ).Apply(a => {
                    var args = new GetServiceQuotaArgs();
                    a[0].Set(args, nameof(args.QuotaCode));
                    a[1].Set(args, nameof(args.QuotaName));
                    a[2].Set(args, nameof(args.ServiceCode));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetServiceQuotaArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
        /// </summary>
        [Input("quotaCode")]
        public string? QuotaCode { get; set; }

        /// <summary>
        /// Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
        /// </summary>
        [Input("quotaName")]
        public string? QuotaName { get; set; }

        /// <summary>
        /// Service code for the quota. Available values can be found with the `aws.servicequotas.getService` data source or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
        /// </summary>
        [Input("serviceCode", required: true)]
        public string ServiceCode { get; set; } = null!;

        public GetServiceQuotaArgs()
        {
        }
    }

    public sealed class GetServiceQuotaOutputArgs
    {
        /// <summary>
        /// Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
        /// </summary>
        [Input("quotaCode")]
        public Input<string>? QuotaCode { get; set; }

        /// <summary>
        /// Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
        /// </summary>
        [Input("quotaName")]
        public Input<string>? QuotaName { get; set; }

        /// <summary>
        /// Service code for the quota. Available values can be found with the `aws.servicequotas.getService` data source or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
        /// </summary>
        [Input("serviceCode", required: true)]
        public Input<string> ServiceCode { get; set; } = null!;

        public GetServiceQuotaOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceQuotaResult
    {
        /// <summary>
        /// Whether the service quota is adjustable.
        /// </summary>
        public readonly bool Adjustable;
        /// <summary>
        /// Amazon Resource Name (ARN) of the service quota.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Default value of the service quota.
        /// </summary>
        public readonly double DefaultValue;
        /// <summary>
        /// Whether the service quota is global for the AWS account.
        /// </summary>
        public readonly bool GlobalQuota;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string QuotaCode;
        public readonly string QuotaName;
        public readonly string ServiceCode;
        /// <summary>
        /// Name of the service.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Current value of the service quota.
        /// </summary>
        public readonly double Value;

        [OutputConstructor]
        private GetServiceQuotaResult(
            bool adjustable,

            string arn,

            double defaultValue,

            bool globalQuota,

            string id,

            string quotaCode,

            string quotaName,

            string serviceCode,

            string serviceName,

            double value)
        {
            Adjustable = adjustable;
            Arn = arn;
            DefaultValue = defaultValue;
            GlobalQuota = globalQuota;
            Id = id;
            QuotaCode = quotaCode;
            QuotaName = quotaName;
            ServiceCode = serviceCode;
            ServiceName = serviceName;
            Value = value;
        }
    }
}
