// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package WindowsFileSystemSelfManagedActiveDirectory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type WindowsFileSystemSelfManagedActiveDirectory struct {
	// A list of up to two IP addresses of DNS servers or domain controllers in the self-managed AD directory. The IP addresses need to be either in the same VPC CIDR range as the file system or in the private IP version 4 (IPv4) address ranges as specified in [RFC 1918](https://tools.ietf.org/html/rfc1918).
	DnsIps []string `pulumi:"dnsIps"`
	// The fully qualified domain name of the self-managed AD directory. For example, `corp.example.com`.
	DomainName string `pulumi:"domainName"`
	// The name of the domain group whose members are granted administrative privileges for the file system. Administrative privileges include taking ownership of files and folders, and setting audit controls (audit ACLs) on files and folders. The group that you specify must already exist in your domain. Defaults to `Domain Admins`.
	FileSystemAdministratorsGroup *string `pulumi:"fileSystemAdministratorsGroup"`
	// The fully qualified distinguished name of the organizational unit within your self-managed AD directory that the Windows File Server instance will join. For example, `OU=FSx,DC=yourdomain,DC=corp,DC=com`. Only accepts OU as the direct parent of the file system. If none is provided, the FSx file system is created in the default location of your self-managed AD directory. To learn more, see [RFC 2253](https://tools.ietf.org/html/rfc2253).
	OrganizationalUnitDistinguishedName *string `pulumi:"organizationalUnitDistinguishedName"`
	// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	Password string `pulumi:"password"`
	// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	Username string `pulumi:"username"`
}

type WindowsFileSystemSelfManagedActiveDirectoryInput interface {
	pulumi.Input

	ToWindowsFileSystemSelfManagedActiveDirectoryOutput() WindowsFileSystemSelfManagedActiveDirectoryOutput
	ToWindowsFileSystemSelfManagedActiveDirectoryOutputWithContext(context.Context) WindowsFileSystemSelfManagedActiveDirectoryOutput
}

type WindowsFileSystemSelfManagedActiveDirectoryArgs struct {
	// A list of up to two IP addresses of DNS servers or domain controllers in the self-managed AD directory. The IP addresses need to be either in the same VPC CIDR range as the file system or in the private IP version 4 (IPv4) address ranges as specified in [RFC 1918](https://tools.ietf.org/html/rfc1918).
	DnsIps pulumi.StringArrayInput `pulumi:"dnsIps"`
	// The fully qualified domain name of the self-managed AD directory. For example, `corp.example.com`.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The name of the domain group whose members are granted administrative privileges for the file system. Administrative privileges include taking ownership of files and folders, and setting audit controls (audit ACLs) on files and folders. The group that you specify must already exist in your domain. Defaults to `Domain Admins`.
	FileSystemAdministratorsGroup pulumi.StringPtrInput `pulumi:"fileSystemAdministratorsGroup"`
	// The fully qualified distinguished name of the organizational unit within your self-managed AD directory that the Windows File Server instance will join. For example, `OU=FSx,DC=yourdomain,DC=corp,DC=com`. Only accepts OU as the direct parent of the file system. If none is provided, the FSx file system is created in the default location of your self-managed AD directory. To learn more, see [RFC 2253](https://tools.ietf.org/html/rfc2253).
	OrganizationalUnitDistinguishedName pulumi.StringPtrInput `pulumi:"organizationalUnitDistinguishedName"`
	// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	Password pulumi.StringInput `pulumi:"password"`
	// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	Username pulumi.StringInput `pulumi:"username"`
}

func (WindowsFileSystemSelfManagedActiveDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsFileSystemSelfManagedActiveDirectory)(nil)).Elem()
}

func (i WindowsFileSystemSelfManagedActiveDirectoryArgs) ToWindowsFileSystemSelfManagedActiveDirectoryOutput() WindowsFileSystemSelfManagedActiveDirectoryOutput {
	return i.ToWindowsFileSystemSelfManagedActiveDirectoryOutputWithContext(context.Background())
}

func (i WindowsFileSystemSelfManagedActiveDirectoryArgs) ToWindowsFileSystemSelfManagedActiveDirectoryOutputWithContext(ctx context.Context) WindowsFileSystemSelfManagedActiveDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsFileSystemSelfManagedActiveDirectoryOutput)
}

