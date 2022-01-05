### 简介

- 该项目主要为扩展Ingress功能，而设计的CRD资源
- HTTPRoute:七层负载配置，Ingress-Controller会监听此资源变化并应用到Nginx配置
- TCPRoute:四层负载配置，Ingress-Controller会监听此资源变化并应用到Nginx配置
- UDPRoute:四层负载皮遏制，Ingress-Controller会监听此资源变化并应用到Nginx配置
- 该项目使用kubeBuilder3.2.0工具生成

#### 修改CRD字段后

- 使用命令`make manifests` 生成CRD资源
- 使用命令`make controller-gen` 生成Controllers代码

#### 使用文档

- [httproute](docs/httproute.md)
- [tcproute](docs/tcproute.md)
- [udproute](docs/udproute.md)