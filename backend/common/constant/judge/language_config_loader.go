package judge

type languageConfig struct {
	Language       string
	SrcName        string
	ExeName        string
	MaxCpuTime     int64
	MaxRealTime    int64
	MaxMemory      int64
	CompileCommand []string
	RunCommand     []string
	CompileEnvs    []string
	RunEnvs        []string
}

var (
	defaultEnv = []string{
		"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
		"LANG=en_US.UTF-8",
		"LANGUAGE=en_US:en",
		"HOME=/w"}

	python3Env = []string{"LANG=en_US.UTF-8",
		"LANGUAGE=en_US:en", "PYTHONIOENCODING=utf-8"}

	golangCompileEnv = []string{
		"GOCACHE=/w", "GOPATH=/w/go", "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
		"LANG=en_US.UTF-8", "LANGUAGE=en_US:en"}

	golangRunEnv = []string{
		"GOCACHE=off", "GODEBUG=madvdontneed=1",
		"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
		"LANG=en_US.UTF-8", "LANGUAGE=en_US:en"}

	cLanguage = languageConfig{
		Language:       "c",
		SrcName:        "main.c",
		ExeName:        "main",
		RunEnvs:        defaultEnv,
		CompileEnvs:    defaultEnv,
		CompileCommand: []string{"/usr/bin/gcc", "-DONLINE_JUDGE", "-w", "-fmax-errors=1", "-std=c11", "main.c", "-lm", "-o", "main"},
		RunCommand:     []string{"/w/main"},
		MaxCpuTime:     3 * 1000 * 1000000,
		MaxRealTime:    10 * 1000 * 1000000,
		MaxMemory:      256 * 1024 * 1024,
	}

	cppLanguage = languageConfig{
		Language:       "cpp",
		SrcName:        "main.cpp",
		ExeName:        "main",
		RunEnvs:        defaultEnv,
		CompileEnvs:    defaultEnv,
		CompileCommand: []string{"/usr/bin/g++", "-DONLINE_JUDGE", "-w", "-fmax-errors=1", "-std=c++14", "main.cpp", "-lm", "-o", "main"},
		RunCommand:     []string{"/w/main"},
		MaxCpuTime:     10 * 1000 * 1000000,
		MaxRealTime:    30 * 1000 * 1000000,
		MaxMemory:      512 * 1024 * 1024,
	}

	javaLanguage = languageConfig{
		Language:       "java",
		SrcName:        "Main.java",
		ExeName:        "Main.jar",
		RunEnvs:        defaultEnv,
		CompileEnvs:    defaultEnv,
		CompileCommand: []string{"/bin/bash", "-c", "javac -encoding utf-8 Main.java && jar -cvf Main.jar *.class"},
		RunCommand:     []string{"/usr/bin/java", "-Dfile.encoding=UTF-8", "-cp", "/w/Main.jar", "Main"},
		MaxCpuTime:     10 * 1000 * 1000000,
		MaxRealTime:    20 * 1000 * 1000000,
		MaxMemory:      512 * 1024 * 1024,
	}

	pythonLanguage = languageConfig{
		Language:       "python",
		SrcName:        "main.py",
		ExeName:        "__pycache__/main.cpython-37.pyc",
		RunEnvs:        python3Env,
		CompileEnvs:    python3Env,
		CompileCommand: []string{"/usr/bin/python3.7", "-m", "py_compile", "/w/main.py"},
		RunCommand:     []string{"/usr/bin/python3.7", "/w/__pycache__/main.cpython-37.pyc"},
		MaxCpuTime:     5 * 1000 * 1000000,
		MaxRealTime:    10 * 1000 * 1000000,
		MaxMemory:      256 * 1024 * 1024,
	}

	golangLanguage = languageConfig{
		Language:       "go",
		SrcName:        "main.go",
		ExeName:        "main",
		RunEnvs:        golangRunEnv,
		CompileEnvs:    golangCompileEnv,
		CompileCommand: []string{"/usr/bin/go", "build", "-o", "main", "main.go"},
		RunCommand:     []string{"/w/main"},
		MaxCpuTime:     10 * 1000 * 1000000,
		MaxRealTime:    5 * 1000 * 1000000,
		MaxMemory:      512 * 1024 * 1024,
	}

	javascriptLanguage = languageConfig{
		Language:    "javascript",
		SrcName:     "main.js",
		ExeName:     "main.js",
		RunEnvs:     defaultEnv,
		CompileEnvs: defaultEnv,
		RunCommand:  []string{"/usr/bin/jsv8/d8", "/w/main.js"},
		MaxCpuTime:  5 * 1000 * 1000000,
		MaxRealTime: 10 * 1000 * 1000000,
		MaxMemory:   256 * 1024 * 1024,
	}

	LanguageConfigMap = map[string]languageConfig{
		"c":          cLanguage,
		"cpp":        cppLanguage,
		"java":       javaLanguage,
		"python":     pythonLanguage,
		"go":         golangLanguage,
		"javascript": javascriptLanguage,
	}
)
