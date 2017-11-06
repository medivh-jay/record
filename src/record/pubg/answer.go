package pubg

import (
	"github.com/robertkrimen/otto"
	"record/log"
	"regexp"
	"strings"
)

func getJsAnswer(html string) string {

	log.Info("问答开始")

	reg := regexp.MustCompile(`(?i:setTimeout\(function\(\){)([\w\W]*)}, 4000\);`)
	result := reg.Find([]byte(html))
	javascript := string(result)

	// 替换所有无用代码
	javascript = strings.Replace(javascript, "<script type=\"text/javascript\">", "", -1)
	javascript = strings.Replace(javascript, "//<![CDATA[", "", -1)
	javascript = strings.Replace(javascript, "//]]>", "", -1)
	javascript = strings.Replace(javascript, "</script>", "", -1)
	javascript = strings.Replace(javascript, "t.length", "15", -1)
	javascript = strings.Replace(javascript, "setTimeout(function(){", "", -1)
	javascript = strings.Replace(javascript, "}, 4000);", "", -1)
	javascript = strings.Replace(javascript, "t = document.createElement('div');", "", -1)
	javascript = strings.Replace(javascript, `t.innerHTML="<a href='/'>x</a>";`, "", -1)
	javascript = strings.Replace(javascript, `t = t.firstChild.href;r = t.match(/https?:\/\//)[0];";`, "", -1)
	javascript = strings.Replace(javascript, `t = t.substr(r.length); t = t.substr(0,t.length-1);`, "", -1)
	javascript = strings.Replace(javascript, `t = t.substr(r.length); t = t.substr(0,t.length-1);`, "", -1)
	javascript = strings.Replace(javascript, `a = document.getElementById('jschl-answer');`, "", -1)
	javascript = strings.Replace(javascript, `f = document.getElementById('challenge-form');`, "", -1)
	javascript = strings.Replace(javascript, `t = t.firstChild.href;`, "", -1)
	javascript = strings.Replace(javascript, `r = t.match(/https?:\/\//)[0];`, "", -1)
	javascript = strings.Replace(javascript, `t = t.substr(r.length);`, "", -1)
	javascript = strings.Replace(javascript, `t = t.substr(0,t.length-1);`, "", -1)
	javascript = strings.Replace(javascript, `t = t.substr(0,15-1);`, "", -1)
	javascript = strings.Replace(javascript, `a.value`, "result", -1)
	javascript = strings.Replace(javascript, `f.action += location.hash;`, "", -1)
	javascript = strings.Replace(javascript, `f.submit();`, "", -1)
	javascript = strings.Replace(javascript, `'; 121'`, "", -1)

	vm := otto.New()
	value, _ := vm.Run(javascript)

	log.Info("结果是： " + value.String())

	return value.String()
}
