package kubectl

import (
	"reflect"
	"testing"

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

func  TestStructuredGenerate(t *testing.T) {
	tests := []struct {
		Name      string
		Images    []string
		expected  *appsv1beta1.Deployment
		expectErr bool
	}{
		{
			Name:   "dep",
			Images: []string{},
			expected: &appsv1beta1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "dep",
					Labels: map[string]string{},
				},
				Spec: appsv1beta1.DeploymentSpec{
					Replicas: newInt32(1),
					Selector: &metav1.LabelSelector{},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{},
						},
						Spec: v1.PodSpec{},
					},
				},
			},
			expectErr: true,
		},
		{
			Name:   "",
			Images: []string{"image1"},
			expected: &appsv1beta1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "",
					Labels: map[string]string{},
				},
				Spec: appsv1beta1.DeploymentSpec{
					Replicas: newInt32(1),
					Selector: &metav1.LabelSelector{},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{},
						},
						Spec: v1.PodSpec{},
					},
				},
			},
			expectErr: true,
		},
		{
			Name:   "dep1",
			Images: []string{"image1"},
			expected: &appsv1beta1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "image1",
					Labels: map[string]string{"app": "image1"},
				},
				Spec: appsv1beta1.DeploymentSpec{
					Replicas: newInt32(1),
					Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "image1"}},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{"app": "image1contain"},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{{
								Name:  "image1",
								Image: "image1",
							},},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			Name:   "dep1",
			Images: []string{"image1contain/image2"},
			expected: &appsv1beta1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "image1contain",
					Labels: map[string]string{"app": "image1contain"},
				},
				Spec: appsv1beta1.DeploymentSpec{
					Replicas: newInt32(1),
					Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "image1contain"}},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{"app": "image1contain"},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{{
								Name:  "image1contain",
								Image: "image1contain",
							},},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			Name:   "dep1:dep2",
			Images: []string{"image1contain/image2"},
			expected: &appsv1beta1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "dep1",
					Labels: map[string]string{"app": "dep1"},
				},
				Spec: appsv1beta1.DeploymentSpec{
					Replicas: newInt32(1),
					Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "dep1"}},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{"app": "dep1"},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{{
								Name:  "dep1",
								Image: "image1contain/image2",
							},},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			Name:   "dep1@dep2",
			Images: []string{"image1contain/image2"},
			expected: &appsv1beta1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "dep1",
					Labels: map[string]string{"app": "dep1"},
				},
				Spec: appsv1beta1.DeploymentSpec{
					Replicas: newInt32(1),
					Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "dep1"}},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{"app": "dep1"},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{{
								Name:  "dep1",
								Image: "image1contain/image2",
							},},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			Name:   "dep1@dep2:dep3",
			Images: []string{"image1contain/image2"},
			expected: &appsv1beta1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "dep1",
					Labels: map[string]string{"app": "dep1"},
				},
				Spec: appsv1beta1.DeploymentSpec{
					Replicas: newInt32(1),
					Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "dep1"}},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{"app": "dep1"},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{{
								Name:  "dep1",
								Image: "image1contain/image2",
							},},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			Name:   "dep1:dep2",
			Images: []string{"image1", "image2"},
			expected: &appsv1beta1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "dep1",
					Labels: map[string]string{"app": "dep1"},
				},
				Spec: appsv1beta1.DeploymentSpec{
					Replicas: newInt32(1),
					Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "dep1"}},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{"app": "dep1"},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{{
								Name:  "dep1",
								Image: "image1",
							}, {
								Name:  "dep1",
								Image: "image2",
							}},
						},
					},
				},
			},
			expectErr: false,
		},
	}

	generator := &DeploymentBasicAppsGeneratorV1{}
	for _, test := range tests {
		obj, err := generator.StructuredGenerate()
		switch {
		case test.expectErr && err != nil:
			continue // loop, since there's no output to check
		case test.expectErr && err == nil:
			t.Errorf("%v: expected error and didn't get one", )
			continue // loop, no expected output object
		case !test.expectErr && err != nil:
			t.Errorf("%v: unexpected error %v", err)
			continue // loop, no output object
		case !test.expectErr && err == nil:
			// do nothing and drop through
		}
		if !reflect.DeepEqual(obj, test.expected) {
			t.Errorf("\nexpected:\n%#v\nsaw:\n%#v", test.expected, obj)
		}
	}
}

func newInt32(val int) *int32{
	p := new(int32)
	*p = int32(val)
	return p
}