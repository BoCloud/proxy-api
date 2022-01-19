#### UDPRoute

- UDPRoute是Kubernetes中的一种自定义资源，在Kubernetes中使用UDPRoute来定义udp路由信息，Ingress Controller会监听该资源的创建、更新、删除等事件并同步更新Nginx配置。当在集群中创建UDPRoute，Ingress Controller会首先对该资源的配置进行检查，检查通过后则UDPRoute创建成功，Ingress Controller接收到UDPRoute的创建成功事件并更新Nginx配置。

- 文件位置： [udproute](../apis/proxy/v1beta1/udproute_types.go)

- 规定：udproute字段注释存在 +unsupported，表示Ingress Controller尚未支持该项配置


#### Spec

| 字段             | 类型                           | 必填 | 描述                                                         | 示例  |
| ---------------- | ------------------------------ | ---- | ------------------------------------------------------------ | ----- |
| IngressClassName | *string                        | 否   | IngressClass的名称，如果为空则使用默认IngressClass，用于指定哪个Controller处理此资源 | nginx |
| Streams          | [[Stream](udproute.md#stream)] | 是   | 多组udp路由配置                                              |       |

##### Stream

- Spec.Streams[0]

| 字段        | 类型                    | 必填 | 描述                            | 示例     |
| ----------- | ----------------------- | ---- | ------------------------------- | -------- |
| Port        | int32                   | 是   | 外部端口                        | 8001     |
| TLS         | *[TLS](udproute.md#tls) | 否   | 用于需要tls认证的tcp通信        | 尚未支持 |
| ServiceName | string                  | 是   | kubernetes中Service名称         | nginx    |
| ServicePort | string                  | 是   | kubernetes中Service的port字段值 | 8081     |

##### TLS

- Spec.Streams[0].TLS

| 字段   | 类型   | 必填 | 描述                                                    | 示例        |
| ------ | ------ | ---- | ------------------------------------------------------- | ----------- |
| Secret | string | 是   | 存储tls证书的secret名称，需要和此CRD在同一个namespace下 | secret-name |

##### Status

| 字段       | 类型  | 必填 | 描述                                  | 示例 |
| ---------- | ----- | ---- | ------------------------------------- | ---- |
| Conditions | array | 否   | 该CRD资源的基本状态信息，由控制器填充 |      |



#### 示例

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: UDPRoute
metadata:
  name: udproute-sample
spec:
  ingressClassName: nginx
  streams:
    - port: 8800
      serviceName: web
      servicePort: 80
```