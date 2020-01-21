// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito
{
    /// <summary>
    /// Provides an AWS Cognito Identity Pool Roles Attachment.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_identity_pool_roles_attachment.html.markdown.
    /// </summary>
    public partial class IdentityPoolRoleAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// An identity pool ID in the format REGION:GUID.
        /// </summary>
        [Output("identityPoolId")]
        public Output<string> IdentityPoolId { get; private set; } = null!;

        /// <summary>
        /// A List of Role Mapping.
        /// </summary>
        [Output("roleMappings")]
        public Output<ImmutableArray<Outputs.IdentityPoolRoleAttachmentRoleMappings>> RoleMappings { get; private set; } = null!;

        /// <summary>
        /// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
        /// </summary>
        [Output("roles")]
        public Output<Outputs.IdentityPoolRoleAttachmentRoles> Roles { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityPoolRoleAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityPoolRoleAttachment(string name, IdentityPoolRoleAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private IdentityPoolRoleAttachment(string name, Input<string> id, IdentityPoolRoleAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityPoolRoleAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityPoolRoleAttachment Get(string name, Input<string> id, IdentityPoolRoleAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityPoolRoleAttachment(name, id, state, options);
        }
    }

    public sealed class IdentityPoolRoleAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An identity pool ID in the format REGION:GUID.
        /// </summary>
        [Input("identityPoolId", required: true)]
        public Input<string> IdentityPoolId { get; set; } = null!;

        [Input("roleMappings")]
        private InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingsArgs>? _roleMappings;

        /// <summary>
        /// A List of Role Mapping.
        /// </summary>
        public InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingsArgs> RoleMappings
        {
            get => _roleMappings ?? (_roleMappings = new InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingsArgs>());
            set => _roleMappings = value;
        }

        /// <summary>
        /// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
        /// </summary>
        [Input("roles", required: true)]
        public Input<Inputs.IdentityPoolRoleAttachmentRolesArgs> Roles { get; set; } = null!;

        public IdentityPoolRoleAttachmentArgs()
        {
        }
    }

    public sealed class IdentityPoolRoleAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An identity pool ID in the format REGION:GUID.
        /// </summary>
        [Input("identityPoolId")]
        public Input<string>? IdentityPoolId { get; set; }

        [Input("roleMappings")]
        private InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingsGetArgs>? _roleMappings;

        /// <summary>
        /// A List of Role Mapping.
        /// </summary>
        public InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingsGetArgs> RoleMappings
        {
            get => _roleMappings ?? (_roleMappings = new InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingsGetArgs>());
            set => _roleMappings = value;
        }

        /// <summary>
        /// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
        /// </summary>
        [Input("roles")]
        public Input<Inputs.IdentityPoolRoleAttachmentRolesGetArgs>? Roles { get; set; }

        public IdentityPoolRoleAttachmentState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class IdentityPoolRoleAttachmentRoleMappingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
        /// </summary>
        [Input("ambiguousRoleResolution")]
        public Input<string>? AmbiguousRoleResolution { get; set; }

        /// <summary>
        /// A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
        /// </summary>
        [Input("identityProvider", required: true)]
        public Input<string> IdentityProvider { get; set; } = null!;

        [Input("mappingRules")]
        private InputList<IdentityPoolRoleAttachmentRoleMappingsMappingRulesArgs>? _mappingRules;

        /// <summary>
        /// The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
        /// </summary>
        public InputList<IdentityPoolRoleAttachmentRoleMappingsMappingRulesArgs> MappingRules
        {
            get => _mappingRules ?? (_mappingRules = new InputList<IdentityPoolRoleAttachmentRoleMappingsMappingRulesArgs>());
            set => _mappingRules = value;
        }

        /// <summary>
        /// The role mapping type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IdentityPoolRoleAttachmentRoleMappingsArgs()
        {
        }
    }

    public sealed class IdentityPoolRoleAttachmentRoleMappingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
        /// </summary>
        [Input("ambiguousRoleResolution")]
        public Input<string>? AmbiguousRoleResolution { get; set; }

        /// <summary>
        /// A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
        /// </summary>
        [Input("identityProvider", required: true)]
        public Input<string> IdentityProvider { get; set; } = null!;

        [Input("mappingRules")]
        private InputList<IdentityPoolRoleAttachmentRoleMappingsMappingRulesGetArgs>? _mappingRules;

        /// <summary>
        /// The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
        /// </summary>
        public InputList<IdentityPoolRoleAttachmentRoleMappingsMappingRulesGetArgs> MappingRules
        {
            get => _mappingRules ?? (_mappingRules = new InputList<IdentityPoolRoleAttachmentRoleMappingsMappingRulesGetArgs>());
            set => _mappingRules = value;
        }

        /// <summary>
        /// The role mapping type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IdentityPoolRoleAttachmentRoleMappingsGetArgs()
        {
        }
    }

    public sealed class IdentityPoolRoleAttachmentRoleMappingsMappingRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The claim name that must be present in the token, for example, "isAdmin" or "paid".
        /// </summary>
        [Input("claim", required: true)]
        public Input<string> Claim { get; set; } = null!;

        /// <summary>
        /// The match condition that specifies how closely the claim value in the IdP token must match Value.
        /// </summary>
        [Input("matchType", required: true)]
        public Input<string> MatchType { get; set; } = null!;

        /// <summary>
        /// The role ARN.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// A brief string that the claim must match, for example, "paid" or "yes".
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public IdentityPoolRoleAttachmentRoleMappingsMappingRulesArgs()
        {
        }
    }

    public sealed class IdentityPoolRoleAttachmentRoleMappingsMappingRulesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The claim name that must be present in the token, for example, "isAdmin" or "paid".
        /// </summary>
        [Input("claim", required: true)]
        public Input<string> Claim { get; set; } = null!;

        /// <summary>
        /// The match condition that specifies how closely the claim value in the IdP token must match Value.
        /// </summary>
        [Input("matchType", required: true)]
        public Input<string> MatchType { get; set; } = null!;

        /// <summary>
        /// The role ARN.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// A brief string that the claim must match, for example, "paid" or "yes".
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public IdentityPoolRoleAttachmentRoleMappingsMappingRulesGetArgs()
        {
        }
    }

    public sealed class IdentityPoolRoleAttachmentRolesArgs : Pulumi.ResourceArgs
    {
        [Input("authenticated")]
        public Input<string>? Authenticated { get; set; }

        [Input("unauthenticated")]
        public Input<string>? Unauthenticated { get; set; }

        public IdentityPoolRoleAttachmentRolesArgs()
        {
        }
    }

    public sealed class IdentityPoolRoleAttachmentRolesGetArgs : Pulumi.ResourceArgs
    {
        [Input("authenticated")]
        public Input<string>? Authenticated { get; set; }

        [Input("unauthenticated")]
        public Input<string>? Unauthenticated { get; set; }

        public IdentityPoolRoleAttachmentRolesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class IdentityPoolRoleAttachmentRoleMappings
    {
        /// <summary>
        /// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
        /// </summary>
        public readonly string? AmbiguousRoleResolution;
        /// <summary>
        /// A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
        /// </summary>
        public readonly string IdentityProvider;
        /// <summary>
        /// The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
        /// </summary>
        public readonly ImmutableArray<IdentityPoolRoleAttachmentRoleMappingsMappingRules> MappingRules;
        /// <summary>
        /// The role mapping type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private IdentityPoolRoleAttachmentRoleMappings(
            string? ambiguousRoleResolution,
            string identityProvider,
            ImmutableArray<IdentityPoolRoleAttachmentRoleMappingsMappingRules> mappingRules,
            string type)
        {
            AmbiguousRoleResolution = ambiguousRoleResolution;
            IdentityProvider = identityProvider;
            MappingRules = mappingRules;
            Type = type;
        }
    }

    [OutputType]
    public sealed class IdentityPoolRoleAttachmentRoleMappingsMappingRules
    {
        /// <summary>
        /// The claim name that must be present in the token, for example, "isAdmin" or "paid".
        /// </summary>
        public readonly string Claim;
        /// <summary>
        /// The match condition that specifies how closely the claim value in the IdP token must match Value.
        /// </summary>
        public readonly string MatchType;
        /// <summary>
        /// The role ARN.
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// A brief string that the claim must match, for example, "paid" or "yes".
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private IdentityPoolRoleAttachmentRoleMappingsMappingRules(
            string claim,
            string matchType,
            string roleArn,
            string value)
        {
            Claim = claim;
            MatchType = matchType;
            RoleArn = roleArn;
            Value = value;
        }
    }

    [OutputType]
    public sealed class IdentityPoolRoleAttachmentRoles
    {
        public readonly string? Authenticated;
        public readonly string? Unauthenticated;

        [OutputConstructor]
        private IdentityPoolRoleAttachmentRoles(
            string? authenticated,
            string? unauthenticated)
        {
            Authenticated = authenticated;
            Unauthenticated = unauthenticated;
        }
    }
    }
}
