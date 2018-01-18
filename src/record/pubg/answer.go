package pubg

import (
	"github.com/robertkrimen/otto"
	"record/log"
	"regexp"
	"strings"
)

var replaceStr = map[string]string{
	`<script type=\"text/javascript\">`:                      ``,
	`//<![CDATA[`:                                            ``,
	`//]]>`:                                                  ``,
	`</script>`:                                              ``,
	`t.length`:                                               `15`,
	`setTimeout(function(){`:                                 ``,
	`}, 4000);`:                                              ``,
	`t = document.createElement('div');`:                     ``,
	`t.innerHTML="<a href='/'>x</a>";`:                       ``,
	`t = t.firstChild.href;r = t.match(/https?:\/\//)[0];";`: ``,
	`t = t.substr(r.length); t = t.substr(0,t.length-1);`:    ``,
	`a = document.getElementById('jschl-answer');`:           ``,
	`f = document.getElementById('challenge-form');`:         ``,
	`t = t.firstChild.href;`:                                 ``,
	`r = t.match(/https?:\/\//)[0];`:                         ``,
	`t = t.substr(r.length);`:                                ``,
	`t = t.substr(0,t.length-1);`:                            ``,
	`t = t.substr(0,15-1);`:                                  ``,
	`a.value`:                                                `result`,
	`f.action += location.hash;`:                             ``,
	`f.submit();`:                                            ``,
	`'; 121'`:                                                ``,
}

func replace(text string) string {

	for key, value := range replaceStr {
		text = strings.Replace(text, key, value, -1)
	}

	return text
}

func getJsAnswer(html string) string {

	log.Info("问答开始")

	reg := regexp.MustCompile(`(?i:setTimeout\(function\(\){)([\w\W]*)}, 4000\);`)
	result := reg.Find([]byte(html))
	javascript := string(result)

	javascript = replace(javascript)
	vm := otto.New()
	value, err := vm.Run(javascript)

	if err != nil {
		log.Info(err)
	}

	log.Info("结果是： " + value.String())

	return value.String()
}
