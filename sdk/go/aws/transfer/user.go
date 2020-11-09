// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package transfer

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a AWS Transfer User resource. Managing SSH keys can be accomplished with the `transfer.SshKey` resource.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/transfer"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooServer, err := transfer.NewServer(ctx, "fooServer", &transfer.ServerArgs{
// 			IdentityProviderType: pulumi.String("SERVICE_MANAGED"),
// 			Tags: pulumi.StringMap{
// 				"NAME": pulumi.String("tf-acc-test-transfer-server"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooRole, err := iam.NewRole(ctx, "fooRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "	\"Version\": \"2012-10-17\",\n", "	\"Statement\": [\n", "		{\n", "		\"Effect\": \"Allow\",\n", "		\"Principal\": {\n", "			\"Service\": \"transfer.amazonaws.com\"\n", "		},\n", "		\"Action\": \"sts:AssumeRole\"\n", "		}\n", "	]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "fooRolePolicy", &iam.RolePolicyArgs{
// 			Role: fooRole.ID(),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "	\"Version\": \"2012-10-17\",\n", "	\"Statement\": [\n", "		{\n", "			\"Sid\": \"AllowFullAccesstoS3\",\n", "			\"Effect\": \"Allow\",\n", "			\"Action\": [\n", "				\"s3:*\"\n", "			],\n", "			\"Resource\": \"*\"\n", "		}\n", "	]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = transfer.NewUser(ctx, "fooUser", &transfer.UserArgs{
// 			ServerId: fooServer.ID(),
// 			UserName: pulumi.String("tftestuser"),
// 			Role:     fooRole.Arn,
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
// Transfer Users can be imported using the `server_id` and `user_name` separated by `/`.
//
// ```sh
//  $ pulumi import aws:transfer/user:User bar s-12345678/test-username
// ```
type User struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of Transfer User
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
	HomeDirectory pulumi.StringPtrOutput `pulumi:"homeDirectory"`
	// Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. documented below.
	HomeDirectoryMappings UserHomeDirectoryMappingArrayOutput `pulumi:"homeDirectoryMappings"`
	// The type of landing directory (folder) you mapped for your users' home directory. Valid values are `PATH` and `LOGICAL`.
	HomeDirectoryType pulumi.StringPtrOutput `pulumi:"homeDirectoryType"`
	// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
	Role pulumi.StringOutput `pulumi:"role"`
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The name used for log in to your SFTP server.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.ServerId == nil {
		return nil, errors.New("missing required argument 'ServerId'")
	}
	if args == nil || args.UserName == nil {
		return nil, errors.New("missing required argument 'UserName'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	var resource User
	err := ctx.RegisterResource("aws:transfer/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("aws:transfer/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Amazon Resource Name (ARN) of Transfer User
	Arn *string `pulumi:"arn"`
	// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
	HomeDirectory *string `pulumi:"homeDirectory"`
	// Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. documented below.
	HomeDirectoryMappings []UserHomeDirectoryMapping `pulumi:"homeDirectoryMappings"`
	// The type of landing directory (folder) you mapped for your users' home directory. Valid values are `PATH` and `LOGICAL`.
	HomeDirectoryType *string `pulumi:"homeDirectoryType"`
	// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
	Policy *string `pulumi:"policy"`
	// Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
	Role *string `pulumi:"role"`
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId *string `pulumi:"serverId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name used for log in to your SFTP server.
	UserName *string `pulumi:"userName"`
}

type UserState struct {
	// Amazon Resource Name (ARN) of Transfer User
	Arn pulumi.StringPtrInput
	// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
	HomeDirectory pulumi.StringPtrInput
	// Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. documented below.
	HomeDirectoryMappings UserHomeDirectoryMappingArrayInput
	// The type of landing directory (folder) you mapped for your users' home directory. Valid values are `PATH` and `LOGICAL`.
	HomeDirectoryType pulumi.StringPtrInput
	// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
	Policy pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
	Role pulumi.StringPtrInput
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name used for log in to your SFTP server.
	UserName pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
	HomeDirectory *string `pulumi:"homeDirectory"`
	// Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. documented below.
	HomeDirectoryMappings []UserHomeDirectoryMapping `pulumi:"homeDirectoryMappings"`
	// The type of landing directory (folder) you mapped for your users' home directory. Valid values are `PATH` and `LOGICAL`.
	HomeDirectoryType *string `pulumi:"homeDirectoryType"`
	// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
	Policy *string `pulumi:"policy"`
	// Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
	Role string `pulumi:"role"`
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId string `pulumi:"serverId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name used for log in to your SFTP server.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
	HomeDirectory pulumi.StringPtrInput
	// Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. documented below.
	HomeDirectoryMappings UserHomeDirectoryMappingArrayInput
	// The type of landing directory (folder) you mapped for your users' home directory. Valid values are `PATH` and `LOGICAL`.
	HomeDirectoryType pulumi.StringPtrInput
	// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
	Policy pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
	Role pulumi.StringInput
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name used for log in to your SFTP server.
	UserName pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}
