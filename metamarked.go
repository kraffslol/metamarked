package metamarked

import (
	"io/ioutil"
	"fmt"
	"strings"
	"regexp"
	"gopkg.in/yaml.v2"
	"errors"
	"os"
)

type Metadata struct {
	Title string
	Route string
}

type Data struct {
	Metadata
	Markdown string
}

func splitInput(content string) ([]string, error) {
	if (!strings.HasPrefix(content, "---")) {
		return nil, errors.New("File doesn't contain metadata")
	}

	re := regexp.MustCompile(`\n(\-{3})`)
	res := re.FindStringIndex(content)

	yamlData := content[4:res[0]]
	markdownData := content[res[1]:len(content)]

	result := []string{}
	result = append(result, yamlData)
	result = append(result, markdownData)

	return result, nil
}

// GetMetaAndMarkdown : Gets YAML Metadata & Markdown from file.
func GetMetaAndMarkdown(path string) (Data, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Data{}, errors.New("File doesn't exist")
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return Data{}, err
	}

	metaAndMarkdown, err := splitInput(string(file))
	if err != nil {
		return Data{}, err
	}

	yamlData := metaAndMarkdown[0]

	data := Data{
		Markdown: metaAndMarkdown[1],
	}

	err = yaml.Unmarshal([]byte(yamlData), &data.Metadata)
	if err != nil {
		return Data{}, err
	}

	fmt.Println(data)
	return data, nil
}