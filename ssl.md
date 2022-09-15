## 自签证书

#### 生成CA:

```
# 生成CA认证机构的证书密钥key
# 需要设置密码，输入两次
openssl> genrsa -des3 -out ca.key 1024

# 去除密钥里的密码(可选)
# 这里需要再输入一次原来设的密码
openssl> rsa -in ca.key -out ca.key

# 用私钥ca.key生成CA认证机构的证书ca.crt
# 其实就是相当于用私钥生成公钥，再把公钥包装成证书
openssl> req -new -x509 -key ca.key -out ca.crt -days 365
# 这个证书ca.crt有的又称为"根证书",因为可以用来认证其他证书
```

#### 生成网站的证书

```
# 生成自己网站的密钥server.key
openssl> genrsa -des3 -out server.key 1024

# 生成自己网站证书的请求文件
# 如果找外面的CA机构认证，也是发个请求文件给他们
# 这个私钥就包含在请求文件中了，认证机构要用它来生成网站的公钥，然后包装成一个证书
openssl> req -new -key server.key -out server.csr

# 使用虚拟的CA认证机构的证书ca.crt，来对自己网站的证书请求文件server.csr进行处理，生成签名后的证书server.crt
# 注意设置序列号和有效期（一般都设1年）
openssl> x509 -req -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt -days 365
```

```
格式	说明
.crt .cer	证书(Certificate)
.key	密钥/私钥(Private Key)
.csr	证书认证签名请求(Certificate signing request)
*.pem	base64编码文本储存格式，可以单独放证书或密钥，也可以同时放两个；base64编码就是两条-------之间的那些莫名其妙的字符
*.der	证书的二进制储存格式(不常用)
```

### iBM

``` 
openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out certificate.pem

openssl x509 -text -noout -in certificate.pem

openssl pkcs12 -inkey key.pem -in certificate.pem -export -out certificate.p12
 
openssl pkcs12 -in certificate.p12 -noout -info
```