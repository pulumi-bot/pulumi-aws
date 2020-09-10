// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeDeploy.Outputs
{

    [OutputType]
    public sealed class DeploymentGroupEc2TagSet
    {
        public readonly ImmutableArray<Outputs.DeploymentGroupEc2TagSetEc2TagFilter> Ec2TagFilters;

        [OutputConstructor]
        private DeploymentGroupEc2TagSet(ImmutableArray<Outputs.DeploymentGroupEc2TagSetEc2TagFilter> ec2TagFilters)
        {
            Ec2TagFilters = ec2TagFilters;
        }
    }
}
