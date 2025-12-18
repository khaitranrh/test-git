module github.com/khaitranrh/test-git/module1

go 1.25.4

replace github.com/khaitranrh/test-git/module2 => ../module2

require github.com/khaitranrh/test-git/module2 v0.0.0-00010101000000-000000000000
