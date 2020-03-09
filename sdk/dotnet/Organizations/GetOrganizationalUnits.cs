// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Organizations
{
    public static partial class GetOrganizationalUnits
    {
        /// <summary>
        /// Get all direct child organizational units under a parent organizational unit. This only provides immediate children, not all children.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/organizations_organizational_units.html.markdown.
        /// </summary>
        public static Task<GetOrganizationalUnitsResult> InvokeAsync(GetOrganizationalUnitsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationalUnitsResult>("aws:organizations/getOrganizationalUnits:getOrganizationalUnits", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetOrganizationalUnitsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The parent ID of the organizational unit.
        /// </summary>
        [Input("parentId", required: true)]
        public string ParentId { get; set; } = null!;

        public GetOrganizationalUnitsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetOrganizationalUnitsResult
    {
        /// <summary>
        /// List of child organizational units, which have the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrganizationalUnitsChildrensResult> Childrens;
        public readonly string ParentId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetOrganizationalUnitsResult(
            ImmutableArray<Outputs.GetOrganizationalUnitsChildrensResult> childrens,
            string parentId,
            string id)
        {
            Childrens = childrens;
            ParentId = parentId;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetOrganizationalUnitsChildrensResult
    {
        /// <summary>
        /// ARN of the organizational unit
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// ID of the organizational unit
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the organizational unit
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetOrganizationalUnitsChildrensResult(
            string arn,
            string id,
            string name)
        {
            Arn = arn;
            Id = id;
            Name = name;
        }
    }
    }
}
