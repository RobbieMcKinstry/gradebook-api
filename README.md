# Rake Commands

Rake comes baked into the Vagrant environment. If you're trying to start related to the API, there's probably a `rake` command for it.

For example, if you want to launch the kubernetes dashboard, you can run the rake task `rake k8s` from inside the Vagrantbox, and then visit `http://localhost:8001/ui` to view the dashboard.

Other rake commands include:

- `rake build`: This command will compile the Go source code. Can be run from the VM or the host.
- `rake`:       This command will regeneate the Goa boilerplate and compile the Go source code. Must be run from within the VM.
- `rake clean`: This deletes the binary created with `rake build`. Can be run from either the VM or the host.
- `rake go_generate`: This command regenerates the Goa boilerplate. Must be run from within the VM.

There are other aliases for these commands, but they're internally a little bit longer than what I have here.

# Vagrant

I'm assuming you have [basic knowledge](https://www.vagrantup.com/docs/getting-started/) of Vagrant. Install Vagrant, then run `vagrant up` to start the VM. If you're shared folders aren't mounting because there's a problem with your Guest Additions, run `vagrant plugin install vagrant-vbguest`. That should solve your problem.

# MySQL

If you want to connect to the database from the command line, you need to tell the MySQL client to use `127.0.0.1` instead of localhost.

	mysql -h 127.0.0.1 -P 3306 -u root -p

The password is `root`. Check out the username and password settings in `.env`, and specified in the `Vagrantfile`.

Copy and paste that command from inside the Vagrant box to connect to MySQL. Then run `use alligrader` to switch to the right database. If you want to see the tables that are created, run `show tables;`. Easy enough, right? If you fucked up something in your database and want a fresh start, run `drop database alligrader; create database alligrader;`.

There's also a command line tool `goose` for switching to the right migration. `goose up` will apply all migrations. `goose down` will unmigrations (or the last migration? I'm not sure. See the goose documentation for more info).

To dump the database contents out into a file, or make a barebones schema:

    mysqldump --compatible=ansi --add-drop-database --skip-comments --add-drop-table --disable-keys --skip-set-charset --host=127.0.0.1 -u root --password=root shaman > mysqldump.sql

# Golang

To run the tests, execute `go test ./spec/...`. The `...` is a special syntax for recursively running tests.

To start the webserve, run `go run main.go`.
To build the binary, run `go build main.go`.

To run the Dredd tests, run `go build main.go` and then run `dredd`. Before submitting a PR, your commit should pass both the `dredd` tests and the `go test` tests. 

I'm going to try to use `rake` to alias these commands. That's on my TODO list.

# Tools for your Host Machine

For your host machine, you should probably install [Virtualbox](https://www.virtualbox.org/wiki/Downloads), [Vagrant](https://www.vagrantup.com/downloads.html) (and Ruby), and [git](https://git-scm.com/downloads).


