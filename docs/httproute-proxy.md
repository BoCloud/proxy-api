#####   Proxy

- 我们知道Nginx代理用户请求转发到真实的后端服务，该文件配置描述的便是如何对代理进行配置

- 在HTTPRoute中Proxy配置分别出现在两个位置 `Spec.Routes[0].Route.Proxy` 和 `Spec.Routes[0].Route.Rules[0].Proxy`，如果在一个HTTRoute中配置都配置了Proxy生效规则如下

- Ingress Controller将会采用就近原则，即优先`Spec.Routes[0].Route.Rules[0].Proxy`，其次`Spec.Routes[0].Route.Proxy`，若不设置则使用默认配置，Ingress Controller采用了更细致的处理方法，根据处理优先级会对Proxy中的所有字段进行合并，然后再应用到Nginx配置



| 字段                 | 类型   | 必填 | 描述 | 示例 |
| -------------------- | ------ | ---- | ---- | ---- |
| BodySize             | string | 否   |      |      |
| ConnectTimeout       | int    | 否   |      |      |
| SendTimeout          | int    | 否   |      |      |
| ReadTimeout          | int    | 否   |      |      |
| BuffersNumber        | int    | 否   |      |      |
| BufferSize           | string | 否   |      |      |
| CookieDomain         | string | 否   |      |      |
| CookiePath           | string | 否   |      |      |
| NextUpstream         | string | 否   |      |      |
| NextUpstreamTimeout  | int    | 否   |      |      |
| NextUpstreamTries    | int    | 否   |      |      |
| ProxyRedirectFrom    | string | 否   |      |      |
| ProxyRedirectTo      | string | 否   |      |      |
| RequestBufering      | string | 否   |      |      |
| ProxyBuffering       | string | 否   |      |      |
| ProxyHTTPVersion     | string | 否   |      |      |
| ProxyMaxTempFileSize | string | 否   |      |      |



#### Proxy默认值

