package v1alpha1

type User struct {
	Name   string   `json:"name" protobuf:"bytes,1,opt,name=name"`
	Groups []string `json:"groups" protobuf:"bytes,2,rep,name=groups"`
}