func (i WindowsFileSystemSelfManagedActiveDirectoryArgs) ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutput() WindowsFileSystemSelfManagedActiveDirectoryPtrOutput {
	return i.ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i WindowsFileSystemSelfManagedActiveDirectoryArgs) ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutputWithContext(ctx context.Context) WindowsFileSystemSelfManagedActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsFileSystemSelfManagedActiveDirectoryOutput).ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutputWithContext(ctx)
}

type WindowsFileSystemSelfManagedActiveDirectoryPtrInput interface {
	pulumi.Input

	ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutput() WindowsFileSystemSelfManagedActiveDirectoryPtrOutput
	ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutputWithContext(context.Context) WindowsFileSystemSelfManagedActiveDirectoryPtrOutput
}

type windowsFileSystemSelfManagedActiveDirectoryPtrType WindowsFileSystemSelfManagedActiveDirectoryArgs

func WindowsFileSystemSelfManagedActiveDirectoryPtr(v *WindowsFileSystemSelfManagedActiveDirectoryArgs) WindowsFileSystemSelfManagedActiveDirectoryPtrInput {	return (*windowsFileSystemSelfManagedActiveDirectoryPtrType)(v)
}

func (*windowsFileSystemSelfManagedActiveDirectoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsFileSystemSelfManagedActiveDirectory)(nil)).Elem()
}

func (i *windowsFileSystemSelfManagedActiveDirectoryPtrType) ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutput() WindowsFileSystemSelfManagedActiveDirectoryPtrOutput {
	return i.ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i *windowsFileSystemSelfManagedActiveDirectoryPtrType) ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutputWithContext(ctx context.Context) WindowsFileSystemSelfManagedActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsFileSystemSelfManagedActiveDirectoryPtrOutput)
}

type WindowsFileSystemSelfManagedActiveDirectoryOutput struct { *pulumi.OutputState }

func (WindowsFileSystemSelfManagedActiveDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsFileSystemSelfManagedActiveDirectory)(nil)).Elem()
}

func (o WindowsFileSystemSelfManagedActiveDirectoryOutput) ToWindowsFileSystemSelfManagedActiveDirectoryOutput() WindowsFileSystemSelfManagedActiveDirectoryOutput {
	return o
}

func (o WindowsFileSystemSelfManagedActiveDirectoryOutput) ToWindowsFileSystemSelfManagedActiveDirectoryOutputWithContext(ctx context.Context) WindowsFileSystemSelfManagedActiveDirectoryOutput {
	return o
}

func (o WindowsFileSystemSelfManagedActiveDirectoryOutput) ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutput() WindowsFileSystemSelfManagedActiveDirectoryPtrOutput {
	return o.ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutputWithContext(context.Background())
}

func (o WindowsFileSystemSelfManagedActiveDirectoryOutput) ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutputWithContext(ctx context.Context) WindowsFileSystemSelfManagedActiveDirectoryPtrOutput {
	return o.ApplyT(func(v WindowsFileSystemSelfManagedActiveDirectory) *WindowsFileSystemSelfManagedActiveDirectory {
		return &v
	}).(WindowsFileSystemSelfManagedActiveDirectoryPtrOutput)
}
// A list of up to two IP addresses of DNS servers or domain controllers in the self-managed AD directory. The IP addresses need to be either in the same VPC CIDR range as the file system or in the private IP version 4 (IPv4) address ranges as specified in [RFC 1918](https://tools.ietf.org/html/rfc1918).
func (o WindowsFileSystemSelfManagedActiveDirectoryOutput) DnsIps() pulumi.StringArrayOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) []string { return v.DnsIps }).(pulumi.StringArrayOutput)
}

// The fully qualified domain name of the self-managed AD directory. For example, `corp.example.com`.
func (o WindowsFileSystemSelfManagedActiveDirectoryOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) string { return v.DomainName }).(pulumi.StringOutput)
}

// The name of the domain group whose members are granted administrative privileges for the file system. Administrative privileges include taking ownership of files and folders, and setting audit controls (audit ACLs) on files and folders. The group that you specify must already exist in your domain. Defaults to `Domain Admins`.
func (o WindowsFileSystemSelfManagedActiveDirectoryOutput) FileSystemAdministratorsGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) *string { return v.FileSystemAdministratorsGroup }).(pulumi.StringPtrOutput)
}

