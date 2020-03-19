// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticBeanstalk
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ID of an [elastic beanstalk hosted zone](http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region).
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elastic_beanstalk_hosted_zone.html.markdown.
        /// </summary>
        [Obsolete("Use GetHostedZone.InvokeAsync() instead")]
        public static Task<GetHostedZoneResult> GetHostedZone(GetHostedZoneArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHostedZoneResult>("aws:elasticbeanstalk/getHostedZone:getHostedZone", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetHostedZone
    {
        /// <summary>
        /// Use this data source to get the ID of an [elastic beanstalk hosted zone](http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region).
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elastic_beanstalk_hosted_zone.html.markdown.
        /// </summary>
        public static Task<GetHostedZoneResult> InvokeAsync(GetHostedZoneArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHostedZoneResult>("aws:elasticbeanstalk/getHostedZone:getHostedZone", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetHostedZoneArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The region you'd like the zone for. By default, fetches the current region.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetHostedZoneArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetHostedZoneResult
    {
        /// <summary>
        /// The region of the hosted zone.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetHostedZoneResult(
            string? region,
            string id)
        {
            Region = region;
            Id = id;
        }
    }
}
