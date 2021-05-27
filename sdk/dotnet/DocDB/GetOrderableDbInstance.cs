// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DocDB
{
    public static class GetOrderableDbInstance
    {
        /// <summary>
        /// Information about DocumentDB orderable DB instances.
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
        ///         var test = Output.Create(Aws.DocDB.GetOrderableDbInstance.InvokeAsync(new Aws.DocDB.GetOrderableDbInstanceArgs
        ///         {
        ///             Engine = "docdb",
        ///             EngineVersion = "3.6.0",
        ///             LicenseModel = "na",
        ///             PreferredInstanceClasses = 
        ///             {
        ///                 "db.r5.large",
        ///                 "db.r4.large",
        ///                 "db.t3.medium",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrderableDbInstanceResult> InvokeAsync(GetOrderableDbInstanceArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrderableDbInstanceResult>("aws:docdb/getOrderableDbInstance:getOrderableDbInstance", args ?? new GetOrderableDbInstanceArgs(), options.WithVersion());

        public static Output<GetOrderableDbInstanceResult> Invoke(GetOrderableDbInstanceOutputArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetOrderableDbInstanceOutputArgs();
            return Pulumi.Output.All(
                args.Engine.Box(),
                args.EngineVersion.Box(),
                args.InstanceClass.Box(),
                args.LicenseModel.Box(),
                args.PreferredInstanceClasses.ToList().Box(),
                args.Vpc.Box()
            ).Apply(a => {
                    var args = new GetOrderableDbInstanceArgs();
                    a[0].Set(args, nameof(args.Engine));
                    a[1].Set(args, nameof(args.EngineVersion));
                    a[2].Set(args, nameof(args.InstanceClass));
                    a[3].Set(args, nameof(args.LicenseModel));
                    a[4].Set(args, nameof(args.PreferredInstanceClasses));
                    a[5].Set(args, nameof(args.Vpc));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetOrderableDbInstanceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// DB engine. Default: `docdb`
        /// </summary>
        [Input("engine")]
        public string? Engine { get; set; }

        /// <summary>
        /// Version of the DB engine.
        /// </summary>
        [Input("engineVersion")]
        public string? EngineVersion { get; set; }

        /// <summary>
        /// DB instance class. Examples of classes are `db.r5.12xlarge`, `db.r5.24xlarge`, `db.r5.2xlarge`, `db.r5.4xlarge`, `db.r5.large`, `db.r5.xlarge`, and `db.t3.medium`. (Conflicts with `preferred_instance_classes`.)
        /// </summary>
        [Input("instanceClass")]
        public string? InstanceClass { get; set; }

        /// <summary>
        /// License model. Default: `na`
        /// </summary>
        [Input("licenseModel")]
        public string? LicenseModel { get; set; }

        [Input("preferredInstanceClasses")]
        private List<string>? _preferredInstanceClasses;

        /// <summary>
        /// Ordered list of preferred DocumentDB DB instance classes. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned. (Conflicts with `instance_class`.)
        /// </summary>
        public List<string> PreferredInstanceClasses
        {
            get => _preferredInstanceClasses ?? (_preferredInstanceClasses = new List<string>());
            set => _preferredInstanceClasses = value;
        }

        /// <summary>
        /// Enable to show only VPC.
        /// </summary>
        [Input("vpc")]
        public bool? Vpc { get; set; }

        public GetOrderableDbInstanceArgs()
        {
        }
    }

    public sealed class GetOrderableDbInstanceOutputArgs
    {
        /// <summary>
        /// DB engine. Default: `docdb`
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Version of the DB engine.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// DB instance class. Examples of classes are `db.r5.12xlarge`, `db.r5.24xlarge`, `db.r5.2xlarge`, `db.r5.4xlarge`, `db.r5.large`, `db.r5.xlarge`, and `db.t3.medium`. (Conflicts with `preferred_instance_classes`.)
        /// </summary>
        [Input("instanceClass")]
        public Input<string>? InstanceClass { get; set; }

        /// <summary>
        /// License model. Default: `na`
        /// </summary>
        [Input("licenseModel")]
        public Input<string>? LicenseModel { get; set; }

        [Input("preferredInstanceClasses")]
        private InputList<string>? _preferredInstanceClasses;

        /// <summary>
        /// Ordered list of preferred DocumentDB DB instance classes. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned. (Conflicts with `instance_class`.)
        /// </summary>
        public InputList<string> PreferredInstanceClasses
        {
            get => _preferredInstanceClasses ?? (_preferredInstanceClasses = new InputList<string>());
            set => _preferredInstanceClasses = value;
        }

        /// <summary>
        /// Enable to show only VPC.
        /// </summary>
        [Input("vpc")]
        public Input<bool>? Vpc { get; set; }

        public GetOrderableDbInstanceOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOrderableDbInstanceResult
    {
        /// <summary>
        /// Availability zones where the instance is available.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        public readonly string? Engine;
        public readonly string EngineVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceClass;
        public readonly string? LicenseModel;
        public readonly ImmutableArray<string> PreferredInstanceClasses;
        public readonly bool Vpc;

        [OutputConstructor]
        private GetOrderableDbInstanceResult(
            ImmutableArray<string> availabilityZones,

            string? engine,

            string engineVersion,

            string id,

            string instanceClass,

            string? licenseModel,

            ImmutableArray<string> preferredInstanceClasses,

            bool vpc)
        {
            AvailabilityZones = availabilityZones;
            Engine = engine;
            EngineVersion = engineVersion;
            Id = id;
            InstanceClass = instanceClass;
            LicenseModel = licenseModel;
            PreferredInstanceClasses = preferredInstanceClasses;
            Vpc = vpc;
        }
    }
}
