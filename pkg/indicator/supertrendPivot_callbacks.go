// Code generated by "callbackgen -type Supertrend"; DO NOT EDIT.

package indicator

import ()

func (inc *PivotSupertrend) OnUpdate(cb func(value float64)) {
	inc.UpdateCallbacks = append(inc.UpdateCallbacks, cb)
}

func (inc *PivotSupertrend) EmitUpdate(value float64) {
	for _, cb := range inc.UpdateCallbacks {
		cb(value)
	}
}
