package main

import (
	"bytes"
	"github.com/hktalent/scan4all/projectdiscovery/nuclei_Yaml"
	"log"
	_ "net/http/pprof"
	"os"
	"sync"
)

/*
1、✓ POC 不落地验证、测试：黑盒模式，最小成本，不修改nuclei代码对情况下
思路：
1）、漏洞扫描管控平台 开启PoC http get服务
安全设计：/PoCs/[为每个节点特有的，具有实效性的key]/具体的pocid.yaml
2）、设置nuclei调用参数：options.Templates 对应命令行中的 -t,格式如下
nuclei -duc -u http://192.168.10.31:8888 -t "http://127.0.0.1:8088/goSwaggerAPI.yaml,http://127.0.0.1:8088/checkGoDebug.yaml"
2、多实例测试


*/

func DoNuclei(buf *bytes.Buffer, wg *sync.WaitGroup, oOpts *map[string]interface{}) {
	defer wg.Done()
	xx := make(chan bool)
	go nuclei_Yaml.RunNuclei(buf, xx, oOpts)
	log.Println("xxx start")
	<-xx
}

/*
1、排除gov
2、排除蜜罐
*/
func main() {
	os.Setenv("enableNuclei", "true")
	if true {
		//go func() {
		//	//szTip = "Since you started http://127.0.0.1:6060/debug/pprof/ with -debug, close the program with: control + C"
		//	fmt.Println("debug info: \nopen http://127.0.0.1:6060/debug/pprof/\n\ngo tool pprof -seconds=10 -http=:9999 http://localhost:6060/debug/pprof/heap")
		//	http.ListenAndServe(":6060", nil)
		//}()
		buf := bytes.Buffer{}
		var wg sync.WaitGroup
		wg.Add(1)
		buf.WriteString("http://192.168.10.31:8888\n")
		pwd, _ := os.Getwd()
		m1 := map[string]interface{}{"UpdateTemplates": false, "Templates": []string{pwd + "/config/nuclei-templates"}, "TemplatesDirectory": pwd + "/config/nuclei-templates", "NoUpdateTemplates": true}
		go DoNuclei(&buf, &wg, &m1)
		wg.Add(1)
		go DoNuclei(&buf, &wg, &m1)
		wg.Add(1)
		go DoNuclei(&buf, &wg, &m1)
		wg.Wait()
	}
	//oUrl, err := url.Parse("173.82.115.38:80")
	//if nil == err {
	//	szK := oUrl.Scheme + "//" + oUrl.Hostname()
	//	log.Println(szK)
	//} else {
	//	log.Println(err)
	//}
	//args := []string{"key111key111key111key111", "dataxxxxxxdataxxxxxxdataxxxxxxdataxxxxxxdataxxxxxx"}
	//key := args[0]
	//value := args[1]
	//Content := []byte(value)
	//block, _ := aes.NewCipher([]byte(key))
	//blockSize := block.BlockSize()
	//n := blockSize - len(Content)%blockSize
	//temp := bytes.Repeat([]byte{byte(n)}, n)
	//Content = append(Content, temp...)
	//
	//iv := uuid.NewV4().Bytes()
	//fmt.Println(len(iv), iv)
	//blockMode := cipher.NewCBCEncrypter(block, iv)
	//cipherText := make([]byte, len(Content))
	//blockMode.CryptBlocks(cipherText, Content)
	//xx5 := append(iv[:], cipherText[:]...)

}
