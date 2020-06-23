// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SSM Document resource
//
// > **NOTE on updating SSM documents:** Only documents with a schema version of 2.0
// or greater can update their content once created, see [SSM Schema Features](http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-ssm-docs.html#document-schemas-features). To update a document with an older
// schema version you must recreate the resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = ssm.NewDocument(ctx, "foo", &ssm.DocumentArgs{
// 			Content:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "  {\n", "    \"schemaVersion\": \"1.2\",\n", "    \"description\": \"Check ip configuration of a Linux instance.\",\n", "    \"parameters\": {\n", "\n", "    },\n", "    \"runtimeConfig\": {\n", "      \"aws:runShellScript\": {\n", "        \"properties\": [\n", "          {\n", "            \"id\": \"0.aws:runShellScript\",\n", "            \"runCommand\": [\"ifconfig\"]\n", "          }\n", "        ]\n", "      }\n", "    }\n", "  }\n", "\n")),
// 			DocumentType: pulumi.String("Command"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Permissions
//
// The permissions attribute specifies how you want to share the document. If you share a document privately,
// you must specify the AWS user account IDs for those people who can use the document. If you share a document
// publicly, you must specify All as the account ID.
//
// The permissions mapping supports the following:
//
// * `type` - The permission type for the document. The permission type can be `Share`.
// * `accountIds` - The AWS user accounts that should have access to the document. The account IDs can either be a group of account IDs or `All`.
type Document struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources DocumentAttachmentsSourceArrayOutput `pulumi:"attachmentsSources"`
	// The JSON or YAML content of the document.
	Content pulumi.StringOutput `pulumi:"content"`
	// The date the document was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The default version of the document.
	DefaultVersion pulumi.StringOutput `pulumi:"defaultVersion"`
	// The description of the document.
	Description pulumi.StringOutput `pulumi:"description"`
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat pulumi.StringPtrOutput `pulumi:"documentFormat"`
	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType pulumi.StringOutput `pulumi:"documentType"`
	// The document version.
	DocumentVersion pulumi.StringOutput `pulumi:"documentVersion"`
	// The sha1 or sha256 of the document content
	Hash pulumi.StringOutput `pulumi:"hash"`
	// "Sha1" "Sha256". The hashing algorithm used when hashing the content.
	HashType pulumi.StringOutput `pulumi:"hashType"`
	// The latest version of the document.
	LatestVersion pulumi.StringOutput `pulumi:"latestVersion"`
	// The name of the document.
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS user account of the person who created the document.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The parameters that are available to this document.
	Parameters DocumentParameterArrayOutput `pulumi:"parameters"`
	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions pulumi.StringMapOutput `pulumi:"permissions"`
	// A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
	PlatformTypes pulumi.StringArrayOutput `pulumi:"platformTypes"`
	// The schema version of the document.
	SchemaVersion pulumi.StringOutput `pulumi:"schemaVersion"`
	// "Creating", "Active" or "Deleting". The current status of the document.
	Status pulumi.StringOutput `pulumi:"status"`
	// A map of tags to assign to the object.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType pulumi.StringPtrOutput `pulumi:"targetType"`
}

