// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot.Outputs
{

    [OutputType]
    public sealed class TopicRuleErrorActionElasticsearch
    {
        public readonly string Endpoint;
        public readonly string Id;
        public readonly string Index;
        public readonly string RoleArn;
        public readonly string Type;

        [OutputConstructor]
        private TopicRuleErrorActionElasticsearch(
            string endpoint,

            string id,

            string index,

            string roleArn,

            string type)
        {
            Endpoint = endpoint;
            Id = id;
            Index = index;
            RoleArn = roleArn;
            Type = type;
        }
    }
}
