package forms

import (
	"strconv"
	"strings"

	coreV1 "k8s.io/api/core/v1"
)

type DeploymentCreateForm struct {
	Namespace string `form:"namespace"`
	Name      string `form:"name"`
	Replicas  int32  `form:"replicas"`
	Image     string `form:"image"`
	Labels    string `form:"labels"`
	Ports     string `form:"ports"`
}

func (f *DeploymentCreateForm) GetLabels() map[string]string {
	labelsMap := make(map[string]string)
	labels := strings.Split(f.Labels, "\n")
	for _, label := range labels {
		values := strings.SplitN(label, ":", 2)
		if len(values) != 2 {
			continue
		}
		labelsMap[strings.TrimSpace(values[0])] = strings.TrimSpace(values[1])
	}
	return labelsMap
}

func (f *DeploymentCreateForm) GetSelectors() map[string]string {
	selectors := f.GetLabels()
	selectors["app"] = f.Name
	return selectors
}
func (f *DeploymentCreateForm) GetPorts() []coreV1.ContainerPort {
	portList := make([]coreV1.ContainerPort, 0, 5)
	ports := strings.Split(f.Ports, "\n")
	for _, port := range ports {
		values := strings.SplitN(port, ",", 3)
		if len(values) != 3 {
			continue
		}
		intPort, err := strconv.Atoi(values[1])
		if err != nil {
			continue
		}
		protocol := coreV1.ProtocolTCP
		if strings.Compare(strings.ToLower(values[0]), "tcp") != 0 {
			protocol = coreV1.ProtocolUDP
		}
		portList = append(portList, coreV1.ContainerPort{
			Name:          strings.TrimSpace(values[2]),
			ContainerPort: int32(intPort),
			Protocol:      protocol,
		})
	}

	return portList
}

func (f *DeploymentCreateForm) GetImageName() string {
	// 全部为应为字母数字和:
	pos := strings.Index(f.Image, ":")
	if pos > 0 {
		return f.Image[:pos]
	}
	return f.Image
}
