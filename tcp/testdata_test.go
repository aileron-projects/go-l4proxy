package tcp

import (
	"os"
	"path/filepath"
	"testing"
)

// test data created by:
// openssl req -new -x509 -nodes -days 36500 -subj '/CN=test' -keyout key.pem -out cert.pem

var certPem = `
-----BEGIN CERTIFICATE-----
MIIDATCCAemgAwIBAgIUHqIQ6bKyKGuLkX3MeascRp7a2DQwDQYJKoZIhvcNAQEL
BQAwDzENMAsGA1UEAwwEdGVzdDAgFw0yNTA1MTYyMjE3NDRaGA8yMTI1MDQyMjIy
MTc0NFowDzENMAsGA1UEAwwEdGVzdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAMdY2WHLQO4OFCwylOmqZ+qgXRXbffIiUit2iVGbN7NnP/Dsw2zjnhyo
9qVIT6gyvpy8/WWAhNynlvHrqVWClj15CrmG1A87KFIY0btzXAQ0VlROcVa7zm8X
o7+4etEU1NiyXF1FQE7kdn/okdVm2TRthkPqmPaJIUNIQ8YoNsCpKwf6znguMztr
gmD4aIrEqExlY2YhJc81Cs0/uns6sabUgdab1chMuSRAAmj8gj7xtNk4v17Zkqgs
PYyl6nFTSzo+eZvJAFsX3uhN5lSxtq2eIm9CPpafOpN+9A973eGFjyQFMB1xYFZ8
sTMPYZIlbAd8IjH4JKyT5PE+FAoeJ38CAwEAAaNTMFEwHQYDVR0OBBYEFE0flNIK
MkxkrVUh/1hn/U8quULeMB8GA1UdIwQYMBaAFE0flNIKMkxkrVUh/1hn/U8quULe
MA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBABO7xImrHi+jyze8
9vdhzl1qwE7obwTtznDH57iUY9sZIT0H5OpLYWr4ZkpI1GSmI0sZcYCBnJ/P/pbS
5b3PHnNVtuESjrDMh8ZwtkYyeiocFgi79bkoYNs8Lk5Fd/XfGtpPdB7V1AcavFLK
WeWndsUi1YvRco/BjB1oWk5SqudPWn43vW4Iyd2WAXjwtV1FrrRxPgBFL+w2Reow
0MHT3w6cOgP8Q5Rpm1Di0RP4CmxZUSR3i2S/98UZbpxv3hRsGFrKUiBegcnWQGKa
3owGH8FMFwZkkcOpxX+5eHSMDWammzTjsOepQg6Krzi7ZjIx5tzGCkEh7lF5QcyV
QGufxek=
-----END CERTIFICATE-----
`

var keyPem = `
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDHWNlhy0DuDhQs
MpTpqmfqoF0V233yIlIrdolRmzezZz/w7MNs454cqPalSE+oMr6cvP1lgITcp5bx
66lVgpY9eQq5htQPOyhSGNG7c1wENFZUTnFWu85vF6O/uHrRFNTYslxdRUBO5HZ/
6JHVZtk0bYZD6pj2iSFDSEPGKDbAqSsH+s54LjM7a4Jg+GiKxKhMZWNmISXPNQrN
P7p7OrGm1IHWm9XITLkkQAJo/II+8bTZOL9e2ZKoLD2MpepxU0s6PnmbyQBbF97o
TeZUsbatniJvQj6WnzqTfvQPe93hhY8kBTAdcWBWfLEzD2GSJWwHfCIx+CSsk+Tx
PhQKHid/AgMBAAECggEAdNScqlTd2+bCadjHL17ggkCz7WBzItp8f957wikO7wTf
E4sfSYzpGHDqBJopnTEIR4H3nGpC87MAow3zcxsShDqRT4Q2eiq4iczj9lE4p3oo
8kX65NMyvCcFoBI1YJ4t1EZMHZm9g78ft/0Mox4gxDpAS0ONnEdoCC+tENB2rlgs
MuHqLtsLsTQLjkeQZ96wKaV+xPibyfJP8hJzUgizFhCyZCXBqNzjYcJ3CyYQH7EC
zn2/jgE1IMvPHiFON0Teaxys45dRuQ9g8aINAbpzjL0sxITXXdRAbZfgv1EULtyC
rSvTYvhOpkkiNMddx4thXKh2JK+N1aBPcWH0KUKQoQKBgQD+BSYCHyI6RYUs0okj
CuomLXS2ktBNfnFBVYMGQl6KZZ8ygbwDU7HN6wbVoxsJZK/I8lt+M6sIyqK1M0B8
+gp8i7iSYuFm1nGN39sWI1XyUV/AAo54tmt/VL0uK0HQuhTfYCUX2beizejILWcP
osMkSaESyYnsStGRktGo7O31LwKBgQDI5pxFYz4roIPMhQk7F6F0a/2z3gtKIAdk
UMbmI/it3xRQiESX1O6oV2/xgLWnY9S+ZwfRT1QHY6cju7JvFO1IwompjhHrxr8w
0FXDMM8STQCfStpPtKeiyQXBOzbV4XjSMv4eC2hPA/T79U+TvzbRVfYTPv0rgPGz
ZSHzxh3+sQKBgQCcWWrlj2gv2a54wjVqfkNUsFHQNllD+XmYLxvwdFVgdeAg5D1n
BtK/3dNdF4GCrJiN5K5v5Tl2tdKkGSGuSvC9a/p1D6uQ8pj+LDTbUXCLL47xziEH
J7DOtMekbTebU42ZV8I9BeLDsF2BKIDw3MujwDYGLEpaSErpgSbPBNYKPQKBgDgh
rtH4S5uW6pxPI2geTx8fOTBzPsZEMqyFoT0WGdpCtQNXP4sIlHH2lDFu29JcI0nb
saR6vKif4EgsnktJFo84w4vbRQiSbELwQrYZwhGD8EORf3O7rwbdkb0OlWCm8cIR
O00btNY27dEkFkrapy9QDEQOWmA4U8/i1tysox7RAoGAWMTENg97usCg27BOlN7V
TGAj/98/uASCTxNUVh2vWUgebOkJcXfHHT/eDWFSLvdDmkF03CxB5/49SOHTGv4C
HGWDrEvA2nj7RAqfiRxFJKJhijWvhN+uB1LTGyhhnpIX1oKwqD5GZZFgF+vnu5kh
9MUFYSOtwGKSSJ51uTQs/cE=
-----END PRIVATE KEY-----
`

func createCertKey(t *testing.T) (cert, key string) {
	dir := t.TempDir()
	cert = filepath.Join(dir, "cert.pem")
	key = filepath.Join(dir, "key.pem")
	if err := os.WriteFile(cert, []byte(certPem), os.ModePerm); err != nil {
		t.Error(err)
	}
	if err := os.WriteFile(key, []byte(keyPem), os.ModePerm); err != nil {
		t.Error(err)
	}
	return cert, key
}
