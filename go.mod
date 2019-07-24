module github.com/michelia/ured

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190506204251-e1dfcc566284
	golang.org/x/net => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190507053917-2953c62de483
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190506145303-2d16b83fe98c
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190513163551-3ee3066db522
)

require (
	github.com/mediocregopher/radix/v3 v3.3.1
	github.com/michelia/ulog v1.0.4
)