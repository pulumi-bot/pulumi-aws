// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetSecurityGroup
    {
        /// <summary>
        /// `aws.ec2.SecurityGroup` provides details about a specific Security Group.
        /// 
        /// This resource can prove useful when a module accepts a Security Group id as
        /// an input variable and needs to, for example, determine the id of the
        /// VPC that the security group belongs to.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example shows how one might accept a Security Group id as a variable
        /// and use this data source to obtain the data necessary to create a subnet.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        /// "TODO: // TODO config"        var selected = Output.Create(Aws.Ec2.GetSecurityGroup.InvokeAsync(new Aws.Ec2.GetSecurityGroupArgs
        ///         {
        ///             Id = securityGroupId,
        ///         }));
        ///         var subnet = new Aws.Ec2.Subnet("subnet", new Aws.Ec2.SubnetArgs
        ///         {
        ///             CidrBlock = "10.0.1.0/24",
        ///             VpcId = selected.Apply(selected =&gt; selected.VpcId),
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSecurityGroupResult> InvokeAsync(GetSecurityGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupResult>("aws:ec2/getSecurityGroup:getSecurityGroup", args ?? new GetSecurityGroupArgs(), options.WithVersion());
    }


    public sealed class GetSecurityGroupArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetSecurityGroupFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetSecurityGroupFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetSecurityGroupFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The id of the specific security group to retrieve.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the field to filter by, as defined by
        /// [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSecurityGroups.html).
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags, each pair of which must exactly match
        /// a pair on the desired security group.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// The id of the VPC that the desired security group belongs to.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetSecurityGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecurityGroupResult
    {
        /// <summary>
        /// The computed ARN of the security group.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The description of the security group.
        /// </summary>
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetSecurityGroupFilterResult> Filters;
        public readonly string Id;
        public readonly string Name;
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string VpcId;

        [OutputConstructor]
        private GetSecurityGroupResult(
            string arn,

            string description,

            ImmutableArray<Outputs.GetSecurityGroupFilterResult> filters,

            string id,

            string name,

            ImmutableDictionary<string, object> tags,

            string vpcId)
        {
            Arn = arn;
            Description = description;
            Filters = filters;
            Id = id;
            Name = name;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
