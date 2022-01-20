#####   Proxy

- 当Nginx接收到客户端发来的请求后，对请求进行解析、路由规则匹配来决定转发到哪个后端服务，在浏览器请求视角看来Nginx是一个代理服务，既然作为一个代理服务，自然就有很多的可配置的代理选项，因此我们在HTTPRoute结构体中增加可选配置字段
- Proxy出现在HTTPRoute的两个位置，分别是 `Spec.Routes[0].Proxy` 和 `Spec.Routes[0].Rules[0].Proxy`，可以为不同的路径分别配置代理

| 字段                 | 类型   | 必填 | 描述                                                    | 默认值        |
| :------------------- | ------ | ---- | ------------------------------------------------------- | :------------ |
| BodySize             | string | 否   | 最大请求体大小设置<br>`client_max_body_size 2m;`        | 1m            |
| ConnectTimeout       | int    | 否   | upstream连接超时时间设置<br>`proxy_connect_timeout 6s;` | 5             |
| SendTimeout          | int    | 否   | upstream发送超时时间设置<br>`proxy_send_timeout 80s;`   | 60            |
| ReadTimeout          | int    | 否   | updatream读取超时时间设置<br>`proxy_read_timeout 80s;`  | 60            |
| BuffersNumber        | int    | 否   | `proxy_buffers 15 4k;`                                  | 4             |
| BufferSize           | string | 否   | `proxy_buffer_size 4k;`                                 | 4k            |
| CookieDomain         | string | 否   | `proxy_cookie_domain off;`                              | off           |
| CookiePath           | string | 否   | `proxy_cookie_path off;`                                | off           |
| NextUpstream         | string | 否   | `proxy_next_upstream error timeout;`                    | error timeout |
| NextUpstreamTimeout  | int    | 否   | `proxy_next_upstream_timeout 0;`                        | 0             |
| NextUpstreamTries    | int    | 否   | `proxy_next_upstream_tries 3;`                          | 3             |
| ProxyRedirectFrom    | string | 否   | `proxy_redirect off;`                                   | off           |
| ProxyRedirectTo      | string | 否   | `proxy_redirect ProxyRedirectFrom ProxyRedirectTo;`     | off           |
| RequestBufering      | string | 否   | `proxy_request_buffering on;`                           | on            |
| ProxyBuffering       | string | 否   | `proxy_buffering off;`                                  | off           |
| ProxyHTTPVersion     | string | 否   | `proxy_http_version 1.1;`                               | 1.1           |
| ProxyMaxTempFileSize | string | 否   | `proxy_max_temp_file_size 1024m;`                       | 1024m         |



##### Proxy生效规则

- Ingress Controller将会采用就近原则，即优先`Spec.Routes[0].Rules[0].Proxy`，然后`Spec.Routes[0].Proxy`，若不设置则使用默认配置
- Ingress Controller会根据就近原则对Proxy中的所有字段进行一一合并，在配置到Nginx中

```
+-----------------+               +-----------------------+               +--------------------------------+          +-----------+
|                 |-------------->|                       |-------------->|                                |--------->|           |
|  defualt Proxy  |  Inheritance  |  Spec.Routes[0].Proxy |  Inheritance  |  Spec.Routes[0].Rules[0].Proxy |  config  |   Nginx   |  
|                 |-------------->|                       |-------------->|                                |--------->|           |
+-----------------+               +-----------------------+               +--------------------------------+          +-----------+
```



##### 示例

- HTTPRoute

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-proxy
spec:
  ingressClassName: nginx
  routes:
    - host: proxy.demo.com
      protocol: http
      # 未配置的采用默认值
      proxy:
        bodySize: "2m"     # 未被重写该配置生效
        connectTimeout: 6
        sendTimeout: 69
        readTimeout: 69
      rules:
        - path: /user
          pathType: exact
          proxy:
            sendTimeout: 80  # 重写上层sendTimeout
            readTimeout: 80
            buffersNumber: 15
            cookieDomain: "off"
          backends:
            - name: service
              port: 80
```

- nginx.conf

```nginx
	## start server proxy.demo.com
	server {
		server_name proxy.demo.com ;
		
		listen 80  ;
		listen [::]:80  ;
		listen 443  ssl http2 ;
		listen [::]:443  ssl http2 ;
		
		location = /user {
			
			client_max_body_size                    2m;

			# Custom headers to proxied server
			proxy_connect_timeout                   6s;
			proxy_send_timeout                      80s;
			proxy_read_timeout                      80s;
			
			proxy_buffering                         off;
			proxy_buffer_size                       4k;
			proxy_buffers                           15 4k;
			
			proxy_max_temp_file_size                1024m;
			
			proxy_request_buffering                 on;
			proxy_http_version                      1.1;
			
			proxy_cookie_domain                     off;
			proxy_cookie_path                       off;
			
			# In case of errors try the next upstream server before returning an error
			proxy_next_upstream                     error timeout;
			proxy_next_upstream_timeout             0;
			proxy_next_upstream_tries               3;
			
			proxy_pass http://upstream_balancer;
			
			proxy_redirect                          off;
		}
	}
```

------

​																					  [跳转HTTRoute](httproute.md)

