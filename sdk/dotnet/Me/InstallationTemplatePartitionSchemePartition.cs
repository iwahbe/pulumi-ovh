// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me
{
    /// <summary>
    /// Use this resource to create a partition in the partition scheme of a custom installation template available for dedicated servers.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mytemplate = new Ovh.Me.InstallationTemplate("mytemplate", new()
    ///     {
    ///         BaseTemplateName = "centos7_64",
    ///         TemplateName = "mytemplate",
    ///         DefaultLanguage = "fr",
    ///     });
    /// 
    ///     var scheme = new Ovh.Me.InstallationTemplatePartitionScheme("scheme", new()
    ///     {
    ///         TemplateName = mytemplate.TemplateName,
    ///         Priority = 1,
    ///     });
    /// 
    ///     var root = new Ovh.Me.InstallationTemplatePartitionSchemePartition("root", new()
    ///     {
    ///         TemplateName = scheme.TemplateName,
    ///         SchemeName = scheme.Name,
    ///         Mountpoint = "/",
    ///         Filesystem = "ext4",
    ///         Size = 400,
    ///         Order = 1,
    ///         Type = "primary",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The resource can be imported using the `template_name`, `scheme_name`, `mountpoint` of the cluster, separated by "/" E.g.,
    /// 
    ///  bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:Me/installationTemplatePartitionSchemePartition:InstallationTemplatePartitionSchemePartition root template_name/scheme_name/mountpoint
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Me/installationTemplatePartitionSchemePartition:InstallationTemplatePartitionSchemePartition")]
    public partial class InstallationTemplatePartitionSchemePartition : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Partition filesystem. Enum with possibles values:
        /// - btrfs
        /// - ext3
        /// - ext4
        /// - ntfs
        /// - reiserfs
        /// - swap
        /// - ufs
        /// - xfs
        /// - zfs
        /// </summary>
        [Output("filesystem")]
        public Output<string> Filesystem { get; private set; } = null!;

        /// <summary>
        /// partition mount point.
        /// </summary>
        [Output("mountpoint")]
        public Output<string> Mountpoint { get; private set; } = null!;

        /// <summary>
        /// step or order. specifies the creation order of the partition on the disk
        /// </summary>
        [Output("order")]
        public Output<int> Order { get; private set; } = null!;

        /// <summary>
        /// raid partition type. Enum with possible values: 
        /// - raid0
        /// - raid1
        /// - raid10
        /// - raid5
        /// - raid6
        /// </summary>
        [Output("raid")]
        public Output<string> Raid { get; private set; } = null!;

        /// <summary>
        /// The partition scheme name.
        /// </summary>
        [Output("schemeName")]
        public Output<string> SchemeName { get; private set; } = null!;

        /// <summary>
        /// size of partition in MB, 0 =&gt; rest of the space.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The template name of the partition scheme.
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;

        /// <summary>
        /// partition type. Enum with possible values:
        /// - lv
        /// - primary
        /// - logical
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The volume name needed for proxmox distribution
        /// </summary>
        [Output("volumeName")]
        public Output<string> VolumeName { get; private set; } = null!;


        /// <summary>
        /// Create a InstallationTemplatePartitionSchemePartition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstallationTemplatePartitionSchemePartition(string name, InstallationTemplatePartitionSchemePartitionArgs args, CustomResourceOptions? options = null)
            : base("ovh:Me/installationTemplatePartitionSchemePartition:InstallationTemplatePartitionSchemePartition", name, args ?? new InstallationTemplatePartitionSchemePartitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstallationTemplatePartitionSchemePartition(string name, Input<string> id, InstallationTemplatePartitionSchemePartitionState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Me/installationTemplatePartitionSchemePartition:InstallationTemplatePartitionSchemePartition", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InstallationTemplatePartitionSchemePartition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstallationTemplatePartitionSchemePartition Get(string name, Input<string> id, InstallationTemplatePartitionSchemePartitionState? state = null, CustomResourceOptions? options = null)
        {
            return new InstallationTemplatePartitionSchemePartition(name, id, state, options);
        }
    }

    public sealed class InstallationTemplatePartitionSchemePartitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Partition filesystem. Enum with possibles values:
        /// - btrfs
        /// - ext3
        /// - ext4
        /// - ntfs
        /// - reiserfs
        /// - swap
        /// - ufs
        /// - xfs
        /// - zfs
        /// </summary>
        [Input("filesystem", required: true)]
        public Input<string> Filesystem { get; set; } = null!;

        /// <summary>
        /// partition mount point.
        /// </summary>
        [Input("mountpoint", required: true)]
        public Input<string> Mountpoint { get; set; } = null!;

        /// <summary>
        /// step or order. specifies the creation order of the partition on the disk
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        /// <summary>
        /// raid partition type. Enum with possible values: 
        /// - raid0
        /// - raid1
        /// - raid10
        /// - raid5
        /// - raid6
        /// </summary>
        [Input("raid")]
        public Input<string>? Raid { get; set; }

        /// <summary>
        /// The partition scheme name.
        /// </summary>
        [Input("schemeName", required: true)]
        public Input<string> SchemeName { get; set; } = null!;

        /// <summary>
        /// size of partition in MB, 0 =&gt; rest of the space.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// The template name of the partition scheme.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        /// <summary>
        /// partition type. Enum with possible values:
        /// - lv
        /// - primary
        /// - logical
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The volume name needed for proxmox distribution
        /// </summary>
        [Input("volumeName")]
        public Input<string>? VolumeName { get; set; }

        public InstallationTemplatePartitionSchemePartitionArgs()
        {
        }
        public static new InstallationTemplatePartitionSchemePartitionArgs Empty => new InstallationTemplatePartitionSchemePartitionArgs();
    }

    public sealed class InstallationTemplatePartitionSchemePartitionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Partition filesystem. Enum with possibles values:
        /// - btrfs
        /// - ext3
        /// - ext4
        /// - ntfs
        /// - reiserfs
        /// - swap
        /// - ufs
        /// - xfs
        /// - zfs
        /// </summary>
        [Input("filesystem")]
        public Input<string>? Filesystem { get; set; }

        /// <summary>
        /// partition mount point.
        /// </summary>
        [Input("mountpoint")]
        public Input<string>? Mountpoint { get; set; }

        /// <summary>
        /// step or order. specifies the creation order of the partition on the disk
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// raid partition type. Enum with possible values: 
        /// - raid0
        /// - raid1
        /// - raid10
        /// - raid5
        /// - raid6
        /// </summary>
        [Input("raid")]
        public Input<string>? Raid { get; set; }

        /// <summary>
        /// The partition scheme name.
        /// </summary>
        [Input("schemeName")]
        public Input<string>? SchemeName { get; set; }

        /// <summary>
        /// size of partition in MB, 0 =&gt; rest of the space.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The template name of the partition scheme.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        /// <summary>
        /// partition type. Enum with possible values:
        /// - lv
        /// - primary
        /// - logical
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The volume name needed for proxmox distribution
        /// </summary>
        [Input("volumeName")]
        public Input<string>? VolumeName { get; set; }

        public InstallationTemplatePartitionSchemePartitionState()
        {
        }
        public static new InstallationTemplatePartitionSchemePartitionState Empty => new InstallationTemplatePartitionSchemePartitionState();
    }
}
