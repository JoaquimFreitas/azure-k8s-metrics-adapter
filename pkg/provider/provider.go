// Package provider is the implementation of custom metric and external metric apis
// see https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/custom-metrics-api.md#api-paths
package provider

import (
	"github.com/Azure/azure-k8s-metrics-adapter/pkg/azure/appinsights"
	"github.com/Azure/azure-k8s-metrics-adapter/pkg/azure/external_metrics"
	"github.com/Azure/azure-k8s-metrics-adapter/pkg/metriccache"
	"github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/provider"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/dynamic"
)

type AzureProvider struct {
	appinsightsClient     appinsights.AzureAppInsightsClient
	mapper                apimeta.RESTMapper
	kubeClient            dynamic.Interface
	metricCache           *metriccache.MetricCache
	azureClientFactory    azureexternalmetrics.AzureClientFactory
	defaultSubscriptionID string
}

func NewAzureProvider(defaultSubscriptionID string, mapper apimeta.RESTMapper, kubeClient dynamic.Interface, appinsightsClient appinsights.AzureAppInsightsClient, azureClientFactory azureexternalmetrics.AzureClientFactory, metricCache *metriccache.MetricCache) provider.MetricsProvider {
	return &AzureProvider{
		defaultSubscriptionID: defaultSubscriptionID,
		mapper:                mapper,
		kubeClient:            kubeClient,
		appinsightsClient:     appinsightsClient,
		metricCache:           metricCache,
		azureClientFactory:    azureClientFactory,
	}
}
