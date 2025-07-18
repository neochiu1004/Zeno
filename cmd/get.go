package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func getCMDs() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Archive the web!",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
			}
		},
	}

	getCMDsFlags(getCmd)
	getHQCmdFlags(getHQCmd)

	getCmd.AddCommand(getURLCmd)
	getCmd.AddCommand(getHQCmd)

	return getCmd
}

func getCMDsFlags(getCmd *cobra.Command) {
	getCmd.PersistentFlags().String("user-agent", "", "User agent to use when requesting URLs.")
	getCmd.PersistentFlags().String("job", "", "Job name to use, will determine the path for the persistent queue, seencheck database, and WARC files.")
	getCmd.PersistentFlags().IntP("workers", "w", 1, "Number of concurrent workers to run.")
	getCmd.PersistentFlags().Int("max-concurrent-assets", 1, "Max number of concurrent assets to fetch PER worker. E.g. if you have 100 workers and this setting at 8, Zeno could do up to 800 concurrent requests at any time.")
	getCmd.PersistentFlags().Int("max-hops", 0, "Maximum number of hops to execute.")
	getCmd.PersistentFlags().String("cookies", "", "File containing cookies that will be used for requests.")
	getCmd.PersistentFlags().Bool("disable-seencheck", false, "Disable the (remote or local) seencheck that avoid re-crawling of URIs.")
	getCmd.PersistentFlags().Bool("api", false, "Enable API")
	getCmd.PersistentFlags().Int("api-port", 9090, "Port to listen on for the API.")
	getCmd.PersistentFlags().Int("max-redirect", 20, "Specifies the maximum number of redirections to follow for a resource.")
	getCmd.PersistentFlags().Int("max-css-jump", 10, "Specifies the maximum number of CSS @import jumps to follow for a resource.")
	getCmd.PersistentFlags().Int("max-retry", 5, "Number of retry if error happen when executing HTTP request.")
	getCmd.PersistentFlags().Int("http-timeout", -1, "Number of seconds to wait before timing out a request. Note: this will CANCEL large files download.")
	getCmd.PersistentFlags().Int("http-read-deadline", 60, "Number of seconds to wait before timing out a (blocking) read.")
	getCmd.PersistentFlags().StringSlice("domains-crawl", []string{}, "Naive domains, full URLs or regexp to match against any URL to determine hop behaviour for outlinks. If an outlink URL is matched it will be queued to crawl with a hop of 0. This flag helps crawling entire domains while doing non-focused crawls.")
	getCmd.PersistentFlags().StringSlice("disable-html-tag", []string{}, "Specify HTML tag to not extract assets from")
	getCmd.PersistentFlags().Bool("capture-alternate-pages", false, "If turned on, <link> HTML tags with \"alternate\" values for their \"rel\" attribute will be archived.")
	getCmd.PersistentFlags().StringSlice("exclude-host", []string{}, "Exclude a specific host from the crawl, note that it will not exclude the domain if it is encountered as an asset for another web page.")
	getCmd.PersistentFlags().StringSlice("include-host", []string{}, "Only crawl specific hosts, note that it will not include the domain if it is encountered as an asset for another web page.")
	getCmd.PersistentFlags().StringSlice("include-string", []string{}, "Only crawl URLs containing this string.")
	getCmd.PersistentFlags().Int("crawl-time-limit", 0, "Number of seconds until the crawl will automatically set itself into the finished state.")
	getCmd.PersistentFlags().Int("crawl-max-time-limit", 0, "Number of seconds until the crawl will automatically panic itself. Default to crawl-time-limit + (crawl-time-limit / 10)")
	getCmd.PersistentFlags().StringSlice("exclude-string", []string{}, "Discard any (discovered) URLs containing this string.")
	getCmd.PersistentFlags().StringSlice("exclusion-file", []string{}, "File containing regex to apply on URLs for exclusion. If the path start with http or https, it will be treated as a URL of a file to download.")
	getCmd.PersistentFlags().Float64("min-space-required", 0, "Minimum space required in GB to continue the crawl. Default will be 50GB * (total disk space / 256GB) if total disk space is less than 256GB, else 50GB.")
	getCmd.PersistentFlags().Bool("strict-regex", false, "If turned on, the xurls `strict` regex setting will be used. Otherwise a looser regex will be used.")

	// Network flags
	getCmd.PersistentFlags().String("proxy", "", "Proxy to use when requesting pages.")
	getCmd.PersistentFlags().Bool("random-local-ip", false, "Use random local IP for requests. (will be ignored if a proxy is set)")
	getCmd.PersistentFlags().Bool("disable-ipv4", false, "Disable IPv4 for requests.")
	getCmd.PersistentFlags().Bool("disable-ipv6", false, "Disable IPv6 for requests.")
	getCmd.PersistentFlags().Bool("ipv6-anyip", false, "Use AnyIP kernel feature for requests. (only IPv6, need --random-local-ip)")

	// Rate limiting flags
	getCmd.PersistentFlags().Bool("disable-rate-limit", false, "Disable the Token Bucket rate limiting.")
	getCmd.PersistentFlags().Float64("rate-limit-capacity", 150, "Bucket capacity for each host.")
	getCmd.PersistentFlags().Float64("rate-limit-refill-rate", 50, "Ideal requests per second for each host.")
	getCmd.PersistentFlags().Duration("rate-limit-cleanup-frequency", time.Duration(5*time.Minute), "How often to run cleanup of stale buckets that are not accessed in the duration.")

	// WARC flags
	getCmd.PersistentFlags().String("warc-prefix", "ZENO", "Prefix to use when naming the WARC files.")
	getCmd.PersistentFlags().String("warc-operator", "", "Contact informations of the crawl operator to write in the Warc-Info record in each WARC file.")
	getCmd.PersistentFlags().String("warc-cdx-dedupe-server", "", "Identify the server to use CDX deduplication. This also activates CDX deduplication on.")
	getCmd.PersistentFlags().Bool("warc-on-disk", false, "Do not use RAM to store payloads when recording traffic to WARCs, everything will happen on disk (usually used to reduce memory usage).")
	getCmd.PersistentFlags().Int("warc-pool-size", 1, "Number of concurrent WARC files to write.")
	getCmd.PersistentFlags().Int("warc-queue-size", -1, "Number of WARC records to queue before blocking the workers. Default is the --warc-pool-size.")
	getCmd.PersistentFlags().String("warc-temp-dir", "", "Custom directory to use for WARC temporary files.")
	getCmd.PersistentFlags().Bool("disable-local-dedupe", false, "Disable local URL agnostic deduplication.")
	getCmd.PersistentFlags().Bool("cert-validation", false, "Enables certificate validation on HTTPS requests.")
	getCmd.PersistentFlags().Bool("disable-assets-capture", false, "Disable assets capture.")
	getCmd.PersistentFlags().Int("warc-dedupe-size", 1024, "Minimum size to deduplicate WARC records with revisit records.")
	getCmd.PersistentFlags().String("warc-cdx-cookie", "", "Pass custom cookie during CDX requests. Example: 'cdx_auth_token=test_value'")
	getCmd.PersistentFlags().Int("warc-size", 1024, "Size of the WARC files in MB.")
	getCmd.PersistentFlags().IntSlice("warc-discard-status", []int{429}, "HTTP status codes to discard from WARC files. By default, 429 is always discarded.")
	getCmd.PersistentFlags().Bool("async-warc-write", false, "Write WARC records asynchronously. EXPERIMENTAL - may cause OOMs, lost data, or other unknown/unpredicted issues. No support will be provided for this feature.")

	// Logging flags
	getCmd.PersistentFlags().Bool("tui", false, "Display a terminal user interface.")
	getCmd.PersistentFlags().String("tui-log-level", "info", "Log level for the TUI.")
	getCmd.PersistentFlags().Bool("no-log-file", false, "Disable log file output.")
	getCmd.PersistentFlags().String("log-file-output-dir", "", "Directory to write log files to.")
	getCmd.PersistentFlags().String("log-file-prefix", "ZENO", "Prefix to use when naming the log files. Default is : `ZENO`, without '-'")
	getCmd.PersistentFlags().String("log-file-level", "info", "Log level for the log file.")
	getCmd.PersistentFlags().String("log-file-rotation", "1h", "Log file rotation period. Default is : `1h`. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.")

	// Profiling flags
	getCmd.PersistentFlags().String("pyroscope-address", "", "Pyroscope server address. Setting this flag will enable profiling.")

	// Prometheus and metrics flags
	getCmd.PersistentFlags().Bool("prometheus", false, "Export metrics in Prometheus format. (implies --api)")
	getCmd.PersistentFlags().String("prometheus-prefix", "zeno_", "String used as a prefix for the exported Prometheus metrics.")

	// Consul flags
	getCmd.PersistentFlags().String("consul-address", "", "Consul address to use for service registration.")
	getCmd.PersistentFlags().String("consul-port", "8500", "Consul port to use for service registration.")
	getCmd.PersistentFlags().String("consul-acl-token", "", "Consul ACL token to use for service registration.")
	getCmd.PersistentFlags().Bool("consul-register", false, "Register Zeno in Consul via the API. (useful when Zeno is running on host and not containerized)")
	getCmd.PersistentFlags().StringSlice("consul-register-tags", []string{}, "Tags to use when registering Zeno in Consul with `--consul-register`.")

	// Alias support
	// As cobra doesn't support aliases natively (couldn't find a way to do it), we have to do it manually
	// This is a workaround to allow users to use `--hops` instead of `--max-hops` for example
	// Aliases shouldn't be used as proper flags nor declared in the config struct
	// Aliases should be marked as deprecated to inform the user base
	// Aliases values should be copied to the proper flag in the config/config.go:handleFlagsAliases() function
	getCmd.PersistentFlags().Int("hops", 0, "Maximum number of hops to execute.")
	getCmd.PersistentFlags().MarkDeprecated("hops", "use --max-hops instead")
	getCmd.PersistentFlags().MarkHidden("hops")

	getCmd.PersistentFlags().Uint("ca", 8, "Max number of concurrent assets to fetch PER worker. E.g. if you have 100 workers and this setting at 8, Zeno could do up to 800 concurrent requests at any time.")
	getCmd.PersistentFlags().MarkDeprecated("ca", "use --max-concurrent-assets")
	getCmd.PersistentFlags().MarkHidden("ca")
}
