##### Cors

- HTTPRoute跨域配置
- Cors出现在HTTPRoute的两个位置，分别是 `Spec.Routes[0].Cors` 和 `Spec.Routes[0].Rules[0].Cors`

| 字段                 | 类型   | 必填 | 描述                                                         | 默认值                                                       |
| -------------------- | ------ | ---- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| CorsAllowOrigin      | string | 否   | `Access-Control-Allow-Origin: *`                             | "*"                                                          |
| CorsAllowMethods     | string | 否   | `Access-Control-Allow-Methods: Options`                      | "GET, PUT, POST, DELETE, PATCH, OPTIONS"                     |
| CorsAllowHeaders     | string | 否   | `Access-Control-Allow-Headers: DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization` | "DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization" |
| CrosAllowCredentials | bool   | 否   | `Access-Control-Allow-Credentials: true`                     | true                                                         |
| CorsExposeHeaders    | string | 否   | `Access-Control-Expose-Headers: `                            | ""                                                           |
| CorsMaxAge           | int    | 否   | `Access-Control-Max-Age: 1700`                               | 1728000                                                      |



##### 生效规则

- Ingress Controller将会采用就近原则，即优先`Spec.Routes[0].Rules[0].Cors`，然后`Spec.Routes[0].Cors`
- 当未配置Cors时，该功能是直接关闭的；当配置了一项Cors参数时，其他参数将使用默认值
- Ingress Controller会根据就近原则对Proxy中的所有字段进行一一合并，在配置到Nginx中

```
+-----------------+               +-----------------------+               +--------------------------------+          +-----------+
|                 |-------------->|                       |-------------->|                                |--------->|           |
|  defualt Cors   |  Inheritance  |  Spec.Routes[0].Cors  |  Inheritance  |  Spec.Routes[0].Rules[0].Cors  |  config  |   Nginx   |  
|                 |-------------->|                       |-------------->|                                |--------->|           |
+-----------------+               +-----------------------+               +--------------------------------+          +-----------+
```



##### 示例

- HTTPRoute

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-cors
spec:
  ingressClassName: nginx
  routes:
    - host: cors.demo.com
      protocol: http
      # 未配置的采用默认值
      cors:
        corsAllowMethods: "Options"
      rules:
        - path: /user
          pathType: exact
          cors:
            corsMaxAge: 1700
          backends:
            - name: service
              port: 80
```

- nginx.conf

```nginx
	## start server cors.demo.com
	server {
		server_name cors.demo.com ;
		
		listen 80  ;
		listen [::]:80  ;
		listen 443  ssl http2 ;
		listen [::]:443  ssl http2 ;
		
		location = /user {
			# Cors Preflight methods needs additional options and different Return Code
			if ($request_method = 'OPTIONS') {
				more_set_headers 'Access-Control-Allow-Origin: *';
				
				more_set_headers 'Access-Control-Allow-Methods: Options';
				more_set_headers 'Access-Control-Allow-Headers: DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization';
				
				more_set_headers 'Access-Control-Max-Age: 1700';
				more_set_headers 'Content-Type: text/plain charset=UTF-8';
				more_set_headers 'Content-Length: 0';
				return 204;
			}
			
			more_set_headers 'Access-Control-Allow-Origin: *';
			proxy_pass http://upstream_balancer;
		}
	}
```

