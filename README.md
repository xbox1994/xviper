A more convenient and powerful go configuration reader based on [viper](https://github.com/spf13/viper).

#### Enhancement
* more convenient init configuration, only use url to locate data source(e.g.: etcd://xxx:2379/path)
* support etcd v3 and watcher mechanism, consul polling watcher mechanism
* three load failed handle strategy (load last valid serialized config data / panic directly / retry N times)

#### TODO
* support more reader, only support file, etcdv3, consul now

#### Dependence
[viper](https://github.com/spf13/viper)  
[etcdv3]()  