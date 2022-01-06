#### HTTPRoute

- HTTPRoute是Kubernetes中的一种自定义资源，在Kubernetes中使用HTTPRoute来定义http路由信息，Ingress Controller会监听该资源的创建、更新、删除等事件并同步更新Nginx配置。当在集群中创建HTTPRoute，Ingress Controller会首先对该资源的配置进行检查，检查通过后则HTTPRoute创建成功，Ingress Controller接收到HTTPRoute的创建成功事件并更新Nginx配置。

- Nginx对七层协议的功能支持非常强大，也因此HTTPRoute的设计相对TCPRoute/UDPRoute复杂的多，并且在充分考虑了Nginx功能特性情况下扩展支持多种灰度发布方案。

- 文件位置： [httproute](../apis/proxy/v1beta1/httproute_types.go)

- 规定：文件中字段注释存在 +unsupported，表示Ingress Controller尚未支持该项配置



#### 资源定义

##### Spec

- httproute支持配置多个域名、域名下多个path、每个path又有不同的配置且每个path支持不同的灰度服务，所以其配置非常复杂
- 我们拆成几部分来讲述各个字段的功能含义，如是静态配置会具体到该配置会应用到Nginx的哪个配置字段

| 字段                      | 类型                   | 必填 | 描述                                                         | 示例           |
| ------------------------- | ---------------------- | ---- | ------------------------------------------------------------ | -------------- |
| Spec.IngressClassName     | *string                | 否   | IngressClass的名称，如果为空则使用默认IngressClass，用于指定哪个Controller处理此资源 | nginx          |
| Spec.Routes               | array[Route]           | 是   | 多组http路由配置                                             |                |
| Spec.Routes[0].Host       | string                 | 否   | 域名FQDN                                                     | demo.nginx.com |
| Spec.Routes[0].Protocol   | string                 | 否   | Nginx转发给后端使用的协议，可选值HTTP、HTTPS、GRPC，默认值HTTP；尚未支持GRPC | HTTP           |
| Spec.Routes[0].TLS        | *struct[TLS]           | 否   | Host的tls证书，用于Nginx验证请求来源是否合法                 |                |
| Spec.Routes[0].TLS.Secret | string                 | 是   | 存储tls证书的secret名称，需要和此CRD在同一个namespace下      | tls-cert       |
| Spec.Routes[0].Proxy      | *struct[Proxy]         | 否   | [proxy](httproute-proxy.md)                                  |                |
| Spec.Routes[0].Cors       | *struct[Cors]          | 否   | [cors](httproute-cors.md)                                    |                |
| Spec.Routes[0].Rules      | *struct[HTTPRouteRule] | 是   | 自定义http路由                                               |                |
| Spec.Routes[0].Options    | map[string]string      | 否   | [options](httproute-options.md)                              |                |
| Status.Hostname           | string                 | 否   | 该CRD资源的基本状态信息，由控制器填充                        | 尚未支持       |

##### Route  

- Spec.Routes[0]

| 字段     | 类型                   | 必填 | 描述                                                         | 示例           |
| -------- | ---------------------- | ---- | ------------------------------------------------------------ | -------------- |
| Host     | string                 | 否   | 域名FQDN                                                     | demo.nginx.com |
| Protocol | string                 | 否   | Nginx转发给后端使用的协议，可选值HTTP、HTTPS、GRPC，默认值HTTP；尚未支持GRPC | HTTP           |
| TLS      | *struct[TLS]           | 否   | Host的tls证书，用于Nginx验证请求来源是否合法                 |                |
| Proxy    | *struct[Proxy]         | 否   | [proxy](httproute-proxy.md)                                  |                |
| Cors     | *struct[Cors]          | 否   | [cors](httproute-cors.md)                                    |                |
| Rules    | *struct[HTTPRouteRule] | 是   | 自定义http路由                                               |                |
| Options  | map[string]string      | 否   | [options](httproute-options.md)                              |                |

##### TLS

- Spec.Routes[0].TLS

| 字段   | 类型   | 必填 | 描述                                                    | 示例 |
| ------ | ------ | ---- | ------------------------------------------------------- | ---- |
| Secret | string | 是   | 存储tls证书的secret名称，需要和此CRD在同一个namespace下 |      |

##### HTTPRouteRule

- Spec.Routes[0].Rules[0]

- 定义http路由信息，应用于Nginx location阶段

| 字段           | 类型                    | 必填 | 描述                                          | 示例     |
| -------------- | ----------------------- | ---- | --------------------------------------------- | -------- |
| Path           | string                  | 是   | 路径                                          | /user    |
| PathType       | string                  | 否   | 路径匹配方式exact、prefix、regex，默认prefix  | exact    |
| Rewrite        | string                  | 否   | 路径重写，将接收到的path重写转发到后端        | /abc     |
| Proxy          | *struct[Proxy]          | 否   | [proxy](httproute-proxy.md)                   |          |
| Cors           | *struct[Cors]           | 否   | [cors](httproute-cors.md)                     |          |
| RateLimit      | struct[ReteLimit]       | 否   | [ratelimit](httproute-ratelimit.md)                |          |
| Options        | map[string]string       | 否   | [options](httproute-options.md)               | 尚未支持 |
| Backends       | array[Backend]          | 否   | [backend](httproute-backend.md)               |          |
| DefaultBackend | *struct[DefaultBackend] | 否   | [defaultbackend](httproute-defaultbackend.md) |          |



#### 示例

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-sample
spec:
  ingressClassName: nginx
  routes:
    - host: ingress.demo.com
      tls:
        secret: tls-secret
      protocol: http
      rules:
        - path: /user
          rateLimit:
            connections: 3
            rpm: 3
            rps: 5
          pathType: exact
          cors:
            corsMaxAge: 10000
          backends:
            - name: service
              port: 80
        - path: /app
          rateLimit:
            connections: 3
            rpm: 3
            rps: 5
          pathType: exact
          backends:
            - name: service2
              port: 8080

---
apiVersion: v1
kind: Secret
metadata:
  name: tls-secret
  namespace: default
data:
  tls.crt: string
  tls.key: string
type: kubernetes.io/tls

```