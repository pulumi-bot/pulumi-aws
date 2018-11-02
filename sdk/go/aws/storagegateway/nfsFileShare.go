// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an AWS Storage Gateway NFS File Share.
type NfsFileShare struct {
	s *pulumi.ResourceState
}

// NewNfsFileShare registers a new resource with the given unique name, arguments, and options.
func NewNfsFileShare(ctx *pulumi.Context,
	name string, args *NfsFileShareArgs, opts ...pulumi.ResourceOpt) (*NfsFileShare, error) {
	if args == nil || args.ClientLists == nil {
		return nil, errors.New("missing required argument 'ClientLists'")
	}
	if args == nil || args.GatewayArn == nil {
		return nil, errors.New("missing required argument 'GatewayArn'")
	}
	if args == nil || args.LocationArn == nil {
		return nil, errors.New("missing required argument 'LocationArn'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["clientLists"] = nil
		inputs["defaultStorageClass"] = nil
		inputs["gatewayArn"] = nil
		inputs["guessMimeTypeEnabled"] = nil
		inputs["kmsEncrypted"] = nil
		inputs["kmsKeyArn"] = nil
		inputs["locationArn"] = nil
		inputs["nfsFileShareDefaults"] = nil
		inputs["objectAcl"] = nil
		inputs["readOnly"] = nil
		inputs["requesterPays"] = nil
		inputs["roleArn"] = nil
		inputs["squash"] = nil
	} else {
		inputs["clientLists"] = args.ClientLists
		inputs["defaultStorageClass"] = args.DefaultStorageClass
		inputs["gatewayArn"] = args.GatewayArn
		inputs["guessMimeTypeEnabled"] = args.GuessMimeTypeEnabled
		inputs["kmsEncrypted"] = args.KmsEncrypted
		inputs["kmsKeyArn"] = args.KmsKeyArn
		inputs["locationArn"] = args.LocationArn
		inputs["nfsFileShareDefaults"] = args.NfsFileShareDefaults
		inputs["objectAcl"] = args.ObjectAcl
		inputs["readOnly"] = args.ReadOnly
		inputs["requesterPays"] = args.RequesterPays
		inputs["roleArn"] = args.RoleArn
		inputs["squash"] = args.Squash
	}
	inputs["arn"] = nil
	inputs["fileshareId"] = nil
	s, err := ctx.RegisterResource("aws:storagegateway/nfsFileShare:NfsFileShare", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NfsFileShare{s: s}, nil
}

// GetNfsFileShare gets an existing NfsFileShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNfsFileShare(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NfsFileShareState, opts ...pulumi.ResourceOpt) (*NfsFileShare, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["clientLists"] = state.ClientLists
		inputs["defaultStorageClass"] = state.DefaultStorageClass
		inputs["fileshareId"] = state.FileshareId
		inputs["gatewayArn"] = state.GatewayArn
		inputs["guessMimeTypeEnabled"] = state.GuessMimeTypeEnabled
		inputs["kmsEncrypted"] = state.KmsEncrypted
		inputs["kmsKeyArn"] = state.KmsKeyArn
		inputs["locationArn"] = state.LocationArn
		inputs["nfsFileShareDefaults"] = state.NfsFileShareDefaults
		inputs["objectAcl"] = state.ObjectAcl
		inputs["readOnly"] = state.ReadOnly
		inputs["requesterPays"] = state.RequesterPays
		inputs["roleArn"] = state.RoleArn
		inputs["squash"] = state.Squash
	}
	s, err := ctx.ReadResource("aws:storagegateway/nfsFileShare:NfsFileShare", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NfsFileShare{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NfsFileShare) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NfsFileShare) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Amazon Resource Name (ARN) of the NFS File Share.
func (r *NfsFileShare) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The list of clients that are allowed to access the file gateway. The list must contain either valid IP addresses or valid CIDR blocks. Set to `["0.0.0.0/0"]` to not limit access. Minimum 1 item. Maximum 100 items.
func (r *NfsFileShare) ClientLists() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["clientLists"])
}

// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
func (r *NfsFileShare) DefaultStorageClass() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultStorageClass"])
}

// ID of the NFS File Share.
func (r *NfsFileShare) FileshareId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fileshareId"])
}

// Amazon Resource Name (ARN) of the file gateway.
func (r *NfsFileShare) GatewayArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["gatewayArn"])
}

// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
func (r *NfsFileShare) GuessMimeTypeEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["guessMimeTypeEnabled"])
}

// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
func (r *NfsFileShare) KmsEncrypted() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["kmsEncrypted"])
}

// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
func (r *NfsFileShare) KmsKeyArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["kmsKeyArn"])
}

// The ARN of the backed storage used for storing file data.
func (r *NfsFileShare) LocationArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["locationArn"])
}

// Nested argument with file share default values. More information below.
func (r *NfsFileShare) NfsFileShareDefaults() *pulumi.Output {
	return r.s.State["nfsFileShareDefaults"]
}

// Access Control List permission for S3 bucket objects. Defaults to `private`.
func (r *NfsFileShare) ObjectAcl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["objectAcl"])
}

// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
func (r *NfsFileShare) ReadOnly() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["readOnly"])
}

// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
func (r *NfsFileShare) RequesterPays() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["requesterPays"])
}

// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
func (r *NfsFileShare) RoleArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["roleArn"])
}

// Maps a user to anonymous user. Defaults to `RootSquash`. Valid values: `RootSquash` (only root is mapped to anonymous user), `NoSquash` (no one is mapped to anonymous user), `AllSquash` (everyone is mapped to anonymous user)
func (r *NfsFileShare) Squash() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["squash"])
}

// Input properties used for looking up and filtering NfsFileShare resources.
type NfsFileShareState struct {
	// Amazon Resource Name (ARN) of the NFS File Share.
	Arn interface{}
	// The list of clients that are allowed to access the file gateway. The list must contain either valid IP addresses or valid CIDR blocks. Set to `["0.0.0.0/0"]` to not limit access. Minimum 1 item. Maximum 100 items.
	ClientLists interface{}
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass interface{}
	// ID of the NFS File Share.
	FileshareId interface{}
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn interface{}
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled interface{}
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted interface{}
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
	KmsKeyArn interface{}
	// The ARN of the backed storage used for storing file data.
	LocationArn interface{}
	// Nested argument with file share default values. More information below.
	NfsFileShareDefaults interface{}
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl interface{}
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly interface{}
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays interface{}
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn interface{}
	// Maps a user to anonymous user. Defaults to `RootSquash`. Valid values: `RootSquash` (only root is mapped to anonymous user), `NoSquash` (no one is mapped to anonymous user), `AllSquash` (everyone is mapped to anonymous user)
	Squash interface{}
}

// The set of arguments for constructing a NfsFileShare resource.
type NfsFileShareArgs struct {
	// The list of clients that are allowed to access the file gateway. The list must contain either valid IP addresses or valid CIDR blocks. Set to `["0.0.0.0/0"]` to not limit access. Minimum 1 item. Maximum 100 items.
	ClientLists interface{}
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass interface{}
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn interface{}
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled interface{}
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted interface{}
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
	KmsKeyArn interface{}
	// The ARN of the backed storage used for storing file data.
	LocationArn interface{}
	// Nested argument with file share default values. More information below.
	NfsFileShareDefaults interface{}
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl interface{}
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly interface{}
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays interface{}
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn interface{}
	// Maps a user to anonymous user. Defaults to `RootSquash`. Valid values: `RootSquash` (only root is mapped to anonymous user), `NoSquash` (no one is mapped to anonymous user), `AllSquash` (everyone is mapped to anonymous user)
	Squash interface{}
}
