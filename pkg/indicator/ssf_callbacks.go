// Code generated by "callbackgen -type SSF"; DO NOT EDIT.

package indicator

import ()

func (inc *SSF) OnUpdate(cb func(value float64)) {
	inc.UpdateCallbacks = append(inc.UpdateCallbacks, cb)
}

func (inc *SSF) EmitUpdate(value float64) {
	for _, cb := range inc.UpdateCallbacks {
		cb(value)
	}
}
