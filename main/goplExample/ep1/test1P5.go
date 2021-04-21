package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	HTTP_PREFIX = "http://"
	TMP_PATH    = "/Users/cb-000008/GoStuff/src/goTest/tmp/"
)

func main() {
	for _, url := range os.Args[1:] {
		if len(url) > 0 {
			url, err := addHttpPrefixIfNot(url) //1p5:2
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error：%v by %s\n", err, url)
				continue
			}
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fentch : %v\n", err)
			os.Exit(1)
		}

		respStatus := resp.Status
		fmt.Printf("url: %s status: %s \n", url, respStatus)
		/*b,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr,"fentch : reading %s : %v \n",url,err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)*/
		///////////1p5:1////////////////
		/*b,err := os.Create(TMP_PATH + strconv.FormatInt(time.Now().Unix(),10)+".txt")
		if err != nil{
			fmt.Printf("Error:Create dest file faild! : %s",err)
			os.Exit(1)
		}*/                               //创建临时文件
		wt := bufio.NewWriter(os.Stdout)  //直接写到标准输出里
		n, cerr := io.Copy(wt, resp.Body) //io.Copy因为以一块32KB的缓冲区顺序拷贝，所以相较于ioutil.RedAll，在读取大文件时不会占用过大内存
		resp.Body.Close()
		if cerr != nil {
			fmt.Printf("Error: Copy Faild : %s \n", cerr)
			os.Exit(1)
		}
		//////////////////////////////
		fmt.Println("write", n, "hi")
	}
}

func addHttpPrefixIfNot(s string) (string, error) {
	var err error
	ret := s
	if !strings.HasPrefix(s, HTTP_PREFIX) {
		sArr := strings.Split(s, "://")
		switch len(sArr) {
		case 0:
			err = errors.New("Empty sArr!")
			return "", err
		case 1:
			ret = HTTP_PREFIX + sArr[0]
		case 2:
			ret = HTTP_PREFIX + sArr[1]
		default:
			err = errors.New("Unexpected sArr length! \n")
			return "", err
		}
	}
	fmt.Printf("%s\n", ret)
	return ret, err
}
