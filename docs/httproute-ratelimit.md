##### RateLimit

- RateLimit出现在HTTPRoute中的`Spec.Routes[0].Rules[0].RateLimit`位置，可以为某个访问路径独立配置Rate Limiting策略
- 如果同时配置三种限流策略，最先达到的限流策略会先行生效

| 字段           | 类型   | 必填 | 描述                                                         | 默认值 |
| -------------- | ------ | ---- | ------------------------------------------------------------ | ------ |
| Connections    | int    | 否   | 基于客户端IP限制一个IP的最大连接数<br>`limit_conn httproute_conn 3;` | 0      |
| RPM            | int    | 否   | 基于客户端IP限制一个IP的最大每秒请求次数<br>`limit_req zone=httproute_rps burst=25 nodelay;` | 0      |
| RPS            | int    | 否   | 基于客户端IP限制一个IP的最大每分钟请求次数<br>`limit_req zone=httproute_rpm burst=15 nodelay;` | 0      |
| LimitRate      | int    | 否   | 未开放此配置                                                 | 0      |
| LimitRateAfter | int    | 否   | 未开放此配置                                                 | 0      |
| Whitelist      | string | 否   | 未开放此配置                                                 | ""     |



##### 示例

- HTTPRoute

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-ratelimit
spec:
  ingressClassName: nginx
  routes:
    - host: ratelimit.demo.com
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
          backends:
            - name: service
              port: 80
        - path: /app
          rateLimit:
            connections: 3
          pathType: exact
          backends:
            - name: service2
              port: 8080
```

- nginx.conf

```nginx
http {
	
	limit_conn_zone $limit_xxx zone=default_httproute-ratelimit_238a801f-0658-4709-9f8f-3ac112ca3a59_conn:5m;
	limit_req_zone $limit_xxx zone=default_httproute-ratelimit_238a801f-0658-4709-9f8f-3ac112ca3a59_rpm:5m rate=3r/m;
	limit_req_zone $limit_xxx zone=default_httproute-ratelimit_238a801f-0658-4709-9f8f-3ac112ca3a59_rps:5m rate=5r/s;
	
	## start server ratelimit.demo.com
	server {
		server_name ratelimit.demo.com ;
		
		listen 80  ;
		listen [::]:80  ;
		listen 443  ssl http2 ;
		listen [::]:443  ssl http2 ;
		
		location = /user {
			limit_conn default_httproute-ratelimit_238a801f-0658-4709-9f8f-3ac112ca3a59_conn 3;
			limit_req zone=default_httproute-ratelimit_238a801f-0658-4709-9f8f-3ac112ca3a59_rps burst=25 nodelay;
			limit_req zone=default_httproute-ratelimit_238a801f-0658-4709-9f8f-3ac112ca3a59_rpm burst=15 nodelay;
			proxy_pass http://upstream_balancer;
		}
		
		location = /app {
			limit_conn default_httproute-ratelimit_238a801f-0658-4709-9f8f-3ac112ca3a59_conn 3;
			proxy_pass http://upstream_balancer;
		}
	}
	## end server ratelimit.demo.com
}
```

------

​																					  [跳转HTTRoute](httproute.md)
