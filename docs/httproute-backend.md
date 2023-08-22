##### Backend

- 一个访问路径可以配置多个发布服务，HTTPRoute支持设置基于权重的灰度发布和基于请求的灰度发布，关于灰度发布详情，请[参考](httproute-backend.md#灰度发布)
- 一个k8s服务对应多个endpoints，即多个pod。HTTPRoute支持为每一个独立的服务配置会话保持策略，开启了会话保持的服务会根据请求的来源活请求cookie分辨客户端，并始终将同一客户端请求发送到固定的一个后端pod。关于会话保持详情请[参考](httproute-backend.md#会话保持)
- Backend出现在HTTPRoute的`Spec.Routes[0].Rules[0].Backends[0]`位置

| 字段                  | 类型                                                        | 必填 | 描述                                                         | 示例     |
| --------------------- | ----------------------------------------------------------- | ---- | ------------------------------------------------------------ | -------- |
| Name                  | string                                                      | 是   | Service名称                                                  | echo     |
| Port                  | *int32                                                      | 是   | Service端口                                                  | 80       |
| Weight                | *int32                                                      | 否   | 灰度权重                                                     | 25       |
| Matches               | [[HTTPMatch](httproute-backend.md#httpmatch)]               | 否   | 匹配规则                                                     |          |
| Strategy              | string                                                      | 否   | 会话保持策略，在同一个service中生效<br>支持`ip`、 `cookie`、 `ewma`、`round_robin`<br>默认为round_robin | ip       |
| ChangeCookieOnFailure | bool                                                        | 否   | 当会话保持策略为`cookie`<br>该配置项表示cookie失效后，更换其它pod endpoint | true     |
| FailTimeOutSeconds    | *int                                                        | 否   |                                                              | 尚未支持 |
| MaxFails              | *int                                                        | 否   |                                                              | 尚未支持 |
| MaxConns              | *int                                                        | 否   |                                                              | 尚未支持 |
| Keepalive             | *int                                                        | 否   |                                                              | 尚未支持 |
| Options               | [Options](httproute-options.md#specroutes0backends0options) | 否   | 详细信息见Options文档                                        |          |



###### HTTPMatch

- HTTPMatch出现在HTTPRoute的 `Spec.Routes[0].Rules[0].Backends[0].HTTPMatch`位置

| 字段     | 类型   | 必填 | 描述                                                         | 示例   |
| -------- | ------ | ---- | ------------------------------------------------------------ | ------ |
| Type     | string | 是   | 支持两种类型 `header`和`cookie`，顾名思义`header`类型需要请求header中携带对应的`key:value` | header |
| GroupId  | int    | 是   | 当配置多组HTTPMatch时，其中GroupId相同的为一组，同组内为`与`关系，不同组间为`或`关系<br>当nginx进行路由转发时，请求需要满足同组内所有条件才会选择该service | 2      |
| Key      | string | 是   | 请求中需要在`Type(header/cookie)`中携带的`Key`               | name   |
| Operator | stirng | 是   | 支持`exact: (key eq value)`和 `regex: (key regex value)`两种匹配方式 | exact  |
| Value    | strng  | 是   | 解析请求，根据`Type(header/cookie)` 和 `Key`获取该`Value`值  | test   |



##### 灰度发布

- 当backends中存在>1个service时，该HTTPRoute要么是基于权重的灰度发布，要么是基于请求的灰度发布，假设weight和HTTPMatch均不配置则Ingress Controller会拒绝该HTTPRoute的创建

###### 基于权重的灰度发布策略

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-weight
spec:
  ingressClassName: nginx
  routes:
  - host: weight.demo.com
    rules:
    - path: /weight
      pathType: prefix
      backends:
      - name: echo
        port: 8080
        weight: 20
      - name: echo-v1
        port: 8080
        weight: 40
      - name: echo-v2
        port: 8080
        weight: 80
      - name: echo-v3
        port: 8080
```

- 解析

```
1. 当请求访问/weight时，nginx根据backends weight比例转发请求，比如流量到达echo-v1的比率为 40/(20+40+80)
2. 在backends中存在一个weight不为0，则灰度发布策略便是基于权重
3. 如果存在一个service的weight为0或者未设置，则流量永远不会到达该服务
4. 至少存在一个service的weight>0，这个在HTTPRoute创建阶段便进行了检查，如果不满足则创建不成功
```

###### 基于请求的灰度发布策略

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-header
spec:
  ingressClassName: nginx
  routes:
  - host: header.demo.com
    rules:
    - path: /header
      pathType: prefix
      backends:
      - name: echo
        port: 80
      - name: echo-v1
        port: 80
        matches:
          - groupId: 1
            type: header
            key: v1
            operator: exact
            value: "true"
      - name: echo-v2
        port: 80
        matches:
          - groupId: 1
            type: header
            key: user-id
            operator: regex
            value: "^[0-9]+3$"
          - groupId: 1
            type: cookie
            key: gender
            operator: exact
            value: male
          - groupId: 2
            type: header
            key: canary
            operator: exact
            value: "true"
```

- 解析

```shell
1. 默认灰度发布服务为基于请求特征的灰度服务
2. 在基于请求特征的灰度发布服务中，必须存在一个没有任何匹配条件的service作为主服务，这样当所有匹配条件都不满足时，则流量会转发到该服务
3. matches中匹配条件可以写多组，请求必须满足所有匹配条件才会选择此服务
4. 如果一个请求被多个service所匹配，则第一个匹配到的serivce会被选择，按照yaml中自上而下的顺序
5. 上述示例中主服务为echo，所有匹配条件都不满足则流量进入此服务
6. 当请求满足(header_v1:true)时，流量进入echo-v1服务
7. echo-v2中匹配条件分成两组，请求满足(header_canary:true)or(cookie_gender:male and header_user-id:regex:^[0-9]+3$)则流量进入此服务
```

###### 会话保持

- 基于sticky cookie的会话保持，负载均衡器会在某个服务第一次响应客户端请求时，为客户端设置一个sticky cookie，客户端下次请求这个服务时会带上这个cookie，负载均衡器根据sticky cookie的内容从服务的pod组里取出与第一次响应相同的pod，达到会话保持的效果
- 当服务里的pod出现错误时，负载均衡客户会根据sticky cookie仍然将请求发送到这个pod，为避免此种情形，可选择开启`changeCookieOnFailure: true`，当固定后端pod失败时，为该客户端重设一个新的sticky cookie，后续该客户端的请求将发送到其它pod。

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-session
spec:
  ingressClassName: nginx
  routes:
  - host: session.demo.com
    rules:
    - path: /session
      pathType: prefix
      backends:
      - name: echo
        port: 8080
        strategy: ip
    - path: /cookie
      pathType: exact
      backends:
        - name: echo-v1
          port: 80
          strategy: cookie
          changeCookieOnFailure: true
```

- 解析

```shell
1. 支持四种会话保持策略设置 ip、cookie、ewma、round_robin，默认为round_robin
2. 会话保持策略在一个service内生效，即如果灰度发布情况优先选择一个灰度服务，然后再服务内多个pod ip实现会话亲和性
3. 当基于cookie会话保持时，支持配置changeCookieOnFailure: true，即cookie失效后更换其他pod ip
```

------

​																					  [跳转HTTRoute](httproute.md)

