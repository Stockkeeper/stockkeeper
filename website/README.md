# Stockkeeper: Website

### Installation

This step will install the necessary NodeJS dependencies for developing on the website. A list of NodeJS dependencies can be found in the file `package.json`.

```
$ npm install
```

### Local Development

This step will start a local development server and open the website with your default browser. When the server detects a change in your code, it will automatically re-build the website and refresh your browser.

```
$ npm run start
```

### Build

This step will build the website, producing the static content in a `/build` directory.

```
$ npm run build
```

### Deployment

This step will build _and_ deploy the website to the public.

Once the build is complete, the website will be pushed to the remote branch `gh-pages`. Then, GitHub Pages will host the website at the public domain `https://stockkeeper.io`.

To deploy, you must have `push` permissions on the this GitHub repository `https://github.com/Stockkeeper/stockkeeper`.


```
$ GIT_USER=<Your GitHub username> npm run deploy
```
