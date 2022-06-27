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

	v1 "github.com/openshift/api/route/v1"
	routev1 "github.com/openshift/client-go/route/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c Openshift) GetAllRoutes(namespace string) (*v1.RouteList, error) {
	cluster, err := routev1.NewForConfig(c.clusterClient())
	if err != nil {
		log.Println(err)
		return &v1.RouteList{}, err
	}

	routeList, err := cluster.RouteV1().Routes(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return &v1.RouteList{}, err
	}

	return routeList, nil

}

func (c Openshift) GetRoute(name string, namespace string) (*v1.Route, error) {
	cluster, err := routev1.NewForConfig(c.clusterClient())
	if err != nil {
		log.Println(err)
		return &v1.Route{}, err
	}

	route, err := cluster.RouteV1().Routes(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return &v1.Route{}, err
	}

	return route, nil
}
