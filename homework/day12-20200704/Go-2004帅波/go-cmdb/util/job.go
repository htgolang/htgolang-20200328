package util


type JsonJob struct {
	Class string `json:"_class"`
	Actions []Actions `json:"actions"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
	DisplayNameOrNull interface{} `json:"displayNameOrNull"`
	FullDisplayName string `json:"fullDisplayName"`
	FullName string `json:"fullName"`
	Name string `json:"name"`
	URL string `json:"url"`
	Buildable bool `json:"buildable"`
	Builds []interface{} `json:"builds"`
	Color string `json:"color"`
	FirstBuild interface{} `json:"firstBuild"`
	HealthReport []interface{} `json:"healthReport"`
	InQueue bool `json:"inQueue"`
	KeepDependencies bool `json:"keepDependencies"`
	LastBuild interface{} `json:"lastBuild"`
	LastCompletedBuild interface{} `json:"lastCompletedBuild"`
	LastFailedBuild interface{} `json:"lastFailedBuild"`
	LastStableBuild interface{} `json:"lastStableBuild"`
	LastSuccessfulBuild interface{} `json:"lastSuccessfulBuild"`
	LastUnstableBuild interface{} `json:"lastUnstableBuild"`
	LastUnsuccessfulBuild interface{} `json:"lastUnsuccessfulBuild"`
	NextBuildNumber int `json:"nextBuildNumber"`
	Property []Property `json:"property"`
	QueueItem interface{} `json:"queueItem"`
	ConcurrentBuild bool `json:"concurrentBuild"`
	Disabled bool `json:"disabled"`
	DownstreamProjects []interface{} `json:"downstreamProjects"`
	LabelExpression interface{} `json:"labelExpression"`
	Scm Scm `json:"scm"`
	UpstreamProjects []interface{} `json:"upstreamProjects"`
}
type DefaultParameterValue struct {
	Class string `json:"_class"`
	Value string `json:"value"`
}
type ParameterDefinitions struct {
	Class string `json:"_class"`
	DefaultParameterValue DefaultParameterValue `json:"defaultParameterValue"`
	Description string `json:"description"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type Actions struct {
	Class string `json:"_class,omitempty"`
	ParameterDefinitions []ParameterDefinitions `json:"parameterDefinitions,omitempty"`
}
type Property struct {
	Class string `json:"_class"`
	ParameterDefinitions []ParameterDefinitions `json:"parameterDefinitions,omitempty"`
}
type Scm struct {
	Class string `json:"_class"`
}


func NewJsonJob() *JsonJob{
	return &JsonJob{}
}
