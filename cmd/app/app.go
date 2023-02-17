package main

import (
	"citizenship/internal/clients/app"
	_ "golang.org/x/net/context"
)

func main() {
	//	logger:=logger.NewLogger()
	//	ctx:=context.Background()
	app := app.NewApp()
	app.Run()
}

// download url		[ ]
// parse url		[ ]
// download pdf		[ ]
// parse pdf		[ ]

// Ctrl+L - select whole line
// Alt+J - find next such word with multiple cursor
// Ctrl+Shift+Arrow - move line up/down
// Ctrl+/ - (un)comment selection
// Ctrl+D - duplicate line
// Ctrl+Alt+L - format

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
