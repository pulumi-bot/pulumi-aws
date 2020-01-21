// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package sshKey

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a AWS Transfer User SSH Key resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/transfer_ssh_key.html.markdown.
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