// The fully qualified distinguished name of the organizational unit within your self-managed AD directory that the Windows File Server instance will join. For example, `OU=FSx,DC=yourdomain,DC=corp,DC=com`. Only accepts OU as the direct parent of the file system. If none is provided, the FSx file system is created in the default location of your self-managed AD directory. To learn more, see [RFC 2253](https://tools.ietf.org/html/rfc2253).
func (o WindowsFileSystemSelfManagedActiveDirectoryOutput) OrganizationalUnitDistinguishedName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) *string { return v.OrganizationalUnitDistinguishedName }).(pulumi.StringPtrOutput)
}

// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
func (o WindowsFileSystemSelfManagedActiveDirectoryOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) string { return v.Password }).(pulumi.StringOutput)
}

// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
func (o WindowsFileSystemSelfManagedActiveDirectoryOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) string { return v.Username }).(pulumi.StringOutput)
}

type WindowsFileSystemSelfManagedActiveDirectoryPtrOutput struct { *pulumi.OutputState}

func (WindowsFileSystemSelfManagedActiveDirectoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsFileSystemSelfManagedActiveDirectory)(nil)).Elem()
}

func (o WindowsFileSystemSelfManagedActiveDirectoryPtrOutput) ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutput() WindowsFileSystemSelfManagedActiveDirectoryPtrOutput {
	return o
}

func (o WindowsFileSystemSelfManagedActiveDirectoryPtrOutput) ToWindowsFileSystemSelfManagedActiveDirectoryPtrOutputWithContext(ctx context.Context) WindowsFileSystemSelfManagedActiveDirectoryPtrOutput {
	return o
}

func (o WindowsFileSystemSelfManagedActiveDirectoryPtrOutput) Elem() WindowsFileSystemSelfManagedActiveDirectoryOutput {
	return o.ApplyT(func (v *WindowsFileSystemSelfManagedActiveDirectory) WindowsFileSystemSelfManagedActiveDirectory { return *v }).(WindowsFileSystemSelfManagedActiveDirectoryOutput)
}

// A list of up to two IP addresses of DNS servers or domain controllers in the self-managed AD directory. The IP addresses need to be either in the same VPC CIDR range as the file system or in the private IP version 4 (IPv4) address ranges as specified in [RFC 1918](https://tools.ietf.org/html/rfc1918).
func (o WindowsFileSystemSelfManagedActiveDirectoryPtrOutput) DnsIps() pulumi.StringArrayOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) []string { return v.DnsIps }).(pulumi.StringArrayOutput)
}

// The fully qualified domain name of the self-managed AD directory. For example, `corp.example.com`.
func (o WindowsFileSystemSelfManagedActiveDirectoryPtrOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) string { return v.DomainName }).(pulumi.StringOutput)
}

// The name of the domain group whose members are granted administrative privileges for the file system. Administrative privileges include taking ownership of files and folders, and setting audit controls (audit ACLs) on files and folders. The group that you specify must already exist in your domain. Defaults to `Domain Admins`.
func (o WindowsFileSystemSelfManagedActiveDirectoryPtrOutput) FileSystemAdministratorsGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) *string { return v.FileSystemAdministratorsGroup }).(pulumi.StringPtrOutput)
}

// The fully qualified distinguished name of the organizational unit within your self-managed AD directory that the Windows File Server instance will join. For example, `OU=FSx,DC=yourdomain,DC=corp,DC=com`. Only accepts OU as the direct parent of the file system. If none is provided, the FSx file system is created in the default location of your self-managed AD directory. To learn more, see [RFC 2253](https://tools.ietf.org/html/rfc2253).
func (o WindowsFileSystemSelfManagedActiveDirectoryPtrOutput) OrganizationalUnitDistinguishedName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) *string { return v.OrganizationalUnitDistinguishedName }).(pulumi.StringPtrOutput)
}

// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
func (o WindowsFileSystemSelfManagedActiveDirectoryPtrOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) string { return v.Password }).(pulumi.StringOutput)
}

// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
func (o WindowsFileSystemSelfManagedActiveDirectoryPtrOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func (v WindowsFileSystemSelfManagedActiveDirectory) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WindowsFileSystemSelfManagedActiveDirectoryOutput{})
	pulumi.RegisterOutputType(WindowsFileSystemSelfManagedActiveDirectoryPtrOutput{})
}
