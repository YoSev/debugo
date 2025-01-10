package debugo

import "fmt"

func (l *Debugger) Extend(namespace string) *Debugger {
	d := new(fmt.Sprintf("%s:%s", l.namespace, namespace), l.options)
	d.applyOptions()

	// ensure to keep color for extended namespaces
	d.color = l.color

	return d
}
