// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Organizations.Outputs
{

    [OutputType]
    public sealed class GetOrganizationNonMasterAccountResult
    {
        public readonly string Arn;
        public readonly string Email;
        public readonly string Id;
        public readonly string Name;
        public readonly string Status;

        [OutputConstructor]
        private GetOrganizationNonMasterAccountResult(
            string arn,

            string email,

            string id,

            string name,

            string status)
        {
            Arn = arn;
            Email = email;
            Id = id;
            Name = name;
            Status = status;
        }
    }
}
