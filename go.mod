module github.com/mi4tin/go-chassis-gate

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181001203147-e3636079e1a4
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20181116152217-5ac8a444bdc5
	golang.org/x/tools => github.com/golang/tools v0.0.0-20181219222714-6e267b5cc78e
)

require (
	github.com/emicklei/go-restful v2.11.2+incompatible
	github.com/go-chassis/go-archaius v1.2.0
	github.com/go-chassis/go-chassis v1.8.2
	github.com/go-chassis/go-restful-swagger20 v1.0.2
	github.com/go-mesh/openlogging v1.0.1
	github.com/go-yaml/yaml v2.1.0+incompatible
	golang.org/x/sys v0.0.0-20191018095205-727590c5006e // indirect
)
