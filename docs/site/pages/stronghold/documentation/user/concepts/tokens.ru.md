---
title: "Токены"
permalink: ru/stronghold/documentation/user/concepts/tokens.html
lang: ru
description: Tokens are a core auth method in Stronghold. Concepts and important features.
---

{% alert level="warning" %}Внутренняя структура токенов нестабильна. Структура токенов не документирована и может меняться. Скрипты и автоматизация, которые полагаются на внутреннюю структуру токена, могут перестать работать.
{% endalert %}

Токены могут использоваться напрямую или использузя [методы аутентификации](../../auth).
С их помощью можно динамически создавать токены на основе внешних идентификаторов.

Все внешние механизмы аутентификации привязываются к динамически
создаваемым токенам. Эти токены обладают всеми теми же свойствами, что и обычный токен, созданный вручную.

Внутри Stronghold токены отображаются на информацию. Самая важная информация, сопоставленная
с токеном является набор из одной или нескольких прикрепленных [политик](../policy). Эти политики определяют, что владельцу токена
разрешено делать в Stronghold. Другая отображаемая информация включает
метаданные, которые можно просматривать и которые добавляются в журнал аудита, например
время создания, время последнего обновления и другие.

## Формат токенов

Токены состоят из _префикса_ и _тела_.

- Префикс_ указывает на тип токена:

  Тип токена | Префикс
  -- | --
  Служебные токены | s.
  Пакетные токены | b.
  Токены восстановления | r.

- Тело_ представляет собой случайно сгенерированную строку из 24 или более символов
  [Base62](https://en.wikipedia.org/wiki/Base62) строка.

Ожидается, что токен будет соответствовать следующему regexp: `[sbr]\.[a-zA-Z0-9]{24,}`.

Примеры:

```shell
b.n6keuKu5Q6pXhaIcfnC9cFNd
r.JaKnR2AIHNk3fC4SGyyyDVoQ9O
s.raPGTZdARXdY0KvHcWSpp5wWZIHNT
```

## Типы токенов

Существует два типа токенов: токены `service` и токены `batch`. Далее
содержится подробная информация об их различиях, но сначала полезно разобраться
в других концепциях токенов. Функции, описанные в следующих разделах,
относятся к сервисным токенам, а их их применимость к пакетным токенам
рассматривается позже.

## Хранилище токенов

Часто в документации или в справочных данных упоминается «хранилище токенов».
Это то же самое, что и аутентификация [`token` backend](../../auth/token).
Это особый бэкэнд, поскольку он отвечает за создание и хранение токенов, и
не может быть отключен. Это также единственный метод аутентификации,
который не имеет возможности входа в систему - все действия требуют наличия аутентифицированных токенов.

## Root токены

Root токены - это токены, к которым прикреплена политика `root`. Root
токены могут делать в Stronghold все, что угодно. _Все_. Кроме того, они являются единственным
тип токенов в Stronghold, который может быть настроен так, чтобы срок его действия никогда не истекал
без необходимости продления. Как следствие, создание корневых токенов намеренно затруднено; на самом деле
существует только три способа создания корневых токенов:

1. Root-токен, сгенерированный во время `d8 stronghold operator init` - этот токен не имеет
   срок действия
2. Используя другой Root токен. Root токен, имеющий срок действия не может создать Root токен без срока действия
3. Используя `d8 stronghold operator generate-root`, для чего нужно иметь кворум unseal-ключей

## Иерархия токенов и токены-сироты

Обычно, когда держатель токена создает новые токены, эти токены создаются
как дочерние элементы исходного токена; токены, которые они создают, будут дочерними элементами этих токенов;
и так далее. Когда родительский токен аннулируется, все его дочерние токены - и все
их аренды - также аннулируются. Это гарантирует, что пользователь не сможет избежать
отзыва, просто создав бесконечное дерево дочерних токенов.

Часто такое поведение нежелательно, поэтому пользователи с соответствующим доступом могут создавать
`orphan` токены. У этих токенов нет родителя - они являются корнем своего собственного
дерева токенов. Эти токены-сироты могут быть созданы:

1. С помощью доступа `write` к эндпоинту `auth/token/create-orphan`
2. Имея `sudo` или `root` доступ к `auth/token/create`
   и установив параметр `no_parent` в значение `true`
3. Через роли хранилища токенов
4. С помошью входа через какой либо иной, (не-`token`) метод аутентификации

Пользователи с соответствующими правами могут также использовать эндпоинт `auth/token/revoke-orphan`,
который отзывает заданный токен, но вместо того, чтобы отзывать дочерние
токены дерева, делает их токенами-сиротами.
Используйте с осторожностью!

## Идентификаторы доступа к токену

При создании токенов также создается и возвращается идентификатор доступа к токену.
Идентификатор доступа - это значение, которое действует как ссылка на токен и может быть использовано только для
для выполнения ограниченных действий:

1. Поиск свойств токена (не включая фактический идентификатор токена)
2. Поиск возможностей токена в заданном пути
3. Обновление токена
4. Отзыв токена

Логи аудита могут быть настроены на то, чтобы не обфусцировать идентификатор доступа токенов в журналах аудита.
Это дает возможность быстро отозвать токены в случае чрезвычайной ситуации.
Однако это также означает, что журналы аудита могут быть использованы для проведения более масштабной
атаки типа «отказ в обслуживании».

Наконец, единственный способ «получить список токенов» - это команда `auth/token/accessors`.
которая на самом деле предоставляет список идентификаторов токенов. Хотя это все еще
опасный эндпоинт (поскольку перечисление всех идентификаторов означает, что они могут быть
быть использованы для отзыва всех токенов), она также предоставляет возможность аудита и отзыва
активных в данный момент токенов.

## Время жизни токена, периодические токены и явные максимальные TTL

Каждый токен, не являющийся корневым, имеет связанное с ним время жизни (TTL),
которое представляет собой текущий период действия с момента создания токена или
последнего обновления, в зависимости от того, что наступило позже.
Корневые токены могут иметь связанный TTL, но TTL может быть и 0, что указывает на то,
что срок действия токена никогда не истекает.
После того как TTL истечет, токен больше не будет функционировать - он и связанные с ним
аренды, аннулируются.

Если токен возобновляемый, можно попросить Stronghold продлить срок действия токена
используя команду `d8 stronghold token renew` или обратившись в соответствующий эндпоинт API.
В это время в игру вступают различные факторы. Результат зависит от того, является ли
токен является периодическим доступным для создания пользователями `root`/`sudo`,
ролями хранилища токенов или некоторыми методами аутентификации),
имеет явное максимальное TTL или не имеет ни того, ни другого.

