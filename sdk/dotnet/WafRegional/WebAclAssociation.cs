// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafRegional
{
    /// <summary>
    /// Manages an association with WAF Regional Web ACL.
    /// 
    /// &gt; **Note:** An Application Load Balancer can only be associated with one WAF Regional WebACL.
    /// 
    /// ## Application Load Balancer Association Example
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var ipset = new Aws.WafRegional.IpSet("ipset", new Aws.WafRegional.IpSetArgs
    ///         {
    ///             IpSetDescriptors = 
    ///             {
    ///                 new Aws.WafRegional.Inputs.IpSetIpSetDescriptorArgs
    ///                 {
    ///                     Type = "IPV4",
    ///                     Value = "192.0.7.0/24",
    ///                 },
    ///             },
    ///         });
    ///         var fooRule = new Aws.WafRegional.Rule("fooRule", new Aws.WafRegional.RuleArgs
    ///         {
    ///             MetricName = "tfWAFRule",
    ///             Predicates = 
    ///             {
    ///                 new Aws.WafRegional.Inputs.RulePredicateArgs
    ///                 {
    ///                     DataId = ipset.Id,
    ///                     Negated = false,
    ///                     Type = "IPMatch",
    ///                 },
    ///             },
    ///         });
    ///         var fooWebAcl = new Aws.WafRegional.WebAcl("fooWebAcl", new Aws.WafRegional.WebAclArgs
    ///         {
    ///             DefaultAction = new Aws.WafRegional.Inputs.WebAclDefaultActionArgs
    ///             {
    ///                 Type = "ALLOW",
    ///             },
    ///             MetricName = "foo",
    ///             Rules = 
    ///             {
    ///                 new Aws.WafRegional.Inputs.WebAclRuleArgs
    ///                 {
    ///                     Action = new Aws.WafRegional.Inputs.WebAclRuleActionArgs
    ///                     {
    ///                         Type = "BLOCK",
    ///                     },
    ///                     Priority = 1,
    ///                     RuleId = fooRule.Id,
    ///                 },
    ///             },
    ///         });
    ///         var fooVpc = new Aws.Ec2.Vpc("fooVpc", new Aws.Ec2.VpcArgs
    ///         {
    ///             CidrBlock = "10.1.0.0/16",
    ///         });
    ///         var available = Output.Create(Aws.GetAvailabilityZones.InvokeAsync());
    ///         var fooSubnet = new Aws.Ec2.Subnet("fooSubnet", new Aws.Ec2.SubnetArgs
    ///         {
    ///             AvailabilityZone = available.Apply(available =&gt; available.Names[0]),
    ///             CidrBlock = "10.1.1.0/24",
    ///             VpcId = fooVpc.Id,
    ///         });
    ///         var bar = new Aws.Ec2.Subnet("bar", new Aws.Ec2.SubnetArgs
    ///         {
    ///             AvailabilityZone = available.Apply(available =&gt; available.Names[1]),
    ///             CidrBlock = "10.1.2.0/24",
    ///             VpcId = fooVpc.Id,
    ///         });
    ///         var fooLoadBalancer = new Aws.Alb.LoadBalancer("fooLoadBalancer", new Aws.Alb.LoadBalancerArgs
    ///         {
    ///             @internal = true,
    ///             Subnets = 
    ///             {
    ///                 fooSubnet.Id,
    ///                 bar.Id,
    ///             },
    ///         });
    ///         var fooWebAclAssociation = new Aws.WafRegional.WebAclAssociation("fooWebAclAssociation", new Aws.WafRegional.WebAclAssociationArgs
    ///         {
    ///             ResourceArn = fooLoadBalancer.Arn,
    ///             WebAclId = fooWebAcl.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## API Gateway Association Example
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var ipset = new Aws.WafRegional.IpSet("ipset", new Aws.WafRegional.IpSetArgs
    ///         {
    ///             IpSetDescriptors = 
    ///             {
    ///                 new Aws.WafRegional.Inputs.IpSetIpSetDescriptorArgs
    ///                 {
    ///                     Type = "IPV4",
    ///                     Value = "192.0.7.0/24",
    ///                 },
    ///             },
    ///         });
    ///         var fooRule = new Aws.WafRegional.Rule("fooRule", new Aws.WafRegional.RuleArgs
    ///         {
    ///             MetricName = "tfWAFRule",
    ///             Predicates = 
    ///             {
    ///                 new Aws.WafRegional.Inputs.RulePredicateArgs
    ///                 {
    ///                     DataId = ipset.Id,
    ///                     Negated = false,
    ///                     Type = "IPMatch",
    ///                 },
    ///             },
    ///         });
    ///         var fooWebAcl = new Aws.WafRegional.WebAcl("fooWebAcl", new Aws.WafRegional.WebAclArgs
    ///         {
    ///             DefaultAction = new Aws.WafRegional.Inputs.WebAclDefaultActionArgs
    ///             {
    ///                 Type = "ALLOW",
    ///             },
    ///             MetricName = "foo",
    ///             Rules = 
    ///             {
    ///                 new Aws.WafRegional.Inputs.WebAclRuleArgs
    ///                 {
    ///                     Action = new Aws.WafRegional.Inputs.WebAclRuleActionArgs
    ///                     {
    ///                         Type = "BLOCK",
    ///                     },
    ///                     Priority = 1,
    ///                     RuleId = fooRule.Id,
    ///                 },
    ///             },
    ///         });
    ///         var testRestApi = new Aws.ApiGateway.RestApi("testRestApi", new Aws.ApiGateway.RestApiArgs
    ///         {
    ///         });
    ///         var testResource = new Aws.ApiGateway.Resource("testResource", new Aws.ApiGateway.ResourceArgs
    ///         {
    ///             ParentId = testRestApi.RootResourceId,
    ///             PathPart = "test",
    ///             RestApi = testRestApi.Id,
    ///         });
    ///         var testMethod = new Aws.ApiGateway.Method("testMethod", new Aws.ApiGateway.MethodArgs
    ///         {
    ///             Authorization = "NONE",
    ///             HttpMethod = "GET",
    ///             ResourceId = testResource.Id,
    ///             RestApi = testRestApi.Id,
    ///         });
    ///         var testMethodResponse = new Aws.ApiGateway.MethodResponse("testMethodResponse", new Aws.ApiGateway.MethodResponseArgs
    ///         {
    ///             HttpMethod = testMethod.HttpMethod,
    ///             ResourceId = testResource.Id,
    ///             RestApi = testRestApi.Id,
    ///             StatusCode = "400",
    ///         });
    ///         var testIntegration = new Aws.ApiGateway.Integration("testIntegration", new Aws.ApiGateway.IntegrationArgs
    ///         {
    ///             HttpMethod = testMethod.HttpMethod,
    ///             IntegrationHttpMethod = "GET",
    ///             ResourceId = testResource.Id,
    ///             RestApi = testRestApi.Id,
    ///             Type = "HTTP",
    ///             Uri = "http://www.example.com",
    ///         });
    ///         var testIntegrationResponse = new Aws.ApiGateway.IntegrationResponse("testIntegrationResponse", new Aws.ApiGateway.IntegrationResponseArgs
    ///         {
    ///             HttpMethod = testIntegration.HttpMethod,
    ///             ResourceId = testResource.Id,
    ///             RestApi = testRestApi.Id,
    ///             StatusCode = testMethodResponse.StatusCode,
    ///         });
    ///         var testDeployment = new Aws.ApiGateway.Deployment("testDeployment", new Aws.ApiGateway.DeploymentArgs
    ///         {
    ///             RestApi = testRestApi.Id,
    ///         });
    ///         var testStage = new Aws.ApiGateway.Stage("testStage", new Aws.ApiGateway.StageArgs
    ///         {
    ///             Deployment = testDeployment.Id,
    ///             RestApi = testRestApi.Id,
    ///             StageName = "test",
    ///         });
    ///         var association = new Aws.WafRegional.WebAclAssociation("association", new Aws.WafRegional.WebAclAssociationArgs
    ///         {
    ///             ResourceArn = testStage.Arn,
    ///             WebAclId = fooWebAcl.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class WebAclAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
        /// </summary>
        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the WAF Regional WebACL to create an association.
        /// </summary>
        [Output("webAclId")]
        public Output<string> WebAclId { get; private set; } = null!;


        /// <summary>
        /// Create a WebAclAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebAclAssociation(string name, WebAclAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:wafregional/webAclAssociation:WebAclAssociation", name, args ?? new WebAclAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebAclAssociation(string name, Input<string> id, WebAclAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:wafregional/webAclAssociation:WebAclAssociation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WebAclAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebAclAssociation Get(string name, Input<string> id, WebAclAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new WebAclAssociation(name, id, state, options);
        }
    }

    public sealed class WebAclAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
        /// </summary>
        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        /// <summary>
        /// The ID of the WAF Regional WebACL to create an association.
        /// </summary>
        [Input("webAclId", required: true)]
        public Input<string> WebAclId { get; set; } = null!;

        public WebAclAssociationArgs()
        {
        }
    }

    public sealed class WebAclAssociationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
        /// </summary>
        [Input("resourceArn")]
        public Input<string>? ResourceArn { get; set; }

        /// <summary>
        /// The ID of the WAF Regional WebACL to create an association.
        /// </summary>
        [Input("webAclId")]
        public Input<string>? WebAclId { get; set; }

        public WebAclAssociationState()
        {
        }
    }
}
