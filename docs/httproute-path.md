##### path

- HTTPRoute

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-path
spec:
  ingressClassName: nginx
  routes:
    - host: path.demo.com
      tls:
        secret: tls-secret
      protocol: http
      rules:
        - path: /exact
          pathType: exact
          backends:
            - name: service
              port: 80
        - path: /user
          pathType: exact
          rewrite: /something
          backends:
            - name: service
              port: 80
        - path: /app
          pathType: prefix
          backends:
            - name: service
              port: 80
        - path: /something(/|$)(.*)
          pathType: regex
          rewrite: "/$2"
          backends:
            - name: service
              port: 80
    - host: path2.demo.com
      protocol: http
      rules:
        - path: /exact
          pathType: exact
          backends:
            - name: service
              port: 80
```

- nginx.conf

```nginx
	## start server path.demo.com
	server {
		server_name path.demo.com ;
		
		listen 80  ;
		listen [::]:80  ;
		listen 443  ssl http2 ;
		listen [::]:443  ssl http2 ;
		
		location ~* "^/something(/|$)(.*)" {
			rewrite "(?i)/something(/|$)(.*)" /$2 break;
			proxy_pass http://upstream_balancer;
		}
		
		location ~* "^/exact" {
			proxy_pass http://upstream_balancer;
		}
		
		location ~* "^/user" {
			rewrite "(?i)/user" /something break;
			proxy_pass http://upstream_balancer;
		}
		
		location ~* "^/app/" {
			rewrite "(?i)/app/" /app break;
			proxy_pass http://upstream_balancer;
		}
		
		location ~* "^/app" {
			proxy_pass http://upstream_balancer;
		}
	}
	## end server path.demo.com
	
	## start server path2.demo.com
	server {
		server_name path2.demo.com ;
		
		listen 80  ;
		listen [::]:80  ;
		listen 443  ssl http2 ;
		listen [::]:443  ssl http2 ;
		
		location = /exact {
			proxy_pass http://upstream_balancer;
		}
	}
```

##### 解析

- 根据上述示例，解析如下

| 路径                 | 类型   | 重写       | 描述                                                         | 备注                                             |
| -------------------- | ------ | ---------- | ------------------------------------------------------------ | ------------------------------------------------ |
| /exact               | exact  |            | `location ~* "^/exact"`                                      | Rules中存在一个Rewrite，所有匹配类型变成正则匹配 |
| /user                | exact  | /something | `location ~* "^/user"` <br>`rewrite "(?i)/user" /something break;` | 变成正则匹配，重写代理路径                       |
| /app                 | prefix |            | `location ~* "^/app/"`  `rewrite "(?i)/app/" /app break;`<br>`location ~* "^/app" ` | 前缀匹配                                         |
| /something(/\|$)(.*) | regex  | /$2        | `location ~* /something(/$)(.*)`<br>`rewrite "(?i)/something(/$)(.*)" /$2 break;` | 正则匹配，按照`()`分组，`$2`代表`(.*)`部分       |
| /exact               | exact  |            | `location = /exact`                                          | Rules没有Rewrite，生成的绝对匹配                 |

------

​																					  [跳转HTTRoute](httproute.md)
