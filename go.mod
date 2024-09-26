module github.com/aws-observability/aws-otel-collector

go 1.22.0

toolchain go1.22.2

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/confmap/provider/s3provider v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/awsproxy v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver v0.109.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver v0.109.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.109.0
	go.opentelemetry.io/collector/confmap v1.15.0
	go.opentelemetry.io/collector/confmap/provider/envprovider v1.15.0
	go.opentelemetry.io/collector/confmap/provider/fileprovider v1.15.0
	go.opentelemetry.io/collector/confmap/provider/httpprovider v0.109.0
	go.opentelemetry.io/collector/confmap/provider/httpsprovider v0.109.0
	go.opentelemetry.io/collector/confmap/provider/yamlprovider v0.109.0
	go.opentelemetry.io/collector/exporter v0.109.0
	go.opentelemetry.io/collector/exporter/debugexporter v0.109.0
	go.opentelemetry.io/collector/exporter/loggingexporter v0.109.0
	go.opentelemetry.io/collector/exporter/otlpexporter v0.109.0
	go.opentelemetry.io/collector/exporter/otlphttpexporter v0.109.0
	go.opentelemetry.io/collector/extension v0.109.0
	go.opentelemetry.io/collector/extension/ballastextension v0.108.1
	go.opentelemetry.io/collector/extension/zpagesextension v0.109.0
	go.opentelemetry.io/collector/featuregate v1.15.0
	go.opentelemetry.io/collector/otelcol v0.109.0
	go.opentelemetry.io/collector/processor v0.109.0
	go.opentelemetry.io/collector/processor/batchprocessor v0.109.0
	go.opentelemetry.io/collector/processor/memorylimiterprocessor v0.109.0
	go.opentelemetry.io/collector/receiver v0.109.0
	go.opentelemetry.io/collector/receiver/otlpreceiver v0.109.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/sys v0.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	cloud.google.com/go/auth v0.7.3 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.4 // indirect
	cloud.google.com/go/compute/metadata v0.5.2 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.13.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.10.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5 v5.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4 v4.3.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.2 // indirect
	github.com/Code-Hex/go-generics-cache v1.5.1 // indirect
	github.com/DataDog/agent-payload/v5 v5.0.132 // indirect
	github.com/DataDog/datadog-agent/comp/core/config v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/core/flare/builder v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/core/flare/types v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/core/hostname/hostnameinterface v0.56.2 // indirect
	github.com/DataDog/datadog-agent/comp/core/log v0.56.2 // indirect
	github.com/DataDog/datadog-agent/comp/core/secrets v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/core/telemetry v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/def v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/logs/agent/config v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/otelcol/logsagentpipeline v0.56.2 // indirect
	github.com/DataDog/datadog-agent/comp/otelcol/logsagentpipeline/logsagentpipelineimpl v0.56.2 // indirect
	github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/exporter/logsagentexporter v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/metricsclient v0.56.2 // indirect
	github.com/DataDog/datadog-agent/comp/trace/compression/def v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/comp/trace/compression/impl-gzip v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/collector/check/defaults v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/config/env v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/config/mock v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/config/model v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/config/setup v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/config/utils v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/auditor v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/client v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/diagnostic v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/message v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/metrics v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/pipeline v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/processor v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/sds v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/sender v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/sources v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/status/statusinterface v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/logs/status/utils v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/proto v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/status/health v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/telemetry v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/trace v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/backoff v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/util/cgroups v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/executable v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/filesystem v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/fxutil v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/hostname/validate v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/http v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/util/log v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/optional v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/pointer v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/scrubber v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/startstop v0.56.2 // indirect
	github.com/DataDog/datadog-agent/pkg/util/statstracker v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/system v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/system/socket v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/util/winutil v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-agent/pkg/version v0.58.0-rc.5 // indirect
	github.com/DataDog/datadog-api-client-go/v2 v2.29.0 // indirect
	github.com/DataDog/datadog-go/v5 v5.5.0 // indirect
	github.com/DataDog/dd-sensitive-data-scanner/sds-go/go v0.0.0-20240906125328-9899f57f27be // indirect
	github.com/DataDog/go-sqllexer v0.0.14 // indirect
	github.com/DataDog/go-tuf v1.1.0-0.5.2 // indirect
	github.com/DataDog/gohai v0.0.0-20230524154621-4316413895ee // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/inframetadata v0.20.0 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes v0.20.0 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/logs v0.20.0 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/metrics v0.19.0 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/quantile v0.19.0 // indirect
	github.com/DataDog/sketches-go v1.4.6 // indirect
	github.com/DataDog/viper v1.13.5 // indirect
	github.com/DataDog/zstd v1.5.6 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.24.2 // indirect
	github.com/IBM/sarama v1.43.3 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/Showmax/go-fqdn v1.0.0 // indirect
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/alecthomas/units v0.0.0-20240626203959-61d1e3462e30 // indirect
	github.com/apache/thrift v0.20.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/aws/aws-sdk-go v1.55.5 // indirect
	github.com/aws/aws-sdk-go-v2 v1.31.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.5 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.27.37 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.35 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.18 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.18 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.61.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.31.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.31.1 // indirect
	github.com/aws/smithy-go v1.21.0 // indirect
	github.com/benbjohnson/clock v1.3.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
	github.com/briandowns/spinner v1.23.1 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/checkpoint-restore/go-criu/v5 v5.3.0 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/cilium/ebpf v0.12.3 // indirect
	github.com/cncf/xds/go v0.0.0-20240905190251-b4127c9b8d78 // indirect
	github.com/containerd/cgroups/v3 v3.0.3 // indirect
	github.com/containerd/console v1.0.4 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/ttrpc v1.2.5 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cyphar/filepath-securejoin v0.2.5 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dennwc/varint v1.0.0 // indirect
	github.com/digitalocean/godo v1.118.0 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/docker v27.1.2+incompatible // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.7.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elastic/go-grok v0.3.1 // indirect
	github.com/elastic/lunes v0.1.0 // indirect
	github.com/emicklei/go-restful/v3 v3.12.1 // indirect
	github.com/envoyproxy/go-control-plane v0.12.1-0.20240621013728-1eb8caab5155 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/euank/go-kmsg-parser v2.0.0+incompatible // indirect
	github.com/expr-lang/expr v1.16.9 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-resty/resty/v2 v2.13.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.1.0 // indirect
	github.com/go-zookeeper/zk v1.0.4 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/cadvisor v0.49.1-0.20240628164550-89f779d86055 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.12.5 // indirect
	github.com/gophercloud/gophercloud v1.13.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/grafana/regexp v0.0.0-20240518133315-a468a5bfb3bc // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/hashicorp/consul/api v1.29.4 // indirect
	github.com/hashicorp/cronexpr v1.1.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/nomad/api v0.0.0-20240923191526-31e466921377 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/hectane/go-acl v0.0.0-20230122075934-ca0b05cb1adb // indirect
	github.com/hetznercloud/hcloud-go/v2 v2.10.2 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/ionos-cloud/sdk-go/v6 v6.1.11 // indirect
	github.com/jaegertracing/jaeger v1.60.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/klauspost/compress v1.17.10 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/kolo/xmlrpc v0.0.0-20220921171641-a4b6fa1dd06b // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/go-syslog/v4 v4.1.0 // indirect
	github.com/leodido/ragel-machinery v0.0.0-20190525184631-5f46317e436b // indirect
	github.com/lightstep/go-expohisto v1.0.0 // indirect
	github.com/linode/linodego v1.37.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20240909124753-873cd0166683 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/miekg/dns v1.1.62 // indirect
	github.com/mistifyio/go-zfs v2.1.2-0.20190413222219-f784269be439+incompatible // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20231216201459-8508981c8b6c // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/sys/mountinfo v0.7.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/mrunalp/fileutils v0.5.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/containerinsight v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/kafka v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azure v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx v0.109.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin v0.109.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/opencontainers/runc v1.1.14 // indirect
	github.com/opencontainers/runtime-spec v1.2.0 // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/openshift/api v0.0.0-20240919193929-2669d1ebc910 // indirect
	github.com/openshift/client-go v0.0.0-20240918182115-6a8ead8397fd // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	github.com/outcaste-io/ristretto v0.2.3 // indirect
	github.com/ovh/go-ovh v1.6.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/prometheus/client_golang v1.20.4 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.57.0 // indirect
	github.com/prometheus/common/sigv4 v0.1.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/prometheus/prometheus v0.54.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/relvacode/iso8601 v1.4.0 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.30 // indirect
	github.com/seccomp/libseccomp-golang v0.10.0 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.8.0 // indirect
	github.com/shirou/gopsutil/v3 v3.24.5 // indirect
	github.com/shirou/gopsutil/v4 v4.24.8 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/signalfx/com_signalfx_metrics_protobuf v0.0.3 // indirect
	github.com/signalfx/sapm-proto v0.14.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/stormcat24/protodep v0.1.8 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/tidwall/gjson v1.17.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/tinylru v1.2.1 // indirect
	github.com/tidwall/wal v1.1.7 // indirect
	github.com/tinylib/msgp v1.1.9 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/ua-parser/uap-go v0.0.0-20240611065828-3a4781585db6 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	github.com/vishvananda/netlink v1.2.1 // indirect
	github.com/vishvananda/netns v0.0.4 // indirect
	github.com/vultr/govultr/v2 v2.17.2 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.etcd.io/bbolt v1.3.11 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.109.0 // indirect
	go.opentelemetry.io/collector/client v1.15.0 // indirect
	go.opentelemetry.io/collector/component/componentprofiles v0.109.0 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.109.0 // indirect
	go.opentelemetry.io/collector/config/configauth v0.109.0 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.15.0 // indirect
	go.opentelemetry.io/collector/config/configgrpc v0.109.0 // indirect
	go.opentelemetry.io/collector/config/confighttp v0.109.0 // indirect
	go.opentelemetry.io/collector/config/confignet v0.109.0 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.15.0 // indirect
	go.opentelemetry.io/collector/config/configretry v1.15.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.109.0 // indirect
	go.opentelemetry.io/collector/config/configtls v1.15.0 // indirect
	go.opentelemetry.io/collector/config/internal v0.109.0 // indirect
	go.opentelemetry.io/collector/connector v0.109.0 // indirect
	go.opentelemetry.io/collector/connector/connectorprofiles v0.109.0 // indirect
	go.opentelemetry.io/collector/consumer v0.109.0 // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.109.0 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.109.0 // indirect
	go.opentelemetry.io/collector/exporter/exporterprofiles v0.109.0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.109.0 // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.109.0 // indirect
	go.opentelemetry.io/collector/extension/extensioncapabilities v0.109.0 // indirect
	go.opentelemetry.io/collector/internal/globalgates v0.109.0 // indirect
	go.opentelemetry.io/collector/pdata v1.15.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.109.0 // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.109.0 // indirect
	go.opentelemetry.io/collector/processor/processorprofiles v0.109.0 // indirect
	go.opentelemetry.io/collector/receiver/receiverprofiles v0.109.0 // indirect
	go.opentelemetry.io/collector/semconv v0.109.0 // indirect
	go.opentelemetry.io/collector/service v0.109.0 // indirect
	go.opentelemetry.io/contrib/config v0.9.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.29.0 // indirect
	go.opentelemetry.io/contrib/zpages v0.54.0 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.5.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.51.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.5.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.29.0 // indirect
	go.opentelemetry.io/otel/log v0.5.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/sdk v1.29.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.5.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/dig v1.18.0 // indirect
	go.uber.org/fx v1.22.2 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/exp v0.0.0-20240909161429-701f63a606c0 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/oauth2 v0.22.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/term v0.24.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	golang.org/x/tools v0.25.0 // indirect
	gonum.org/v1/gonum v0.15.1 // indirect
	google.golang.org/api v0.188.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/grpc v1.66.2 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gopkg.in/zorkian/go-datadog-api.v2 v2.30.0 // indirect
	k8s.io/api v0.31.1 // indirect
	k8s.io/apimachinery v0.31.1 // indirect
	k8s.io/client-go v0.31.1 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240903163716-9e1beecbcb38 // indirect
	k8s.io/utils v0.0.0-20240921022957-49e7df575cb6 // indirect
	sigs.k8s.io/controller-runtime v0.19.0 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/4433
exclude github.com/StackExchange/wmi v1.2.0

// https://github.com/google/flatbuffers/issues/6466
exclude github.com/google/flatbuffers v1.12.0

// see https://github.com/aws-observability/aws-otel-collector/issues/977
exclude github.com/docker/distribution v2.8.0+incompatible

replace github.com/go-chi/chi/v4 => github.com/go-chi/chi v4.0.0+incompatible

// see https://github.com/ionos-cloud/sdk-go/issues/27
exclude github.com/ionos-cloud/sdk-go/v6 v6.0.5851

// https://github.com/go-openapi/spec/issues/156
exclude github.com/go-openapi/spec v0.20.5

// excluded as DataDog/agent-payload/v5 v5.0.59 has been removed from source directory but is still present in proxy package
exclude github.com/DataDog/agent-payload/v5 v5.0.59 // indirect

// openshift removed all tags from their repo, use the pseudoversion from the release-3.9 branch HEAD
exclude github.com/openshift/api v3.9.0+incompatible

// checksun mismatch issue for cadvisor - https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/32381
exclude github.com/google/cadvisor v0.49.1
