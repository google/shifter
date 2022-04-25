# Shifter Web UI

The Shifter Web UI is written in Vue3 with Vite as the Development server.

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```

### Run in Docker Container with [ESLint](https://eslint.org/)

```sh
docker run -p 8085:8080 -e SHIFTER_SERVER_ENDPOINT="https://api.shifter.cloud/api/v1/" images.shifter.cloud/shifter-ui:latest
```
