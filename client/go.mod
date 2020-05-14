module github.com/deranjer/pickleit/client

go 1.14

replace github.com/amlwwalker/pickleit => ../../pickleit //alias for local development

require (
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/Sereal/Sereal v0.0.0-20200430150152-3c99d16fbeb1 // indirect
	github.com/amlwwalker/pickleit v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	google.golang.org/appengine v1.6.6 // indirect
)
