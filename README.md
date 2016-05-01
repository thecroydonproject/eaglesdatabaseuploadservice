
<p align="center">
  <img src="madeinhot.jpg" width="100"/>
</p>

# dbuplaod


# What is it?

Simple programme that reads json data from datasourcefolder and uploads it into a postgres db.
the database is runnin on a Vagrant machine which have to be started first.  vagrant up builds
the db vm.  boostrap.sh provisions the vm and creates db and user called cpfc in the public schema.

# Progame logic

1) read all files in datasource dir

2) read the content of each file into []byte

3) for each file

    a) unmarshall []byte into json

2) create postgres db table

3) upload content of jsonfile



# how do i use it

1) clone the repo

``
$mkdir temp
$cd temp
$ git clone https://github.com/thecroydonproject/eaglesdatabaseuploadservice.git
``


2)  set-up env variable (db connection string)

  include this in your ~/.profile
```
  export DATABASE_URL_CPFC="postgres://cpfc:cpfc@localhost:15432/cpfc"
```
3) start vagrant machine (postgres db)
```
$ cd postgres
$ vagrant up
```
4) run

``
$ go build && ./dbuplaod
``




## To check the database

```
$ vargrant ssh                                                               // ssh into vm
$ sudo -su postgres                                                      //change user
$ psql                                                                          //run psql
$ \c cpfc                                                                       // connect to cpfc database
cpfc # \dt                                                                     //describe tables in cpfc database
cpfc# \d cpfc                                                                 //describe cpfc table
cpfc# select * from cpfc                                                //check content of cpfc
cpfcf#\q                                                                    \\exit psql
```

 ## What is Made iN H_oT

  software written by UK Home Office Technology will be badged

  with <p align="center"> <img src="madeinhot.jpg" width="100"/>
       </p> as mark of quality, trust and consistance. This badge
       is not yet declared official.




