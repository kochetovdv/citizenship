package downloader

import (
	"citizenship/internal/order"
	"fmt"
	_ "html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Downloader struct {
	path string
	url  string
}

func NewDownloader(url, path string) *Downloader {
	d := new(Downloader)
	d.path = path
	d.url = url
	return d
}

func (d *Downloader) Download() []order.Order {
	log.Printf("Downloading %s", d.url)

	resp, err := http.Get(d.url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	html := string(body)
	lines := strings.Split(html, "\n")
	var data []order.Order
	for _, line := range lines {
		tmpData, err := d.ExtractData(line)
		if err != nil {
			log.Print(err)
			continue
		}
		data = append(data, tmpData...)
	}

	//	fmt.Printf("%s", body)

	for _, el := range data {
		fmt.Printf("Date:%s\tFilename:%s\tLink:%s\tNumber:%s\n", el.Date, el.Filename, el.Link, el.Number)
	}

// TODO Optimize for just new file checking
	for _, el := range data {
		err := d.DownloadFile(el.Filename, el.Link)
		if err != nil {
			fmt.Println(err)
		}
	}
	return data
}

func (d *Downloader) ExtractData(htmlLine string) ([]order.Order, error) {
	dateRegExp := regexp.MustCompile(`<strong>([^<]+)</strong>`)
	tempDate := dateRegExp.FindStringSubmatch(htmlLine)
	if len(tempDate) < 1 {
		return nil, fmt.Errorf("no date found in %s", htmlLine)
	}
	date := strings.Replace(tempDate[0], "<strong>", "", 1)
	date = strings.Replace(date, "</strong>", "", 1)
	date = strings.TrimSpace(date)

	linksRegExp := regexp.MustCompile(`<a href="([^"]+)">([^<]+)</a>`)
	links := linksRegExp.FindAllStringSubmatch(htmlLine, -1)
	if len(date) < 1 {
		return nil, fmt.Errorf("no date found in %s", links)
	}

	var orders []order.Order
	for _, link := range links {
		filename := strings.Split(link[1], "/")[len(strings.Split(link[1], "/"))-1]
		orders = append(orders, order.Order{
			Date:     date,
			Filename: filename,
			Link:     link[1],
			Number:   link[2],
		})
	}

	return orders, nil
}

func (d *Downloader) DownloadFile(fileName, url string) error {
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

	fmt.Printf("File %s downloaded to %s\n", fileName, filePath)
	return nil
}

/*
   	tmpl, err := template.New("data").Parse(`
   	{{range .}}
   	<li>Data de <strong>{{.Date}}</strong> numărul: <a href="{{.Link}}">{{.CaseNumber}}</a></li>
   	{{end}}
     `)
     if err != nil {
   	  // handle error
     }


     data := []Data{
   	  {Date: "23.12.2022", CaseNumber: "1440P", Link: "http://cetatenie.just.ro/wp-content/uploads/2022/01/Ordin-1440P-23-12-2022-NPE.pdf"},
   	  // add more data here
     }


     err = tmpl.Execute(os.Stdout, data)
     if err != nil {
   	  // handle error
     }
*/

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
