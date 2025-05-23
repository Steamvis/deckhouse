---
title: "Метод OIDC"
permalink: ru/stronghold/documentation/user/auth/oidc.html
lang: ru
---

### Метод аутентификации JWT/OIDC

Метод `jwt auth` может быть использован для аутентификации в Deckhouse Stronghold с помощью OIDC или путем предоставления JWT.

Метод OIDC позволяет выполнять аутентификацию через настроенный провайдер OIDC с помощью веб-браузера пользователя. Этот метод может быть инициирован из пользовательского интерфейса Deckhouse Stronghold или из командной строки. В качестве альтернативы может быть предоставлен непосредственно JWT. JWT криптографически проверяется с помощью локально предоставленных ключей, либо, если настроен, для получения соответствующих ключей может быть использован сервис OIDC Discovery. Выбор метода настраивается для каждой роли.

Оба метода позволяют дополнительно обрабатывать данные утверждений в JWT. В документе рассмотрены концепции, общие для обоих методов, а также примеры использования OIDC и JWT.

#### Аутентификация в OIDC

В данном разделе рассматривается настройка и использование ролей OIDC. Предполагается базовое знакомство с концепциями OIDC. Поток Authorization Code использует расширение Proof Key for Code Exchange (PKCE).

Deckhouse Stronghold включает два встроенных потока авторизации OIDC: пользовательский интерфейс Deckhouse Stronghold UI и CLI с использованием логина vault.

##### Перенаправление URI

Важной частью конфигурации ролей OIDC является правильная настройка URI перенаправления. Это должно быть сделано как в Deckhouse Stronghold, так и в провайдере OIDC, причем эти настройки должны совпадать. URI перенаправления задаются для роли с помощью параметра `allowed_redirect_uris`. Для настройки потоков Deckhouse Stronghold UI и CLI существуют различные URI перенаправления, поэтому в зависимости от установки необходимо настроить один или оба.

##### CLI

Если планируется поддержка аутентификации через `d8 stronghold login -method=oidc`, необходимо задать URI перенаправления на localhost. Обычно это может быть: `http://localhost:8250/oidc/callback`. При необходимости при входе в систему через CLI можно указать другой хост и/или порт прослушивания, при этом URI с этим хостом/портом должен совпадать с одним из настроенных перенаправляемых URI. Эти же URI "localhost" должны быть добавлены и в провайдер.

#### Deckhouse Stronghold UI

Для входа в систему через Deckhouse Stronghold сам UI не требует настройки и конфигурируется автоматически при включении Deckhouse Stronghold

##### Вход в систему OIDC (Deckhouse Stronghold UI)

Выберите метод входа в систему "OIDC".

При необходимости введите имя роли.

Нажмите кнопку "Войти с помощью провайдера OIDC" и завершите аутентификацию с помощью настроенного провайдера.

##### Вход в систему OIDC (CLI)

Для входа в систему CLI по умолчанию используется путь `/oidc_deckhouse`. Если данный метод аутентификации был включен по другому пути, укажите в CLI путь `-path=/my-path`.

```shell
d8 stronghold login -method=oidc -path=oidc_deckhouse role=test
Complete the login via your OIDC provider. Launching browser to:
https://myco.auth0.com/authorize?redirect_uri=http%3A%2F%2Flocalhost%3A8400%2Foidc%2Fcallback&client_id=r3qXc2bix9eF...
```

Браузер откроется по сгенерированному URL-адресу для завершения входа в систему провайдера. URL может быть введен вручную, если браузер не может быть открыт автоматически.

Автоматический запуск браузера по умолчанию на URL переключается при помощи `skip_browser` (по умолчанию: "false") при входе в систему.

Слушатель обратного вызова может быть настроен с помощью следующих необязательных параметров. Обычно их установка не требуется:

* mount (по умолчанию "oidc")
* listenaddress (по умолчанию "localhost")
* port (по умолчанию 8250)
* callbackhost (по умолчанию "localhost")
* callbackmethod (по умолчанию "http")
* callbackport (по умолчанию - это значение, заданное для порта). Это значение используется в параметре `redirect_uri`, тогда как `port` - это порт сервера localhost, который принимает запросы. В некоторых случаях эти два параметра могут различаться.

##### Конфигурация провайдера OIDC

Способ аутентификации OIDC успешно протестирован с рядом провайдеров. Полное руководство по настройке приложений OAuth/OIDC не входит в документацию Deckhouse Stronghold.

Некоторые советы по устранению неполадок в конфигурации OIDC и настройке OIDC представлены ниже:

* Если параметр роли (например, `bound_claims`) требует значения карты (map), его нельзя установить отдельно с помощью Deckhouse Stronghold CLI. В таких случаях запишите конфигурацию в виде одного JSON-объекта:

  ```shell
  d8 stronghold write auth/oidc/role/demo -<<EOF
  {
  "user_claim": "sub",
  "bound_audiences": "abc123",
  "role_type": "oidc",
  "policies": "demo",
  "ttl": "1h",
  "bound_claims": { "groups": ["mygroup/mysubgroup"] }
  }
  EOF
  ```

* Проследите за выводом журнала Deckhouse Stronghold, в котором содержится важная информация о сбоях проверки OIDC.

* Убедитесь, что URI перенаправления корректны в Deckhouse Stronghold и на провайдере. Они должны точно совпадать. Проверьте: http/https, 127.0.0.1/localhost, номера портов, наличие слэшей в конце.

* Единственной конфигурацией утверждения, которая требуется для роли, является user_claim. Когда станет известно, что аутентификация работает, добавьте дополнительные привязки утверждений и копирование метаданных.

* bound_audiences не является обязательным для ролей OIDC и обычно не требуется. Идентификаторы клиентов OIDC используют client_id для определения аудитории, и проверка OIDC предполагает это.

* Уточните у провайдера диапазоны для получения необходимой информации. Часто требуется запрашивать диапазоны "профиль" и "группы", которые можно добавить, установив для роли значение `oidc_scopes="profile,groups"`.

* Если в журналах появляются ошибки, связанные с заявками, внимательно изучите документацию поставщика, чтобы понять, как он называет и структурирует заявки. В зависимости от провайдера, сконструируйте простой запрос `curl implicit grant` для получения JWT, который можно просмотреть.

Пример декодирования JWT (в примере он расположен в поле `access_token` JSON-ответа) представлен ниже:

  ```shell
  cat jwt.json | jq -r .access_token | cut -d. -f2 | base64 -D
  ```

* В Deckhouse Stronghold доступна ролевая опция `verbose_oidc_logging`, которая будет записывать полученный OIDC-токен в журналы сервера, если включено ведение журнала на уровне отладки. Это может быть полезно при отладке настройки провайдера и проверке того, что полученные претензии соответствуют ожиданиям. Поскольку данные утверждений записываются в журнал дословно и могут содержать конфиденциальную информацию, эту опцию не следует использовать в производстве.
