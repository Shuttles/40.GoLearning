
func stopAllGoRoutines(mig *tosv1alpha1.Mig) error {
	// use namespaced mig crd name as channel name
	prefix := getGroutineChanName(mig.Namespace, mig.Name)

	for _, chanName := range gm.List() {
		if hasPrefix(chanName, prefix, chanSeparator) {
			if err := gm.StopGoroutine(chanName, grmgr.TriggerByIntervention); err != nil {
				return err
			}
		}
	}
	return nil
}