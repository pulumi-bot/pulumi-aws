// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetDedicatedHost
    {
        /// <summary>
        /// Use this data source to get information about the host when allocating an EC2 Dedicated Host.
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
        ///         var test = new Aws.Ec2.DedicatedHost("test", new Aws.Ec2.DedicatedHostArgs
        ///         {
        ///             AutoPlacement = "on",
        ///             AvailabilityZone = "us-west-1a",
        ///             HostRecovery = "on",
        ///             InstanceType = "c5.18xlarge",
        ///         });
        ///         var testData = test.Id.Apply(id =&gt; Aws.Ec2.GetDedicatedHost.InvokeAsync(new Aws.Ec2.GetDedicatedHostArgs
        ///         {
        ///             HostId = id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDedicatedHostResult> InvokeAsync(GetDedicatedHostArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDedicatedHostResult>("aws:ec2/getDedicatedHost:getDedicatedHost", args ?? new GetDedicatedHostArgs(), options.WithVersion());

        public static Output<GetDedicatedHostResult> Invoke(GetDedicatedHostOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.HostId.Box(),
                args.Tags.ToDict().Box()
            ).Apply(a => {
                    var args = new GetDedicatedHostArgs();
                    a[0].Set(args, nameof(args.HostId));
                    a[1].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetDedicatedHostArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The host ID.
        /// </summary>
        [Input("hostId", required: true)]
        public string HostId { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetDedicatedHostArgs()
        {
        }
    }

    public sealed class GetDedicatedHostOutputArgs
    {
        /// <summary>
        /// The host ID.
        /// </summary>
        [Input("hostId", required: true)]
        public Input<string> HostId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetDedicatedHostOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDedicatedHostResult
    {
        public readonly string AutoPlacement;
        public readonly string AvailabilityZone;
        /// <summary>
        /// The number of cores on the Dedicated Host.
        /// </summary>
        public readonly int Cores;
        /// <summary>
        /// The host ID.
        /// </summary>
        public readonly string HostId;
        public readonly string HostRecovery;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The instance family supported by the Dedicated Host. For example, m5.
        /// * `instance_type` -The instance type supported by the Dedicated Host. For example, m5.large. If the host supports multiple instance types, no instanceType is returned.
        /// </summary>
        public readonly string InstanceFamily;
        public readonly string InstanceState;
        public readonly string InstanceType;
        /// <summary>
        /// The instance family supported by the Dedicated Host. For example, m5.
        /// </summary>
        public readonly int Sockets;
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The total number of vCPUs on the Dedicated Host.
        /// </summary>
        public readonly int TotalVcpus;

        [OutputConstructor]
        private GetDedicatedHostResult(
            string autoPlacement,

            string availabilityZone,

            int cores,

            string hostId,

            string hostRecovery,

            string id,

            string instanceFamily,

            string instanceState,

            string instanceType,

            int sockets,

            ImmutableDictionary<string, string> tags,

            int totalVcpus)
        {
            AutoPlacement = autoPlacement;
            AvailabilityZone = availabilityZone;
            Cores = cores;
            HostId = hostId;
            HostRecovery = hostRecovery;
            Id = id;
            InstanceFamily = instanceFamily;
            InstanceState = instanceState;
            InstanceType = instanceType;
            Sockets = sockets;
            Tags = tags;
            TotalVcpus = totalVcpus;
        }
    }
}
