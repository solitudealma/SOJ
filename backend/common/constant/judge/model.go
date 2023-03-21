package judge

// CmdFile defines file from multiple source including local / memory / cached or pipe collector
type CmdFile struct {
	Src     *string `json:"src"`
	Content *string `json:"content"`
	FileID  *string `json:"fileId"`
	Name    *string `json:"name"`
	Max     *int64  `json:"max"`
	Pipe    bool    `json:"pipe"`
	Symlink *string `json:"symlink"`
}

// Cmd defines command and limits to start a program using in envexec
type Cmd struct {
	Args              []string           `json:"args"`
	Env               []string           `json:"env,omitempty"`
	Files             []*CmdFile         `json:"files,omitempty"`
	Tty               bool               `json:"tty,omitempty"`
	CpuLimit          uint64             `json:"cpuLimit"`
	RealCPULimit      uint64             `json:"realCpuLimit"`
	ClockLimit        uint64             `json:"clockLimit"`
	MemoryLimit       uint64             `json:"memoryLimit"`
	StackLimit        uint64             `json:"stackLimit"`
	ProcLimit         uint64             `json:"procLimit"`
	CpuRateLimit      uint64             `json:"cpuRateLimit"`
	CpuSetLimit       string             `json:"cpuSetLimit"`
	StrictMemoryLimit bool               `json:"strictMemoryLimit"`
	CopyIn            map[string]CmdFile `json:"copyIn"`
	CopyOut           []string           `json:"copyOut"`
	CopyOutCached     []string           `json:"copyOutCached"`
	CopyOutMax        uint64             `json:"copyOutMax"`
	CopyOutDir        string             `json:"copyOutDir"`
}

type FileError struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Message string `json:"message,omitempty"`
}

// Result defines single command result
type Result struct {
	Status     string            `json:"status"`
	ExitStatus int               `json:"exitStatus"`
	Error      string            `json:"error,omitempty"`
	Time       uint64            `json:"time"`
	Memory     uint64            `json:"memory"`
	RunTime    uint64            `json:"runTime"`
	Files      map[string]string `json:"files,omitempty"`
	FileIds    map[string]string `json:"fileIds,omitempty"`
	FileError  []FileError       `json:"fileError,omitempty"`
	Buffs      map[string][]byte `json:"-"`
}
