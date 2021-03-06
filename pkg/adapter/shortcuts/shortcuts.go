package shortcuts

import "github.com/containers/libpod/libpod"

// GetPodsByContext gets pods whether all, latest, or a slice of names/ids
func GetPodsByContext(all, latest bool, pods []string, runtime *libpod.Runtime) ([]*libpod.Pod, error) {
	var outpods []*libpod.Pod
	if all {
		return runtime.GetAllPods()
	}
	if latest {
		p, err := runtime.GetLatestPod()
		if err != nil {
			return nil, err
		}
		outpods = append(outpods, p)
		return outpods, nil
	}
	for _, p := range pods {
		pod, err := runtime.LookupPod(p)
		if err != nil {
			return nil, err
		}
		outpods = append(outpods, pod)
	}
	return outpods, nil
}

// GetContainersByContext gets pods whether all, latest, or a slice of names/ids
func GetContainersByContext(all, latest bool, names []string, runtime *libpod.Runtime) ([]*libpod.Container, error) {
	var ctrs = []*libpod.Container{}

	if all {
		return runtime.GetAllContainers()
	}

	if latest {
		c, err := runtime.GetLatestContainer()
		if err != nil {
			return nil, err
		}
		ctrs = append(ctrs, c)
		return ctrs, nil
	}

	for _, c := range names {
		ctr, err := runtime.LookupContainer(c)
		if err != nil {
			return nil, err
		}
		ctrs = append(ctrs, ctr)
	}
	return ctrs, nil
}
