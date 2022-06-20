package config

import (
	"github.com/philchia/agollo/v4"
	"github.com/yunkeCN/apollo-golang/v2/setup"
)

func watchUpdateConfig() {
	if len(setup.WatchConfigFields) == 0 {
		return
	}

	go func() {
		defer func() {
			recover()
		}()

		apolloClient.OnUpdate(func(event *agollo.ChangeEvent) {
			if event.Namespace != currentNamespace {
				return
			}

			for apolloKeyName, value := range event.Changes {
				if value.ChangeType != agollo.MODIFY {
					continue
				}
				if field, ok := setup.WatchConfigFields[apolloKeyName]; ok {
					if !field.Value.CanSet() {
						continue
					}
					_ = field.SetNewValue(value.NewValue)
				}
			}
		})
	}()
}
