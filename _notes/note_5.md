# ООП и ораганизация объектов в GO

## Какие принципы ООП реализованы в GO?

В языке GO реализованы все принципы ООП, но уникальным способом.

1. **Инкапсуляция** реализуется через строчные и заглавные буквы. Методы с заглавной буквы доступны вне пакета.

2. **Абстракция** достигается с помощью интерфейсов, который представляет набор сигнатур. Реализация интерфейсов происходит неявно.

3. **Полиморфизм**. Способность функции обрабатывать данные разных типов.

Go достигает полиморфизма через интерфейсы. Любой тип, реализующий интерфейс, может быть использован везде, где ожидается этот интерфейс

```golang
type Interface interface {
    Method(in string) (out string)
}

type Struct struct {}

func (s Struct) Method(in string) (out string) { ... }

type OtherStruct struct {}

func (s OtherStruct) Method(in string) (out string) { ... }

//------------------------//
func main() {
    var ex1 Interface = Struct{}
    var ex2 Interface = OtherStruct{}
    
    ex1.Mathod()
    ex2.Method()

}
```

4. Go не поддерживает классическое наследование. Вместо этого язык использует композицию через встраивание структур (struct embedding)

```golang
type Hero struct {
    team string
}

func (h Hero) SayTeam() {
    fmt.Println("My Team is", h.team)
}

type Batman struct {
    Hero         // встраивание (embedding)
    name string
}

func (b Batman) SayImBatman() {
    fmt.Printf("I'm %s and I'm Batman\n", b.name)
}
```

## Что такое интерфейс?

Программная синтаксическая структура, определяющая отношение между объектами, которые разделяют определенное поведенческое множество и не связаны никак иначе. 

```golang
type SomeInterface interface {
    SomeMethod(in string) (out string, err error)
}
```

## Какие есть свойства интерфейсов?

1. Интерфейс - это конкретная структура
```golang
type iface struct {
    tab *itab
    data unsafe.Pointer
}
```

2. Создать не nil интерфейс нельзя
3. За каждым не-nil интерфейсом стоит конкретный тип


## Best practice использования интерфейсов в GO

1. Описываются там, где используются
2. Должны содержать только то, что используется


## Какие интерфейсы стандартной библиотеки вы знаете?

пакет io (Read, Write, ..), error


