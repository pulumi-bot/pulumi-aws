// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Cognito Identity Pool Roles Attachment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cognito"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		mainIdentityPool, err := cognito.NewIdentityPool(ctx, "mainIdentityPool", &cognito.IdentityPoolArgs{
// 			IdentityPoolName:               pulumi.String("identity pool"),
// 			AllowUnauthenticatedIdentities: pulumi.Bool(false),
// 			SupportedLoginProviders: pulumi.StringMap{
// 				"graph.facebook.com": pulumi.String("7346241598935555"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		authenticatedRole, err := iam.NewRole(ctx, "authenticatedRole", &iam.RoleArgs{
// 			AssumeRolePolicy: mainIdentityPool.ID().ApplyT(func(id string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Federated\": \"cognito-identity.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRoleWithWebIdentity\",\n", "      \"Condition\": {\n", "        \"StringEquals\": {\n", "          \"cognito-identity.amazonaws.com:aud\": \"", id, "\"\n", "        },\n", "        \"ForAnyValue:StringLike\": {\n", "          \"cognito-identity.amazonaws.com:amr\": \"authenticated\"\n", "        }\n", "      }\n", "    }\n", "  ]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "authenticatedRolePolicy", &iam.RolePolicyArgs{
// 			Role:   authenticatedRole.ID(),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Action\": [\n", "        \"mobileanalytics:PutEvents\",\n", "        \"cognito-sync:*\",\n", "        \"cognito-identity:*\"\n", "      ],\n", "      \"Resource\": [\n", "        \"*\"\n", "      ]\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cognito.NewIdentityPoolRoleAttachment(ctx, "mainIdentityPoolRoleAttachment", &cognito.IdentityPoolRoleAttachmentArgs{
// 			IdentityPoolId: mainIdentityPool.ID(),
// 			RoleMappings: cognito.IdentityPoolRoleAttachmentRoleMappingArray{
// 				&cognito.IdentityPoolRoleAttachmentRoleMappingArgs{
// 					IdentityProvider:        pulumi.String("graph.facebook.com"),
// 					AmbiguousRoleResolution: pulumi.String("AuthenticatedRole"),
// 					Type:                    pulumi.String("Rules"),
// 					MappingRules: cognito.IdentityPoolRoleAttachmentRoleMappingMappingRuleArray{
// 						&cognito.IdentityPoolRoleAttachmentRoleMappingMappingRuleArgs{
// 							Claim:     pulumi.String("isAdmin"),
// 							MatchType: pulumi.String("Equals"),
// 							RoleArn:   authenticatedRole.Arn,
// 							Value:     pulumi.String("paid"),
// 						},
// 					},
// 				},
// 			},
// 			Roles: pulumi.StringMap{
// 				"authenticated": authenticatedRole.Arn,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Cognito Identity Pool Roles Attachment can be imported using the Identity Pool id, e.g.
//
// ```sh
//  $ pulumi import aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment example <identity-pool-id>
// ```
type IdentityPoolRoleAttachment struct {
	pulumi.CustomResourceState

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId pulumi.StringOutput `pulumi:"identityPoolId"`
	// A List of Role Mapping.
	RoleMappings IdentityPoolRoleAttachmentRoleMappingArrayOutput `pulumi:"roleMappings"`
	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles pulumi.StringMapOutput `pulumi:"roles"`
}

// NewIdentityPoolRoleAttachment registers a new resource with the given unique name, arguments, and options.
func NewIdentityPoolRoleAttachment(ctx *pulumi.Context,
	name string, args *IdentityPoolRoleAttachmentArgs, opts ...pulumi.ResourceOption) (*IdentityPoolRoleAttachment, error) {
	if args == nil || args.IdentityPoolId == nil {
		return nil, errors.New("missing required argument 'IdentityPoolId'")
	}
	if args == nil || args.Roles == nil {
		return nil, errors.New("missing required argument 'Roles'")
	}
	if args == nil {
		args = &IdentityPoolRoleAttachmentArgs{}
	}
	var resource IdentityPoolRoleAttachment
	err := ctx.RegisterResource("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityPoolRoleAttachment gets an existing IdentityPoolRoleAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityPoolRoleAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityPoolRoleAttachmentState, opts ...pulumi.ResourceOption) (*IdentityPoolRoleAttachment, error) {
	var resource IdentityPoolRoleAttachment
	err := ctx.ReadResource("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityPoolRoleAttachment resources.
type identityPoolRoleAttachmentState struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId *string `pulumi:"identityPoolId"`
	// A List of Role Mapping.
	RoleMappings []IdentityPoolRoleAttachmentRoleMapping `pulumi:"roleMappings"`
	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles map[string]string `pulumi:"roles"`
}

type IdentityPoolRoleAttachmentState struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId pulumi.StringPtrInput
	// A List of Role Mapping.
	RoleMappings IdentityPoolRoleAttachmentRoleMappingArrayInput
	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles pulumi.StringMapInput
}

func (IdentityPoolRoleAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityPoolRoleAttachmentState)(nil)).Elem()
}

type identityPoolRoleAttachmentArgs struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId string `pulumi:"identityPoolId"`
	// A List of Role Mapping.
	RoleMappings []IdentityPoolRoleAttachmentRoleMapping `pulumi:"roleMappings"`
	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles map[string]string `pulumi:"roles"`
}

// The set of arguments for constructing a IdentityPoolRoleAttachment resource.
type IdentityPoolRoleAttachmentArgs struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId pulumi.StringInput
	// A List of Role Mapping.
	RoleMappings IdentityPoolRoleAttachmentRoleMappingArrayInput
	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles pulumi.StringMapInput
}

func (IdentityPoolRoleAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityPoolRoleAttachmentArgs)(nil)).Elem()
}
