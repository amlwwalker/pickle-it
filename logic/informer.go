package logic

// this should enumerate certain message types that the front end can retrieve
// over a channel. the manager will output certain message types at certain times.

// An Op is a type that is used to describe what type
// of event has occurred during the management process.
type OpCode uint32

type OperatingMessage struct {
	Code        OpCode
	data        string
	CustomField string
}

func (op *OperatingMessage) Custom() string {
	if op.CustomField != "" {
		return op.CustomField
	} else {
		return op.data
	}
	return "???"
}

// Ops
const (
	Op_NewDiff OpCode = iota
	Op_NewFile
	Op_NewBase
	Op_WatchCommencing
	Op_WatchStopped
	Op_Message
	Op_EnablingPlugin
	Op_PluginEnabled
	Op_PluginError
	Op_None
)

var ops = map[OpCode]OperatingMessage{
	Op_NewDiff:         OperatingMessage{Op_NewDiff, "New diff created", ""},
	Op_NewFile:         OperatingMessage{Op_NewFile, "New file created", ""},
	Op_NewBase:         OperatingMessage{Op_NewBase, "New base created", ""},
	Op_WatchCommencing: OperatingMessage{Op_WatchCommencing, "File watching has started", ""},
	Op_WatchStopped:    OperatingMessage{Op_WatchStopped, "File watching has stopped", ""},
	Op_Message:         OperatingMessage{Op_Message, "Custom message attached - ", ""},
	Op_EnablingPlugin:  OperatingMessage{Op_EnablingPlugin, "Enabling Plugin - ", ""},
	Op_PluginEnabled:   OperatingMessage{Op_PluginEnabled, "Plugin Enabled", ""},
	Op_PluginError:     OperatingMessage{Op_PluginError, "Error enabling plugin", ""},
	Op_None:            OperatingMessage{Op_None, "No error code known", ""},
}

// String prints the string version of the Op consts
func (e OpCode) String() string {
	if op, found := ops[e]; found {
		return op.data
	}
	return "???"
}

func (e OpCode) Retrieve() OperatingMessage {
	if op, found := ops[e]; found {
		return op
	}
	return ops[Op_None]
}
