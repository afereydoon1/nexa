package lang

import (
	"encoding/json"
	"os"
	"strings"
)

type Translator struct {
	messages map[string]map[string]string
}

func NewTranslator() *Translator {
	return &Translator{
		messages: map[string]map[string]string{
			"en": load("internal/infra/lang/en.json"),
			"fa": load("internal/infra/lang/fa.json"),
		},
	}
}

func (t *Translator) Translate(lang, code string) string {

	lang = normalizeLang(lang)

	if msg, ok := t.messages[lang][code]; ok {
		return msg
	}

	return code
}

func normalizeLang(lang string) string {

	if strings.Contains(lang, ",") {
		lang = strings.Split(lang, ",")[0]
	}

	if strings.Contains(lang, "-") {
		lang = strings.Split(lang, "-")[0]
	}

	if lang == "" {
		lang = "en"
	}

	return strings.ToLower(lang)
}

func load(path string) map[string]string {

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var data map[string]string

	if err := json.Unmarshal(file, &data); err != nil {
		panic(err)
	}

	return data
}
