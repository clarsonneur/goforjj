package main

// __MYPLUGIN: {{ range $ObjectName, $Object := .Yaml.Objects }}\
// Object {{ $ObjectName }} groups structure

// Groups structure

// __MYPLUGIN: {{   range $GroupName, $Group := $Object.Groups }}\
type DataStruct struct { // __MYPLUGIN: type {{ go_vars $GroupName }}Struct struct {
// __MYPLUGIN: {{     range $FlagName, $Flag := $Group.Flags }}\
// __MYPLUGIN: {{       if (eq $Flag.Type "[]string")}}\
	Data2 []string `json:"data-data2"` // __MYPLUGIN: 	{{ go_vars $FlagName }} []string `json:"{{ $GroupName }}-{{ $FlagName }}"` // {{ $Flag.Help }}
// __MYPLUGIN: {{       end }}\
// __MYPLUGIN: {{       if (or (not $Flag.Type) (eq $Flag.Type "string"))}}\
	Data1 string   `json:"data-data1"` // __MYPLUGIN: 	{{ go_vars $FlagName }} string `json:"{{ $GroupName }}-{{ $FlagName }}"` // {{ $Flag.Help }}
// __MYPLUGIN: {{       end }}\
// __MYPLUGIN: {{     end }}\
}

// __MYPLUGIN: {{   end }}\

// Object Instance structures

type AppInstanceStruct struct { // __MYPLUGIN: type {{ go_vars $ObjectName}}InstanceStruct struct {
// __MYPLUGIN: {{   range $ParamName, $Opts := $Object.Flags }}\
// __MYPLUGIN: {{     if eq $Opts.Type "[]string"}}\
	Param2 []string `json:"param2"` // __MYPLUGIN: 	{{ go_vars $ParamName }} []string `json:"{{ $ParamName }}"` // {{ $Opts.Help }}
// __MYPLUGIN: {{     end }}\
// __MYPLUGIN: {{     if or (not $Opts.Type) (eq $Opts.Type "string")}}\
	Param1 string   `json:"param1"` // __MYPLUGIN: 	{{ go_vars $ParamName }} string `json:"{{ $ParamName }}"` // {{ $Opts.Help }}
// __MYPLUGIN: {{     end }}\
// __MYPLUGIN: {{   end }}\
// __MYPLUGIN: {{   if $Object.Groups }}\

	// Groups

// __MYPLUGIN: {{   end }}\
// __MYPLUGIN: {{   range $GroupName, $Group := $Object.Groups }}\
	DataStruct // __MYPLUGIN: 	{{ go_vars $GroupName }}Struct
// __MYPLUGIN: {{   end }}\
}

// __MYPLUGIN: {{ end }}\

// ************************
// Create request structure
// ************************

type CreateReq struct {
	Forj struct {
// __MYPLUGIN: {{ range $FlagName, $FlagOpts := .Yaml.Tasks.common }}\
		ForjjInstanceName string `json:"forjj-instance-name"` // __MYPLUGIN: 		{{ go_vars $FlagName }} string `json:"{{ $FlagName }}"`
		ForjjSourceMount  string `json:"forjj-source-mount"`  // __MYPLUGIN: {{ end }}\
// __MYPLUGIN: {{ range $FlagName, $FlagOpts := .Yaml.Tasks.create }}\
// __MYPLUGIN: 		{{ go_vars $FlagName }} string `json:"{{ $FlagName }}"`
// __MYPLUGIN: {{ end }}\
	}
	Objects CreateArgReq
}

type CreateArgReq struct {
// __MYPLUGIN: {{ range $Objectname, $ObjectOpts := .Yaml.Objects }}\
	App map[string]AppInstanceStruct // __MYPLUGIN: 	{{ go_vars $Objectname}} map[string]{{ go_vars $Objectname}}InstanceStruct `json:"{{$Objectname}}"` // Object details
// __MYPLUGIN: {{ end }}\
}

// ************************
// Update request structure
// ************************

type UpdateReq struct {
	Forj struct {
// __MYPLUGIN: {{ range $FlagName, $FlagOpts := .Yaml.Tasks.common }}\
		ForjjInstanceName string `json:"forjj-instance-name"` // __MYPLUGIN: 		{{ go_vars $FlagName }} string `json:"{{ $FlagName }}"`
		ForjjSourceMount  string `json:"forjj-source-mount"`  // __MYPLUGIN: {{ end }}\
		// __MYPLUGIN: {{ range $FlagName, $FlagOpts := .Yaml.Tasks.update }}\
		// __MYPLUGIN: 		{{ go_vars $FlagName }} string `json:"{{ $FlagName }}"`
		// __MYPLUGIN: {{ end }}\
	}
	Objects UpdateArgReq
}

type UpdateArgReq struct {
// __MYPLUGIN: {{ range $Objectname, $Opts := .Yaml.Objects }}\
	App map[string]AppInstanceStruct // __MYPLUGIN: 	{{ go_vars $Objectname}} map[string]{{ go_vars $Objectname}}InstanceStruct `json:"{{$Objectname}}"` // Object details
	// __MYPLUGIN: {{ end }}\
}

// **************************
// Maintain request structure
// **************************

type MaintainReq struct {
	Forj struct {
// __MYPLUGIN: {{ range $FlagName, $FlagOpts := .Yaml.Tasks.common }}\
		ForjjInstanceName string `json:"forjj-instance-name"` // __MYPLUGIN: 		{{ go_vars $FlagName }} string `json:"{{ $FlagName }}"`
		ForjjSourceMount  string `json:"forjj-source-mount"`  // __MYPLUGIN: {{ end }}\
// __MYPLUGIN: {{ range $FlagName, $FlagOpts := .Yaml.Tasks.maintain }}\
// __MYPLUGIN: 		{{ go_vars $FlagName }} string `json:"{{ $FlagName }}"`
// __MYPLUGIN: {{ end }}\
	}
	Objects MaintainArgReq
}

type MaintainArgReq struct {
// __MYPLUGIN: {{ range $Objectname, $Opts := .Yaml.Objects }}\
// __MYPLUGIN: {{   if object_has_secure $Opts }}\
	App map[string]AppMaintainStruct `json:"app"` // __MYPLUGIN: 	{{ go_vars $Objectname}} map[string]{{ go_vars $Objectname}}MaintainStruct `json:"{{$Objectname}}"` // Object details
	// __MYPLUGIN: {{   end }}\
	// __MYPLUGIN: {{ end }}\
}

// __MYPLUGIN: {{ range $Objectname, $ObjectOpts := .Yaml.Objects }}\
// __MYPLUGIN: {{   if object_has_secure $ObjectOpts }}\
type AppMaintainStruct struct { // __MYPLUGIN: type {{ go_vars $Objectname}}MaintainStruct struct {
	// __MYPLUGIN: {{ range $ParamName, $Opts := $ObjectOpts.Flags }}\
	// __MYPLUGIN: {{   if $Opts.Options.Secure }}\
	Token string // __MYPLUGIN: 	{{ go_vars $ParamName }} string `json:"{{ $ParamName }}"` // {{ $Opts.Help }}
	// __MYPLUGIN: {{   end }}\
	// __MYPLUGIN: {{ end }}\
}

// __MYPLUGIN: {{   end }}\
// __MYPLUGIN: {{ end }}\

// YamlDesc has been created from your '{{.Yaml.Name}}.yaml' file.
const YamlDesc = "{{ escape .Yaml_data}}"
