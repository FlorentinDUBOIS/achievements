# achievements

> Startup week-end achievements website

## Installation

### Dependencies

To compile and use this project, you will need some dependencies which are :

* GNU Make
* Go language
* Go dep
* Node.js
* Npm

### Install from sources

Firstly, clone the project into the `$GOPATH`:

```bash
git clone git@github.com:FlorentinDUBOIS/achievements.git $GOPATH/src/github.com/FlorentinDUBOIS/achievements
```

Now, go into the fresh install:

```bash
cd $GOPATH/src/github.com/FlorentinDUBOIS/achievements
```

From here, we have to build the user interface and the back-end.

#### Compile the backend

Install Go dependencies:

```bash
make dep
```

Build a release of the project:

```bash
make release
```

A release is now available in directory `dist`.

#### Compile the user interface

For detailed explanation on how things work, checkout the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).

Go into the ui directory:

```bash
cd ui
```

Install node.js dependencies

```bash
npm install
```

Finaly, build the user interface:

```bash
npm run build
```

Now, the user interface is statically compiled in the folder `dist`. This is the folder which should be served.
