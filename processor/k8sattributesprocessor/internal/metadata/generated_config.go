// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
)

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for k8sattributes resource attributes.
type ResourceAttributesConfig struct {
	ContainerID               ResourceAttributeConfig `mapstructure:"container.id"`
	ContainerImageName        ResourceAttributeConfig `mapstructure:"container.image.name"`
	ContainerImageRepoDigests ResourceAttributeConfig `mapstructure:"container.image.repo_digests"`
	ContainerImageTag         ResourceAttributeConfig `mapstructure:"container.image.tag"`
	K8sClusterUID             ResourceAttributeConfig `mapstructure:"k8s.cluster.uid"`
	K8sContainerName          ResourceAttributeConfig `mapstructure:"k8s.container.name"`
	K8sCronjobName            ResourceAttributeConfig `mapstructure:"k8s.cronjob.name"`
	K8sDaemonsetName          ResourceAttributeConfig `mapstructure:"k8s.daemonset.name"`
	K8sDaemonsetUID           ResourceAttributeConfig `mapstructure:"k8s.daemonset.uid"`
	K8sDeploymentName         ResourceAttributeConfig `mapstructure:"k8s.deployment.name"`
	K8sDeploymentUID          ResourceAttributeConfig `mapstructure:"k8s.deployment.uid"`
	K8sJobName                ResourceAttributeConfig `mapstructure:"k8s.job.name"`
	K8sJobUID                 ResourceAttributeConfig `mapstructure:"k8s.job.uid"`
	K8sNamespaceName          ResourceAttributeConfig `mapstructure:"k8s.namespace.name"`
	K8sNodeName               ResourceAttributeConfig `mapstructure:"k8s.node.name"`
	K8sNodeUID                ResourceAttributeConfig `mapstructure:"k8s.node.uid"`
	K8sPodHostname            ResourceAttributeConfig `mapstructure:"k8s.pod.hostname"`
	K8sPodIP                  ResourceAttributeConfig `mapstructure:"k8s.pod.ip"`
	K8sPodName                ResourceAttributeConfig `mapstructure:"k8s.pod.name"`
	K8sPodStartTime           ResourceAttributeConfig `mapstructure:"k8s.pod.start_time"`
	K8sPodUID                 ResourceAttributeConfig `mapstructure:"k8s.pod.uid"`
	K8sReplicasetName         ResourceAttributeConfig `mapstructure:"k8s.replicaset.name"`
	K8sReplicasetUID          ResourceAttributeConfig `mapstructure:"k8s.replicaset.uid"`
	K8sStatefulsetName        ResourceAttributeConfig `mapstructure:"k8s.statefulset.name"`
	K8sStatefulsetUID         ResourceAttributeConfig `mapstructure:"k8s.statefulset.uid"`
	ServiceInstanceID         ResourceAttributeConfig `mapstructure:"service.instance.id"`
	ServiceName               ResourceAttributeConfig `mapstructure:"service.name"`
	ServiceNamespace          ResourceAttributeConfig `mapstructure:"service.namespace"`
	ServiceVersion            ResourceAttributeConfig `mapstructure:"service.version"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		ContainerID: ResourceAttributeConfig{
			Enabled: false,
		},
		ContainerImageName: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerImageRepoDigests: ResourceAttributeConfig{
			Enabled: false,
		},
		ContainerImageTag: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sClusterUID: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sContainerName: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sCronjobName: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sDaemonsetName: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sDaemonsetUID: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sDeploymentName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sDeploymentUID: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sJobName: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sJobUID: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sNamespaceName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sNodeName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sNodeUID: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sPodHostname: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sPodIP: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sPodName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sPodStartTime: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sPodUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sReplicasetName: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sReplicasetUID: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sStatefulsetName: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sStatefulsetUID: ResourceAttributeConfig{
			Enabled: false,
		},
		ServiceInstanceID: ResourceAttributeConfig{
			Enabled: false,
		},
		ServiceName: ResourceAttributeConfig{
			Enabled: false,
		},
		ServiceNamespace: ResourceAttributeConfig{
			Enabled: false,
		},
		ServiceVersion: ResourceAttributeConfig{
			Enabled: false,
		},
	}
}
