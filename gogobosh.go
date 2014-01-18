package gogobosh

type DirectorConfig struct {
	targetURL string
	username  string
	password  string
}

type Director struct {
	Name string
	URL string
	Version string
	User string
	UUID string
	CPI string
	DNSEnabled bool
	DNSDomainName string
	CompiledPackageCacheEnabled bool
	CompiledPackageCacheProvider string
	SnapshotsEnabled bool
}

type VMStatus struct {
	JobName string
	Index int
	JobState string
	VMCid string
	AgentID string
	ResourcePool string
	ResurrectionPaused bool
	IPs []string
	DNSs []string
	CPUUser float64
	CPUSys float64
	CPUWait float64
	MemoryPercent float64
	MemoryKb int
	SwapPercent float64
	SwapKb int
	DiskPersistentPercent float64
}

func NewDirector(targetURL string, username string, password string) (director Director) {
	config := DirectorConfig{}
	config.targetURL = targetURL
	config.username = username
	config.password = password
	
	director = Director{}
	
	return
}