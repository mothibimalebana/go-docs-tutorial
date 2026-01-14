module example/hello

go 1.25.5

replace example/greetings => ../greetings

require example/greetings v0.0.0-00010101000000-000000000000

require golang.org/x/example/hello v0.0.0-20250915201037-7f05d217867b // indirect
