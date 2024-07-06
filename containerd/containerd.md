## containerd 架构
![](./images/containerd-architecture.png)
## containerd 目录结构
```shell
obitoquilt@ubuntu-k3s:/var/lib/rancher/k3s/agent/containerd$ tree /var/lib/rancher/k3s/agent/containerd -L 2
/var/lib/rancher/k3s/agent/containerd
├── bin
├── containerd.log
├── io.containerd.content.v1.content
│   ├── blobs
│   └── ingest
├── io.containerd.grpc.v1.introspection
│   └── uuid
├── io.containerd.metadata.v1.bolt
│   └── meta.db
├── io.containerd.runtime.v1.linux
├── io.containerd.runtime.v2.task
├── io.containerd.snapshotter.v1.btrfs
│   ├── active
│   ├── snapshots
│   └── view
├── io.containerd.snapshotter.v1.fuse-overlayfs
│   └── snapshots
├── io.containerd.snapshotter.v1.native
│   └── snapshots
├── io.containerd.snapshotter.v1.overlayfs
│   ├── metadata.db
│   └── snapshots
├── io.containerd.snapshotter.v1.stargz
│   ├── snapshotter
│   └── stargz
├── lib
└── tmpmounts
```

## 拉取镜像时 containerd 目录文件变化
```shell
$ k3s crictl pull alpine:3.20.1

obitoquilt@ubuntu-k3s:~/k3s$ inotifywait -m -r -e create -e modify -e delete /var/lib/rancher/k3s/agent | grep -v log
Setting up watches.  Beware: since -r was given, this may take a while!
Watches established.
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/ CREATE,ISDIR e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ MODIFY startedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ CREATE updatedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ MODIFY updatedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ CREATE total
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ MODIFY total
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ CREATE data
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ MODIFY data
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ DELETE ref
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ DELETE startedat
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ DELETE updatedat
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426/ DELETE total
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/ DELETE,ISDIR e76113e3d92727912318f58eacdb1bbfc27d096139e55134d703135afc9cb426
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/ CREATE,ISDIR eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ CREATE ref
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ MODIFY ref
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ CREATE startedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ MODIFY startedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ CREATE updatedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ MODIFY updatedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ CREATE total
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ MODIFY total
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ CREATE data
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ MODIFY data
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ DELETE ref
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ DELETE startedat
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ DELETE updatedat
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9/ DELETE total
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/ DELETE,ISDIR eb37adc8510dc1eee844e94c64c18ab6e05c501b228bd7f7580ed744e24b1ec9
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/ CREATE,ISDIR 813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ CREATE ref
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ MODIFY ref
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ CREATE startedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ MODIFY startedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ CREATE updatedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ MODIFY updatedat.tmp
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ CREATE total
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ MODIFY total
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ CREATE data
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ MODIFY data
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ DELETE ref
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ DELETE startedat
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ DELETE updatedat
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1/ DELETE total
/var/lib/rancher/k3s/agent/containerd/io.containerd.content.v1.content/ingest/ DELETE,ISDIR 813901937ca45648ac5dddb49842259a1e6e8be03802caf6a58a307e66b8c6e1
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
/var/lib/rancher/k3s/agent/containerd/io.containerd.metadata.v1.bolt/ MODIFY meta.db
```

## 参考
1. [Containerd是如何存储容器镜像和数据的](https://blog.frognew.com/2021/06/relearning-container-09.html)