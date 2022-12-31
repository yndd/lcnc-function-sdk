package fn

import (
	"encoding/json"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Resources struct {
	// holds the input KRM resources with the key being GVK in string format
	Input map[string][]runtime.RawExtension `json:"input,omitempty" yaml:"input,omitempty"`
	// holds the output KRM resources with the key being GVK in string format
	Output map[string][]runtime.RawExtension `json:"output,omitempty" yaml:"output,omitempty"`
	// holds the conditional KRM resources with the key being GVK in string format
	Conditions map[string][]runtime.RawExtension `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

// An Object is a Kubernetes object.
type Object interface {
	metav1.Object
	runtime.Object
}

func (r *Resources) AddIntput(o Object) error {
	gvkString := GetGVKString(o)
	_, ok := r.Input[gvkString]
	if !ok {
		r.Input[gvkString] = []runtime.RawExtension{}
	}
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	r.Input[gvkString] = append(r.Input[gvkString], runtime.RawExtension{Raw: b})
	return nil
}

func (r *Resources) AddOutput(o Object) error {
	gvkString := GetGVKString(o)
	_, ok := r.Output[gvkString]
	if !ok {
		r.Output[gvkString] = []runtime.RawExtension{}
	}
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	r.Output[gvkString] = append(r.Output[gvkString], runtime.RawExtension{Raw: b})
	return nil
}

func (r *Resources) AddCondition(o Object) error {
	gvkString := GetGVKString(o)
	_, ok := r.Conditions[gvkString]
	if !ok {
		r.Conditions[gvkString] = []runtime.RawExtension{}
	}
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	r.Conditions[gvkString] = append(r.Conditions[gvkString], runtime.RawExtension{Raw: b})
	return nil
}

func (r *Resources) AddUniqueIntput(o Object) error {
	gvkString := GetGVKString(o)
	_, ok := r.Input[gvkString]
	if !ok {
		r.Input[gvkString] = []runtime.RawExtension{}
	}
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	present, idx := isPresent(r.Input[gvkString], o)
	if !present {
		r.Input[gvkString] = append(r.Input[gvkString], runtime.RawExtension{Raw: b})
	} else {
		// overwrite
		r.Input[gvkString][idx] = runtime.RawExtension{Raw: b}
	}
	return nil
}

func (r *Resources) AddUniqueOutput(o Object) error {
	gvkString := GetGVKString(o)
	_, ok := r.Output[gvkString]
	if !ok {
		r.Output[gvkString] = []runtime.RawExtension{}
	}
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	present, idx := isPresent(r.Output[gvkString], o)
	if !present {
		r.Output[gvkString] = append(r.Output[gvkString], runtime.RawExtension{Raw: b})
	} else {
		// overwrite
		r.Output[gvkString][idx] = runtime.RawExtension{Raw: b}
	}
	return nil
}

func (r *Resources) AddUniqueCondition(o Object) error {
	gvkString := GetGVKString(o)
	_, ok := r.Conditions[gvkString]
	if !ok {
		r.Conditions[gvkString] = []runtime.RawExtension{}
	}
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	present, idx := isPresent(r.Conditions[gvkString], o)
	if !present {
		r.Conditions[gvkString] = append(r.Conditions[gvkString], runtime.RawExtension{Raw: b})
	} else {
		// overwrite
		r.Conditions[gvkString][idx] = runtime.RawExtension{Raw: b}
	}
	return nil
}


func GetGVKString(o Object) string {
	gvk := o.GetObjectKind().GroupVersionKind()
	return GVKToString(&gvk)
}

func GVKToString(gvk *schema.GroupVersionKind) string {
	return fmt.Sprintf("%s.%s.%s", gvk.Kind, gvk.Version, gvk.Group)
}

func StringToGVK(s string) *schema.GroupVersionKind {
	var gvk *schema.GroupVersionKind
	if strings.Count(s, ".") >= 2 {
		s := strings.SplitN(s, ".", 3)
		gvk = &schema.GroupVersionKind{Group: s[2], Version: s[1], Kind: s[0]}
	}
	return gvk
}

func isPresent(slice []runtime.RawExtension, o Object) (bool, int) {
	for idx, v := range slice {
		u := &unstructured.Unstructured{}
		if err := json.Unmarshal(v.Raw, u); err != nil {
			return false, 0
		}
		if u.GetName() == o.GetName() && u.GetNamespace() == o.GetNamespace() {
			return true, idx
		}
	}
	return false, 0
}
