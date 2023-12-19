// Code generated by "go.opentelemetry.io/collector/cmd/builder". DO NOT EDIT.

package main

import (
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/receiver"
	forwardconnector "go.opentelemetry.io/collector/connector/forwardconnector"
	countconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/countconnector"
	datadogconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/datadogconnector"
	exceptionsconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/exceptionsconnector"
	routingconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector"
	servicegraphconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector"
	spanmetricsconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector"
	debugexporter "go.opentelemetry.io/collector/exporter/debugexporter"
	loggingexporter "go.opentelemetry.io/collector/exporter/loggingexporter"
	otlpexporter "go.opentelemetry.io/collector/exporter/otlpexporter"
	otlphttpexporter "go.opentelemetry.io/collector/exporter/otlphttpexporter"
	alibabacloudlogserviceexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"
	awscloudwatchlogsexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter"
	awsemfexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"
	awskinesisexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter"
	awss3exporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"
	awsxrayexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"
	azuredataexplorerexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"
	azuremonitorexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"
	carbonexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter"
	clickhouseexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"
	cassandraexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/cassandraexporter"
	coralogixexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"
	datadogexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"
	datasetexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter"
	dynatraceexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter"
	elasticsearchexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"
	f5cloudexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/f5cloudexporter"
	fileexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"
	googlecloudexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter"
	googlecloudpubsubexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter"
	googlemanagedprometheusexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlemanagedprometheusexporter"
	honeycombmarkerexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/honeycombmarkerexporter"
	influxdbexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/influxdbexporter"
	instanaexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/instanaexporter"
	kafkaexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"
	loadbalancingexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"
	logicmonitorexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter"
	logzioexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter"
	lokiexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/lokiexporter"
	mezmoexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/mezmoexporter"
	opencensusexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opencensusexporter"
	opensearchexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter"
	prometheusexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"
	prometheusremotewriteexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"
	pulsarexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"
	sapmexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter"
	sentryexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter"
	signalfxexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"
	skywalkingexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/skywalkingexporter"
	splunkhecexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"
	sumologicexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter"
	syslogexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/syslogexporter"
	tanzuobservabilityexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tanzuobservabilityexporter"
	tencentcloudlogserviceexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tencentcloudlogserviceexporter"
	zipkinexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter"
	zpagesextension "go.opentelemetry.io/collector/extension/zpagesextension"
	ballastextension "go.opentelemetry.io/collector/extension/ballastextension"
	healthcheckextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	jaegerremotesampling "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling"
	ecsobserver "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"
	ecstaskobserver "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecstaskobserver"
	hostobserver "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver"
	k8sobserver "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"
	dockerobserver "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/dockerobserver"
	opampextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampextension"
	pprofextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	batchprocessor "go.opentelemetry.io/collector/processor/batchprocessor"
	memorylimiterprocessor "go.opentelemetry.io/collector/processor/memorylimiterprocessor"
	attributesprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	cumulativetodeltaprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor"
	datadogprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/datadogprocessor"
	deltatorateprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor"
	filterprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	groupbyattrsprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor"
	groupbytraceprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor"
	k8sattributesprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"
	metricsgenerationprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor"
	metricstransformprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"
	probabilisticsamplerprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"
	redactionprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor"
	resourcedetectionprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
	resourceprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	routingprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor"
	servicegraphprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/servicegraphprocessor"
	spanmetricsprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanmetricsprocessor"
	sumologicprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/sumologicprocessor"
	spanprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"
	tailsamplingprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"
	remotetapprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/remotetapprocessor"
	otlpreceiver "go.opentelemetry.io/collector/receiver/otlpreceiver"
	zipkinreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"
	filelogreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"
)

