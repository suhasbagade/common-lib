package httpclient

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	apihost      = "dataservice.accuweather.com"
	apiurlscheme = "https"
)

var apikey = os.Getenv("weatherapikey")

func Getbaseurl() url.URL {

	var querryparam = "apikey=" + apikey

	Baseurl := url.URL{
		Scheme:   apiurlscheme,
		Host:     apihost,
		RawQuery: querryparam,
	}

	return Baseurl
}

func Getrequest(requesturl string) []byte {
	client := http.Client{}
	res, err := client.Get(requesturl)
	if err != nil {
		log.Fatalln(err)
	}
	resbody, errr := ioutil.ReadAll(res.Body)
	if errr != nil {
		log.Fatalln(errr)
	}

	defer res.Body.Close()

	return (resbody)

}
