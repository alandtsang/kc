package action

type Action int

const (
	ActionCreate Action = iota
	ActionExpose
	ActionRun
	ActionSet
	ActionExplain
	ActionGet
	ActionEdit
	ActionDelete
	ActionRollout
	ActionScale
	ActionAutoscale
	ActionCertificate
	ActionClusterInfo
	ActionTop
	ActionCordon
	ActionUncordon
	ActionDrain
	ActionTaint
	ActionDescribe
	ActionLogs
	ActionAttach
	ActionExec
	ActionPortForward
	ActionProxy
	ActionCp
	ActionAuth
	ActionApply
	ActionPatch
	ActionReplace
	ActionWait
	ActionConvert
	ActionLabel
	ActionAnnotate
	ActionCompletion
	ActionAlpha
	ActionApiResources
	ActionApiVersions
	ActionConfig
	ActionPlugin
	ActionVersion
)
