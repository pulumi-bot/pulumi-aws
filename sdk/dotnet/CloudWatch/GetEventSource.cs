// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    public static class GetEventSource
    {
        /// <summary>
        /// Use this data source to get information about an EventBridge Partner Event Source. This data source will only return one partner event source. An error will be returned if multiple sources match the same name prefix.
        /// 
        /// &gt; **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
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
        ///         var examplepartner = Output.Create(Aws.CloudWatch.GetEventSource.InvokeAsync(new Aws.CloudWatch.GetEventSourceArgs
        ///         {
        ///             NamePrefix = "aws.partner/examplepartner.com",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEventSourceResult> InvokeAsync(GetEventSourceArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEventSourceResult>("aws:cloudwatch/getEventSource:getEventSource", args ?? new GetEventSourceArgs(), options.WithVersion());

        public static Output<GetEventSourceResult> Invoke(GetEventSourceOutputArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetEventSourceOutputArgs();
            return Pulumi.Output.All(
                args.NamePrefix.Box()
            ).Apply(a => {
                    var args = new GetEventSourceArgs();
                    a[0].Set(args, nameof(args.NamePrefix));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetEventSourceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifying this limits the results to only those partner event sources with names that start with the specified prefix
        /// </summary>
        [Input("namePrefix")]
        public string? NamePrefix { get; set; }

        public GetEventSourceArgs()
        {
        }
    }

    public sealed class GetEventSourceOutputArgs
    {
        /// <summary>
        /// Specifying this limits the results to only those partner event sources with names that start with the specified prefix
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        public GetEventSourceOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEventSourceResult
    {
        /// <summary>
        /// The ARN of the partner event source
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The name of the SaaS partner that created the event source
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the event source
        /// </summary>
        public readonly string Name;
        public readonly string? NamePrefix;
        /// <summary>
        /// The state of the event source (`ACTIVE` or `PENDING`)
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetEventSourceResult(
            string arn,

            string createdBy,

            string id,

            string name,

            string? namePrefix,

            string state)
        {
            Arn = arn;
            CreatedBy = createdBy;
            Id = id;
            Name = name;
            NamePrefix = namePrefix;
            State = state;
        }
    }
}
