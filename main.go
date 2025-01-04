package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/miekg/dns"
)

func main() {
	// Парсинг аргументов командной строки
	domain := flag.String("domain", "", "Доменное имя для запроса")
	typeArg := flag.String("type", "A", "Тип DNS-запроса (A, AAAA, MX, etc.)")
	server := flag.String("server", "127.0.0.1", "DNS-сервер для отправки запроса")
	port := flag.String("port", "53", "Порт DNS-сервера")
	secret := flag.String("secret", "", "Секретный ключ для идентификации клиента")
	flag.Parse()

	if *domain == "" || *typeArg == "" {
		fmt.Println("Использование: ./dns-query-tool -domain <домен> -type <тип> -server <сервер> -port <порт> -secret <секрет>")
		os.Exit(1)
	}

	// Преобразование типа DNS-запроса
	dnsType, ok := dns.StringToType[*typeArg]
	if !ok {
		log.Fatalf("Недопустимый тип запроса: %s", *typeArg)
	}

	// Формирование адреса сервера с портом
	serverAddr := net.JoinHostPort(*server, *port)

	// Создание DNS-запроса
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(*domain), dnsType)
	m.RecursionDesired = true

	// Добавление секрета в OPT-секцию (EDNS)
	if *secret != "" {
		opt := &dns.OPT{
			Hdr: dns.RR_Header{
				Name:   ".",
				Rrtype: dns.TypeOPT,
				Class:  dns.DefaultMsgSize,
			},
		}
		opt.Option = append(opt.Option, &dns.EDNS0_LOCAL{
			Code: 0xFACA, // Локальный код
			Data: []byte(*secret),
		})
		m.Extra = append(m.Extra, opt)
	}

	// Создание клиента и отправка запроса
	c := new(dns.Client)
	r, rtt, err := c.Exchange(m, serverAddr)
	if err != nil {
		log.Fatalf("Ошибка запроса: %v", err)
	}

	// Вывод результата
	fmt.Printf(";; Query time: %v\n", rtt)
	fmt.Printf(";; SERVER: %s\n", serverAddr)
	fmt.Printf(";; MSG SIZE  rcvd: %d\n\n", r.Len())

	fmt.Println(r.String())
}
