package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
)

//func main() {
//	resp, err := http.Get(
//		"http://www.zhenai.com/zhenghun")
//	if err != nil {
//		panic(err)
//	}
//
//	defer resp.Body.Close()
//	if resp.StatusCode != http.StatusOK {
//		fmt.Println("Error, status code",
//			resp.StatusCode)
//		return
//	}
//
//	all, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		panic(err)
//	}
//	//fmt.Printf("%s\n", all)
//	printCityList(all)
//}
//func printCityList(contents []byte) {
//	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
//	matches := re.FindAllSubmatch(contents, -1)
//	for _, m := range matches {
//		fmt.Printf("City: %s, URL: %s\n ", m[2], m[1])
//	}
//	fmt.Printf("Matches found: %d\n", len(matches))
//}

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
