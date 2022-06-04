# CV Maker

Allows you to generate a CV using a information defined in a json file.

## Generating CV

**Prerequisites**

* `resume.json` file
* `profile.jpg` image

Creates a json file called, `resume.json`, and put this json file in the root directory of this project.

```
make run
```

This will generate a file `resume.pdf` with your CV.
Take a look at the file `example.json` to see an example of a valid `resume.json` file.

# CV example

Checkout a CV example [here](https://mega.nz/file/gZdFCI4K#l3rB7rohKi2Rc4ZSsSyfl0003y5oBa8O0hD08jMirrk).


# TODO

* Add unit tests
* Add more documentation about the format of the json configuration file.
* Make the script more customizable.
* Make a simple website to allow the generation of this cv.

# Downside

Currently I only have one template.
