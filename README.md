A more convenient and powerful go configuration reader based on [viper](https://github.com/spf13/viper).

#### Enhancement
* more convenient init configuration, only use url to locate data source
* support etcd v3 and watcher mechanism, consul polling watcher mechanism
* three load failed handle strategy (load last valid serialized config data / panic directly / retry N times)

#### Url pattern
* File: relative path(`file://./config/package.json`), absolute path(`file:///config/package.json`)
* Etcdv3: etcdv3://www.etcd-server.com:2379/config.json
* Consul: consul:///www.consul-servre.com/config.json

#### TODO
* support more reader, only support file, etcdv3, consul now

#### Dependence
[viper](https://github.com/spf13/viper)  
[etcdv3]()  