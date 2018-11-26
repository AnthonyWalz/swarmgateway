package main
import (
        "net/http"
        "os/exec"
        "strings"
        "net/http/httputil"
        "net/url"
)
func main() {
        http.HandleFunc("/", serve)
        http.ListenAndServe(":80", nil)
}
func serve(w http.ResponseWriter, r *http.Request) {
        temp := exec.Command("dig", r.Host, "cname", "+short")
        e, err := temp.Output()
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        domain := string(e)
        key := ""
        if domain == ""{
                parts := strings.Split(string(r.Host),".")
                key = parts[0] + parts[1]
        }else{
                if last := len(domain) - 1; last >= 0 && domain[last] == '.' {
                        domain = domain[:last]
                }
                parts := strings.Split(domain,".")
                key = parts[0] + parts[1]
        }
        target :=  "https:/" + "/swarm-gateways.net/bzz:/" + key
        serveReverseProxy(target,w,r)
}
func serveReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
        url, _ := url.Parse(target)
        proxy := httputil.NewSingleHostReverseProxy(url)
        r.URL.Host = url.Host
        r.URL.Scheme = url.Scheme
        r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
        r.Host = url.Host
        proxy.ServeHTTP(w, r)
}