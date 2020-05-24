# Car Jacked Client

The Client side application for Car Jacked

## Running the Client Locally

Before proceeding make sure to be in the client directory

### Manual Setup

Make sure to make nodejs and yarn installed globally before proceeding.

```bash
  # Install all packages excluding dev dependencies
  yarn install --production

  # Start Development Server
  yarn develop

```

### Using Docker

#### Build Locally and Run

```bash
  docker build -t carjacked-client .
  docker run -p 8000:8000 -t carjacked-client
```

#### Run Directly from Docker Hub Image

```bash
  docker run -p 8000:8000 -t d3vd/carjacked-client
```

## Libraries and Packages Used

- Frontend Framework - [React.js](https://reactjs.org/)
- Static File Generator - [Gatsby.js](https://www.gatsbyjs.org/)
- UI Framework - [Ant Design](https://ant.design/)
- CSS Preprocessor - [Sass](https://ant.design/)
- HTTP Client - [Axios](https://www.npmjs.com/package/axios)
