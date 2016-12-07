package

import ("crypto/tls")

type originInfo struct {
  Hostname string,
  Port int
}

func getOrigin() originInfo {
  oOriginInfo := originInfo{}
  return oOriginInfo
}
