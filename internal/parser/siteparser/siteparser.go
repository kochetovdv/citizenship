package siteparser

import (
	"citizenship/internal/order"
	//"citizenship/pkg/regulars"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// SiteParser is an structure for downloading orders from a website
type SiteParser struct {
	path string
	url  string
}

// NewSiteParser creates a new SiteParser
func NewSiteParser(url, path string) *SiteParser {
	s := new(SiteParser)
	s.path = path
	s.url = url
	return s
}

// Site parsing
func (s *SiteParser) Parse(url string) *order.Orders {
	response, _ := s.connect()
	listOfOrders := s.extractData(response)

	listOfOrders.Print()
	listOfOrders.Statistics()
	// TODO Optimize for just new file checking
	for _, el := range listOfOrders.Orders {
			err := s.DownloadFile(el.Filename, el.Link)
			if err != nil {
				log.Printf("Error with downloading: %s", err)
			}
		}
		return listOfOrders 
}

// TODO add error handling, change return data type to Orders.
// TODO Replace to pkg/web-services
func (s *SiteParser) connect() ([]byte, error) {
	// TODO after refactoring, change to use pkg/web-services
	// response, err:= connect(s.url)
	log.Printf("Connecting: %s", s.url)

	resp, err := http.Get(s.url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	// TODO add error handling, change return
	if resp.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	// TODO add error handling, change return
	if err != nil {
		log.Print(err)
	}

	return body, nil
}

func (d *SiteParser) extractData(body []byte) *order.Orders {
	orders := order.NewOrders()

	html := string(body)
	htmlLines := strings.Split(html, "\n")
	var htmlLi []string

	for _, line := range htmlLines {
		if strings.Contains(line, "<li>") {
			htmlLi = append(htmlLi, line)
		}
	}

	for _, line := range htmlLi {
		dateRegExp := regexp.MustCompile(`<strong>([^<]+)</strong>`)

		date := dateRegExp.FindString(line)
		date = strings.Replace(date, "<strong>", "", -1)
		date = strings.Replace(date, "</strong>", "", -1)

		linksRegExp := regexp.MustCompile(`<a href="([^"]+)">([^<]+)</a>`)
		links := linksRegExp.FindAllStringSubmatch(line, -1)

		for _, link := range links {
			filename := filepath.Base(link[1])
			order := order.NewOrder(
				date,
				filename,
				link[1],
				link[2],
			)
			orders.Add(order)
		}
	}

	return orders

}

/*
// Create orders from links
func (d *SiteParser) CreateOrders(links []string) ([]order.Order, error) {

	//TODO Use model order.Orders in this block (new, add, etc)
	var data []order.Order
	for _, line := range lines {
		tmpData, err := s.ExtractData(line)
		if err != nil {
			log.Print(err)
			continue
		}
		data = append(data, tmpData...)
	}

	for _, el := range data {
		log.Printf("Date:%s\tFilename:%s\tLink:%s\tNumber:%s\n", el.Date, el.Filename, el.Link, el.Number)
	}

}*/

// TODO replace to pkg/downloader/downloader.go
func (d *SiteParser) DownloadFile(fileName, url string) error {
	// Create the directory if it does not exist
	if _, err := os.Stat(d.path); os.IsNotExist(err) {
		err = os.MkdirAll(d.path, 0755)
		if err != nil {
			return err
		}
	}

	// Check if the file already exists in the path
	filePath := filepath.Join(d.path, fileName)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		log.Printf("File %s already exists\n", fileName)
		return nil
	}

	// Download the file from the URL
	log.Printf("Starting download file from %s\n", url)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create a new file and save the response body to it
	log.Printf("Saving file %s\n", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	log.Printf("File %s downloaded to %s\n", fileName, filePath)
	return nil
}

/*
Есть статичная страница HTML, на которой имеются li элементы следующего вида:
<li>Data de <strong>23.12.2022 </strong>numărul:
<a href="http://cetatenie.just.ro/wp-content/uploads/2022/01/Ordin-1440P-23-12-2022-NPE.pdf">1440P</a></li>
Дата есть всегда в тегах strong. Количество ссылок может быть любым - от 1 до бесконечности.
Текст 1440P в данном примере - номер дела, каждая ссылка относится именно к номеру дела.
У одного номера дела - одна ссылка. Номера дел относятся к датам - у одной даты может быть много дел.
Напиши компилируемый код программы на golang, который парсит страницу html и складывает в массив структур даты,
номера дел и ссылки на эти номера дел.
*/

/*
напиши компилируемый код на golang, который получает текст сайта и ищет в нем данные, складывая их в структуру. Каждый экземпляр данных должен содержать дату, имя файла PDF, ссылку скачивания и номер в формате "ЧислоP".
Данные выглядят следующим образом:
<li>Data de <strong>03.02.2023 </strong>numărul: <a href="http://cetatenie.just.ro/wp-content/uploads/2022/01/Ordin-132-art-11-288-pers.pdf">132P</a> <a href="http://cetatenie.just.ro/wp-content/uploads/2022/01/Ordin-134-art-11-101-pers-.pdf">134P</a> <a href="http://cetatenie.just.ro/wp-content/uploads/2022/01/Ordin-135-art-11-293-pers-.pdf">135P</a></li>
На одну дату может приходиться минимум одна связка файл+ссылка+номер, но их может быть и много, надо это предусмотреть
*/
