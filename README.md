# Cassie

A simple cli utility to change the case style in your document. It supports several languages

![Go](https://img.shields.io/badge/Go-0074cc)
![JavaScript](https://img.shields.io/badge/JavaScript-f7df1e)
![C++](https://img.shields.io/badge/C%2B%2B-00599c)
![PHP](https://img.shields.io/badge/PHP-777bb3)
![C](https://img.shields.io/badge/C-004482)
![Java](https://img.shields.io/badge/Java-007396)
![Python](https://img.shields.io/badge/Python-3776ab)
![Rust](https://img.shields.io/badge/Rust-000000)
![Perl](https://img.shields.io/badge/Perl-0298c3)
![Bash](https://img.shields.io/badge/Bash-4e9a06)
![Ruby](https://img.shields.io/badge/Ruby-cc0000)
![Swift](https://img.shields.io/badge/Swift-f05138)
![Kotlin](https://img.shields.io/badge/Kotlin-7f52ff)
![TypeScript](https://img.shields.io/badge/TypeScript-3178c6)
![Objective-C](https://img.shields.io/badge/Objective--C-7a1f7d)
![Lua](https://img.shields.io/badge/Lua-2c2d72)
![R](https://img.shields.io/badge/R-276dc3)
![Scala](https://img.shields.io/badge/Scala-DC322F)
![Haskell](https://img.shields.io/badge/Haskell-5e5086)
![Elixir](https://img.shields.io/badge/Elixir-6e4a7e)
![Julia](https://img.shields.io/badge/Julia-9558b2)
![F#](https://img.shields.io/badge/F%23-7b4f96)
![Dart](https://img.shields.io/badge/Dart-00b4b4)
![Shell](https://img.shields.io/badge/Shell-1e1e1e)
![Groovy](https://img.shields.io/badge/Groovy-4298b8)


## Usage

```bash
go run main.go <file-name> <options> 
```

### Example

```bash
go run main.go test_assets/camel.js cs
```
```js
class user_account_test {
    constructor(user_name, user_email) {
      this.user_name = user_name;
      this.user_email = user_email;
      this.is_logged_in = false;
    }
  
    log_in() {
      this.is_logged_in = true;
      console.log(`${this.user_name} has logged in.`);
    }
  
    log_out() {
      this.is_logged_in = false;
      console.log(`${this.user_name} has logged out.`);
    }
  
    get_user_info() {
      return {
        name: this.user_name,
        email: this.user_email,
        status: this.is_logged_in ? "online" : "offline"
      };
    }
  }
```

### Options
It's very simple, it uses BSD-like syntax

```bash
ps [from pascal to snake]
```
There are three options

```text
p : PascalCase
s : snake_case
c : camelCase
```

## Languages 

The program will detect the language of the input file, and perform a case conversion

> **Warning:**  
>
> If you are using standard or other librsries they'll be changed as well. 
>
>
> **std::for_each(..)**
>
> _will become_ 
>
> **std::forEach(..)**
>
> So, make sure you are changing the case to the one most appropriate for your language.