// NewDocument registers a new resource with the given unique name, arguments, and options.
func NewDocument(ctx *pulumi.Context,
	name string, args *DocumentArgs, opts ...pulumi.ResourceOption) (*Document, error) {
	if args == nil || args.Content == nil {
		return nil, errors.New("missing required argument 'Content'")
	}
	if args == nil || args.DocumentType == nil {
		return nil, errors.New("missing required argument 'DocumentType'")
	}
	if args == nil {
		args = &DocumentArgs{}
	}
	var resource Document
	err := ctx.RegisterResource("aws:ssm/document:Document", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocument gets an existing Document resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocument(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentState, opts ...pulumi.ResourceOption) (*Document, error) {
	var resource Document
	err := ctx.ReadResource("aws:ssm/document:Document", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Document resources.
type documentState struct {
	Arn *string `pulumi:"arn"`
	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources []DocumentAttachmentsSource `pulumi:"attachmentsSources"`
	// The JSON or YAML content of the document.
	Content *string `pulumi:"content"`
	// The date the document was created.
	CreatedDate *string `pulumi:"createdDate"`
	// The default version of the document.
	DefaultVersion *string `pulumi:"defaultVersion"`
	// The description of the document.
	Description *string `pulumi:"description"`
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat *string `pulumi:"documentFormat"`
	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType *string `pulumi:"documentType"`
	// The document version.
	DocumentVersion *string `pulumi:"documentVersion"`
	// The sha1 or sha256 of the document content
	Hash *string `pulumi:"hash"`
	// "Sha1" "Sha256". The hashing algorithm used when hashing the content.
	HashType *string `pulumi:"hashType"`
	// The latest version of the document.
	LatestVersion *string `pulumi:"latestVersion"`
	// The name of the document.
	Name *string `pulumi:"name"`
	// The AWS user account of the person who created the document.
	Owner *string `pulumi:"owner"`
	// The parameters that are available to this document.
	Parameters []DocumentParameter `pulumi:"parameters"`
	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions map[string]string `pulumi:"permissions"`
	// A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
	PlatformTypes []string `pulumi:"platformTypes"`
	// The schema version of the document.
	SchemaVersion *string `pulumi:"schemaVersion"`
	// "Creating", "Active" or "Deleting". The current status of the document.
	Status *string `pulumi:"status"`
	// A map of tags to assign to the object.
	Tags map[string]interface{} `pulumi:"tags"`
	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType *string `pulumi:"targetType"`
}

type DocumentState struct {
	Arn pulumi.StringPtrInput
	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources DocumentAttachmentsSourceArrayInput
	// The JSON or YAML content of the document.
	Content pulumi.StringPtrInput
	// The date the document was created.
	CreatedDate pulumi.StringPtrInput
	// The default version of the document.
	DefaultVersion pulumi.StringPtrInput
	// The description of the document.
	Description pulumi.StringPtrInput
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat pulumi.StringPtrInput
	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType pulumi.StringPtrInput
	// The document version.
	DocumentVersion pulumi.StringPtrInput
	// The sha1 or sha256 of the document content
	Hash pulumi.StringPtrInput
	// "Sha1" "Sha256". The hashing algorithm used when hashing the content.
	HashType pulumi.StringPtrInput
	// The latest version of the document.
	LatestVersion pulumi.StringPtrInput
	// The name of the document.
	Name pulumi.StringPtrInput
	// The AWS user account of the person who created the document.
	Owner pulumi.StringPtrInput
	// The parameters that are available to this document.
	Parameters DocumentParameterArrayInput
	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions pulumi.StringMapInput
	// A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
	PlatformTypes pulumi.StringArrayInput
	// The schema version of the document.
	SchemaVersion pulumi.StringPtrInput
	// "Creating", "Active" or "Deleting". The current status of the document.
	Status pulumi.StringPtrInput
	// A map of tags to assign to the object.
	Tags pulumi.MapInput
	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType pulumi.StringPtrInput
}

func (DocumentState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentState)(nil)).Elem()
}

type documentArgs struct {
	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources []DocumentAttachmentsSource `pulumi:"attachmentsSources"`
	// The JSON or YAML content of the document.
	Content string `pulumi:"content"`
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat *string `pulumi:"documentFormat"`
	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType string `pulumi:"documentType"`
	// The name of the document.
	Name *string `pulumi:"name"`
	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions map[string]string `pulumi:"permissions"`
	// A map of tags to assign to the object.
	Tags map[string]interface{} `pulumi:"tags"`
	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType *string `pulumi:"targetType"`
}

// The set of arguments for constructing a Document resource.
type DocumentArgs struct {
	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources DocumentAttachmentsSourceArrayInput
	// The JSON or YAML content of the document.
	Content pulumi.StringInput
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat pulumi.StringPtrInput
	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType pulumi.StringInput
	// The name of the document.
	Name pulumi.StringPtrInput
	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions pulumi.StringMapInput
	// A map of tags to assign to the object.
	Tags pulumi.MapInput
	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType pulumi.StringPtrInput
}

func (DocumentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentArgs)(nil)).Elem()
}
