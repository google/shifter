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

	v1 "github.com/openshift/api/template/v1"
	templatev1 "github.com/openshift/client-go/template/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c Openshift) GetAllTemplates(namespace string) (*v1.TemplateList, error) {
	cluster, err := templatev1.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.TemplateList{}, err
	}

	templates, err := cluster.TemplateV1().Templates(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		lib.CLog("error", "Getting Templates from Namespace: "+namespace, err)
		return &v1.TemplateList{}, err
	}
	lib.CLog("debug", "Getting Templates from Namespace: "+namespace)
	return templates, err
}

func (c Openshift) GetTemplate(name string, namespace string) (*v1.Template, error) {
	cluster, err := templatev1.NewForConfig(c.clusterClient())
	if err != nil {
		lib.CLog("error", "Unable to connect to cluster", err)
		return &v1.Template{}, err
	}

	template, err := cluster.TemplateV1().Templates(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		lib.CLog("error", "Getting Template with Name: "+name+" from Namespace: "+namespace, err)
		return &v1.Template{}, err
	}
	lib.CLog("info", "Getting Template with Name: "+name+" from Namespace: "+namespace)
	return template, err
}
