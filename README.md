# dockit [![GoDoc](https://godoc.org/github.com/zepouet/dockit?status.svg)](https://godoc.org/github.com/treeptik/docker-watchdog) [![Build Status](https://travis-ci.org/treeptik/docker-monitor.svg)](https://travis-ci.org/treeptik/docker-monitor)
dockit is a toolkit with many useful commands to help devops people when using [Docker](http://www.docker.com) containers and images.

## Type implemented

### *Heimdall* (Heimdallr) as your docker gardian
- [ ] **wtf** General information about environment
- [ ] **detail** Userfriendly container info display
- [ ] **tree** Display with a tree the relation between containers and volumes
- [ ] **orphean** List the orphean data
- [ ] **bubbles** Bubbles visualization sort by images & size
- [ ] **neo** Matrice visualization by link & volume
- [ ] **hyperview** : Introspect the container at runtime to display process nature. Can be helped by friends like Jmxtrans.

### *Proteus* (Πρωτεύς) as your storekeeper 
- [ ] **save** backup your volume as tar.gz
- [ ] **load** restore your volume from tar.gz
- [ ] **backup** backup your volume with btrfs support 
- [ ] **clone** clone your volume with btrfs support 

## Install and Launch

### Download & Compile
```
go get github.com/zepouet/dockit
glide install
./install.sh
```

### Configure & Launch

dockit using one system variable for its configuration.
```
DOCKER_HOST #default value : unix:///var/run/docker.sock
```
dockit start usage
```
Usage: 
  dockit [flags]
  dockit [command]

Available Commands: 
  version     docker-watchdog version
  help        Help about any command

Flags:
  -d, --debug=false: Run dockit server in "debug" mode
  -h, --help=false: help for dockit
  -p, --port=8080: dockit server port


Use "docker-viz watchdog [command]" for more information about a command
```

# Nice articles and tutorials

Thanks for everybody named after for their works
* Nicolas Deloof https://github.com/ndeloof/docker-gc


