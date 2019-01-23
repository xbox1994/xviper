A more convenient and powerful go configuration reader based on [viper](https://github.com/spf13/viper).

#### Enhancement
* more convenient init configuration, only use url to locate data source
* support etcd v3 and watcher mechanism, consul long polling watcher mechanism
* three load failed handle strategy (load last valid serialized config data / panic directly / retry N times)

#### Url pattern
* File: file://./config/package.json(relative path); file:///config/package.json(absolute path)
* Etcdv3: etcdv3://127.0.0.1:2379/config.json
* Consul: consul://127.0.0.1:8500/config.json

#### TODO
* support more reader, only support file, etcdv3, consul now

#### How to use
example/main.go