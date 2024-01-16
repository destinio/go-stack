# Go Stack

## Backend

### fiber
- https://gofiber.io/
- https://github.com/gofiber/fiber 

**Install**
`go get -u github.com/gofiber/fiber/v3`

```go
package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// static files
	app.Static("/", "./public")

	app.Get("/", func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString("<h1>Hello, World!</h1>")
	})

	app.Listen(":1337")
}
```

## Database

### Gorm (ORM)
- https://gorm.io/

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

```go
// for code examples see:
// https://gorm.io/docs/#Quick-Start

db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
if err != nil {
  panic("failed to connect database")
}
```

## Frontend

### HTMX
- https://htmx.org/

```html
<script src="https://unpkg.com/htmx.org@1.9.10"></script>
<!-- have a button POST a click via AJAX -->
<button hx-post="/clicked" hx-swap="outerHTML">
  Click Me
</button>
```

### Fiber Templates
- https://docs.gofiber.io/template/html/
- https://docs.gofiber.io/template/html/TEMPLATES_CHEATSHEET
- https://docs.gofiber.io/guide/templates/
- https://github.com/gofiber/template

### Tailwind CSS
- https://tailwindcss.com/

```bash
npx tailwindcss init
```

```css
/* ./styles/tailwind.css */

@tailwind base;
@tailwind components;
@tailwind utilities;
```

```js
// ./tailwind.config.js
module.exports = {
  content: ["./views/**/*.html"], // add other files types here
  theme: {
    extend: {},
  },
  plugins: [],
}
```

**scripts**

```makefile
# ./makefile

tw-b:
	@npx tailwindcss -i ./styles/tailwind.css -o ./public/main.css

tw-w:
	@npx tailwindcss -i ./styles/tailwind.css -o ./public/main.css --watch
```
