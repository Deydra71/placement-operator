/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package placement

import (
	placementv1 "github.com/openstack-k8s-operators/placement-operator/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
)

// getVolumes - service volumes
func getVolumes(instance *placementv1.PlacementAPI) []corev1.Volume {
	var scriptsVolumeDefaultMode int32 = 0755
	var config0640AccessMode int32 = 0640

	volumes := []corev1.Volume{
		{
			Name: "scripts",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					DefaultMode: &scriptsVolumeDefaultMode,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: instance.Name + "-scripts",
					},
				},
			},
		},
		{
			Name: "config-data",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					DefaultMode: &config0640AccessMode,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: instance.Name + "-config-data",
					},
				},
			},
		},
		{
			Name: "config-data-merged",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{Medium: ""},
			},
		},
	}

	if instance.Spec.TLS != nil && instance.Spec.TLS.Service != nil {
		tls := instance.Spec.TLS
		for _, service := range tls.Service {
			serviceVolumes := service.CreateVolumes()
			volumes = append(volumes, serviceVolumes...)
		}
	}

	if instance.Spec.TLS != nil && instance.Spec.TLS.Ca != nil {
		ca := instance.Spec.TLS.Ca
		caVolumes := ca.CreateVolumes()
		volumes = append(volumes, caVolumes...)
	}

	return volumes
}

// getInitVolumeMounts - general init task VolumeMounts
func getInitVolumeMounts(instance *placementv1.PlacementAPI) []corev1.VolumeMount {
	volumeMounts := []corev1.VolumeMount{
		{
			Name:      "scripts",
			MountPath: "/usr/local/bin/container-scripts",
			ReadOnly:  true,
		},
		{
			Name:      "config-data",
			MountPath: "/var/lib/config-data/default",
			ReadOnly:  true,
		},
		{
			Name:      "config-data-merged",
			MountPath: "/var/lib/config-data/merged",
			ReadOnly:  false,
		},
	}

	// Question: Do we need the tls during the inialization phase?
	if instance.Spec.TLS != nil && instance.Spec.TLS.Service != nil {
		tls := instance.Spec.TLS
		for _, service := range tls.Service {
			serviceVolumeMounts := service.CreateVolumeMounts()
			volumeMounts = append(volumeMounts, serviceVolumeMounts...)
		}
	}

	if instance.Spec.TLS != nil && instance.Spec.TLS.Ca != nil {
		ca := instance.Spec.TLS.Ca
		caVolumeMounts := ca.CreateVolumeMounts()
		volumeMounts = append(volumeMounts, caVolumeMounts...)
	}

	return volumeMounts
}

// getVolumeMounts - general VolumeMounts
func getVolumeMounts(instance *placementv1.PlacementAPI) []corev1.VolumeMount {
	volumeMounts := []corev1.VolumeMount{
		{
			Name:      "scripts",
			MountPath: "/usr/local/bin/container-scripts",
			ReadOnly:  true,
		},
		{
			Name:      "config-data-merged",
			MountPath: "/var/lib/config-data/merged",
			ReadOnly:  false,
		},
		{
			Name:      "config-data-merged",
			MountPath: "/var/lib/kolla/config_files/config.json",
			SubPath:   "placement-api-config.json",
			ReadOnly:  true,
		},
	}

	if instance.Spec.TLS != nil && instance.Spec.TLS.Service != nil {
		tls := instance.Spec.TLS
		for _, service := range tls.Service {
			serviceVolumeMounts := service.CreateVolumeMounts()
			volumeMounts = append(volumeMounts, serviceVolumeMounts...)
		}
	}

	if instance.Spec.TLS != nil && instance.Spec.TLS.Ca != nil {
		ca := instance.Spec.TLS.Ca
		caVolumeMounts := ca.CreateVolumeMounts()
		volumeMounts = append(volumeMounts, caVolumeMounts...)
	}

	return volumeMounts
}
