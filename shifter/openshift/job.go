/*
copyright 2019 google llc
licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at
    http://www.apache.org/licenses/license-2.0
unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package openshift

import (
	"context"
	"shifter/lib"

	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (c Openshift) GetAllJobs(namespace string) (*v1.JobList, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.JobList{}, err
	}

	jobList, err := cluster.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting Jobs from Namespace: "+namespace, err)
		return &v1.JobList{}, err
	}
	lib.CLog("debug", "Getting Jobs from Namespace: "+namespace)
	return jobList, err
}

func (c Openshift) GetJob(name string, namespace string) (*v1.Job, error) {
	cluster, err := kubernetes.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.Job{}, err
	}

	job, err := cluster.BatchV1().Jobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting Job with Name: "+name+" from Namespace: "+namespace, err)
		return &v1.Job{}, err
	}
	lib.CLog("info", "Getting Job with Name: "+name+" from Namespace: "+namespace)
	return job, err
}
