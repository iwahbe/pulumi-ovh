// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a custom installation template available for dedicated servers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Me.GetInstallationTemplate(ctx, &me.GetInstallationTemplateArgs{
//				TemplateName: "mytemplate",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupInstallationTemplate(ctx *pulumi.Context, args *LookupInstallationTemplateArgs, opts ...pulumi.InvokeOption) (*LookupInstallationTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstallationTemplateResult
	err := ctx.Invoke("ovh:Me/getInstallationTemplate:getInstallationTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstallationTemplate.
type LookupInstallationTemplateArgs struct {
	// This template name
	TemplateName string `pulumi:"templateName"`
}

// A collection of values returned by getInstallationTemplate.
type LookupInstallationTemplateResult struct {
	// List of all language available for this template.
	AvailableLanguages []string `pulumi:"availableLanguages"`
	// This distribution is new and, although tested and functional, may still display odd behaviour.
	Beta bool `pulumi:"beta"`
	// This template bit format (32 or 64).
	BitFormat int `pulumi:"bitFormat"`
	// Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation).
	Category       string                                 `pulumi:"category"`
	Customizations []GetInstallationTemplateCustomization `pulumi:"customizations"`
	// The default language of this template.
	DefaultLanguage string `pulumi:"defaultLanguage"`
	// is this distribution deprecated.
	Deprecated bool `pulumi:"deprecated"`
	// information about this template.
	Description string `pulumi:"description"`
	// the distribution this template is based on.
	Distribution string `pulumi:"distribution"`
	// this template family type (bsd,linux,solaris,windows).
	Family string `pulumi:"family"`
	// Filesystems available (btrfs,ext3,ext4,ntfs,reiserfs,swap,ufs,xfs,zfs).
	Filesystems []string `pulumi:"filesystems"`
	// This distribution supports hardware raid configuration through the OVHcloud API.
	HardRaidConfiguration bool `pulumi:"hardRaidConfiguration"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Date of last modification of the base image.
	LastModification string                                       `pulumi:"lastModification"`
	LvmReady         bool                                         `pulumi:"lvmReady"`
	PartitionSchemes []GetInstallationTemplatePartitionSchemeType `pulumi:"partitionSchemes"`
	// This distribution supports installation using the distribution's native kernel instead of the recommended OVHcloud kernel.
	SupportsDistributionKernel bool `pulumi:"supportsDistributionKernel"`
	// This distribution supports RTM software.
	SupportsRtm bool `pulumi:"supportsRtm"`
	// This distribution supports the microsoft SQL server.
	SupportsSqlServer bool   `pulumi:"supportsSqlServer"`
	TemplateName      string `pulumi:"templateName"`
}

func LookupInstallationTemplateOutput(ctx *pulumi.Context, args LookupInstallationTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupInstallationTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstallationTemplateResult, error) {
			args := v.(LookupInstallationTemplateArgs)
			r, err := LookupInstallationTemplate(ctx, &args, opts...)
			var s LookupInstallationTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstallationTemplateResultOutput)
}

// A collection of arguments for invoking getInstallationTemplate.
type LookupInstallationTemplateOutputArgs struct {
	// This template name
	TemplateName pulumi.StringInput `pulumi:"templateName"`
}

func (LookupInstallationTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstallationTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getInstallationTemplate.
type LookupInstallationTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupInstallationTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstallationTemplateResult)(nil)).Elem()
}

func (o LookupInstallationTemplateResultOutput) ToLookupInstallationTemplateResultOutput() LookupInstallationTemplateResultOutput {
	return o
}

func (o LookupInstallationTemplateResultOutput) ToLookupInstallationTemplateResultOutputWithContext(ctx context.Context) LookupInstallationTemplateResultOutput {
	return o
}

// List of all language available for this template.
func (o LookupInstallationTemplateResultOutput) AvailableLanguages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) []string { return v.AvailableLanguages }).(pulumi.StringArrayOutput)
}

// This distribution is new and, although tested and functional, may still display odd behaviour.
func (o LookupInstallationTemplateResultOutput) Beta() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) bool { return v.Beta }).(pulumi.BoolOutput)
}

// This template bit format (32 or 64).
func (o LookupInstallationTemplateResultOutput) BitFormat() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) int { return v.BitFormat }).(pulumi.IntOutput)
}

// Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation).
func (o LookupInstallationTemplateResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o LookupInstallationTemplateResultOutput) Customizations() GetInstallationTemplateCustomizationArrayOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) []GetInstallationTemplateCustomization {
		return v.Customizations
	}).(GetInstallationTemplateCustomizationArrayOutput)
}

// The default language of this template.
func (o LookupInstallationTemplateResultOutput) DefaultLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) string { return v.DefaultLanguage }).(pulumi.StringOutput)
}

// is this distribution deprecated.
func (o LookupInstallationTemplateResultOutput) Deprecated() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) bool { return v.Deprecated }).(pulumi.BoolOutput)
}

// information about this template.
func (o LookupInstallationTemplateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) string { return v.Description }).(pulumi.StringOutput)
}

// the distribution this template is based on.
func (o LookupInstallationTemplateResultOutput) Distribution() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) string { return v.Distribution }).(pulumi.StringOutput)
}

// this template family type (bsd,linux,solaris,windows).
func (o LookupInstallationTemplateResultOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) string { return v.Family }).(pulumi.StringOutput)
}

// Filesystems available (btrfs,ext3,ext4,ntfs,reiserfs,swap,ufs,xfs,zfs).
func (o LookupInstallationTemplateResultOutput) Filesystems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) []string { return v.Filesystems }).(pulumi.StringArrayOutput)
}

// This distribution supports hardware raid configuration through the OVHcloud API.
func (o LookupInstallationTemplateResultOutput) HardRaidConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) bool { return v.HardRaidConfiguration }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstallationTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Date of last modification of the base image.
func (o LookupInstallationTemplateResultOutput) LastModification() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) string { return v.LastModification }).(pulumi.StringOutput)
}

func (o LookupInstallationTemplateResultOutput) LvmReady() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) bool { return v.LvmReady }).(pulumi.BoolOutput)
}

func (o LookupInstallationTemplateResultOutput) PartitionSchemes() GetInstallationTemplatePartitionSchemeTypeArrayOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) []GetInstallationTemplatePartitionSchemeType {
		return v.PartitionSchemes
	}).(GetInstallationTemplatePartitionSchemeTypeArrayOutput)
}

// This distribution supports installation using the distribution's native kernel instead of the recommended OVHcloud kernel.
func (o LookupInstallationTemplateResultOutput) SupportsDistributionKernel() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) bool { return v.SupportsDistributionKernel }).(pulumi.BoolOutput)
}

// This distribution supports RTM software.
func (o LookupInstallationTemplateResultOutput) SupportsRtm() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) bool { return v.SupportsRtm }).(pulumi.BoolOutput)
}

// This distribution supports the microsoft SQL server.
func (o LookupInstallationTemplateResultOutput) SupportsSqlServer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) bool { return v.SupportsSqlServer }).(pulumi.BoolOutput)
}

func (o LookupInstallationTemplateResultOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstallationTemplateResult) string { return v.TemplateName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstallationTemplateResultOutput{})
}
