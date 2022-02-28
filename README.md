### 简介

- 该项目主要为扩展Ingress功能，而设计的CRD资源
- HTTPRoute: 七层负载配置，兼容原生Ingress功能并进行了极大扩展
- TCPRoute: 四层负载配置，支持TCP服务
- UDPRoute: 四层负载配置，支持UDP服务
- 该项目使用kubeBuilder3.2.0工具生成

#### 修改CRD字段后

- 使用命令`make manifests` 生成CRD资源
- 使用命令`make controller-gen` 生成controllers代码

- 使用命令`make legacy` 生成v1beta1版本的CRD资源,支持低于v1.16版本k8s

#### 使用文档

- [httproute](docs/httproute.md)
- [tcproute](docs/tcproute.md)
- [udproute](docs/udproute.md)