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
	"log"

	v1 "github.com/openshift/api/project/v1"
	projectv1 "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c Openshift) GetAllProjects() (*v1.ProjectList, error) {
	cluster, err := projectv1.NewForConfig(c.clusterClient())
	if err != nil {
		log.Println(err)
		return &v1.ProjectList{}, err
	}

	projectList, err := cluster.Projects().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return &v1.ProjectList{}, err
	}

	return projectList, nil
}

func (c Openshift) GetProject(name string) (*v1.Project, error) {
	cluster, err := projectv1.NewForConfig(c.clusterClient())
	if err != nil {
		log.Println(err)
		return &v1.Project{}, err
	}

	project, err := cluster.Projects().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return &v1.Project{}, err
	}

	return project, nil
}