## Общий случай

В общем случае, когда для токена не задан ни период, ни явное максимальное значение
TTL, время жизни токена с момента его создания будет сравниваться с максимальным TTL.
Это максимальное значение TTL генерируется динамически и может меняться от обновления
к обновлению, поэтому значение не может быть отображено при поиске информации о токене.
Оно основано на комбинации факторов:

1. Максимальный TTL по умолчанию в системе составляет 32 дня, но может быть изменен.
2. Максимальный TTL, установленный для монтирования с помощью `mount tuning`.
   Это значение может перекрывать максимальный TTL системы - оно может быть длиннее или короче,
   и если это значение установлено, то оно будет соблюдаться.
3. Значение, предложенное методом аутентификации, который выдал токен. Оно может
   быть настроено для каждой роли, группы или пользователя. Это значение может быть меньше,
   чем максимальное TTL монтирования (или, если оно не задано, максимальное TTL системы),
   но не может быть больше.

Обратите внимание, что значения в (2) и (3) могут измениться в любой момент времени,
поэтому окончательное определение текущего допустимого максимального TTL производится
в момент обновления с использованием текущих значений. Также поэтому важно всегда убеждаться,
что TTL, возвращаемый в результате операции обновления, находится в допустимом диапазоне;
если это значение не расширяется, то, скорее всего, TTL токена не может быть расширен за
пределы его текущего значения, и клиент может захотеть пройти повторную аутентификацию и
получить новый токен. Однако за пределами прямого взаимодействия с оператором Stronghold
никогда не будет отзывать токен до истечения возвращенного TTL.

### Явные максимальные значения TTL

Для токенов может быть явно задано максимальное значение TTL. Это значение становится жестким
ограничением на время жизни токена - независимо от того, каковы значения в (1), (2) и (3)
из общего случая, токен не может жить дольше этого явно установленного значения. Это работет
даже при использовании периодических токенов, чтобы избежать обычного механизма TTL.

### Периодические токены

В некоторых случаях отзыв токена может быть проблематичным - например, если долго работающий
сервис должен поддерживать пул SQL-соединений в течение длительного времени. В этом случае
можно использовать периодический токен. Периодические токены можно создать несколькими способами:

