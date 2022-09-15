package agent

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/dollarkillerx/go-proxy-manager/internal/utils"
	"github.com/dollarkillerx/go-proxy-manager/proto/agent"
	"github.com/dollarkillerx/go-proxy-manager/proto/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	token = ""
)

func init() {
	token = os.Getenv("Token")
	if token == "" {
		token = "f32e4a2f-0aec-4e5b-b4ab-e88ab8a68353"
	}
}

func TestAgent(t *testing.T) {
	dial, err := grpc.Dial(":8501",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(utils.RPCAuth(token, nil)))
	if err != nil {
		log.Fatalln(err)
	}
	defer dial.Close()

	client := agent.NewAgentClient(dial)
	client.AddTask(context.TODO(), &agent.AddTaskReq{
		Domain:    "d1.baidu.com",
		ProxyType: common.ProxyType_ReverseProxyEnum,
		ReverseProxy: &common.ReverseProxy{
			Target:           "https://www.baidu.com",
			PathSubstitution: true,
		},
		EnableSsl: true,
		Certificate: &common.Certificate{
			PublicKey:  PublicKey,
			PrivateKey: PrivateKey,
		},
	})

}

var PrivateKey = `
-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDobjv/6Qb9+coL
xjGRzqjgxjzZZdmnk1ecY7xKii4i8xGKN6oObj5TcBCf1zEp+biPmWfo0grf9FW8
UQex0JPP19Qkeqy03i8w1ZzxWJ235e+8LjDZl8dqtkhwrqHBefKs6lcAZv5RPIEU
tfS0nM9lXbPTqfG6d5/z1/EMyvPsOOYESRKEDcKm5HqGJjKBUWzgkEwRj9vcoJWq
ju7ocKtyAGss8UCD0M9508x8yZWuDi44kERWWoFdusOwqVHd0CrPH0t87vjXNGrC
j/1hex4zbFao//vC2c55KFk34jzGuvWjmJ/vzFbyhhEx81Y7q1RrAIpJKKx2qwj6
ymcQDWTpAgMBAAECggEBAIO4FOCOGpxLtqi9hNHwn9vh/BHnmlVRNtE15DsJ2+OK
hGIj8YeECmYWmBU9SmtTTbhzY2OZxPft67hw4z9tyJZ85YhWzREi6PNv2yUa5yaU
dCvcsWpejXB4JM1W6exPQ6uikfN0aqN2+l/WTfsc8purYoGNSaQ2rrrtoXlyepFf
CKVdy+aHO9jUfMkZr/wCbJzvdtXSOBQOC/qQLoEv/zzWIkKwYFoAMw50ps8wPSjg
EuFVcWjoxah0t/8UmuLHBbh2vjQFx9gyrDJzXveP5NSAyQ60MBfRFae+yjw+vuZd
5kVO8UzG7bcBGwpWoJ34f9CyimeHnCM9WK4ZukvILJECgYEA/i30YvSZHqSxhUZo
IENJAMpk/V6tHwi8ZIAG45bqePf2rTSPnJr/OGERsEYWUmx/HxvNsj4+9kH7fApr
JDpwA3NGuTxRcXdL7hWhmto4uKGWNBFIucC/wB+B/Koni3c5J982awDxhFs/JaiV
KEmkjZSKdBIfvisCP4sHGRZrkh8CgYEA6hhnCiFJVlbzekUTAtiGsdi2S9KO0ywh
859nwDRkQQBGBO7oAIhl3JX2Nhqxvm/s4tKzf3CRqz+EtxgbwXPUVgNWuCvozKzl
yES25suk6A/ovCsoqseBKartWBf+KmVyxeMeCEwETOJVmfv/gbP7tVJZxvXsz3N9
XM8rtu3Jd/cCgYBYR8m7nMyZsgXeDigYHIZ8ec89mOZ0Auq71SOPZFknjqfkXH8M
m2DNShmEqbMTCY/VVCjLIYdorF4WJOA1gv5olF8Z8vMuf4qltGmeiPBuX+D1UJN7
wJBYwBi7krWNYOk/Ce2ymG9J0w9JmxXDGmDGs7Kqwai4ueNcAoOCmSDfbQKBgQCf
A0ZyvhuX12Z6BpOcWkJcai2iIXu8+/xwBTave6ch9DKbxUC4o+41QXRGWimkh1is
1pzxEyEP5wyaZnjsFmcEqi3s2n6/ES7gs7Rd37S7oZrgvpxYLT6SdDXWoi2W1OR6
gQT0c9Zz90ZYW7G6g+yVxUOQf5qFbBiaw2sZTjqNDwKBgBlCKU4iQNTNpJFc0pyQ
qFRzFBtBWeDR7o5BaSWVC1EqP3dI+5uXxq56Ty7HoBqjwRO84h/dW4XXgORZMD0d
j4SEqcS/FyhjOdSwV9oiIZSF7ObkfgBQpizZheE86LB4W3DjuNR4kjwzAMfUrVzU
WrHvf68ZGOjoIkYBI9M+y4tx
-----END PRIVATE KEY-----
`

var PublicKey = `
-----BEGIN CERTIFICATE-----
MIIDlzCCAn+gAwIBAgIUaKguHsbsqklQjSlCqSPcfATIPVwwDQYJKoZIhvcNAQEL
BQAwWzELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDEUMBIGA1UEAwwLKi5iYWlkdS5jb20w
HhcNMjIwOTE1MDM1NjE5WhcNMjMwOTE1MDM1NjE5WjBbMQswCQYDVQQGEwJBVTET
MBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0cyBQ
dHkgTHRkMRQwEgYDVQQDDAsqLmJhaWR1LmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBAOhuO//pBv35ygvGMZHOqODGPNll2aeTV5xjvEqKLiLzEYo3
qg5uPlNwEJ/XMSn5uI+ZZ+jSCt/0VbxRB7HQk8/X1CR6rLTeLzDVnPFYnbfl77wu
MNmXx2q2SHCuocF58qzqVwBm/lE8gRS19LScz2Vds9Op8bp3n/PX8QzK8+w45gRJ
EoQNwqbkeoYmMoFRbOCQTBGP29yglaqO7uhwq3IAayzxQIPQz3nTzHzJla4OLjiQ
RFZagV26w7CpUd3QKs8fS3zu+Nc0asKP/WF7HjNsVqj/+8LZznkoWTfiPMa69aOY
n+/MVvKGETHzVjurVGsAikkorHarCPrKZxANZOkCAwEAAaNTMFEwHQYDVR0OBBYE
FAQRq3B373RHbKqJpq8TpCr69IvMMB8GA1UdIwQYMBaAFAQRq3B373RHbKqJpq8T
pCr69IvMMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBADo9SQBW
MFn0JwNKI0v9BtxJwoSs0OQqmw25dB7DBHXN5F3gkGbRwwVVlLEU522EsK1z+ecG
hyzRTHd2BGXs1DLOo1IWPSiWy+4C7MddTc7heEBVq03wfnt7kFEEFQpcb0MXlTdh
FISvLJcbQ1/ZJzv0N0AMBKgJtiE2qH+17PNwpybVqs3Rlcy1chMeTO156gIS5ew3
yTg+IEJjtI0Vymbs0EI65B8I5Ef0cplRaSUYSQRIp8qQuO6/qrSOMhNUmCikYisr
/lUeOSG7smwrPxk0BeNMfAt4lGXF6EJxshKFxXauI37C+C5PSwjqKE+3kcSOP7uX
hx4jq15oQrDdLgg=
-----END CERTIFICATE-----
`
