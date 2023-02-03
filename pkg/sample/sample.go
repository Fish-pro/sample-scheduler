package sample

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "sample-scheduler"

func New(obj runtime.Object, f framework.Handle) (framework.Plugin, error) {
	args := &Args{}
	klog.V(3).Infof("get plugin config args: %+v", args)
	return &Sample{
		args:   args,
		handle: f,
	}, nil
}

type Args struct {
	FavoriteColor  string `json:"favorite_color,omitempty"`
	FavoriteNumber int    `json:"favorite_number,omitempty"`
	ThanksTo       string `json:"thanks_to,omitempty"`
}

type Sample struct {
	args   *Args
	handle framework.Handle
}

func (s *Sample) Name() string {
	return Name
}

func (s *Sample) PreFilter(ctx context.Context, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("prefilter pod: %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}

func (s *Sample) Filter(ctx context.Context, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("filter pod: %v, node: %v", pod.Name, nodeName)
	return framework.NewStatus(framework.Success, "")
}

func (s *Sample) PreBind(ctx context.Context, pod *v1.Pod, nodeName string) *framework.Status {
	nodeList, err := s.handle.SnapshotSharedLister().NodeInfos().List()
	if err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("failed to list node: %w", err))
	}
	for _, node := range nodeList {
		if node.Node().Name == nodeName {
			klog.V(3).Infof("prebind node info: %+v", node.Node())
			return framework.NewStatus(framework.Success, "")
		}
	}
	return framework.NewStatus(framework.Error, fmt.Sprintf("prebind get node info error: %+v", nodeName))
}
