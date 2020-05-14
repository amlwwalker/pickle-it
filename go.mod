module github.com/deranjer/pickleit

go 1.14

require (
	github.com/amlwwalker/fdelta v0.0.0-20200513211915-3b53ff25eff6
	github.com/amlwwalker/pickleit v0.0.0-00010101000000-000000000000
	github.com/apsdehal/go-logger v0.0.0-20190515212710-b0d6ccfee0e6
	github.com/asdine/storm v2.1.2+incompatible
	github.com/atrox/homedir v1.0.0
	github.com/boltdb/bolt v1.3.1
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/everdev/mack v0.0.0-20200226161639-15be3d47cc54
	github.com/go-vgo/robotgo v0.0.0-20200509164312-55a4babb8ca7
	github.com/jroimartin/gocui v0.4.0
	github.com/kalafut/imohash v1.0.0
	github.com/kr/binarydist v0.1.0
	github.com/radovskyb/watcher v1.0.7
	github.com/schollz/progressbar/v3 v3.3.3
	github.com/therecipe/qt v0.0.0-20200126204426-5074eb6d8c41
	github.com/ttacon/emoji v0.0.0-20140807004100-e1647f8352b4
	go.etcd.io/bbolt v1.3.4 // indirect
)

replace github.com/amlwwalker/pickleit => ../pickleit //alias for local development
