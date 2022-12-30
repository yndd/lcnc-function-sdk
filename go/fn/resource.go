package fn

import (
	"encoding/json"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Resources struct {
	// holds the input KRM resources with the key being GVK in string format
	Input map[string][]string `json:"input,omitempty" yaml:"input,omitempty"`
	// holds the output KRM resources with the key being GVK in string format
	Output map[string][]string `json:"output,omitempty" yaml:"output,omitempty"`
	// holds the conditional KRM resources with the key being GVK in string format
	Conditions map[string][]string `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

func (r *Resources) AddOutput(o *unstructured.Unstructured) error {
	apiversionSplit := strings.Split(o.GetAPIVersion(), "/")
	gvkString := GVKToString(&schema.GroupVersionKind{
		Group:   apiversionSplit[0],
		Version: apiversionSplit[1],
		Kind:    o.GetKind(),
	})
	_, ok := r.Output[gvkString]
	if !ok {
		r.Output[gvkString] = []string{}
	}
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	r.Output[gvkString] = append(r.Output[gvkString], string(b))
	return nil
}

func (r *Resources) AddCondition(o *unstructured.Unstructured) error {
	gvkString := GetGVKString(o)
	_, ok := r.Conditions[gvkString]
	if !ok {
		r.Conditions[gvkString] = []string{}
	}
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	r.Conditions[gvkString] = append(r.Conditions[gvkString], string(b))
	return nil
}

func GetGVKString(o *unstructured.Unstructured) string {
	apiversionSplit := strings.Split(o.GetAPIVersion(), "/")
	return GVKToString(&schema.GroupVersionKind{
		Group:   apiversionSplit[0],
		Version: apiversionSplit[1],
		Kind:    o.GetKind(),
	})
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
