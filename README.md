# go-gofr-template
a boilerplate microservice template using golang and gofr.

This repo is intended to be a template to build different microservices.

### setup

copy configs/.sample.local.env to configs/.env, then change settings as appropriate for your local setup

open migrations/20240720174600_create_user_table_and_seeds.go. Scroll down to SeedUsers() and change
the seed users to have the values you want.
```shell
	seedUsers := [][]string{
		{"god", "almighty", "admin@johnscode.com", "god"},
		{"dev", "johnscode", "dev@johnscode.com", "dev"},
		{"john", "code", "john@johnscode.com", "j"},
	}
```
This repo is intended to be a template to build different microservices. I have used a user table as an example.
Depending on your process, you may choose to handle seeding your database differently, particularly a users
table. It is bad practice to include passwords in a repository and will cause problems if your service is 
subject to certain compliance regimes.

### setting up database seeds and migrations

See the gofr docs on [migrations]()https://gofr.dev/docs/advanced-guide/handling-data-migrations

in short,

- create a go file in the migrations folder with a name in form YYYYMMDDHHMMSS_what_your_doing.go
```shell
    touch 20240226153000_do_some_stuff.go
```
- add a function to encapsulate what you are doing:

 ```
    func doSomeStuffWithDB() migration.Migrate {
	    return migration.Migrate{
		    UP: func(d migration.Datasource) error {
			    _, err := d.SQL.Exec(someSqlQuery)
			    if err != nil {
				    r   eturn err
			    }
			    return nil
		    },
	    }
    }
 ```
- modify the map in migrations/all.go
```
    func All() map[int64]migration.Migrate {
        return map[int64]migration.Migrate{
            20240226153000: doSomeStuffWithDB(),
        }
    }
```