func components() (otelcol.Factories, error) {
	var err error
	factories := otelcol.Factories{}

	factories.Extensions, err = extension.MakeFactoryMap(
		zpagesextension.NewFactory(),
		ballastextension.NewFactory(),
		healthcheckextension.NewFactory(),
		jaegerremotesampling.NewFactory(),
		ecsobserver.NewFactory(),
		ecstaskobserver.NewFactory(),
		hostobserver.NewFactory(),
		k8sobserver.NewFactory(),
		dockerobserver.NewFactory(),
		opampextension.NewFactory(),
		pprofextension.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Receivers, err = receiver.MakeFactoryMap(
		otlpreceiver.NewFactory(),
		zipkinreceiver.NewFactory(),
		filelogreceiver.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Exporters, err = exporter.MakeFactoryMap(
		debugexporter.NewFactory(),
		loggingexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
		alibabacloudlogserviceexporter.NewFactory(),
		awscloudwatchlogsexporter.NewFactory(),
		awsemfexporter.NewFactory(),
		awskinesisexporter.NewFactory(),
		awss3exporter.NewFactory(),
		awsxrayexporter.NewFactory(),
		azuredataexplorerexporter.NewFactory(),
		azuremonitorexporter.NewFactory(),
		carbonexporter.NewFactory(),
		clickhouseexporter.NewFactory(),
		cassandraexporter.NewFactory(),
		coralogixexporter.NewFactory(),
		datadogexporter.NewFactory(),
		datasetexporter.NewFactory(),
		dynatraceexporter.NewFactory(),
		elasticsearchexporter.NewFactory(),
		f5cloudexporter.NewFactory(),
		fileexporter.NewFactory(),
		googlecloudexporter.NewFactory(),
		googlecloudpubsubexporter.NewFactory(),
		googlemanagedprometheusexporter.NewFactory(),
		honeycombmarkerexporter.NewFactory(),
		influxdbexporter.NewFactory(),
		instanaexporter.NewFactory(),
		kafkaexporter.NewFactory(),
		loadbalancingexporter.NewFactory(),
		logicmonitorexporter.NewFactory(),
		logzioexporter.NewFactory(),
		lokiexporter.NewFactory(),
		mezmoexporter.NewFactory(),
		opencensusexporter.NewFactory(),
		opensearchexporter.NewFactory(),
		prometheusexporter.NewFactory(),
		prometheusremotewriteexporter.NewFactory(),
		pulsarexporter.NewFactory(),
		sapmexporter.NewFactory(),
		sentryexporter.NewFactory(),
		signalfxexporter.NewFactory(),
		skywalkingexporter.NewFactory(),
		splunkhecexporter.NewFactory(),
		sumologicexporter.NewFactory(),
		syslogexporter.NewFactory(),
		tanzuobservabilityexporter.NewFactory(),
		tencentcloudlogserviceexporter.NewFactory(),
		zipkinexporter.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Processors, err = processor.MakeFactoryMap(
		batchprocessor.NewFactory(),
		memorylimiterprocessor.NewFactory(),
		attributesprocessor.NewFactory(),
		cumulativetodeltaprocessor.NewFactory(),
		datadogprocessor.NewFactory(),
		deltatorateprocessor.NewFactory(),
		filterprocessor.NewFactory(),
		groupbyattrsprocessor.NewFactory(),
		groupbytraceprocessor.NewFactory(),
		k8sattributesprocessor.NewFactory(),
		metricsgenerationprocessor.NewFactory(),
		metricstransformprocessor.NewFactory(),
		probabilisticsamplerprocessor.NewFactory(),
		redactionprocessor.NewFactory(),
		resourcedetectionprocessor.NewFactory(),
		resourceprocessor.NewFactory(),
		routingprocessor.NewFactory(),
		servicegraphprocessor.NewFactory(),
		spanmetricsprocessor.NewFactory(),
		sumologicprocessor.NewFactory(),
		spanprocessor.NewFactory(),
		tailsamplingprocessor.NewFactory(),
		remotetapprocessor.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Connectors, err = connector.MakeFactoryMap(
		forwardconnector.NewFactory(),
		countconnector.NewFactory(),
		datadogconnector.NewFactory(),
		exceptionsconnector.NewFactory(),
		routingconnector.NewFactory(),
		servicegraphconnector.NewFactory(),
		spanmetricsconnector.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	return factories, nil
}
