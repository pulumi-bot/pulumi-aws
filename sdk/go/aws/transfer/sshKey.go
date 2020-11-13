// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package transfer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a AWS Transfer User SSH Key resource.
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
// 		exampleServer, err := transfer.NewServer(ctx, "exampleServer", &transfer.ServerArgs{
// 			IdentityProviderType: pulumi.String("SERVICE_MANAGED"),
// 			Tags: pulumi.StringMap{
// 				"NAME": pulumi.String("tf-acc-test-transfer-server"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "	\"Version\": \"2012-10-17\",\n", "	\"Statement\": [\n", "		{\n", "		\"Effect\": \"Allow\",\n", "		\"Principal\": {\n", "			\"Service\": \"transfer.amazonaws.com\"\n", "		},\n", "		\"Action\": \"sts:AssumeRole\"\n", "		}\n", "	]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleUser, err := transfer.NewUser(ctx, "exampleUser", &transfer.UserArgs{
// 			ServerId: exampleServer.ID(),
// 			UserName: pulumi.String("tftestuser"),
// 			Role:     exampleRole.Arn,
// 			Tags: pulumi.StringMap{
// 				"NAME": pulumi.String("tftestuser"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = transfer.NewSshKey(ctx, "exampleSshKey", &transfer.SshKeyArgs{
// 			ServerId: exampleServer.ID(),
// 			UserName: exampleUser.UserName,
// 			Body:     pulumi.String("... SSH key ..."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "exampleRolePolicy", &iam.RolePolicyArgs{
// 			Role: exampleRole.ID(),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "	\"Version\": \"2012-10-17\",\n", "	\"Statement\": [\n", "		{\n", "			\"Sid\": \"AllowFullAccesstoS3\",\n", "			\"Effect\": \"Allow\",\n", "			\"Action\": [\n", "				\"s3:*\"\n", "			],\n", "			\"Resource\": \"*\"\n", "		}\n", "	]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SshKey struct {
	pulumi.CustomResourceState

	// The public key portion of an SSH key pair.
	Body pulumi.StringOutput `pulumi:"body"`
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// The name of the user account that is assigned to one or more servers.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewSshKey registers a new resource with the given unique name, arguments, and options.
func NewSshKey(ctx *pulumi.Context,
	name string, args *SshKeyArgs, opts ...pulumi.ResourceOption) (*SshKey, error) {
	if args == nil || args.Body == nil {
		return nil, errors.New("missing required argument 'Body'")
	}
	if args == nil || args.ServerId == nil {
		return nil, errors.New("missing required argument 'ServerId'")
	}
	if args == nil || args.UserName == nil {
		return nil, errors.New("missing required argument 'UserName'")
	}
	if args == nil {
		args = &SshKeyArgs{}
	}
	var resource SshKey
	err := ctx.RegisterResource("aws:transfer/sshKey:SshKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSshKey gets an existing SshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SshKeyState, opts ...pulumi.ResourceOption) (*SshKey, error) {
	var resource SshKey
	err := ctx.ReadResource("aws:transfer/sshKey:SshKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SshKey resources.
type sshKeyState struct {
	// The public key portion of an SSH key pair.
	Body *string `pulumi:"body"`
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId *string `pulumi:"serverId"`
	// The name of the user account that is assigned to one or more servers.
	UserName *string `pulumi:"userName"`
}

type SshKeyState struct {
	// The public key portion of an SSH key pair.
	Body pulumi.StringPtrInput
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId pulumi.StringPtrInput
	// The name of the user account that is assigned to one or more servers.
	UserName pulumi.StringPtrInput
}

func (SshKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshKeyState)(nil)).Elem()
}

type sshKeyArgs struct {
	// The public key portion of an SSH key pair.
	Body string `pulumi:"body"`
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId string `pulumi:"serverId"`
	// The name of the user account that is assigned to one or more servers.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a SshKey resource.
type SshKeyArgs struct {
	// The public key portion of an SSH key pair.
	Body pulumi.StringInput
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId pulumi.StringInput
	// The name of the user account that is assigned to one or more servers.
	UserName pulumi.StringInput
}

func (SshKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshKeyArgs)(nil)).Elem()
}

type SshKeyInput interface {
	pulumi.Input

	ToSshKeyOutput() SshKeyOutput
	ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput
}

func (SshKey) ElementType() reflect.Type {
	return reflect.TypeOf((*SshKey)(nil)).Elem()
}

func (i SshKey) ToSshKeyOutput() SshKeyOutput {
	return i.ToSshKeyOutputWithContext(context.Background())
}

func (i SshKey) ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshKeyOutput)
}

type SshKeyOutput struct {
	*pulumi.OutputState
}

func (SshKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshKeyOutput)(nil)).Elem()
}

func (o SshKeyOutput) ToSshKeyOutput() SshKeyOutput {
	return o
}

func (o SshKeyOutput) ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SshKeyOutput{})
}
