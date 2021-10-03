package clidocstool

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var (
	cmd    *cobra.Command
	subCmd *cobra.Command
	flags  *flag.FlagSet
)

//nolint:errcheck
func init() {
	cmd = &cobra.Command{Use: "do [OPTIONS] arg1 arg2"}
	subCmd = &cobra.Command{Use: "sub [OPTIONS] arg1 arg2", Run: func(cmd *cobra.Command, args []string) {}}

	flags = subCmd.Flags()
	flags.Bool("push", false, "Shorthand for --output=type=registry")
	flags.Bool("load", false, "Shorthand for --output=type=docker")
	flags.StringArrayP("tag", "t", []string{}, "Name and optionally a tag in the 'name:tag' format")
	flags.SetAnnotation("tag", "docs.external.url", []string{"https://docs.docker.com/engine/reference/commandline/build/#tag-an-image--t"})
	flags.StringArray("build-arg", []string{}, "Set build-time variables")
	flags.SetAnnotation("build-arg", "docs.external.url", []string{"https://docs.docker.com/engine/reference/commandline/build/#set-build-time-variables---build-arg"})
	flags.StringP("file", "f", "", "Name of the Dockerfile (Default is 'PATH/Dockerfile')")
	flags.SetAnnotation("file", "docs.external.url", []string{"https://docs.docker.com/engine/reference/commandline/build/#specify-a-dockerfile--f"})
	flags.StringArray("label", []string{}, "Set metadata for an image")
	flags.StringArray("cache-from", []string{}, "External cache sources (eg. user/app:cache, type=local,src=path/to/dir)")
	flags.StringArray("cache-to", []string{}, "Cache export destinations (eg. user/app:cache, type=local,dest=path/to/dir)")
	flags.String("target", "", "Set the target build stage to build.")
	flags.SetAnnotation("target", "docs.external.url", []string{"https://docs.docker.com/engine/reference/commandline/build/#specifying-target-build-stage---target"})
	flags.StringSlice("allow", []string{}, "Allow extra privileged entitlement, e.g. network.host, security.insecure")
	flags.StringArray("platform", []string{}, "Set target platform for build")
	flags.StringArray("secret", []string{}, "Secret file to expose to the build: id=mysecret,src=/local/secret")
	flags.StringArray("ssh", []string{}, "SSH agent socket or keys to expose to the build (format: `default|<id>[=<socket>|<key>[,<key>]]`)")
	flags.StringArrayP("output", "o", []string{}, "Output destination (format: type=local,dest=path)")
	// not implemented
	flags.String("network", "default", "Set the networking mode for the RUN instructions during build")
	flags.StringSlice("add-host", []string{}, "Add a custom host-to-IP mapping (host:ip)")
	flags.SetAnnotation("add-host", "docs.external.url", []string{"https://docs.docker.com/engine/reference/commandline/build/#add-entries-to-container-hosts-file---add-host"})
	flags.String("iidfile", "", "Write the image ID to the file")
	// hidden flags
	flags.BoolP("quiet", "q", false, "Suppress the build output and print image ID on success")
	flags.MarkHidden("quiet")
	flags.Bool("squash", false, "Squash newly built layers into a single new layer")
	flags.MarkHidden("squash")
	flags.String("ulimit", "", "Ulimit options")
	flags.MarkHidden("ulimit")
	flags.StringSlice("security-opt", []string{}, "Security options")
	flags.MarkHidden("security-opt")
	flags.Bool("compress", false, "Compress the build context using gzip")

	cmd.AddCommand(subCmd)
}
