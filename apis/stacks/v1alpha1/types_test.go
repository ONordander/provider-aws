/*
Copyright 2019 The Crossplane Authors.

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

package v1alpha1

import (
	"testing"

	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplaneio/crossplane-runtime/pkg/test"
	localtest "github.com/crossplaneio/crossplane/pkg/test"
)

const (
	namespace = "default"
	name      = "test-instance"
)

var (
	ctx = context.TODO()
	c   client.Client

	key = types.NamespacedName{Name: name, Namespace: namespace}
)

func TestMain(m *testing.M) {
	t := test.NewEnv(namespace, SchemeBuilder.SchemeBuilder, localtest.CRDs())
	c = t.StartClient()
	t.StopAndExit(m.Run())
}

func TestStackRequest(t *testing.T) {
	g := NewGomegaWithT(t)

	created := &StackRequest{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: StackRequestSpec{
			Source:  "registry.crossplane.io",
			Package: "testpackage:v0.1",
		},
	}

	// Test Create
	fetched := &StackRequest{}
	g.Expect(c.Create(ctx, created)).NotTo(HaveOccurred())

	g.Expect(c.Get(ctx, key, fetched)).NotTo(HaveOccurred())
	g.Expect(fetched).To(Equal(created))

	// Test Updating the annotations
	updated := fetched.DeepCopy()
	updated.Annotations = map[string]string{"hello": "world"}
	g.Expect(c.Update(ctx, updated)).NotTo(HaveOccurred())

	g.Expect(c.Get(ctx, key, fetched)).NotTo(HaveOccurred())
	g.Expect(fetched).To(Equal(updated))

	// Test Delete
	g.Expect(c.Delete(ctx, fetched)).NotTo(HaveOccurred())
	g.Expect(c.Get(ctx, key, fetched)).To(HaveOccurred())
}

func TestClusterStackRequest(t *testing.T) {
	g := NewGomegaWithT(t)

	created := &ClusterStackRequest{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: ClusterStackRequestSpec{
			Source:  "registry.crossplane.io",
			Package: "testpackage:v0.1",
		},
	}

	// Test Create
	fetched := &ClusterStackRequest{}
	g.Expect(c.Create(ctx, created)).NotTo(HaveOccurred())

	g.Expect(c.Get(ctx, key, fetched)).NotTo(HaveOccurred())
	g.Expect(fetched).To(Equal(created))

	// Test Updating the annotations
	updated := fetched.DeepCopy()
	updated.Annotations = map[string]string{"hello": "world"}
	g.Expect(c.Update(ctx, updated)).NotTo(HaveOccurred())

	g.Expect(c.Get(ctx, key, fetched)).NotTo(HaveOccurred())
	g.Expect(fetched).To(Equal(updated))

	// Test Delete
	g.Expect(c.Delete(ctx, fetched)).NotTo(HaveOccurred())
	g.Expect(c.Get(ctx, key, fetched)).To(HaveOccurred())
}

func TestStack(t *testing.T) {
	g := NewGomegaWithT(t)

	created := &Stack{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: StackSpec{
			AppMetadataSpec: AppMetadataSpec{
				Title:   "myapp",
				Version: "v0.1.0",
				Owners: []ContributorSpec{
					{Name: "dev1", Email: "dev1@foo.com"},
				},
				Company:  "foo-inc.",
				Keywords: []string{"app", "useless", "example"},
				Website:  "https://app.foo.io",
			},
		},
	}

	// Test Create
	fetched := &Stack{}
	g.Expect(c.Create(ctx, created)).NotTo(HaveOccurred())

	g.Expect(c.Get(ctx, key, fetched)).NotTo(HaveOccurred())
	g.Expect(fetched).To(Equal(created))

	// Test Updating the annotations
	updated := fetched.DeepCopy()
	updated.Annotations = map[string]string{"hello": "world"}
	g.Expect(c.Update(ctx, updated)).NotTo(HaveOccurred())

	g.Expect(c.Get(ctx, key, fetched)).NotTo(HaveOccurred())
	g.Expect(fetched).To(Equal(updated))

	// Test Delete
	g.Expect(c.Delete(ctx, fetched)).NotTo(HaveOccurred())
	g.Expect(c.Get(ctx, key, fetched)).To(HaveOccurred())
}

func TestNewCRDList(t *testing.T) {
	g := NewGomegaWithT(t)
	crdList := NewCRDList()
	g.Expect(crdList).NotTo(BeNil())
	g.Expect(crdList.Owned).NotTo(BeNil())
	g.Expect(crdList.DependsOn).NotTo(BeNil())
}