1. Используя доступ `sudo` или `root` токен через эндпоинт `auth/token/create`
2. С помощью ролей хранилища токенов
3. Используя метод аутентификации, который поддерживает их выдачу, например AppRole

В момент выпуска TTL периодического токена будет равен настроенному периоду.
При каждом обновлении TTL будет сбрасываться до этого настроенного периода,
и пока токен успешно обновляется в течение каждого из этих периодов времени,
срок его действия никогда не истечет. За исключением `root` токенов, это единственный способ
для токена в Stronghold иметь неограниченное время жизни.

Идея периодических маркеров заключается в том, что системам и службам легко выполнять
действие относительно часто - например, каждые два часа или даже каждые пять минут.
Поэтому до тех пор, пока система активно обновляет этот токен - другими словами, пока система
жива,- ей разрешается продолжать использовать токен и все связанные с ним аренды.
Однако если система перестанет обновлять токены в течение этого периода (например, если она была выключена),
срок действия токена истечет довольно быстро. Хорошая практика заключается в том,
чтобы этот период был как можно короче, и, вообще говоря, не стоит давать людям периодические токены.

При использовании периодических токенов необходимо знать несколько важных моментов:

- Когда периодический токен создается с помощью роли хранилища токенов,
  во время обновления будет использоваться _текущее_ значение параметра периода роли.
- Токен с периодом и явным максимальным TTL будет действовать как периодический токен,
  но будет отозван по достижении явного максимального TTL

## Токены с привязкой к CIDR

Некоторые токены могут быть привязаны к CIDR, которые ограничивают диапазон клиентских IP-адресов,
которым разрешено их использовать. Это касается всех токенов, кроме неистекающих корневых токенов
(тех, у которых TTL равен нулю). Если срок действия корневого токена истек, на него также распространяется
действие CIDR-привязки.

## Типы токенов в деталях

В настоящее время существует два типа токенов.

### Сервисные токены

Сервисные токены - это то, что пользователи обычно воспринимают как «обычные» токены Stronghold.
Они поддерживают все функции, такие как продление, отзыв, создание дочерних токенов и многое другое.
Соответственно, они тяжелы в создании и отслеживании.

### Пакетные токены

Пакетные токены - это зашифрованные блобы, которые содержат достаточно информации для
использования в действиях Stronghold, но не требуют хранения на диске для их отслеживания.
В результате они чрезвычайно легки и масштабируемы, но лишены большинства гибкостей и
возможностей сервисных токенов.

### Сравнение токенов по типам

Эта справочная таблица описывает разницу в поведении сервсиных и пакетных токенов.

|                                     |           Сервисные токены           |         Пакетные токены          |
| ----------------------------------- | -----------------------------------: | -------------------------------: |
| Могут быть корневыми токенами       | Да                                   | Нет                              |
| Может создавать дочерние токены     | Да                                   | Нет                              |
| Может быть возобновляемым           | Да                                   | Нет                              |
| Может отменяться вручную            | Да                                   | Нет                              |
| Может быть периодическим            | Да                                   | Нет                              |
| Может иметь явный максимальный TTL  | Да                                   | Нет (всегда используется фиксированный TTL)|
| Имеет аксессоры                     | Да                                   | Нет                              |
| Имеет Cubbyhole                     | Да                                   | Нет |
| Отзывается с помощью родителя (если не сирота) | Да                        | Перестает работать               |
| Отзывается с помощью родителя (если не сирота) | Да                        | Перестает работать               |
| Назначение аренды динамических секретов | Самостоятельно                   | Родитель (если не сирота)        |
| Стоимость                           | Тяжеловесная; многократная запись в хранилище при создании токена | Легкая; нет затрат на хранение при создании токена |

### Сервисная и пакетная обработка аренды токенов

#### Сервисные токены

Аренды, созданные сервисными токенами (включая аренды дочерних токенов), отслеживаются вместе с ними
и аннулируются по истечении срока действия токена.

#### Пакетные токены

Аренды, созданные пакетными токенами, ограничены оставшимся TTL пакетных токенов и, если
пакетный токен не является сиротой, отслеживаются его родителем. Они аннулируются, когда
TTL пакетного токена истекает, или когда родитель пакетного токена аннулируется (в этот
момент пакетному токену также отказывают в доступе к Stronghold).
